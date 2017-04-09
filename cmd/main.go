package main

import (
	"fmt"

	"pault.ag/go/maclookup"
	"pault.ag/go/proc-arp"
)

func main() {
	arpTable, err := proc_arp.LoadARPTable()
	if err != nil {
		panic(err)
	}
	for _, entry := range arpTable {
		vendor := maclookup.Lookup(entry.MAC)
		fmt.Printf(
			"%s %s %s\n",
			entry.IP,
			entry.MAC,
			vendor.Name,
		)
	}
}
