package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func ensureDirectoryExists(path string) cli.ExitCoder {
	if _, err := os.Stat(path); err != nil {
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			msg := fmt.Errorf("ディレクトリが作成できませんでした: %w", err)
			return cli.Exit(msg, exitCodeErrWrite.ToInt())
		}
	}

	return nil
}

func showSuccess() {
	fmt.Println("✨ 完了しました")
}
