// ghostdb
package main

import (
	"fmt"
	"github.com/balanceit/ghostdb/cli"
)

func main() {
	fmt.Printf("calling Get() = %v\n", cli.Get())
	fmt.Printf("calling Put() = %v\n", cli.Put())
	fmt.Printf("calling Post() = %v\n", cli.Post())
	fmt.Printf("calling Delete() = %v\n", cli.Delete())
}
