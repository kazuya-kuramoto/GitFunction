package Console

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
)

func Info(message string) {
	fmt.Printf("\x1b[0m%s\x1b[0m\n", message)
}

func Error(message string) {
	fmt.Printf("\x1b[31m%s\x1b[0m\n", message)
	os.Exit(-1)
}

func Question(q string, defaultFlag bool) bool {
	result := true
	fmt.Print(q)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		i := scanner.Text()

		if i == "Y" || i == "y" || i == "yes" || i == "YES" || i == "Yes" {
			break
		} else if i == "N" || i == "n" || i == "No" || i == "no" || i == "NO" {
			result = false
			break
		} else if i == "" {
			result = defaultFlag
			break
		} else {
			fmt.Println("yかnで答えてください。")
			fmt.Print(q)
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return result
}

func Exec(command string) error {
	// コマンドの登録
	cmd := exec.Command("sh", "-c", command)

	// `StdoutPipe()` 経由で`io.ReadCloser`を取得
	display, err := cmd.CombinedOutput()

	// 描画
	fmt.Printf("\x1b[33m$ %s\n\x1b[0m", command)
	if len(display) > 0 && err == nil {
		fmt.Printf("\x1b[36m%s\x1b[0m", display)
	}

	if len(display) > 0 && err != nil {
		fmt.Printf("\x1b[31m%s\x1b[0m", display)
	}

	// EOFエラーをそのまま返さない
	if err == io.EOF {
		return nil
	}

	return err
}
