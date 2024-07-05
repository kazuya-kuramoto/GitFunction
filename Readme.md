# GitFunctions
Git操作を高速化します。

# Setup
```bash
cd $GOPATH
cd src/github.dena.jp/kazuya-kuramoto
git clone git@github.dena.jp:kazuya-kuramoto/GitFunction.git
cd GitFunction
make install
```

# Update
```bash
gf update
```

# Usage

```
-------------------------------------------
 GitFunction v1.11 Usage
-------------------------------------------

   ● ブランチ関連
    ◇ ブランチ切り替え: gf to <branch name>
    ◇ 前のブランチに切り替え: gf to undo
    ◇ developブランチに切り替え: gf to
    ◇ 指定ブランチからブランチを切り替え: gf to <branch name> <from branch name>
    ◇ 現在ブランチ名: gf now
    ◇ 所有ブランチリスト: gf branch
    ◇ いらなくなったブランチを一括削除: gf branch clear
    ◇ リモートブランチにチェックアウト: gf checkout <remote branch name>

  ● サブモジュール関連
    ◇ サブモジュール初期化: gf mod init
    ◇ サブモジュール更新: gf mod update
    ◇ サブモジュールの初期化と更新: gf mod

  ● スタッシュ関連
    ◇ 変更を保存: gf save
    ◇ 最終セーブ地点をロード: gf load

  ● クリーン関連
    ◇ 現在の変更を全て消す: gf clear
    ◇ 現在のブランチを強制的にリモートのブランチに合わせる: gf sync

  ● プッシュ・プル
    ◇ 現在のブランチにプッシュ: gf push
    ◇ 現在のブランチをプル: gf pull

  ● コミット
    ◇ 現在のブランチにコミット: gf commit <commitMessage>

  ● マージ
    ◇ ローカルブランチをマージ: gf merge <branchName>
    ◇ リモートブランチをマージ: gf merge -r <branchName>

```
