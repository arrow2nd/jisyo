package main

import (
	"fmt"
	"os"
	"text/template"

	"github.com/urfave/cli/v2"
)

func listCommand() *cli.Command {
	return &cli.Command{
		Name:    "list",
		Aliases: []string{"l"},
		Usage:   "追加済の辞書を表示",
		Action: func(ctx *cli.Context) error {
			funcMap := template.FuncMap{
				"inc": func(i int) int {
					return i + 1
				},
			}

			temp := `{{range $i, $j := .Jisyos -}}
{{inc $i}}. {{.URL}}
  - hash: {{.SHA256}}
{{end -}}`

			t, err := template.New("list").Funcs(funcMap).Parse(temp)
			if err != nil {
				msg := fmt.Errorf("テンプレートの解析に失敗しました: %w", err)
				return cli.Exit(msg, exitCodeErrTemplate.ToInt())
			}

			if err := t.Execute(os.Stdout, sharedConfig); err != nil {
				msg := fmt.Errorf("出力に失敗しました: %w", err)
				return cli.Exit(msg, exitCodeErrTemplate.ToInt())
			}

			return nil
		},
	}
}
