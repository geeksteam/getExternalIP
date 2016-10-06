# getExternalIP
Get external IP from myexternalip.com.

Useful to get external IP for **Google Cloud Engine** or **AWS** instances.

Using http://ipinfo.io web service.

Example usage:
```go
package main

import (
	"fmt"

	"github.com/geeksteam/getExternalIP"
)

// GetIP getting external IP from web service
func main() {
	// Get only IP of your instance
	fmt.Println("Your external IP is:'" + getExternalIP.GetIP() + "'")
	// Get city of your instance's IP
	fmt.Println("Your IP city is:'" + getExternalIP.GetIPinfo().City + "'")
}

```

Structure of `GetIPinfo()` return:
```go
type IPstruct struct {
	IP      string 
	City    string
	Region  string
	Country string 
	Loc     string 
	Org     string 
}
```
