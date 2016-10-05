# getExternalIP
Get external IP from myexternalip.com.

Example usage:
```go
package main

import (
	"fmt"

	"github.com/geeksteam/getExternalIP"
)

// GetIP getting external IP from some service
func main() {
	fmt.Println("Your external IP is:'" + getExternalIP.GetIP() + "'")
}

```
