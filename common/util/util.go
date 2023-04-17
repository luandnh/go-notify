package util

import (
	"math/rand"
	"strconv"
	"time"
)

func ParsePageSize(param string) int {
	if len(param) < 1 {
		return 10
	}
	pageSize, err := strconv.Atoi(param)
	if err != nil {
		return 10
	}
	return pageSize
}

func ParsePage(param string) []byte {
	if len(param) < 1 {
		return nil
	}
	_, err := strconv.Atoi(param)
	if err != nil {
		return nil
	}
	return []byte(param)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_-.^")

func GenRandomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
