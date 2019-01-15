package qiita

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
	"github.com/uskey512/qiita-url-replace/models/qiita"
)

const (
	QiitaEndPointBase     = "https://qiita.com/api/v2/"
	QiitaTeamEndpointBase = "https://%s.qiita.com/api/v2/"
)

type QiitaClient struct {
	endpointBase string
	authToken    string
}

func NewQiitaClient(authToken string) (*QiitaClient, error) {
	if authToken == "" {
		return nil, errors.New("authToken is empty")
	}

	return &QiitaClient{QiitaEndPointBase, authToken}, nil
}

func NewQiitaTeamClient(team, authToken string) (*QiitaClient, error) {
	if authToken == "" {
		return nil, errors.New("authToken is empty")
	}

	if team == "" {
		return nil, errors.New("team is empty")
	}

	return &QiitaClient{fmt.Sprintf(QiitaTeamEndpointBase, team), authToken}, nil
}

func (q *QiitaClient) GetAuthenticatedUser() (GetAuthenticatedUserResponse, error) {
	url := q.endpointBase + "authenticated_user"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "GetAuthenticatedUser fail")
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", q.authToken))

	client := new(http.Client)
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	raw, _ := ioutil.ReadAll(resp.Body)
	var data qiita.GetAuthenticatedUser

	if err := json.Unmarshal(raw, &data); err != nil {
		return nil, errors.Wrap(err, "Json Unmarshal error : ", err)
	}
	fmt.Println(data[0].Body)

	return data, nil
}
