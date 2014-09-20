package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

var cfUrl string = "https://www.cloudflare.com/api_json.html"

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func makeRequest(method string, parameters url.Values) (body []byte, err error) {
	parameters.Set("tkn", getConfig().CloudflareKey)
	parameters.Set("email", getConfig().CloudflareEmail)
	parameters.Set("a", method)

	resp, err := http.PostForm(cfUrl, parameters)
	if err != nil {
		return []byte{}, err
	}

	return ioutil.ReadAll(resp.Body)
}

func loadAllRecords(zone string, offset int) *CfLoadAllResponse {
	v := url.Values{}
	v.Set("z", zone)
	v.Set("o", strconv.Itoa(offset))

	body, err := makeRequest("rec_load_all", v)
	if err != nil {
		log.Fatal(err)
	}

	data := CfLoadAllResponse{}

	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatal(err)
	}

	return &data
}

func deleteCfSubdomain(subdomain string, zone string) {
	searching := []string{subdomain + "." + zone,
		"_minecraft._tcp." + subdomain + "." + zone}

	id := ""
	offset := 0

	for {
		log.Printf("Searching with offset %d in %s\n", offset, zone)
		output := loadAllRecords(zone, offset)

		for _, record := range output.Response.Recs.Objs {
			if stringInSlice(record["name"].(string), searching) {
				id = record["rec_id"].(string)
			}
		}

		if id != "" || !output.Response.Recs.Has_more {
			break
		}
		offset += 180
	}

	if id == "" {
		log.Println("No subdomain found like that!\n")
	}

	v := url.Values{}
	v.Set("z", zone)
	v.Set("id", id)
	makeRequest("rec_delete", v)
	log.Println("Subdomain deleted from CloudFlare.\n")
}
