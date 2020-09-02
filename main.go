package main

import (
	"flag"
	"fmt"

	"github.com/jsking216/nasomi-price-check/parsers"
)

func main() {
	itemPtr := flag.String("item", "", "ffxi item name")
	flag.Parse()

	itemID, price, vendorErr := parsers.VendorParse(itemPtr)
	if vendorErr != nil {
		fmt.Println(vendorErr)
	}
	fmt.Println(price)
	fmt.Println(itemID)

	ahPrice, ahErr := parsers.AuctionParse(itemID)
	if ahErr != nil {
		fmt.Println(ahErr)
	}

	bazaarPrice, bazaarErr := parsers.BazaarParse(itemPtr)
	if bazaarErr != nil {
		fmt.Println(bazaarErr)
	}

	fmt.Println(ahPrice)
	fmt.Println(bazaarPrice)
}
