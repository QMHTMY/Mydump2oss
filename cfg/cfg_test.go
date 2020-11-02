package cfg

import "testing"

var (
	testFile = "cfg_test.json"
	testItem = Item{
		EndPoint:        "s3.com",
		AccessKeyID:     "s3user",
		SecretAccessKey: "password",
		UseSSL:          true,
	}
)

func TestWriteItem(t *testing.T) {
	err := WriteItem(testFile, testItem)
	if err != nil {
		t.Error(err)
	}
}

func TestReadItem(t *testing.T) {
	item, err := ReadItem(testFile)
	if err != nil {
		t.Error("Error reading item")
	}

	if item != testItem {
		t.Error("Not equal")
	}
}
