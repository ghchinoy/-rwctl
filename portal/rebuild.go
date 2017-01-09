package portal

import (
	"bytes"
	"github.com/ghchinoy/rwctl/control"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"errors"
	"encoding/json"
)

const (
	CMRebuildStyles = "/resources/branding/generatestyles"
)

// RebuildStyles rebuilds the portal's styles
func RebuildStyles(config control.Configuration, theme string, debug bool) error {

	client, _, err := control.LoginToCM(config, debug)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	rebuildStylesURI := config.URL + CMRebuildStyles

	if debug {
		log.Println("Rebuilding styles for theme", theme)
	}
	postdata := url.Values{}
	postdata.Set("theme", theme)

	req, _ := http.NewRequest("POST", rebuildStylesURI, bytes.NewBufferString(postdata.Encode()))
	req.Header.Add("Content-Length", strconv.Itoa(len(postdata.Encode())))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	control.AddCsrfHeader(req, client)
	resp, err := client.Do(req)

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	if resp.StatusCode != 200 {
		if debug {
			log.Println(resp.StatusCode, resp.Status)
		}
		return errors.New("Unable to parse less file. " + resp.Status + " when calling " + rebuildStylesURI)
	}

	var results map[string]interface{}
	err = json.Unmarshal(data, &results)
	status := results["result"]
	log.Printf("Rebuild styles: %s", status)
	return nil

}
