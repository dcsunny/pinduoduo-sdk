package util

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

func HttpPOST(url string, paramsBody io.Reader, result interface{}) error {
	req, _ := http.NewRequest("POST", url, paramsBody)
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return RequestError
	}
	body, _ := ioutil.ReadAll(res.Body)
	err := json.Unmarshal(body, &result)
	if err != nil {

		return err
	}

	return nil
}
