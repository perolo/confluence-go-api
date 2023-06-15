package goconfluence

type AllSpaces struct {
	Results []struct {
		Name        string      `json:"name"`
		Key         string      `json:"key"`
		ID          int         `json:"id"`
		Type        string      `json:"type"`
		HomepageID  int         `json:"homepageId"`
		Icon        interface{} `json:"icon"`
		Status      string      `json:"status"`
		Description interface{} `json:"description"`
	} `json:"results"`
	Links struct {
		Next string `json:"next"`
	} `json:"_links"`
}
