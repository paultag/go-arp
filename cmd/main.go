package main

import (
	"fmt"

	"pault.ag/go/arp"
	"pault.ag/go/maclookup"
)

func main() {
	arpTable, err := arp.LoadARPTable()
	if err != nil {
		panic(err)
	}
	for _, entry := range arpTable {
		vendor, err := maclookup.Lookup(entry.MAC)
		if err != nil {
			panic(err)
		}
		fmt.Printf(
			"%s %s %s\n",
			entry.IP,
			entry.MAC,
			vendor.Name,
		)
	}
}
