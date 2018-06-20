package main

import (
	"fmt"
	"log"

	"github.com/cseeger-epages/confluence-go-api"
)

func main() {
	api, err := goconfluence.NewAPI("https://<your-domain>.atlassian.net", "<username>", "<api-token>")
	if err != nil {
		log.Fatal(err)
	}

	// get label information
	labels, err := api.GetLabels("1234567")
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range labels.Labels {
		fmt.Printf("%+v\n", v)
	}

	// add new label
	labels := []goconfluence.Label{
		goconfluence.Label{
			Prefix: "global",
			Name:   "test-label-api",
		},
	}

	lres, err := api.AddLabels("1234567", &labels)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range lres.Labels {
		fmt.Printf("%+v\n", v)
	}

	// remove label
	_, err = api.DeleteLabel("1234567", "test-label-api")
	if err != nil {
		log.Fatal(err)
	}
}
