package mail

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func doRequest(r *http.Request, v interface{}) error {
	client := &http.Client{}
	resp, err := client.Do(r)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, v)
	if err != nil {
		return err
	}
	return nil
}
