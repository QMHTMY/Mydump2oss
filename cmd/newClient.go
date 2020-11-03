package cmd

import (
	"errors"

	"github.com/QMHTMY/Mydump2oss/cfg"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/spf13/viper"
)

// 获取配置文件名
func getCurConfig() string {
	curConfig := viper.GetString("config")

	if curConfig != "" {
		// 若通过--config指定配置文件位置，则判断是否存在
		if !fileExist(curConfig) {
			err := errors.New("Config file does not exist")
			er(err)
		}
	} else {
		// 否则使用默认值
		curConfig = configPath + fileName
	}

	return curConfig
}

// 构造client
func newClient() *minio.Client {
	// 获取配置信息
	curConfig := getCurConfig()
	item, err := cfg.ReadItem(curConfig)
	if err != nil {
		er(err)
	}

	endPoint = item.EndPoint
	accessKeyID = item.AccessKeyID
	secretAccessKey = item.SecretAccessKey
	useSSL = item.UseSSL

	// 通过配置信息构造并返回对象存储的client
	client, err := minio.New(
		endPoint,
		&minio.Options{
			Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
			Secure: useSSL,
		},
	)
	if err != nil {
		er(err)
	}

	return client
}
