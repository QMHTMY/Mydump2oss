package cmd

import (
	"context"
	"errors"
	"fmt"
	"path/filepath"
	"sync"

	"github.com/minio/minio-go/v7"
	"github.com/spf13/cobra"
)

var (
	// 等待组
	wgrmb sync.WaitGroup
	wgrmo sync.WaitGroup

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
		wgrmb.Add(1)
		go removeBucket(client, ctx, bucket)
	}
	wgrmb.Wait()
}

func removeBucket(client *minio.Client, ctx context.Context, bucket string) {
	defer wgrmb.Done()

	bucket = trimSuffix(bucket, "/")
	checkBucket(client, ctx, bucket)

	err := client.RemoveBucket(ctx, bucket)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Successfully deleted bucket:", bucket)
	}
}

func removeObjectRun(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		er(errors.New(rmoUsage))
	}

	client := newClient()
	ctx := context.Background()

	for _, path2obj := range args {
		wgrmo.Add(1)
		go removeObject(client, ctx, path2obj)
	}
	wgrmo.Wait()
}

func removeObject(client *minio.Client, ctx context.Context, path2obj string) {
	defer wgrmo.Done()

	bucket := filepath.Dir(path2obj)
	checkBucket(client, ctx, bucket)

	obj := filepath.Base(path2obj)
	opts := minio.RemoveObjectOptions{GovernanceBypass: true}
	err := client.RemoveObject(ctx, bucket, obj, opts)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Successfully deleted object:", obj)
	}
}
