package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v2"
)

type jisyo struct {
	URL    string `json:"url"`
	SHA256 string `json:"sha256"`
}

type config struct {
	DirPath string  `json:"dirPath"`
	Jisyos  []jisyo `json:"jisyos"`
}

const configFileName = "jisyo.json"

func defaultConfigPath() (string, cli.ExitCoder) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		msg := fmt.Errorf("ホームディレクトリが取得できませんでした: %w", err)
		return "", cli.Exit(msg, exitCodeErrRead.ToInt())
	}

	configDir := filepath.Join(homeDir, ".config", "jisyo")

	// ディレクトリが無いなら作成
	if _, err := os.Stat(configDir); err != nil {
		if err := os.MkdirAll(configDir, os.ModePerm); err != nil {
			msg := fmt.Errorf("設定ディレクトリが作成できませんでした: %w", err)
			return "", cli.Exit(msg, exitCodeErrWrite.ToInt())
		}
	}

	return filepath.Join(configDir, configFileName), nil
}

func loadConfig(ctx *cli.Context) (*config, cli.ExitCoder) {
	path := ctx.String("config")

	buf, err := os.ReadFile(path)
	if err != nil {
		msg := fmt.Errorf("設定ファイルの読み込みに失敗しました。jisyo init を実行してください: %w", err)
		return nil, cli.Exit(msg, exitCodeErrRead.ToInt())
	}

	out := &config{}
	if err := json.Unmarshal(buf, out); err != nil {
		msg := fmt.Errorf("設定ファイルの解析に失敗しました: %w", err)
		return nil, cli.Exit(msg, exitCodeErrConfig.ToInt())
	}

	return out, nil
}

func saveConfig(ctx *cli.Context, in config) cli.ExitCoder {
	buf := &bytes.Buffer{}

	if err := json.NewEncoder(buf).Encode(in); err != nil {
		msg := fmt.Errorf("設定ファイルの解析に失敗しました: %w", err)
		return cli.Exit(msg, exitCodeErrConfig.ToInt())
	}

	path := ctx.String("config")

	if err := os.WriteFile(path, buf.Bytes(), os.ModePerm); err != nil {
		msg := fmt.Errorf("設定ファイルの書き込みに失敗しました: %w", err)
		return cli.Exit(msg, exitCodeErrWrite.ToInt())
	}

	return nil
}
