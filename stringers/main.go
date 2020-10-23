package main

import (
	"fmt"
	"strings"
)

type IPAddr [4]byte

func (ip IPAddr) String() string {
	ipVals := make([]string, 4)
	for i, v := range ip {
		ipVals[i] = fmt.Sprintf("%v", v)
	}
	return strings.Join(ipVals, ".")
	// return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3]) // Another solution
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
