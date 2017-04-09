package main

import (
	"fmt"

	"pault.ag/go/arp"
)

func main() {
	arpTable, err := arp.LoadARPTable()
	if err != nil {
		panic(err)
	}
	for _, entry := range arpTable {
		fmt.Printf("%s\n", entry)
	}
}
