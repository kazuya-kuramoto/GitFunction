package Controller

import (
	"fmt"
	"gf/src/Core/Command"
	"gf/src/Core/Console"
	"gf/src/Core/Git"
	"os"
)

// ブランチの移動 (ブランチがあれば移動するし、なければ作る)
func To(ctx *Command.Context) {
	// 遷移先ブランチ名
	var branchToName string
	if ctx.OptionCount() > 0 {
		branchToName = ctx.GetBranchName()
	}

	// チェックアウト元のブランチ名がある場合はそこから生やす
	if ctx.OptionCount() == 2 {
		branchFromName := ctx.GetOption(1)
		if !Git.Exists(branchFromName) {
			Console.Error(branchFromName + "は存在しないブランチです。")
			return
		}

		Git.FromCheckout(branchToName, branchFromName)
		return

		// 普通のチェックアウト
	} else if ctx.OptionCount() == 1 {
		// 1個前のブランチへ
		if branchToName == "undo" {
			beforeBranchName := Git.GetCurrentBranchName()
			Git.Checkout(Git.FetchBeforeBranchName())
			Git.WriteBeforeBranchName(beforeBranchName)
			return
		}

		beforeBranchName := Git.GetCurrentBranchName()
		Git.Checkout(branchToName)
		Git.WriteBeforeBranchName(beforeBranchName)
		// developへ
	} else {
		beforeBranchName := Git.GetCurrentBranchName()
		Git.Checkout("develop")
		Git.WriteBeforeBranchName(beforeBranchName)
	}
}

func Now(_ *Command.Context) {
	now := Git.GetCurrentBranchName()
	Console.Info(now)
}

func BranchControl(ctx *Command.Context) {
	// ブランチの表示
	if ctx.OptionCount() == 0 {
		branches := Git.GetAllBranchNames()
		for _, name := range branches {
			Console.Info(name)
		}
		return
	}

	// ブランチの削除
	if ctx.GetOption(0) == "clear" {
		Git.CleanBranches()
	}
}

func Save(_ *Command.Context) {
	Git.AllStash()
	Console.Info("セーブしました。")
}

func Load(_ *Command.Context) {
	Git.StashPop()
	Console.Info("セーブデータをロードしました。")
}

func Clear(_ *Command.Context) {
	Git.ResetHardHead()
	Git.CleanFilesEachDirectory()
	Git.FetchPrune()
	Console.Info("クリーンアップしました。")
}

func Push(_ *Command.Context) {
	currentBranchName := Git.GetCurrentBranchName()
	if Console.Question(currentBranchName+"にpushして良いですか? [Y/n]", true) {
		Console.Info(currentBranchName + "にプッシュ中....")
		Git.PushCurrentBranch()
		Console.Info(currentBranchName + "にプッシュしました。")
	} else {
		Console.Info("中断しました。")
	}
}

func Pull(_ *Command.Context) {
	currentBranchName := Git.GetCurrentBranchName()
	if Console.Question(currentBranchName+"をpullして良いですか? [Y/n]", true) {
		Console.Info(currentBranchName + "をプル中....")
		Git.PullCurrentBranch()
		Console.Info(currentBranchName + "をプルしました。")
	} else {
		Console.Info("中断しました。")
	}
}

func Sync(_ *Command.Context) {
	currentBranchName := Git.GetCurrentBranchName()
	if Console.Question("origin/"+currentBranchName+"に完全に合わせますが良いですか? [Y/n]", true) {
		Git.SyncCurrentBranch()
		Console.Info("シンクロ完了しました。")
	} else {
		Console.Info("中断しました。")
	}
}

func GetBeforeBranchName(_ *Command.Context) {
	Console.Info(Git.FetchBeforeBranchName())
}

func ModInit() {
	Console.Info("サブモジュールを初期化します。")
	Git.SubmoduleInit()
	Console.Info("サブモジュールを初期化完了しました。")
}

func ModUpdate() {
	Console.Info("サブモジュールを更新します。")
	Git.SubmoduleUpdate()
	Console.Info("サブモジュールを更新完了しました。")
}

func Update() {
	fmt.Printf("    ◇◇◇◇◇◇ %s \x1b[35m%s\x1b[0m ◇◇◇◇◇\n", "GitFunction", "Update")
	err := Console.Exec(
		"cd " + os.Getenv("GOPATH") + "/src/gf/ && " +
			"git clean -df && " +
			"git pull origin develop && " +
			"make update")
	if err != nil {
		Console.Error("Updateに失敗しました")
	}
	Console.Info("Updateが正常に終了しました。")
}

func Checkout(ctx *Command.Context) {
	if ctx.OptionCount() == 0 {
		Console.Error("Branch名を指定してください")
		return
	}

	Console.Info("フェッチします")
	Git.Fetch()
	Console.Info("チェックアウトの開始")
	Git.ForceRemoteBranchCheckout(ctx.GetBranchName())
	Git.Checkout(ctx.GetBranchName())
	Console.Info("チェックアウト完了")
}

func Commit(ctx *Command.Context) {
	if ctx.OptionCount() == 0 {
		Console.Error("コミットメッセージが必要です")
		return
	}

	currentBranchName := Git.GetCurrentBranchName()
	if !Console.Question(currentBranchName+"にコミットして良いですか? [Y/n]", true) {
		Console.Info("中断しました")
		return
	}

	Git.Commit(ctx.GetOption(0))

	Console.Info("コミットしました")
}

func Merge(ctx *Command.Context) {
	if ctx.OptionCount() == 0 {
		Console.Error("ブランチ名を指定してください。")
		return
	}

	if ctx.OptionCount() == 1 {
		if ctx.GetOption(0) == "-r" {
			Console.Error("ブランチ名を指定してください")
		}

		Git.LocalMerge(ctx.GetOption(0))
		return
	}

	if ctx.OptionCount() >= 3 {
		Console.Error("パラメータが多すぎます\ngf merge [-r] <branchName>")
		return
	}

	option0 := ctx.GetOption(0)
	option1 := ctx.GetOption(1)
	if option0 != "-r" {
		Console.Error("オプションが違います\ngf merge -r <branchName>")
		return
	}

	Git.RemoteMerge(option1)
}
