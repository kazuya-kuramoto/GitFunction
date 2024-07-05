#!/usr/bin/make -f

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

	echo "ビルドが完了したでございます。"

update:
	go build -o bin/$(APP_NAME) $(BUILD_TARGET)
	cp ./bin/$(APP_NAME) /usr/local/bin/$(APP_NAME)
	echo "アップデートが終了しました。"