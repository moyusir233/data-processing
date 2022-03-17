package conf

import (
	"encoding/json"
	utilApi "gitee.com/moyusir/util/api/util/v1"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"io/ioutil"
	"log"
	"os"
)

// 定义了当前程序服务的用户及其设备的基本信息，由服务中心生成代码时注入
var (
	// Username 该服务对应用户的id
	Username = "test"
)

func initEnv() {
	if username, ok := os.LookupEnv("USERNAME"); ok {
		Username = username
	} else {
		log.Fatalln("The required environment variable USERNAME is missing")
	}
}

func LoadConfig(path string) (*Bootstrap, error) {
	initEnv()

	c := config.New(
		config.WithSource(
			file.NewSource(path),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		return nil, err
	}

	var bc Bootstrap
	if err := c.Scan(&bc); err != nil {
		return nil, err
	}
	return &bc, nil
}

func LoadRegisterInfo(path string) ([]utilApi.DeviceStateRegisterInfo, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var infos []utilApi.DeviceStateRegisterInfo
	err = json.Unmarshal(bytes, &infos)
	if err != nil {
		return nil, err
	}

	return infos, nil
}
