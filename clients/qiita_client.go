package clients

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
	"github.com/uskey512/qiita-url-replace/models"
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

func (q *QiitaClient) GetAuthenticatedUser() (*models.GetAuthenticatedUserResponse, error) {
	url := q.endpointBase + "authenticated_user"
	raw, err := q.executeRequest(http.MethodGet, url)
	if err != nil {
		return nil, err
	}

	data := new(models.GetAuthenticatedUserResponse)
	if err := json.Unmarshal(raw, &data); err != nil {
		return nil, errors.Wrap(err, "Json Unmarshal error ")
	}

	return data, nil
}

func (q *QiitaClient) GetAuthenticatedUserItems(page, perPage int) (*models.GetAuthenticatedUserItemsResponse, error) {
	url := q.endpointBase + "authenticated_user/items?page=1&per_page=100"
	raw, err := q.executeRequest(http.MethodGet, url)
	if err != nil {
		return nil, err
	}

	data := new(models.GetAuthenticatedUserItemsResponse)
	if err := json.Unmarshal(raw, &data); err != nil {
		return nil, errors.Wrap(err, "Json Unmarshal error ")
	}

	return data, nil
}

func (q *QiitaClient) executeRequest(method, url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "http.NewRequest fail")
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", q.authToken))

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("%s %s fail", method, url))
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
