package main

import (
	"github.com/shopspring/decimal"
	"math"
	"math/big"
	"strconv"
)

// Epsilon 自定义精度
const Epsilon float64 = 0.0000001

// 0.2210 ns/op	存在误差，但是可以约定
func Equal1(a, b float64) bool {
	return math.Abs(a-b) < Epsilon
}

// 37.50 ns/op  小数位拆开分别保存，缺点是引用被修改的坑，如果使用计算也要用这个
func Equal2(a, b float64) bool {
	return big.NewFloat(a).Cmp(big.NewFloat(b)) == 0
}

// 1092 ns/op  慢但是安全
func Equal3(a, b float64) bool {
	return decimal.NewFromFloat(a).Cmp(decimal.NewFromFloat(b)) == 0
}

// 75.01 ns/op
func Equal4() {
	//	%f 的调用
	strconv.FormatFloat(1.2345, 'f', -1, 64)
}

// 这个函数的意思是返回目标值最接近的浮点数，如果目标相等就直接返回
// 累计误差如果大于1bit就不能够判断相等
//
//	func Nextafter(x, y float64) (r float64) {
//		switch {
//		case IsNaN(x) || IsNaN(y): // special case
//			r = NaN()
//		case x == y:
//			r = x
//		case x == 0:
//			r = Copysign(Float64frombits(1), y)
//		case (y > x) == (x > 0):
//			r = Float64frombits(Float64bits(x) + 1)
//		default:
//			r = Float64frombits(Float64bits(x) - 1)
//		}
//		return
//	}
func Equal5(a, b float64) bool {
	return math.Nextafter(a, b) == b
}

// 1.548 ns/op  优点：对于非复杂计算，用于Nextafter32可以实现精确比较 缺点是只支持1bit的误差，方案应用不广泛
func Equal6(a, b float64) bool {
	return math.Nextafter32(float32(a), float32(b)) == float32(b)
}
