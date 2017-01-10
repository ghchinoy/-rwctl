package cms

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"github.com/ghchinoy/rwctl/cm"
	"github.com/ghchinoy/rwctl/control"
	"log"
	"net/http"
	"encoding/json"
)

var debug bool
var config control.Configuration

// ListTopLevelCMS outputs the contents from the root of each CMS path
func ListTopLevelCMS(c control.Configuration, d bool) {
	debug = d
	config = c
	listTopLevelCMS()
}

// ListCMSPath lists the contents of a CMS path
func ListCMSPath(c control.Configuration, path string, d bool) {
	debug = d
	config = c
	listCMS(path, 0)
}


// listTopLevelCMS is a convenience method to call listCMS for /content and /resources
func listTopLevelCMS() {
	listCMS("/content", 0)
	listCMS("/resources", 0)
}

// listCMS lists the contents of a CMS path
func listCMS(path string, depth int) (int, int, error) {

	var dircount, filecount int

	if depth == 0 {
		fmt.Println(path)
	}

	cms, err := getCMSPath(path)
	if err != nil {
		//log.Println("An error in getting path: ", err)
		// TODO determine what to do with an error that happens with "/content/api"
		return dircount, filecount, err
	}

	// Output content path
	var sep string

	var partialbuffer bytes.Buffer

	for k, v := range cms.Channel.Items {
		// separator determination
		sep = "├──"
		if k == len(cms.Channel.Items)-1 {
			sep = "└──"
		}
		// indent determination
		var indent string
		if depth > 0 {
			indent = " "
			for i := 0; i < depth; i++ {
				indent = indent + indent
			}
			sep = fmt.Sprintf("|%s%s", indent, sep)
		}

		partialbuffer.WriteString(fmt.Sprintf("%s %s\n", sep, v.Title))
		fmt.Printf("%s %s\n", sep, v.Title)

		if v.Category[0].Value == "folder" {
			dircount++

			descend := path + "/" + v.Title
			depth++
			d, f, err := listCMS(descend, depth)
			if err != nil {
				// eat the error for rendering purposes
				//log.Printf("Unable to retrieve content of path %s. %s", path, err)
				//return dircount, filecount, err
			}
			dircount = dircount + d
			filecount = filecount + f
			depth--

		} else {
			filecount++
		}
	}
	//fmt.Print(partialbuffer.String())

	if depth == 0 {
		fmt.Printf("\n%v directories, %v files\n", dircount, filecount)
	}

	return dircount, filecount, nil
}

func getCMSPath(path string) (cm.ApisResponse, error) {

	var cms cm.ApisResponse

	if debug {
		log.Println("Getting content of: ", path)
	}
	// GET content path
	client, _, err := control.LoginToCM(config, debug)
	if err != nil {
		log.Fatalln(err)
		return cms, err
	}

	url := config.URL + path

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Accept", "application/json")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return cms, err
	}
	if debug {
		log.Printf("%s", bodyBytes)
	}

	err = json.Unmarshal(bodyBytes, &cms)
	if err != nil {
		return cms, err
	}

	return cms, nil
}