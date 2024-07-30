package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

var (
	RequestError = errors.New("request return status error")
)

func HttpPOST(url string, paramsBody io.Reader, result interface{}) error {
	req, err := http.NewRequest("POST", url, paramsBody)
	if err != nil {
		return err
	}
	//http.DefaultClient.Timeout = 30 * time.Second
	res, err := http.DefaultClient.Do(req)
	defer res.Body.Close()
	if err != nil {
		return err
	}

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
