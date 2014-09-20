package main

import (
	"flag"
	"strings"
)

var target string

func init() {
	flag.StringVar(&target, "subdomain", "", "The full subdomain to remove from our records")
}

func getParts() (subdomain string, domain string) {

	parts := strings.Split(target, ".")

	domain = strings.Join(parts[len(parts)-2:], ".")
	subdomain = strings.Join(parts[:len(parts)-2], ".")

	return
}

func main() {
	flag.Parse()

	subdomain, domain := getParts()

	deleteCfSubdomain(subdomain, domain)
	deleteSqlSubdomain(subdomain, domain)
}
