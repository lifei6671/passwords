package passwords

import (

	"math/rand"
	"time"
)

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

//生成指定长度的随机byte数组.
func RandomBytes(n int) []byte {

	b := make([]byte, n)
	rand.Seed(time.Now().UnixNano())
	_, err := rand.Read(b)

	if err != nil {
		src := rand.NewSource(time.Now().UnixNano())
		for i, cache, remain := n - 1, src.Int63(), letterIdxMax; i >= 0; remain-- {
			if remain == 0 {
				cache, remain = src.Int63(), letterIdxMax
			}
			if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
				b[i] = letterBytes[idx]
				i--
			}
			cache >>= letterIdxBits
		}
	}
	return b
}
//获取指定区间的随机数.
func RandomInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min)
}