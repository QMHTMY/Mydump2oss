package cmd

import (
	"errors"

	"github.com/QMHTMY/Mydump2oss/cfg"
	"github.com/spf13/cobra"
)

var (
	// 配置参数
	endPoint        string
	accessKeyID     string
	secretAccessKey string
	useSSL          bool

	// cfg命令
	configCmd = &cobra.Command{
		Use:     "cfg",
		Short:   "Set authentication configurations",
		Long:    `Set configurations for Minio/S3 Cloud Storage, including endPoint，acessKeyID，secretAccessKey，useSSL`,
		Aliases: []string{"config"},
		Example: `  Mydump2oss cfg --endPoint="x" --accessKeyID="x" --secretAccessKey="x" --useSSL="true"`,
		Run:     configRun,
	}
)

func configRun(cmd *cobra.Command, args []string) {
	if endPoint == "" || accessKeyID == "" || secretAccessKey == "" {
		if len(args) < 4 {
			err := errors.New(`Usage: Mydump2oss cfg --endPoint="xx" --accessKeyID="xx" --secretAccessKey="xx" --useSSL="true"`)
			er(err)
		} else {
			endPoint = args[0]
			accessKeyID = args[1]
			secretAccessKey = args[2]
			if args[3] == "true" {
				useSSL = true
			}
		}
	}

	item := cfg.Item{
		EndPoint:        endPoint,
		AccessKeyID:     accessKeyID,
		SecretAccessKey: secretAccessKey,
		UseSSL:          useSSL,
	}

	if err := cfg.WriteItem(configPath+fileName, item); err != nil {
		er(err)
	}
}

func init() {
	configCmd.Flags().StringVar(&endPoint, "endPoint", "", "oss storage endpoint")
	configCmd.Flags().StringVar(&accessKeyID, "accessKeyID", "", "login id")
	configCmd.Flags().StringVar(&secretAccessKey, "secretAccessKey", "", "login key")
	configCmd.Flags().BoolVar(&useSSL, "useSSL", false, "use ssl or not")
}
