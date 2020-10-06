package main

import (
	"fmt"
	"strconv"
	"strings"
)

// IPAddr type is a 4-bytes array for holding 32 bits of IPv4 address
type IPAddr [4]byte

func (ipaddr IPAddr) String() string {
	var octets []string = make([]string, len(ipaddr))
	for i, octet := range ipaddr {
		octets[i] = strconv.Itoa(int(octet))
	}
	return strings.Join(octets, ".")
}

func runEx5() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
