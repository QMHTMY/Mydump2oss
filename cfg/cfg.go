package cfg

import (
	"encoding/json"
	"io/ioutil"
)

// MinIo/S3 Cloud Storage的认证信息
type Item struct {
	EndPoint        string
	AccessKeyID     string
	SecretAccessKey string
	UseSSL          bool
}

// 保存Item为json文件
func WriteItem(filename string, item Item) error {
	b, err := json.Marshal(item)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, b, 0644)
	if err != nil {
		return err
	}

	return nil
}

// 读取json文件并转换为Item
func ReadItem(filename string) (Item, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return Item{}, err
	}

	item := Item{}
	if err := json.Unmarshal(b, &item); err != nil {
		return Item{}, err
	}

	return item, nil
}
