package pkg

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ConvertJson(url string, v interface{}) error{
	body, err := GetJson(url)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &v)
	if err != nil {
		return err
	}

	return nil
}

func GetJson(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil{
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil{
		return nil, err
	}

	return body, nil
}
