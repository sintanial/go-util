package strutil

import (
	"strings"
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"encoding/json"
	"strconv"
)

func IndexOf(search string, list []string) int {
	for i := 0; i < len(list); i++ {
		if list[i] == search {
			return i
		}
	}

	return -1
}

func Contains(search string, list []string) bool {
	for i := 0; i < len(list); i++ {
		if strings.Contains(search, list[i]) {
			return true
		}
	}

	return false
}

func Concat(args ...interface{}) string {
	return Join("", args...)
}

func Join(sep string, args ...interface{}) string {
	var res []string
	for _, arg := range args {
		switch arg := arg.(type) {
		case string: res = append(res, arg)
		case bool: res = append(res, strconv.FormatBool(arg))
		case int: res = append(res, strconv.Itoa(arg))
		case float64: res = append(res, strconv.FormatFloat(arg, 'f', -1, 64))
		}
	}

	return strings.Join(res, sep)
}

func Md5(data string) string {
	uniqid := md5.Sum([]byte(data))
	return hex.EncodeToString(uniqid[:])
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func MarshalJson(i interface{}) string {
	bt, _ := json.Marshal(i)
	return string(bt)
}