package main

import (
	"fmt"

	"github.com/uskey512/qiita-url-replace/services"
)

func getParameter() (services.ConnectionSetting, services.ReplaceSetting) {
	var token, team, srcDomain, dstDomain string

	fmt.Print("アクセストークン : ")
	fmt.Scan(&token)

	fmt.Print("置換前文字列 : ")
	fmt.Scan(&srcDomain)

	fmt.Print("置換後文字列 : ")
	fmt.Scan(&dstDomain)

	fmt.Print("所属チーム(Qiita:Teamの場合のみ) : ")
	fmt.Scan(&team)

	return services.ConnectionSetting{
			AuthToken: token,
			Team:      team,
		}, services.ReplaceSetting{
			SrcDomain: srcDomain,
			DstDomain: dstDomain,
		}
}

func main() {
	c, r := getParameter()
	qc := services.InitService(c)
	services.ReplaceBodyUrlDomain(qc, r)
}
