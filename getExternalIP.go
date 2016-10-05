// GetExternalIP getting IP of the service by myexternalip.com
package GetExternalIP

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// GetIP getting external IP from some service
func GetIP() string {
	GetIPURL := "http://myexternalip.com/raw"
	resp, err := http.Get(GetIPURL)
	if err != nil {
		fmt.Println("Error during get external IP from:" + GetIPURL)
		os.Exit(1)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return strings.TrimSpace(string(body))
}
