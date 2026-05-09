package utils

import (
	"crypto/md5"
	"fmt"
	"sort"
	"strings"
)

func MakeSign(data map[string]string, secret string) string {
	// 1. 获取所有 key 并排序
	keys := make([]string, 0, len(data))
	for k := range data {
		if k != "sign" {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)

	// 2. 拼接 key=value
	params := make([]string, 0, len(data))
	for _, k := range keys {
		params = append(params, fmt.Sprintf("%s=%s", k, data[k]))
	}

	// 3. 用 & 连接
	queryString := strings.Join(params, "&")

	fmt.Printf("signStr: %s secret: %s\n\n", queryString, secret)

	// 4. MD5(queryString + secret)
	hash := md5.Sum([]byte(queryString + secret))
	return fmt.Sprintf("%x", hash)
}

func Pay247VerifyCallback(params map[string]string, mchKey string) bool {
	signature := params["sign"]

	currentSignature := MakeSign(params, mchKey)

	fmt.Printf("signature: %s\ncurrentSignature: %s\n", signature, currentSignature)
	return signature == currentSignature
}
