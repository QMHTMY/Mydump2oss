package cmd

// Mydump2oss 版本信息
var version = "v0.1\nGolang version v1.15\nCopyright (c) 2020 Shieber"

// 云端数据存储位置，此为美国东部1区
var region = "us-east-1"

// 关于配置文件的参数
var (
	configPath string // 配置文件路径
	fileName   string // 配置文件名
	cfgFile    string // --config参数
)

// 配置文件存储的认证信息参数
var (
	endPoint        string
	accessKeyID     string
	secretAccessKey string
	useSSL          bool
)

// 各命令的使用方法
var (
	cfgUsage = `Usage: Mydump2oss cfg --endPoint="xx" --accessKeyID="xx" --secretAccessKey="xx" --useSSL="true"`
	cpUsage  = "Usage: Mydump2oss cp object1... bucket"
	lsUsage  = "Usage: Mydump2oss ls bucket(s)..."
	mbUsage  = "Usage: Mydump2oss mb bucket(s)..."
	mrUsage  = "Usage: Mydump2oss mr localDir bucket"
	rmbUsage = "Usage: Mydump2oss rmb bucket(s) ..."
	rmoUsage = "Usage: Mydump2oss rmo bucket/objects ..."
)

// 常量，数据大小单位K/M/G/T
const (
	_ = iota
	K = 1 << (iota * 10)
	M = 1 << (iota * 10)
	G = 1 << (iota * 10)
	T = 1 << (iota * 10)
)
