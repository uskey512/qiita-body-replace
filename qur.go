package main

import (
	"fmt"

	"github.com/uskey512/qiita-url-replace/services"
)

func getParameter() (services.ConnectionSetting, services.ReplaceSetting) {
	var token, team, srcDomain, dstDomain string

	fmt.Print("トークン : ")
	fmt.Scan(&token)

	fmt.Print("変換前ドメイン : ")
	fmt.Scan(&srcDomain)

	fmt.Print("変更後ドメイン : ")
	fmt.Scan(&dstDomain)

	fmt.Print("所属チーム(Optional) : ")
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
