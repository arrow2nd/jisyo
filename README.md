# jisyo

[![release](https://github.com/arrow2nd/jisyo/actions/workflows/release.yml/badge.svg)](https://github.com/arrow2nd/jisyo/actions/workflows/release.yml)

📚 SKK辞書マネージャ

## できること

- 辞書の管理 (JSON)
- 指定ディレクトリへの一括ダウンロード
- 辞書の追加・削除

## help

```sh
$ jisyo -h

NAME:
   jisyo - 📚 SKK辞書マネージャ

USAGE:
   jisyo [global options] command [command options] [arguments...]

VERSION:
   0.0.5

COMMANDS:
   init, i      設定ファイルを初期化
   download, d  辞書をダウンロード
   add, a       辞書を追加
   remove, r    辞書を削除
   list, l      追加済の辞書を表示
   config, c    設定ファイルをエディタで開く
   help, h      Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --config value, -c value  設定ファイルのパス
   --help, -h                show help
   --version, -v             print the version
```
