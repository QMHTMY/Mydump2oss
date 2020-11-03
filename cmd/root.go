package cmd

import (
	"errors"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// 全局主命令Mydump2oss
	RootCmd = &cobra.Command{
		Use:     "Mydump2oss",
		Short:   "Mydump2oss is a file upload tool",
		Long:    "Mydump2oss, a tool to upload files to MinIo/S3... Cloud Storage",
		Version: version,
	}
)

// 初始化，保存Minio/S3 cloud storage的认证信息到配置文件
func init() {
	// 指定默认配置文件路径和文件名
	home, err := homedir.Dir()
	if err != nil {
		er(err)
	}
	separator := string(os.PathSeparator)
	configPath = home + separator + ".Mydump2oss" + separator
	fileName = "config.json"

	// 若不存在则创建
	if !fileExist(configPath + fileName) {
		mkConfigFile(configPath, fileName)
		err := errors.New(`Please specify config file by --config or initialize config file by: Mydump2oss cfg ...`)
		er(err)
	}

	// 添加 --config flag 以使用非默认配置文件
	viper.AddConfigPath(configPath)
	viper.SetConfigName(fileName)
	RootCmd.PersistentFlags().StringVar(
		&cfgFile,
		"config",
		"",
		"Config file to store authentication info, defauls to $HOME/.Mydump2oss/config.json",
	)
}

// 将子命令加入RootCmd中
func addCommands() {
	RootCmd.AddCommand(configCmd)
	RootCmd.AddCommand(makeBucketCmd)
	RootCmd.AddCommand(listCmd)
	RootCmd.AddCommand(copyCmd)
	RootCmd.AddCommand(mirrorCmd)
	RootCmd.AddCommand(removeBucketCmd)
	RootCmd.AddCommand(removeObjectCmd)
}

// 主函数，供外部直接调用执行
func Execute() {
	addCommands()
	if err := RootCmd.Execute(); err != nil {
		er(err)
	}
}
