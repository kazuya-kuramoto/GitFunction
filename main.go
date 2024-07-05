package main

import (
	"fmt"
	"gf/src/Controller"
	"gf/src/Core/Command"
	"gf/src/Core/Console"
	"os"
	"os/exec"
)

func main() {
	dir, _ := os.Getwd()
	_, _ = exec.Command("GIT_DIR=" + dir).Output()
	_, _ = exec.Command("GIT_WORK_TREE=" + dir).Output()

	context := Command.New(os.Args)

	if context == nil {
		usageOut()
		return
	}

	switch context.GetRouter() {
	case "to":
		Controller.To(context)
		break
	case "now":
		Controller.Now(context)
		break
	case "branch":
		Controller.BranchControl(context)
		break
	case "checkout":
		Controller.Checkout(context)
		break
	case "save":
		Controller.Save(context)
		break
	case "load":
		Controller.Load(context)
		break
	case "clear":
		Controller.Clear(context)
		break
	case "push":
		Controller.Push(context)
		break
	case "pull":
		Controller.Pull(context)
		break
	case "mod":
		if context.OptionCount() == 0 {
			Controller.ModInit()
			Controller.ModUpdate()
			return
		}

		if context.GetOption(0) == "init" {
			Controller.ModInit()
		} else if context.GetOption(0) == "update" {
			Controller.ModUpdate()
		} else {
			usageOut()
		}
		break
	case "sync":
		Controller.Sync(context)
		break
	case "before":
		Controller.GetBeforeBranchName(context)
		break
	case "update":
		Controller.Update()
		break
	case "merge":
		Controller.Merge(context)
	case "commit":
		Controller.Commit(context)
	default:
		usageOut()
		break
	}
}

func usageOut() {
	Console.Info("-------------------------------------------")
	Console.Info(" GitFunction v1.12 Usage")
	Console.Info("-------------------------------------------")

	Console.Info("\n   ● ブランチ関連")
	fmt.Printf("    ◇ %s: \x1b[35m%s\x1b[0m\n", "ブランチ切り替え", "gf to <branch name>")
	fmt.Printf("    ◇ %s: \x1b[35m%s\x1b[0m\n", "前のブランチに切り替え", "gf to undo")
	fmt.Printf("    ◇ %s: \x1b[35m%s\x1b[0m\n", "developブランチに切り替え", "gf to")
	fmt.Printf("    ◇ %s: \x1b[35m%s\x1b[0m\n", "指定ブランチからブランチを切り替え", "gf to <branch name> <from branch name>")
	fmt.Printf("    ◇ %s: \x1b[35m%s\x1b[0m\n", "現在ブランチ名", "gf now")
	fmt.Printf("    ◇ %s: \x1b[35m%s\x1b[0m\n", "所有ブランチリスト", "gf branch")
	fmt.Printf("    ◇ %s: \x1b[35m%s\x1b[0m\n", "いらなくなったブランチを一括削除", "gf branch clear")
	fmt.Printf("    ◇ %s: \x1b[35m%s\x1b[0m\n", "リモートブランチにチェックアウト", "gf checkout <remote branch name>")

	Console.Info("\n  ● サブモジュール関連")
	fmt.Printf("    ◇ %s: \x1b[35m%s\x1b[0m\n", "サブモジュール初期化", "gf mod init")
	fmt.Printf("    ◇ %s: \x1b[35m%s\x1b[0m\n", "サブモジュール更新", "gf mod update")
	fmt.Printf("    ◇ %s: \x1b[35m%s\x1b[0m\n", "サブモジュールの初期化と更新", "gf mod")

	Console.Info("\n  ● スタッシュ関連")
	fmt.Printf("    ◇ %s: \x1b[35m%s\x1b[0m\n", "変更を保存", "gf save")
	fmt.Printf("    ◇ %s: \x1b[35m%s\x1b[0m\n", "最終セーブ地点をロード", "gf load")

	Console.Info("\n  ● クリーン関連")
	fmt.Printf("    ◇ %s: \x1b[35m%s\x1b[0m\n", "現在の変更を全て消す", "gf clear")
	fmt.Printf("    ◇ %s: \x1b[35m%s\x1b[0m\n", "現在のブランチを強制的にリモートのブランチに合わせる", "gf sync")

	Console.Info("\n  ● プッシュ・プル")
	fmt.Printf("    ◇ %s: \x1b[35m%s\x1b[0m\n", "現在のブランチにプッシュ", "gf push")
	fmt.Printf("    ◇ %s: \x1b[35m%s\x1b[0m\n", "現在のブランチをプル", "gf pull")

	Console.Info("\n  ● コミット")
	fmt.Printf("    ◇ %s: \x1b[35m%s\x1b[0m\n", "現在のブランチにコミット", "gf commit <commitMessage>")

	Console.Info("\n  ● マージ")
	fmt.Printf("    ◇ %s: \x1b[35m%s\x1b[0m\n", "ローカルブランチをマージ", "gf merge <branchName>")
	fmt.Printf("    ◇ %s: \x1b[35m%s\x1b[0m\n", "リモートブランチをマージ", "gf merge -r <branchName>")
}
