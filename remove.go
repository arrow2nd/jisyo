package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/urfave/cli/v2"
)

func removeCommand() *cli.Command {
	return &cli.Command{
		Name:        "remove",
		Aliases:     []string{"r"},
		Usage:       "辞書を削除",
		Description: "辞書ファイルも削除されます",
		ArgsUsage:   "<辞書名...>",
		Before: func(ctx *cli.Context) error {
			if ctx.Args().Len() == 0 {
				return cli.Exit("引数エラー: 1つ以上の引数が必要です", exitCodeErrArg.ToInt())
			}

			return nil
		},
		Action: func(ctx *cli.Context) error {
			removeJisyos := ctx.Args().Len()

			for i, removeJisyoName := range ctx.Args().Slice() {
				fmt.Println("🧹 削除を開始します")

				newJisyos := []jisyo{}
				removeJisyoName = strings.TrimSpace(removeJisyoName)

				for _, jisyo := range sharedConfig.Jisyos {
					if jisyo.Name != removeJisyoName {
						newJisyos = append(newJisyos, jisyo)
						continue
					}

					fmt.Printf("[%d/%d] 🗑 %s (%s)\n", i+1, removeJisyos, jisyo.Name, jisyo.URL)

					path := filepath.Join(sharedConfig.DirPath, jisyo.Name)

					// ファイルがなければ削除処理は実行しない
					if _, err := os.Stat(path); err != nil {
						continue
					}

					if err := os.Remove(path); err != nil {
						msg := fmt.Errorf("削除に失敗しました: %w", err)
						return cli.Exit(msg, exitCodeErrWrite.ToInt())
					}
				}

				if len(newJisyos) == len(sharedConfig.Jisyos) {
					msg := fmt.Sprintf("「%s」は登録されていません", removeJisyoName)
					return cli.Exit(msg, exitCodeErr.ToInt())
				}

				sharedConfig.Jisyos = newJisyos
			}

			showSuccess()

			return saveConfig(ctx, *sharedConfig)
		},
	}
}
