package cmd

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/minio/minio-go/v7"
)

// 将文件大小转化为对应单位B/K/M/G/T
func size2string(size float64) string {
	unit := ""
	if size > T {
		size /= T
		unit = "T"
	} else if size > G {
		size /= G
		unit = "G"
	} else if size > M {
		size /= M
		unit = "M"
	} else if size > K {
		size /= K
		unit = "K"
	} else {
		unit = "B"
	}

	sizeStr := strconv.FormatFloat(size, 'f', 2, 64)
	sizeStr += unit

	return sizeStr
}

// 全局错误处理函数
func er(msg interface{}) {
	fmt.Println(msg)
	os.Exit(-1)
}

// 检测MinIo/S3等云服务上的文件桶是否存在
func checkBucket(client *minio.Client, ctx context.Context, bucket string) {
	exist, err := client.BucketExists(ctx, bucket)
	if err != nil {
		er(err)
	} else if !exist {
		err := errors.New("Bucket " + bucket + " does not exist")
		er(err)
	}
}

// 检测文件是否存在
func fileExist(file string) bool {
	_, err := os.Stat(file)
	return err == nil || os.IsExist(err)
}

// 创建新文件
func mkNewFile(file string) {
	fp, err := os.Create(file)
	if err != nil {
		er(err)
	}
	defer fp.Close()

	fmt.Println("Successfully created file:", file)
}

// 在指定位置path处创建新配置文件
func mkConfigFile(path, file string) {
	err := os.MkdirAll(path, 0766)
	if err != nil {
		er(err)
	}

	mkNewFile(path + file)
}

// 去除字符串中的suffix "/"
func trimSuffix(str, suffix string) string {
	if strings.HasSuffix(str, suffix) {
		str = strings.TrimSuffix(str, suffix)
	}
	return str
}
