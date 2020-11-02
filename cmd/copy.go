package cmd

import (
	"context"
	"errors"
	"fmt"
	"path/filepath"

	"github.com/minio/minio-go/v7"
	"github.com/spf13/cobra"
)

var copyCmd = &cobra.Command{
	Use:     "cp obj(s) ... bucket",
	Short:   "Copy local objects to a remote bucket",
	Long:    "Copy local objects to a remote bucket on MinIo/S3 Cloud Storage",
	Aliases: []string{"copy", "upload"},
	Example: "  Mydump2oss cp file.sql mysql_backup",
	Run:     copyRun,
}

func copyRun(cmd *cobra.Command, args []string) {
	length := len(args)
	if length < 2 {
		err := errors.New("Usage: Mydump2oss cp object1... bucket")
		er(err)
	}

	bucket := args[length-1]
	client := newClient()
	ctx := context.Background()
	checkBucket(ctx, client, bucket)

	putObjectOptions := minio.PutObjectOptions{ContentType: "application/octet-stream"}
	for _, obj := range args[:length-1] {
		fileName := filepath.Base(obj)

		_, err := client.FPutObject(ctx, bucket, fileName, obj, putObjectOptions)
		if err != nil {
			er(err)
		} else {
			fmt.Println("Successfully uploaded object:", fileName)
		}
	}
}
