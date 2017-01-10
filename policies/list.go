package policies

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/ghchinoy/rwctl/cm"
	"github.com/ghchinoy/rwctl/control"
)

const (
	// PoliciesGetURI is the policies endpoint defined here http://docs.akana.com/cm/api/policies/m_policies_getPolicies.htm
	PoliciesGetURI = "/api/policies"
)

// PolicyType is a type of policy
type PolicyType int

const (
	// Operational represents an Operational Policy type
	Operational PolicyType = 1 + iota
	// Denial represents a Denial of Service Policy type
	Denial
	// Compliance represents a Compliance Policy type
	Compliance
	// SLA represents a Service Level Policy type
	SLA
)

func sliceContains(a []string, t string) bool {
	var found bool
	for _, v := range a {
		if v == t {
			found = true
		}
	}

	return found

}

// ListPolicies outputs a list of policies and their IDs
// should show a more human readable output
func ListPolicies(types string, config control.Configuration, debug bool) error {
	if debug {
		log.Println("Listing Policies of type:", types)
	}

	client, _, err := control.LoginToCM(config, debug)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	policyTypes := []string{"Operational Policy", "Denial of Service", "Compliance Policy", "Service Level Policy"}

	if sliceContains(policyTypes, types) {
		policyTypes = []string{types}
	}

	for _, policyType := range policyTypes {
		if debug {
			log.Printf("%s\n", policyType)
		}
		url := config.URL + PoliciesGetURI + "?Type=" + url.QueryEscape(policyType)
		//log.Printf("* %s\n", url)

		//client := &http.Client{}
		req, err := http.NewRequest("GET", url, nil)
		req.Header.Add("Accept", "application/json")

		if debug {
			log.Println("Calling", url)
			control.DebugRequestHeader(req)
		}

		resp, err := client.Do(req)

		defer resp.Body.Close()

		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		if debug {
			control.DebugResponseHeader(resp)
		}
		var policies cm.ApisResponse
		err = json.Unmarshal(bodyBytes, &policies)
		if debug {
			log.Println("Found", len(policies.Channel.Items), " policies.")
		}

		fmt.Printf("%v %s Policies.\n", len(policies.Channel.Items), policyType)
		fmt.Println("---------------------------------")
		pattern := "%-45s %s\n"
		fmt.Printf(pattern, "ID", "Title")

		if len(policies.Channel.Items) > 1 {
			//log.Printf("%s", bodyBytes)
			for _, v := range policies.Channel.Items {
				fmt.Printf(pattern, v.Guid.Value, v.Title)
			}
		}

	}

	return nil
}
