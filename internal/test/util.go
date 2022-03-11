package test

import (
	"gitee.com/moyusir/data-processing/internal/conf"
	"os"
)

// 提供给测试程序使用的通用初始化函数
func generalInit() (*conf.Bootstrap, error) {
	err := os.Setenv("USERNAME", "test")
	if err != nil {
		return nil, err
	}

	// 导入配置
	bootstrap, err := conf.LoadConfig("../../configs/config.yaml")
	if err != nil {
		return nil, err
	}

	return bootstrap, nil
}
