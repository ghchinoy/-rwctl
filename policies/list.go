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
	"strings"
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
func ListPolicies(types []string, showinactivepolicies bool, outputformat string, config control.Configuration, debug bool) error {
	if debug {
		log.Println("Listing Policies of type:", types)
	}

	policymap := make(map[string]string)
	policymap["o"] = "Operational Policy"
	policymap["d"] = "Denial of Service"
	policymap["s"] = "Service Level Policy"
	policymap["c"] = "Compliance Policy"

	policyTypes := []string{"Operational Policy", "Denial of Service", "Compliance Policy", "Service Level Policy"}


	if types[0] == "all" { // all
		for _, policyType := range policyTypes {

			err := outputPolicyTypeListing(policyType, showinactivepolicies, outputformat, config, debug)
			if err != nil {
				fmt.Println("Unable to retrieve polices of type", policyType, err.Error())
			}
		}
	} else { // only the ones chosen
		for _, t := range types {
			firstletter := strings.Split(strings.ToLower(t), "")[0]
			err := outputPolicyTypeListing(policymap[firstletter], showinactivepolicies, outputformat, config, debug)
			if err != nil {
				fmt.Println("Unable to retrieve polices of type", t, err.Error())
			}
		}

	}



	return nil
}

func outputPolicyTypeListing(policyType string, showinactivepolicies bool, outputformat string, config control.Configuration, debug bool) error {

	client, _, err := control.LoginToCM(config, debug)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	url := config.URL + PoliciesGetURI + "?Type=" + url.QueryEscape(policyType)
	if showinactivepolicies {
		url = url + "IncludeInactivePolicies=true"
	}
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

	inactive := "."
	if showinactivepolicies {
		inactive = " (Inactive)."
	}

	if outputformat == "json" {
		jsondata, err := json.Marshal(policies.Channel.Items)
		if err != nil { // this would be odd to get here, as this is just remarshalling the API response
			fmt.Println("Can't convert to json JSON")
		}
		fmt.Printf("%s\n", jsondata)

	}  else {
		fmt.Printf("%v %s Policies%s\n", len(policies.Channel.Items), policyType, inactive)
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