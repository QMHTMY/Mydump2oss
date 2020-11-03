package cmd

import (
	"context"
	"errors"
	"fmt"
	"path/filepath"

	"github.com/minio/minio-go/v7"
	"github.com/spf13/cobra"
)

var (
	// rmb删除云端的文件桶
	removeBucketCmd = &cobra.Command{
		Use:     "rmb bucket(s) ...",
		Short:   "Remove remote bucket(s)",
		Long:    "Remove remote bucket(s) on MinIo/S3 Cloud Storage",
		Aliases: []string{"rmbucket", "removebucket"},
		Example: "  Mydump2oss rmb sql_bucket/ code_bucket/",
		Run:     removeBucketRun,
	}

	// rmo删除云端文件桶中文件
	removeObjectCmd = &cobra.Command{
		Use:     "rmo bucket/objs ...",
		Short:   "Remove remote object(s)",
		Long:    "Remove remote object(s) on MinIo/S3 Cloud Storage",
		Aliases: []string{"rm", "rmobject", "removeobject"},
		Example: "  Mydump2oss rmo mysql/data.sql code/src.gz",
		Run:     removeObjectRun,
	}
)

func removeBucketRun(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		er(errors.New(rmbUsage))
	}

	client := newClient()
	ctx := context.Background()
	for _, bucket := range args {
		bucket = trimSuffix(bucket, "/")
		checkBucket(ctx, client, bucket)

		err := client.RemoveBucket(ctx, bucket)
		if err != nil {
			er(err)
		}
		fmt.Println("Successfully removed bucket:", bucket)
	}
}

func removeObjectRun(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		er(errors.New(rmoUsage))
	}

	client := newClient()
	ctx := context.Background()
	opts := minio.RemoveObjectOptions{GovernanceBypass: true}
	for _, path2obj := range args {
		bucket := filepath.Dir(path2obj)
		obj := filepath.Base(path2obj)
		checkBucket(ctx, client, bucket)

		err := client.RemoveObject(ctx, bucket, obj, opts)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Successfully removed object:", obj)
	}
}
