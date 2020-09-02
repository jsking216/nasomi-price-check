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
	} else {
		fmt.Println("Vendor Price: " + price)
	}

	ahPrice, ahErr := parsers.AuctionParse(itemID)
	if ahErr != nil {
		fmt.Println(ahErr)
	} else {
		fmt.Println("Auction House Price: " + ahPrice)
	}

	bazaarPrice, bazaarErr := parsers.BazaarParse(itemPtr)
	if bazaarErr != nil {
		fmt.Println(bazaarErr)
	} else {
		fmt.Println("Bazaar Price: " + bazaarPrice)
	}

	if vendorErr != nil && ahErr != nil && bazaarErr != nil {
		fmt.Println("ERROR: All price requests failed -- are you sure the item name is correct?")
	}

}
