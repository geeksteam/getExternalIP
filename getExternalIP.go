// Package getExternalIP getting IP of the service by myexternalip.com
package getExternalIP

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

const getIPURL = "http://ipinfo.io/json"

// IPstruct ip address info structure
type IPstruct struct {
	IP      string `json:"ip"`
	City    string `json:"city"`
	Region  string `json:"region"`
	Country string `json:"country"`
	Loc     string `json:"loc"`
	Org     string `json:"org"`
}

// GetIP getting external IP from some service
func GetIP() string {
	resp, err := http.Get(getIPURL)
	if err != nil {
		fmt.Println("Error during get external IP info from:" + getIPURL)
		os.Exit(1)
	}
	defer resp.Body.Close()

	var newIP IPstruct
	json.NewDecoder(resp.Body).Decode(&newIP)

	return newIP.IP
}

// GetIPinfo getting external IP full info from some service
func GetIPinfo() IPstruct {
	resp, err := http.Get(getIPURL)
	if err != nil {
		fmt.Println("Error during get external IP info from:" + getIPURL)
		os.Exit(1)
	}
	defer resp.Body.Close()

	var newIP IPstruct
	json.NewDecoder(resp.Body).Decode(&newIP)

	return newIP
}
