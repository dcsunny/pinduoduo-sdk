package util

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"fmt"
)

var (
	RequestError = errors.New("request return status error")
)

func HttpPOST(url string, paramsBody io.Reader, result interface{}) error {
	req, err := http.NewRequest("POST", url, paramsBody)
	if err != nil {
		return err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return RequestError
	}
	body, _ := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println(fmt.Sprintf("json err,json:%s", string(body)))
		return err
	}

	return nil
}
