package Git

import (
	"fmt"
	"gf/src/Core/Console"
	"gf/src/Directory"
	"gf/src/File"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func Checkout(branchName string) {
	// チェックアウト
	if Exists(branchName) {
		err := Console.Exec("git checkout " + branchName)
		if err != nil {
			Console.Error("チェックアウトに失敗しました。")
			return
		}
		Console.Info(branchName + "にチェックアウトしました。")

		// 新規作成
	} else {
		err := Console.Exec("git checkout -b " + branchName)
		if err != nil {
			Console.Error("チェックアウトに失敗しました。")
			return
		}

		Console.Info(branchName + "を新しく作成し、チェックアウトしました。")
	}
}

func FromCheckout(to string, from string) {
	// チェックアウト
	if Exists(to) {
		cmd := exec.Command("git", "checkout", to, from)
		if err := cmd.Start(); err != nil {
			Console.Error("チェックアウトに失敗しました。" + err.Error())
		}

		if err := cmd.Wait(); err != nil {
			Console.Error("チェックアウトに失敗しました。" + err.Error())
		}
		Console.Info(to + "を" + from + "からチェックアウトしました。")

		// 新規作成
	} else {
		cmd := exec.Command("git", "checkout", "-b", to, from)
		if err := cmd.Start(); err != nil {
			Console.Error("チェックアウトに失敗しました。" + err.Error())
		}

		if err := cmd.Wait(); err != nil {
			Console.Error("チェックアウトに失敗しました。" + err.Error())
		}

		Console.Info(to + "を新しく作成し、" + from + "からチェックアウトしました。")
	}
}

func AllStash() {
	cmd := exec.Command("git", "stash", "-u")
	if err := cmd.Start(); err != nil {
		Console.Error("スタッシュに失敗しました。" + err.Error())
	}

	if err := cmd.Wait(); err != nil {
		Console.Error("スタッシュに失敗しました。" + err.Error())
	}

	fmt.Println("一時的変更のあるファイルを全て保管しました。")
}

func StashPop() {
	err := Console.Exec("git stash pop")
	if err != nil {
		Console.Error("Stash Popに失敗しました。")
	}
}

func Exists(branchName string) bool {
	cmd := exec.Command("git", "show-ref", "--verify", "refs/heads/"+branchName)
	if err := cmd.Start(); err != nil {
		return false
	}

	if err := cmd.Wait(); err != nil {
		return false
	}
	return true
}

func GetAllBranchNames() []string {
	data, err := exec.Command("git", "branch").Output()
	if err != nil {
		Console.Error("ブランチ一覧の取得に失敗しました。" + string(data) + err.Error())
	}

	rawBranches := string(data)
	branches := strings.Split(rawBranches, "\n")
	var finalizedBranches []string
	for _, branch := range branches {
		branch = strings.Replace(branch, "*", "", -1)
		branch = strings.Replace(branch, " ", "", -1)
		if branch == "" {
			continue
		}

		finalizedBranches = append(finalizedBranches, branch)
	}

	return finalizedBranches
}

func GetCurrentBranchName() string {
	data, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
	if err != nil {
		Console.Error("現在のブランチ名を取得できません" + string(data) + err.Error())
	}
	return strings.Replace(string(data), "\n", "", -1)
}

func DeleteBranch(branchName string) {
	err := Console.Exec("git branch -D " + branchName)
	if err != nil {
		Console.Error("ブランチの削除に失敗しました。")
	}
}

func ResetHardHead() {
	err := Console.Exec("git reset --hard HEAD")
	if err != nil {
		Console.Error("Resetに失敗しました。")
	}
}

func CleanFilesEachDirectory() {
	err := Console.Exec("git clean -df")
	if err != nil {
		Console.Error("クリーンに失敗しました。")
	}
}

func FetchPrune() {
	err := Console.Exec("git fetch -p")
	if err != nil {
		Console.Error("フェッチに失敗しました。")
	}
}

func PushCurrentBranch() {
	err := Console.Exec("git push origin " + GetCurrentBranchName())
	if err != nil {
		Console.Error("プッシュに失敗しました。")
	}
}

func PullCurrentBranch() {
	err := Console.Exec("git pull origin " + GetCurrentBranchName())
	if err != nil {
		Console.Error("プルに失敗しました。")
	}
}

func Fetch() {
	err := Console.Exec("git fetch origin")
	if err != nil {
		Console.Error("フェチに失敗しました。")
	}
}

func RenameCurrentBranch(newName string) {
	err := Console.Exec("git branch -m " + newName)
	if err != nil {
		Console.Error("ブランチ名の更新に失敗しました。")
	}
}

func SyncCurrentBranch() {

	currentBranchName := GetCurrentBranchName()

	ResetHardHead()

	CleanFilesEachDirectory()

	RenameCurrentBranch("prev-" + GetCurrentBranchName())

	FetchPrune()

	Checkout(currentBranchName)
}

func getBeforeBranchNameSaveFilePath() string {
	dir, _ := os.Executable()
	path := filepath.Join(dir, "../gfconf/before_branch_name.tmp")
	return path
}

func FetchBeforeBranchName() string {
	if !File.Exists(getBeforeBranchNameSaveFilePath()) {
		if !File.Exists(filepath.Dir(getBeforeBranchNameSaveFilePath())) {
			Directory.Make(filepath.Dir(getBeforeBranchNameSaveFilePath()))
		}

		return "develop"
	}

	return File.ReadAllText(getBeforeBranchNameSaveFilePath())
}

func WriteBeforeBranchName(branchName string) {
	if !File.Exists(filepath.Dir(getBeforeBranchNameSaveFilePath())) {
		Directory.Make(filepath.Dir(getBeforeBranchNameSaveFilePath()))
	}

	File.WriteString(getBeforeBranchNameSaveFilePath(), branchName)
}

func SubmoduleInit() {
	out, err := exec.Command("git", "submodule", "init").Output()
	if err != nil {
		Console.Error("サブモジュールの初期化に失敗しました。" + err.Error())
	}
	Console.Info(string(out))
}

func SubmoduleUpdate() {
	out, err := exec.Command("git", "submodule", "update").Output()
	if err != nil {
		Console.Error("サブモジュールの更新に失敗しました。" + err.Error())
	}
	Console.Info(string(out))
}

func CleanBranches() {
	err := Console.Exec("git branch --merged | egrep -v '\\*|develop|master' | xargs git branch -d")
	if err != nil {
		Console.Error("ブランチの清掃処理に失敗しました。")
	}
}

func ForceRemoteBranchCheckout(branchName string) {
	err := Console.Exec("git branch --force " + branchName + " origin/" + branchName)
	if err != nil {
		Console.Error("チェックアウトに失敗しました。")
	}
}

func Commit(message string) {
	err := Console.Exec("git commit -m " + message)
	if err != nil {
		Console.Error("コミットに失敗しました。")
	}
}

func LocalMerge(branchName string) {
	err := Console.Exec("git merge " + branchName)
	if err != nil {
		Console.Error("マージに失敗しました。")
	}
}

func RemoteMerge(branchName string) {
	err := Console.Exec("git merge origin/" + branchName)
	if err != nil {
		Console.Error("マージに失敗しました。")
	}
}
