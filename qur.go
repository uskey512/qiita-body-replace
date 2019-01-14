package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type setting struct {
	token     string
	srcDomain string
	dstDomain string
}

func getParameter() setting {
	var token, srcDomain, dstDomain string

	fmt.Print("token : ")
	fmt.Scan(&token)

	fmt.Print("変換前ドメイン : ")
	fmt.Scan(&srcDomain)

	fmt.Print("変更後ドメイン : ")
	fmt.Scan(&dstDomain)

	return setting{
		token:     token,
		srcDomain: srcDomain,
		dstDomain: dstDomain,
	}
}

func readFromQiita(s setting) {
	url := fmt.Sprintf("https://%s/api/v2/authenticated_user/items", s.dstDomain)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.token))

	client := new(http.Client)
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	raw, _ := ioutil.ReadAll(resp.Body)
	var data qiitaPost

	if err := json.Unmarshal(raw, &data); err != nil {
		fmt.Println("JSON Unmarshal error:", err)
		return
	}
	fmt.Println(data[0].Body)
}

func main() {
	s := getParameter()
	readFromQiita(s)
}