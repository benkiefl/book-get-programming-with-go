// https://golang.org/pkg/net/url/#Parse
package main

import (
	"fmt"
	"net/url"
	"strings"
)

type websites struct {
	name string
	URL  string
}

// storing MyErrors in a []errors slice
type MyErrors []error

// implementing the Error interface for MyErrors
func (me MyErrors) Error() string {
	var s []string
	for _, err := range me {
		s = append(s, err.Error())
	}
	return strings.Join(s, ", ")
}

func main() {

	sites := []websites{
		{name: "Yandex", URL: "https://www.yandex.com"},
		{name: "Google", URL: "https://www.google com"},
		{name: "DuckDuckGo", URL: "https://www.duckduck}go.com"},
		{name: "Bing", URL: "www.bing.com"},
		{name: "Shodan", URL: "https://www.shodan.|o/"},
	}

	err := check_sites(sites)

	if err != nil {
		if errors, ok := err.(MyErrors); ok {
			fmt.Printf("Errors encountered: \n")
			for _, e := range errors {
				// type assertion!
				url_err, _ := e.(*url.Error)
				fmt.Printf("- URL %v has an %v.\n", url_err.URL, url_err.Err)
			}
		}
	}

}

func check_sites(s []websites) error {
	var errs MyErrors
	for i := range s {
		_, err := url.Parse(s[i].URL)
		if err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		return errs
	}
	return nil
}
