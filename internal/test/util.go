package test

import (
	"context"
	"fmt"
	v1 "gitee.com/moyusir/data-processing/api/dataProcessing/v1"
	"gitee.com/moyusir/data-processing/internal/biz"
	"gitee.com/moyusir/data-processing/internal/conf"
	"gitee.com/moyusir/data-processing/internal/data"
	utilApi "gitee.com/moyusir/util/api/util/v1"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-redis/redis/v8"
	"google.golang.org/protobuf/proto"
	"os"
	"testing"
)

// 提供给测试程序使用的通用初始化函数
func generalInit(path string) (*conf.Bootstrap, error) {
	err := os.Setenv("USERNAME", "test")
	if err != nil {
		return nil, err
	}

	// 导入配置
	if path == "" {
		path = "../../configs/config.yaml"
	}
	bootstrap, err := conf.LoadConfig(path)
	if err != nil {
		return nil, err
	}

	return bootstrap, nil
}

func newApp(logger log.Logger, hs *http.Server) *kratos.App {
	var (
		// Name is the name of the compiled software.
		Name string
		// Version is the version of the compiled software.
		Version string
		id, _   = os.Hostname()
	)

	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			hs,
		),
	)
}

func StartDataProcessingServer(t *testing.T, bootstrap *conf.Bootstrap, registerInfo []utilApi.DeviceStateRegisterInfo) (
	v1.ConfigHTTPClient, v1.WarningDetectHTTPClient) {
	// 启动服务器
	app, cleanup, err := initApp(bootstrap.Server, bootstrap.Data, registerInfo, log.NewStdLogger(os.Stdout))
	if err != nil {
		t.Fatal(err)
	}

	done := make(chan struct{})
	go func() {
		defer close(done)
		err := app.Run()
		if err != nil {
			t.Error(err)
		}
	}()
	t.Cleanup(func() {
		// 发起关闭服务器的命令，并等待服务器关闭
		app.Stop()
		cleanup()
		<-done
	})

	// 实例化测试使用的http客户端
	var (
		httpConn                *http.Client
		configHTTPClient        v1.ConfigHTTPClient
		warningDetectHTTPClient v1.WarningDetectHTTPClient
	)
initClient:
	for {
		select {
		case <-done:
			t.Fatal("failed to start server app")
		default:
			httpConn, err = http.NewClient(
				context.Background(),
				http.WithEndpoint("localhost:8000"),
			)
			if err != nil {
				continue
			}
			configHTTPClient = v1.NewConfigHTTPClient(httpConn)
			warningDetectHTTPClient = v1.NewWarningDetectHTTPClient(httpConn)
			break initClient
		}
	}
	t.Cleanup(func() {
		httpConn.Close()
	})
	return configHTTPClient, warningDetectHTTPClient
}

// 保存设备配置的辅助函数
func saveConfig(client *data.Data, deviceID string, config proto.Message) error {
	key := biz.GetDeviceConfigKey(&biz.DeviceGeneralInfo{
		DeviceID:      deviceID,
		DeviceClassID: 0,
	})
	marshal, err := proto.Marshal(config)
	if err != nil {
		return err
	}
	value := fmt.Sprintf("%x", marshal)
	return client.HSet(context.Background(), key, deviceID, value).Err()
}

// 保存设备状态的辅助函数
func saveState(
	client *data.Data,
	deviceID string,
	state proto.Message,
	stamp int64,
	fields map[string]float64) error {
	cmders, err := client.TxPipelined(context.Background(), func(p redis.Pipeliner) error {
		// 保存设备状态
		stateKey := biz.GetDeviceStateKey(0)
		marshal, err := proto.Marshal(state)
		if err != nil {
			return err
		}
		value := fmt.Sprintf("%x", marshal)
		p.ZAdd(context.Background(), stateKey, &redis.Z{
			Score:  float64(stamp),
			Member: value,
		})

		// 保存设备字段
		for k, v := range fields {
			fieldKey := biz.GetDeviceStateFieldKey(&biz.DeviceGeneralInfo{
				DeviceID:      deviceID,
				DeviceClassID: 0,
			}, k)
			label := biz.GetDeviceStateFieldLabel(0, k)

			var args []interface{}
			args = append(args, "TS.ADD")
			args = append(args, fieldKey, stamp, v)

			// 当字段对应的ts未被创建时，以下设置的参数会被用于ts的创建

			// 设置ts的标签
			args = append(args, "LABELS", biz.WarningDetectFieldLabelName, label)
			p.Do(context.Background(), args...)

		}

		return nil
	})

	if err != nil {
		return err
	}
	for _, cmder := range cmders {
		if cmder.Err() != nil {
			return cmder.Err()
		}
	}
	return nil
}
