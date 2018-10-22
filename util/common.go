package util

import (
	"encoding/json"
	"fmt"
	"strings"
)

func NewInt(i int) *int {
	return &i
}

func NewString(s string) *string {
	return &s
}

func NewBool(b bool) *bool {
	return &b
}
func NewInt64(i int64) *int64 {
	return &i
}

func ArrayStringToString(arr []string) string {
	return `["` + strings.Join(arr, `","`) + `"]`
}

func FormatURLParams(params interface{}) (map[string]interface{}, string) {
	paramsURL := ""
	var result map[string]interface{}
	if params != nil {
		j, _ := json.Marshal(params)
		decoder := json.NewDecoder(strings.NewReader(string(j)))
		decoder.UseNumber()
		decoder.Decode(&result)
		if len(result) != 0 {
			for k, v := range result {
				switch v.(type) {
				case []interface{}:
					m, _ := json.Marshal(v)
					mStr := string(m)
					result[k] = mStr
					paramsURL = paramsURL + fmt.Sprintf("&%s=%v", k, mStr)
					break
				default:
					paramsURL = paramsURL + fmt.Sprintf("&%s=%v", k, v)
					break
				}

			}
		}
	}
	if result == nil {
		result = make(map[string]interface{})
	}
	return result, paramsURL
}
