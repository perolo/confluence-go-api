package goconfluence

// AllSpacesQuery defines the query parameters
// Query parameter values https://developer.atlassian.com/cloud/confluence/rest/#api-space-get
type Expandable struct {
	Settings    string `json:"settings"`
	Metadata    string `json:"metadata"`
	Operations  string `json:"operations"`
	LookAndFeel string `json:"lookAndFeel"`
	Identifiers string `json:"identifiers"`
	Permissions string `json:"permissions"`
	Icon        string `json:"icon"`
	Description string `json:"description"`
	Theme       string `json:"theme"`
	History     string `json:"history"`
	Homepage    string `json:"homepage"`
}
type RLinks struct {
	Webui string `json:"webui"`
	Self  string `json:"self"`
}
type SLinks struct {
	Base    string `json:"base"`
	Context string `json:"context"`
	Next    string `json:"next"`
	Self    string `json:"self"`
}
type Stype struct {
	ID         int        `json:"id"`
	Key        string     `json:"key"`
	Name       string     `json:"name"`
	Type       string     `json:"type"`
	Status     string     `json:"status"`
	Expandable Expandable `json:"_expandable"`
	Links      RLinks     `json:"_links"`
}
type AllSpacesQuery struct {
	Results []Stype `json:"results"`
	Start   int     `json:"start"`
	Limit   int     `json:"limit"`
	Size    int     `json:"size"`
	Links   SLinks  `json:"_links"`
}

type AllSpacesOption struct {
	Start  int    `url:"start"`
	Limit  int    `url:"limit"`
	Type   string `url:"type"`
	Status string `url:"status"`
}

// getSpaceEndpoint creates the correct api endpoint

// GetAllSpaces queries content using a query parameters
func (a *API) GetAllSpaces(options AllSpacesOption) (*AllSpaces, error) {
	u := a.endPoint.String() + "api/v2/spaces"
	// ep.RawQuery = addAllSpacesQueryParams(query).Encode()
	endpoint, err := addOptions(u, options)
	if err != nil {
		return nil, err
	}
	return a.SendAllSpacesRequest(endpoint, "GET")
}

func (a *API) GetNextSpaces(link string) (*AllSpaces, error) {
	u := a.endPoint.Scheme + "://" + a.endPoint.Hostname() + link
	return a.SendAllSpacesRequest(u, "GET")
}
