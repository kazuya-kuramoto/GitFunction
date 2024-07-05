########## CONST ##########
# サイレントモードで実行
.SILENT:
# アプリ/コマンド名称
APP_NAME=gf
# ビルドターゲットスクリプトパス
BUILD_TARGET=main.go
# Windowsの場合のOS/アーキテクチャの定数
GOOS_WIN=windows
GOARCH_WIN=386
# MacOSの場合のOS/アーキテクチャの定数
GOOS_MAC=darwin
GOARCH_MAC_AMD64=amd64
GOARCH_MAC_ARM64=arm64

########## COMMAND ##########
install:
	go build -o bin/$(APP_NAME) $(BUILD_TARGET)
	cp ./bin/$(APP_NAME) /usr/local/bin/$(APP_NAME)
	echo "インストールが終了しました。"

build:
    # UNIX
	go build -o bin/$(APP_NAME) $(BUILD_TARGET)

	# WIN
	GOOS=$(GOOS_WIN) \
	GOARCH=$(GOARCH_WIN) \
	go build -o bin/$(APP_NAME).exe $(BUILD_TARGET)

	# MacOS (Intel)
	GOOS=$(GOOS_MAC) \
	GOARCH=$(GOARCH_MAC_AMD64) \
	go build -o bin/$(APP_NAME)-mac-amd64 $(BUILD_TARGET)

	# MacOS (Apple Silicon)
	GOOS=$(GOOS_MAC) \
	GOARCH=$(GOARCH_MAC_ARM64) \
	go build -o bin/$(APP_NAME)-mac-arm64 $(BUILD_TARGET)

	echo "ビルドが完了したでございます。"

update:
	go build -o bin/$(APP_NAME) $(BUILD_TARGET)
	cp ./bin/$(APP_NAME) /usr/local/bin/$(APP_NAME)
	echo "アップデートが終了しました。"
