# GitFunction
Gitのコマンドを省略化して、普段操作をするコマンドの量を75%程度圧縮できます。
様々なコマンドの省コマンド化をしていますので、普段の業務効率を改善できます。

## 利用可能環境
- MAC OS
- Windows
- UNIX

## 免責事項
- このソフトウェアを使用し、データの破損等が発生しいかなる損害が発生したとしてもこのソフトウェアの開発者は責任を取りません。
- MITライセンスとして提供をしており、業務上の利用には専門の部署にて使用の確認を行なってください。
- このプロジェクトは有志による開発であって、利益を獲得することを目的とした活動ではありません。
- Gitのバージョンによってはこのソフトウェアが正しく動作しないことがあります。
- 導入やバグレポートの対応など、サポートは一切いたしません。
- 動作環境、利用可能環境とは動作を保証するものではありませんので、必ず動作の確認は個人の責任として行なってください。
- 転売、再配布、作者を偽る行為はご遠慮ください。

### リリースからのインストール

1. [リリースページ](https://github.com/yourusername/yourrepository/releases)にアクセスします。
2. 最新のリリースを選択します。
3. 適切なバイナリファイルをダウンロードします。
   - Windows: `gf.exe`
   - UNIX/Linux: `gf`
   - MacOS (Intel): `gf-mac-amd64`
   - MacOS (Apple Silicon): `gf-mac-arm64`

#### Windows

1. ダウンロードした`gf.exe`ファイルを適当なディレクトリに配置します。
2. コマンドプロンプトまたはPowerShellを開き、次のコマンドを実行してパスを通します（任意）:

   ```sh
   setx PATH "%PATH%;C:\path\to\directory"
   ```

3. インストールが完了しました。以下のコマンドで動作確認を行ってください:

   ```sh
   gf --version
   ```

#### UNIX/Linux

1. ダウンロードした`gf`ファイルを`/usr/local/bin`にコピーします:

   ```sh
   sudo cp gf /usr/local/bin/
   ```

2. 実行権限を付与します:

   ```sh
   sudo chmod +x /usr/local/bin/gf
   ```

3. インストールが完了しました。以下のコマンドで動作確認を行ってください:

   ```sh
   gf --version
   ```

#### MacOS (Intel)

1. ダウンロードした`gf-mac-amd64`ファイルを`/usr/local/bin`にコピーします:

   ```sh
   sudo cp gf-mac-amd64 /usr/local/bin/gf
   ```

2. 実行権限を付与します:

   ```sh
   sudo chmod +x /usr/local/bin/gf
   ```

3. インストールが完了しました。以下のコマンドで動作確認を行ってください:

   ```sh
   gf --version
   ```

#### MacOS (Apple Silicon)

1. ダウンロードした`gf-mac-arm64`ファイルを`/usr/local/bin`にコピーします:

   ```sh
   sudo cp gf-mac-arm64 /usr/local/bin/gf
   ```

2. 実行権限を付与します:

   ```sh
   sudo chmod +x /usr/local/bin/gf
   ```

3. インストールが完了しました。以下のコマンドで動作確認を行ってください:

   ```sh
   gf --version
   ```

## Update

```bash
gf update
```

## Usage

```
-------------------------------------------
 GitFunction v1.0.0 Usage
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