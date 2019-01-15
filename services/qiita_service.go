package services

import (
	"fmt"
	"os"
	"strings"

	"github.com/uskey512/qiita-url-replace/clients"
	"github.com/uskey512/qiita-url-replace/models"
)

const getPerPage = 100

type ConnectionSetting struct {
	AuthToken string
	Team      string
}

type ReplaceSetting struct {
	SrcDomain string
	DstDomain string
}

func InitService(c ConnectionSetting) *clients.QiitaClient {
	var qc *clients.QiitaClient
	var err error

	if c.Team == "" {
		qc, err = clients.NewQiitaClient(c.AuthToken)
	} else {
		qc, err = clients.NewQiitaTeamClient(c.Team, c.AuthToken)
	}

	if err != nil {
		fmt.Printf("InitService Error %+v", err)
		os.Exit(1)
	}

	return qc
}

func ReplaceBodyUrlDomain(qc *clients.QiitaClient, r ReplaceSetting) {

	page := 1
	itemCount := 0
	replacedItemCount := 0

	fmt.Println("置換対象URL")
	for {
		items, err := qc.GetAuthenticatedUserItems(page, getPerPage)
		if err != nil {
			fmt.Printf("InitService Error %+v", err)
			os.Exit(1)
		}

		for i := 0; i < len((*items).Response); i++ {
			itemCount++
			if strings.Index((*items).Response[i].Body, r.SrcDomain) != -1 {
				item := (*items).Response[i]

				replacedBody := strings.Replace(item.Body, r.SrcDomain, r.DstDomain, -1)
				replacedPatchRequest := models.PatchItemsIdRequest{
					Title: item.Title,
					Body:  replacedBody,
				}

				qc.PatchItemById(item.ID, &replacedPatchRequest)
				replacedItemCount++

				fmt.Printf("replace done. url : %s ", item.URL)
			}
		}
		if len((*items).Response) != getPerPage {
			break
		}
		page++
	}

	fmt.Printf("変換した記事の数 : %d / %d\n", replacedItemCount, itemCount)
}
