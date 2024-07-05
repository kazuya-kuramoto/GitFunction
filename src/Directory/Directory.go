package Directory

import (
	"fmt"
	"os"
)

func Make(path string) {
	if err := os.MkdirAll(path, 0777); err != nil {
		fmt.Println(err)
	}
}

func Delete(path string) {
	if err := os.Remove(path); err != nil {
		fmt.Println(err)
	}
}
