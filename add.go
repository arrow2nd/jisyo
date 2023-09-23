package main

import (
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/urfave/cli/v2"
)

func addCommand() *cli.Command {
	return &cli.Command{
		Name:        "add",
		Aliases:     []string{"a"},
		Usage:       "辞書を追加",
		Description: "辞書のURLを追加します",
		Action: func(ctx *cli.Context) error {
			// 辞書のURL
			prompt := &survey.Input{Message: "辞書のURL"}

			opts := survey.WithValidator(func(ans interface{}) error {
				url, ok := ans.(string)

				if !ok || !strings.HasPrefix(url, "http") {
					return fmt.Errorf("不正なURLです")
				}

				return nil
			})

			url := ""
			if err := survey.AskOne(prompt, &url, opts); err != nil {
				return nil // キャンセル
			}

			// 重複確認
			for _, j := range sharedConfig.Jisyos {
				if j.URL == url {
					return cli.Exit("この辞書は既に追加されています", exitCodeErr.ToInt())
				}
			}

			// ダウンロード

			// ハッシュ値を計算
			hash := ""

			// 保存
			sharedConfig.Jisyos = append(sharedConfig.Jisyos, jisyo{
				URL:    url,
				SHA256: hash,
			})

			return saveConfig(ctx, *sharedConfig)
		},
	}
}
