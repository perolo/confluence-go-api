package goconfluence

import (
	"net/url"
	"strconv"
)

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

// getSpaceEndpoint creates the correct api endpoint
func (a *API) getSpaceEndpoint() (*url.URL, error) {
	return url.ParseRequestURI(a.endPoint.String() + "/rest/api/space")
}

// GetAllSpaces queries content using a query parameters
func (a *API) GetAllSpaces(query AllSpacesQuery) (*AllSpaces, error) {
	ep, err := a.getSpaceEndpoint()
	if err != nil {
		return nil, err
	}
	ep.RawQuery = addAllSpacesQueryParams(query).Encode()
	return a.SendAllSpacesRequest(ep, "GET")
}

// addAllSpacesQueryParams adds the defined query parameters
func addAllSpacesQueryParams(query AllSpacesQuery) *url.Values {

	data := url.Values{}
	if query.Limit != 0 {
		data.Set("limit", strconv.Itoa(query.Limit))
	}
	if query.Start != 0 {
		data.Set("start", strconv.Itoa(query.Start))
	}
	return &data
}
