package main

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

func main() {
	rootCmd := &cobra.Command{
		Use:     "Mydump2oss",
		Short:   "Mydump2oss is a file upload tool",
		Long:    "Mydump2oss, a tool to upload files to MinIo/S3... Cloud Storage",
		Version: "v0.1\nGolang version v1.15\nCopyright (c) 2020 Shieber",
	}
	if err := doc.GenMarkdownTree(rootCmd, "docs/"); err != nil {
		log.Fatal(err)
	}

	configCmd := &cobra.Command{
		Use:     "cfg",
		Short:   "Set Cloud Storage Service authentication configurations",
		Long:    `Set Cloud Storage Service configurations for Minio/S3 Cloud Storage, including endPoint，acessKeyID，secretAccessKey，useSSL`,
		Aliases: []string{"config"},
		Example: `  Mydump2oss cfg --endPoint="x" --accessKeyID="x" --secretAccessKey="x" --useSSL="true"`,
	}
	if err := doc.GenMarkdownTree(configCmd, "docs/"); err != nil {
		log.Fatal(err)
	}

	copyCmd := &cobra.Command{
		Use:     "cp obj(s) ... bucket",
		Short:   "Copy local objects to a remote bucket",
		Long:    "Copy local objects to a remote bucket on MinIo/S3 Cloud Storage",
		Aliases: []string{"copy", "upload"},
		Example: "  Mydump2oss cp file.sql mysql_backup",
	}
	if err := doc.GenMarkdownTree(copyCmd, "docs/"); err != nil {
		log.Fatal(err)
	}

	listCmd := &cobra.Command{
		Use:     "ls bucket(s) ...",
		Short:   "List objects of remote bucket(s)",
		Long:    "List objects in remote bucket(s) on MinIo/S3 Cloud Storage",
		Aliases: []string{"list"},
		Example: "  Mydump2oss ls mysql_bukcet",
	}
	if err := doc.GenMarkdownTree(listCmd, "docs/"); err != nil {
		log.Fatal(err)
	}

	makeBucketCmd := &cobra.Command{
		Use:     "mb bucket(s) ...",
		Short:   "Make remote bucket(s)",
		Long:    "Make remote bucket(s) on MinIo/S3 Cloud Storage",
		Aliases: []string{"mkb", "mkbucket", "makebucket"},
		Example: "  Mydump2oss mb bucket1 bucket2 bucket3",
	}
	if err := doc.GenMarkdownTree(makeBucketCmd, "docs/"); err != nil {
		log.Fatal(err)
	}

	mirrorCmd := &cobra.Command{
		Use:     "mr localdir bucket",
		Short:   "Synchronize local objects to a remote bucket",
		Long:    "Synchronize local objects to a remote bucket on MinIo/S3 Cloud Storage",
		Aliases: []string{"mirror"},
		Example: "  Mydump2oss mr sql_bkp sql_bucket",
	}
	if err := doc.GenMarkdownTree(mirrorCmd, "docs/"); err != nil {
		log.Fatal(err)
	}

	removeBucketCmd := &cobra.Command{
		Use:     "rmb bucket(s) ...",
		Short:   "Remove bucket(s)",
		Long:    "Remove remote bucket(s) on MinIo/S3 Cloud Storage",
		Aliases: []string{"rmbucket", "removebucket"},
		Example: "  Mydump2oss rmb sql_bucket code_bucket",
	}
	if err := doc.GenMarkdownTree(removeBucketCmd, "docs/"); err != nil {
		log.Fatal(err)
	}

	removeObjectCmd := &cobra.Command{
		Use:     "rmo bucket/objs ...",
		Short:   "Remove object(s)",
		Long:    "Remove remote object(s) on MinIo/S3 Cloud Storage",
		Aliases: []string{"rm", "rmobject", "removeobject"},
		Example: "  Mydump2oss rmo sql_bucket/org.sql.gz",
	}
	if err := doc.GenMarkdownTree(removeObjectCmd, "docs/"); err != nil {
		log.Fatal(err)
	}
}
