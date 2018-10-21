package util

import (
	"fmt"
	"sort"

	"strings"

	"crypto/md5"
	"io"
)

func Signature(apiType string, clientID string, clientSecret string, accessToken string, timestamp int64, params map[string]interface{}) string {

	params["type"] = apiType
	params["client_id"] = clientID
	if accessToken != "" {
		params["access_token"] = accessToken
	}
	params["timestamp"] = timestamp
	params["data_type"] = "JSON"
	array := make([]string, 0)
	for k, _ := range params {
		array = append(array, k)
	}
	sort.Strings(array)
	sign := clientSecret
	for _, v := range array {
		sign = sign + v + fmt.Sprint(params[v])
	}
	sign = sign + clientSecret
	sign = MD5(sign)
	return strings.ToUpper(sign)
}

func MD5(src string) string {
	h := md5.New()
	io.WriteString(h, src)
	return fmt.Sprintf("%x", h.Sum(nil))
}
