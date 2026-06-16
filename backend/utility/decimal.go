package utility

import "github.com/shopspring/decimal"

// DivIntToInt 整数除整数，返回整数部分
func DivIntToInt(a, b int) int {
	if b == 0 {
		return 0
	}
	return int(decimal.NewFromInt(int64(a)).Div(decimal.NewFromInt(int64(b))).Round(0).IntPart())
}

// DivIntToFloat2 整数除整数，返回2位小数
func DivIntToFloat2(a, b int) float64 {
	if b == 0 {
		return 0
	}
	res, _ := decimal.NewFromInt(int64(a)).Div(decimal.NewFromInt(int64(b))).Round(2).Float64()
	return res
}

// DivIntToPercent 整数除整数，返回2位小数百分比（乘以100）
func DivIntToPercent(a, b int) float64 {
	if b == 0 {
		return 0
	}
	res, _ := decimal.NewFromInt(int64(a)).Div(decimal.NewFromInt(int64(b))).Mul(decimal.NewFromInt(100)).Round(2).Float64()
	return res
}

// DivLTV 计算 LTV（生命周期价值）
// ARPU = 时间段内充值总金额 / 总新增人数
// LT = 时间段内的天数 / (1 - 留存率)
// betRate 为留存率百分比（如 80 表示 80%）
func DivLTV(arpu, dayNum int, betRate float64) float64 {
	if betRate == 100 {
		return 0
	}
	// 留存率是百分比，需要 100 - xx 再乘 100
	res, _ := decimal.NewFromInt(int64(arpu)).
		Mul(decimal.NewFromInt(int64(dayNum))).
		Div(decimal.NewFromFloat(100 - betRate)).
		Mul(decimal.NewFromInt(100)).
		Round(2).Float64()
	return res
}
