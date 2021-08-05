package util

import (
	"math/rand"
	"time"
)

func RandomString(n int) string {
	var letters = []byte("ssrgdthfythgdfsgvfvdisfivvdiefhurvhuaefmrgjrgjasdkefrggrthvrv")
	result := make([]byte, n)

	// 如果不使用rand.Seed(seed int64)，每次运行，得到的随机数会一样，程序不停止，一直获取的随机数是不一样的
	// 每次运行时rand.Seed(seed int64)，seed的值要不一样，这样生成的随机数才会和上次运行时生成的随机数不一样
	rand.Seed(time.Now().Unix())

	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}

	return string(result)
}
