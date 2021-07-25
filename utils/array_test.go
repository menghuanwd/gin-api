package utils

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHasOne(t *testing.T) {
	assert := assert.New(t)
	input := []string{"a", "b", "c"}

	target := "bc"
	target2 := "b"

	// if HasOne(input, target) {

	// } else {
	// 	t.Error("No exist", target, input)
	// }

	assert.Equal(HasOne(input, target), false)
	assert.Equal(HasOne(input, target2), true)
}
func Benchmark_Division(b *testing.B) {
	for i := 0; i < b.N; i++ { //use b.N for looping
		Division(4, 5)
	}
}

func Benchmark_TimeConsumingFunction(b *testing.B) {
	b.StopTimer() //调用该函数停止压力测试的时间计数

	//做一些初始化的工作,例如读取文件数据,数据库连接之类的,
	//这样这些时间不影响我们测试函数本身的性能

	b.StartTimer() //重新开始时间
	for i := 0; i < b.N; i++ {
		Division(4, 5)
	}
}
func Division(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("除数不能为0")
	}

	return a / b, nil
}
