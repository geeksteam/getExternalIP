// Package getExternalIP getting IP of the service by myexternalip.com
package getExternalIP

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const getIPURL = "http://myexternalip.com/raw"

// GetIP getting external IP from some service
func GetIP() string {
	resp, err := http.Get(getIPURL)
	if err != nil {
		fmt.Println("Error during get external IP from:" + getIPURL)
		os.Exit(1)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return strings.TrimSpace(string(body))
}
