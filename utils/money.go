package utils

import (
	"fmt"
	"github.com/shopspring/decimal"
)

//将元转换为分
var oneHundred, _ = decimal.NewFromString("100")

func ConvertStringYuan2Fen(yuan string) int64 {
	source, err := decimal.NewFromString(yuan)
	if err != nil {
		fmt.Print("转换元到分异常", err.Error())
		return 0
	}
	res := source.Mul(oneHundred)
	return res.IntPart()
}

func ConvertInt64Fen2Yuan(fen int64) string {
	source := decimal.New(fen, 0)
	res := source.DivRound(oneHundred, 2)
	return res.String()
}
