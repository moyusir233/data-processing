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
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"google.golang.org/protobuf/proto"
	"os"
	"testing"
	"time"
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
func saveConfig(client *data.RedisData, deviceID string, config proto.Message) error {
	key := biz.GetDeviceConfigKey(&biz.DeviceGeneralInfo{
		DeviceID:      deviceID,
		DeviceClassID: 1,
	})
	marshal, err := proto.Marshal(config)
	if err != nil {
		return err
	}
	value := fmt.Sprintf("%x", marshal)
	return client.HSet(context.Background(), key, deviceID, value).Err()
}

// 保存设备状态的辅助函数
func saveState(client influxdb2.Client, org, bucket string, time time.Time,
	deviceID string, fields map[string]float64, tags map[string]string) error {
	writeAPI := client.WriteAPI(org, bucket)

	point := write.NewPointWithMeasurement(deviceID).SetTime(time.UTC())
	for k, v := range fields {
		point.AddField(k, v)
	}
	for k, v := range tags {
		point.AddTag(k, v)
	}
	point.SortFields().SortTags()

	writeAPI.WritePoint(point)
	writeAPI.Flush()
	return nil
}

// 清理influxdb的函数
func clearInfluxdb(client *data.InfluxdbData, org, bucket string, deviceClassID int) {
	deleteAPI := client.DeleteAPI()
	start := time.Now().UTC().Add(-1 * time.Hour)
	end := time.Now().UTC()
	predicate := fmt.Sprintf(`deviceClassID="%d"`, deviceClassID)
	deleteAPI.DeleteWithName(context.Background(), org, bucket, start, end, predicate)
}
