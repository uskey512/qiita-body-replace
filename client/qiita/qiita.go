package qiita

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"
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

}
