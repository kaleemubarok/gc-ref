package userauth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func (auth *AuthClient) GetUserInfo(accessToken string) *User {
	client := &http.Client{
		Timeout: auth.timeout,
	}

	req, err := http.NewRequest("GET", auth.host+"/user/info", nil)
	if err != nil {
		return nil
	}
	req.Header.Set("X-Access-Token", accessToken)

	respRaw, err := client.Do(req)
	if err != nil {
		return nil
	}
	defer respRaw.Body.Close()

	if respRaw.StatusCode != 200 {
		return nil
	}

	respByte, err := ioutil.ReadAll(respRaw.Body)
	if err != nil {
		log.Print(err)
		return nil
	}

	var resp struct {
		Err  string `json:"err"`
		Data User   `json:"data"`
	}

	err = json.Unmarshal(respByte, &resp)
	if err != nil {
		return nil
	}

	return &resp.Data
}

func (auth *AuthClient) GetUserByID(accessToken, userID string) *User {
	client := &http.Client{
		Timeout: auth.timeout,
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/user/%s", auth.host, userID), nil)
	if err != nil {
		return nil
	}
	req.Header.Set("X-Access-Token", accessToken)

	respRaw, err := client.Do(req)
	if err != nil {
		return nil
	}
	defer respRaw.Body.Close()

	if err != nil {
		return nil
	}

	if respRaw.StatusCode != 200 {
		return nil
	}

	respByte, err := ioutil.ReadAll(respRaw.Body)
	if err != nil {
		log.Print(err)
		return nil
	}

	var resp struct {
		Err  string `json:"err"`
		Data User   `json:"data"`
	}

	err = json.Unmarshal(respByte, &resp)
	if err != nil {
		return nil
	}

	return &resp.Data
}
