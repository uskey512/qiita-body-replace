package models

type PatchItemsIdRequest struct {
	Body         string `json:"body"`
	Coediting    bool   `json:"coediting"`
	GroupURLName string `json:"group_url_name"`
	Private      bool   `json:"private"`
	Tags         []struct {
		Name     string   `json:"name"`
		Versions []string `json:"versions"`
	} `json:"tags"`
	Title string `json:"title"`
}
