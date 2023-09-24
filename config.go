package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/urfave/cli/v2"
)

type jisyo struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type config struct {
	DirPath string  `json:"dirPath"`
	Jisyos  []jisyo `json:"jisyos"`
}

func (c *config) GetDirPath() (string, cli.ExitCoder) {
	dirPath, exit := expandTilde(c.DirPath)
	if exit != nil {
		return "", exit
	}

	return dirPath, nil
}

const configFileName = "jisyo.json"

func configCommand() *cli.Command {
	return &cli.Command{
		Name:    "config",
		Aliases: []string{"c"},
		Usage:   "設定ファイルをエディタで開く",
		Action: func(ctx *cli.Context) error {
			editor := ctx.String("editor")
			configPath := ctx.String("config")
			cmd := exec.Command(editor, configPath)
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err := cmd.Run(); err != nil {
				msg := fmt.Errorf("エディタの起動に失敗しました: %w", err)
				return cli.Exit(msg, exitCodeErr.ToInt())
			}
			return nil
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "editor",
				Usage:   "エディタ",
				Value:   os.Getenv("EDITOR"),
				Aliases: []string{"e"},
			},
		},
	}
}

func defaultConfigPath() (string, cli.ExitCoder) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		msg := fmt.Errorf("ホームディレクトリが取得できませんでした: %w", err)
		return "", cli.Exit(msg, exitCodeErrRead.ToInt())
	}

	configDir := filepath.Join(homeDir, ".config", "jisyo")

	// ディレクトリが無いなら作成
	if exit := ensureDirectoryExists(configDir); exit != nil {
		return "", exit
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
	bytes, err := json.MarshalIndent(in, "", "  ")

	if err != nil {
		msg := fmt.Errorf("設定ファイルの解析に失敗しました: %w", err)
		return cli.Exit(msg, exitCodeErrConfig.ToInt())
	}

	path := ctx.String("config")

	if err := os.WriteFile(path, bytes, os.ModePerm); err != nil {
		msg := fmt.Errorf("設定ファイルの書き込みに失敗しました: %w", err)
		return cli.Exit(msg, exitCodeErrWrite.ToInt())
	}

	return nil
}
