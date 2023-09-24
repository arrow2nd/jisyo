package main

import "github.com/urfave/cli/v2"

type exitCode int

func (e exitCode) ToInt() int {
	return int(e)
}

const (
	// exitCodeOK : 正常
	exitCodeOK exitCode = iota
	// exitCodeErr : エラー
	exitCodeErr
	// exitCodeErrConfig : 設定ファイルエラー
	exitCodeErrConfig exitCode = iota + 62
	// exitCodeErrRead : 読み込みエラー
	exitCodeErrRead
	// exitCOdeErrWrite : 書き込みエラー
	exitCodeErrWrite
	// exitCodeErrTemplate : テンプレートエラー
	exitCodeErrTemplate
	// exitCodeErrNetwork : ネットワークエラー
	exitCodeErrNetwork
	// exitCodeErrExtract : 展開エラー
	exitCodeErrExtract
	// exitCodeErrArg : 引数エラー
	exitCodeErrArg
)

func exitCancel() cli.ExitCoder {
	return cli.Exit("🚫 キャンセルしました", exitCodeErr.ToInt())
}
