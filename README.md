# protoc plugin作成用のリポジトリ

## 概要
protoc pluginを作成するためのリポジトリです。

## 使い方
1. このリポジトリをクローンします。
2. 新規にプラグインを作成する場合は、以下を実行
    ```shell
    bash new_plugin.sh <plugin_name>
    ```

3. 作成したプラグインのディレクトリに移動し、以下を実行
    ```shell
    make all
    ```

## コマンド
- `make all`: プラグインをビルドし、protocを実行
- `make build`: プラグインをビルド
- `make gen`: protocを実行
- `make clean`: ビルドしたファイルを削除

## 要求環境
- protoc
- go