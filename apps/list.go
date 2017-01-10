package apps

import (
	"github.com/ghchinoy/rwctl/control"
	"io/ioutil"
	"strings"
	"sort"
	"fmt"
	"log"
	"net/http"
	"github.com/ghchinoy/rwctl/cm"
	"encoding/json"
)

const (
	CMListAppsURI   = "/api/search?sortBy=com.soa.sort.order.alphabetical&count=20&start=0&q=type:app"
)

// App is a convenience structure for a CM App
type App struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	ID          string `json:"id"`
	Visibility  string `json:"visibility"`
	Connections int    `json:"connections"`
	Followers   int    `json:"followers"`
	Rating      float32
}

// Apps is a collection of API structs
type Apps []App

// Len is an implementation of sort interface for length of Apps
func (slice Apps) Len() int {
	return len(slice)
}

// Less is an implementation of sort interface for less comparison
func (slice Apps) Less(i, j int) bool {
	return slice[i].Name < slice[j].Name
}

// Swap is an implementation of the sort interface swap function
func (slice Apps) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}


func ListApps(config control.Configuration, debug bool) error {

	if debug {
		log.Println("Listing Apps")
	}

	client, userinfo, err := control.LoginToCM(config, debug)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	url := config.URL + CMListAppsURI

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Accept", "application/json")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if debug {
		log.Printf("%s", bodyBytes)
	}
	var apps cm.ApisResponse
	err = json.Unmarshal(bodyBytes, &apps)
	if debug {
		log.Printf("Found %v Apps", len(apps.Channel.Items))
	}

	var appList Apps

	domainsuffix := strings.Split(userinfo.LoginDomainID, ".")[1]

	for _, v := range apps.Channel.Items {
		var visibility string
		cats := v.Category
		for _, c := range cats {
			if c.Domain == "uddi:soa.com:visibility" {
				visibility = c.Value
			}
		}
		// Shorten Registered Users visibility
		if visibility == "com.soa.visibility.registered.users" {
			visibility = "Registered Users"
		}
		// Remove domain suffix from App GUID
		appguid := strings.Replace(v.Guid.Value, "."+domainsuffix, "", -1)

		appList = append(appList, App{
			Name:        v.Title,
			ID:          appguid,
			Visibility:  visibility,
			Connections: v.Connections,
			Followers:   v.Followers,
			Rating:      v.Rating,
		})
	}
	sort.Sort(appList)
	fmt.Printf("%v apps (suffix: %s)\n", len(appList), domainsuffix)
	// TODO get max length of []App fields and dynamically set the format pattern
	pattern := "%-36s %-20s %-8s %-3v %-3v %-3v\n"
	fmt.Printf(pattern, "ID", "Name", "Vis", "Con", "Fol", "Rat")
	for _, v := range appList {
		fmt.Printf(pattern, v.ID, v.Name, v.Visibility, v.Connections, v.Followers, v.Rating)
	}

	return nil
}