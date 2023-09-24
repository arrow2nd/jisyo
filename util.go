package main

import (
	"fmt"
	"os"
	"strings"

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

func expandTilde(path string) (string, cli.ExitCoder) {
	if strings.HasPrefix(path, "~") {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			msg := fmt.Errorf("ホームディレクトリの取得に失敗しました: %w", err)
			return "", cli.Exit(msg, exitCodeErrRead.ToInt())
		}

		return strings.Replace(path, "~", homeDir, 1), nil
	}

	return path, nil
}

func showSuccess() {
	fmt.Println("✨ 完了しました")
}
