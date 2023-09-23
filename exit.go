package main

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
)
