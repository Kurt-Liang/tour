package cmd

import (
	"log"
	"strings"

	"github.com/Kurt-Liang/tour/internal/word"
	"github.com/spf13/cobra"
)

const (
	ModeUpper = iota + 1
	ModeLower
	ModeUnderscoreToUpperCamelCase
	ModeUnderscoreToLowerCamelCase
	ModeCamelCaseToUnderscore
)

var str string
var mode int8

var desc = strings.Join([]string{
	"該子指令支持各種單字格式轉換，模式如下：",
	"1：全部轉大寫",
	"2：全部轉小寫",
	"3：底線單字轉換成大寫駝峰單字",
	"4：底線單字轉換成小寫駝峰單字",
	"5：駝峰單字轉底線單字",
}, "\n")

var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "單字格式轉換",
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case ModeUpper:
			content = word.ToUpper(str)
		case ModeLower:
			content = word.ToLower(str)
		case ModeUnderscoreToUpperCamelCase:
			content = word.UnderscoreToUpperCamelCase(str)
		case ModeUnderscoreToLowerCamelCase:
			content = word.UnderscoreToLowerCamelCase(str)
		case ModeCamelCaseToUnderscore:
			content = word.CamelCaseToUnderscore(str)
		default:
			log.Fatalf("暫不支援該轉換模式，請執行 word --help 檢視說明文件")
		}

		log.Printf("輸出結果: %s", content)
	},
}

func init() {
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "請輸入單字")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "請輸入單字的轉換模式")
}
