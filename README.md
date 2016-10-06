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
	fmt.Println("Your external IP is:'" + getExternalIP.GetIP() + "'")
}

```
