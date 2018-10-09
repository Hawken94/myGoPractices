package utils

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestPrintCardsType(t *testing.T) {
	printCardsType()
	t.Errorf("xhk")
}

func TestRandomStr(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(randomStr(32))
	}
	t.Error(rand.NewSource(time.Now().UnixNano()).Int63(), ":", letterIdxMask)
	t.Error("random str:", randomStr(32))
}

func randomStr(randomLength int) string {
	return RandStringBytesMaskImprSrc(randomLength)
}

var src = rand.NewSource(time.Now().UnixNano())

const letterBytes = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func RandStringBytesMaskImprSrc(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

// 0hEmALxNxMaZUCnnyL3YqBdqW6ohdiL0

var t1 rune
var t2 byte

func TestFloat64(t *testing.T) {
	var tests = []struct {
		number int
	}{
		{1200}, {1234}, {78000},
	}

	for _, v := range tests {
		// result, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(v.number)/100.00), 64)
		result := math.Trunc(float64(v.number) * float64(1e2) / float64(1e4))
		t.Error("转化后的数字:", result)
	}
}

func TestStrconvFloat64(t *testing.T) {

	r, _ := strconv.ParseFloat("100.00", 64)
	t.Error("转化后的数字:", r, float64(910)/100.00)
}
