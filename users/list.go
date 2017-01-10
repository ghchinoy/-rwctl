package users

import (
	"io/ioutil"
	"fmt"
	"sort"
	"github.com/ryanuber/columnize"
	"github.com/ghchinoy/rwctl/control"

	"net/http"
	"github.com/ghchinoy/rwctl/cm"
	"log"
	"encoding/json"
)

const (
	CMListUsersURI  = "/api/search?sort=asc&sortBy=com.soa.sort.order.title_sort&Federation=false&count=20&start=0&q=type:user"
)

// User is a convenience structure for a CM User
type User struct {
	Name        string
	ProfileName string
	Version     string
	ID          string
	Domain      string
	Email       string
	UserName    string
}

// Users is a collection of API structs
type Users []User

// Len is an implementation of sort interface for length of Users list
func (slice Users) Len() int {
	return len(slice)
}

// Less is an implementation of sort interface for less comparison
func (slice Users) Less(i, j int) bool {
	return slice[i].Name < slice[j].Name
}

// Swap is an implementation of the sort interface swap function
func (slice Users) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

// ListUsers lists users on platform
func ListUsers(config control.Configuration, debug bool) error {
	if debug {
		log.Println("Listing Users")
	}

	client, _, err := control.LoginToCM(config, debug)
	if err != nil {
		return err
	}

	url := config.URL + CMListUsersURI

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Accept", "application/json")
	if debug {
		log.Println("curl command:", control.CURLThis(client, req))
	}
	resp, err := client.Do(req)
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var apis cm.ApisResponse
	err = json.Unmarshal(bodyBytes, &apis)
	if debug {
		log.Printf("Found %v Users", len(apis.Channel.Items))
		fmt.Printf("%s", bodyBytes)
	}
	var userList Users

	for _, v := range apis.Channel.Items {
		userList = append(userList, User{
			ProfileName: v.Title, Name: v.Description, Domain: v.Domain, ID: v.Guid.Value,
			UserName: v.UserName, Email: v.Email,
		})
	}
	sort.Sort(userList)
	fmt.Printf("%v Users\n", len(userList))
	var data []string
	data = append(data, "Name | Email | UserName | Domain | ID")
	for _, v := range userList {
		data = append(data, fmt.Sprintf("%s | %s | %s | %s | %s", v.Name, v.Email, v.UserName, v.Domain, v.ID))
		//fmt.Printf("%-28s %-29s %s @ %s\n", v.Name, v.Email, v.UserName, v.Domain)
	}
	result := columnize.SimpleFormat(data)
	fmt.Println(result)
	//fmt.Printf("%s", bodyBytes)

	return nil
}