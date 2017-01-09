package portal

import (
	"github.com/ghchinoy/rwctl/control"

	"log"
	"net/http"
)

const (
	CMLandingIndex         = "/content/home/landing/index.htm"
	CMInternationalization = "/i18n"
	CMCustomLess           = "/less/custom.less"
	CMFavicon              = "/style/images/favicon.ico"
)

// ResetTheme restores the portal theme to as close to vanilla as possible
func ResetTheme(config control.Configuration, theme string, debug bool) error {

	if debug {
		log.Println("Resetting theme", theme)
	}

	client, _, err := control.LoginToCM(config, debug)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	urls := []string{
		CMLandingIndex,
		"/resources/theme/" + theme + CMInternationalization,
		"/resources/theme/" + theme + CMCustomLess,
		"/resources/theme/" + theme + CMFavicon,
		"/resources/theme/" + theme + "/SOA",
		"/resources/theme/" + theme + "/less",
		"/resources/theme/" + theme + "/style",
	}

	for _, url := range urls {
		urlStr := config.URL + url
		log.Println("Deleting", url)
		err := callDeleteURL(client, urlStr, debug)
		if err != nil {
			return err
		}
	}

	return nil

}

// Used by resetCM, this is called multiple times to delete
// a specific url
func callDeleteURL(client *http.Client, urlStr string, debug bool) error {

	//client := &http.Client{}
	//client.Jar = jar
	req, err := http.NewRequest("DELETE", urlStr, nil)
	control.AddCsrfHeader(req, client)
	resp, err := client.Do(req)
	defer resp.Body.Close()
	//bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if debug {
		if resp.StatusCode != 200 {
			log.Println("Delete:", resp.Status)
		}
	}

	return nil
}
