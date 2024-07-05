package File

import (
	"io/ioutil"
	"os"
)

func ReadAllText(filePath string) string {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		// エラー処理
		panic(err)
	}
	return string(data)
}

func WriteString(filePath string, data string) {
	err := ioutil.WriteFile(filePath, []byte(data), 0644)
	if err != nil {
		panic(err)
	}
}

func Exists(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil
}
