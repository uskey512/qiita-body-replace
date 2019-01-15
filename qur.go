package main

import (
	"fmt"

	"github.com/uskey512/qiita-url-replace/client/qiita"
	"github.com/uskey512/qiita-url-replace/models/qiita"
)

type setting struct {
	token     string
	srcDomain string
	dstDomain string
	team      string
}

func getParameter() setting {
	var token, team, srcDomain, dstDomain string

	fmt.Print("token : ")
	fmt.Scan(&token)

	fmt.Print("変換前ドメイン : ")
	fmt.Scan(&srcDomain)

	fmt.Print("変更後ドメイン : ")
	fmt.Scan(&dstDomain)

	fmt.Print("所属team(Optional) : ")
	fmt.Scan(&team)

	return setting{token, srcDomain, dstDomain, team}
}

func main() {
	s := getParameter()

	var qc QiitaClient
	if s.team == nil {
		qc = qiita.NewQiitaClient(s.token)
	} else {
		qc = qiita.NewQiitaTeamClient(s.team, s.token)
	}

	user, _ := qc.GetAuthenticatedUser()
	fmt.Println(user.ItemsCount)

}
