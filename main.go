package main

import (
	"flag"
	"fmt"

	"github.com/jsking216/nasomi-price-check/parsers"
)

func main() {
	start := time.Now()
	itemPtr := flag.String("item", "", "ffxi item name")
	flag.Parse()

	bazaarCh := make(chan BazaarParseResult)
	go parsers.BazaarParse(itemPtr, bazaarCh)
	
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

	if bazaarCh.Error) != nil {
		fmt.Println(bazaarCh.Error)
	} else {
		fmt.Println("Bazaar Price: " + bazaarCh.Price)
	}

	if vendorErr != nil && ahErr != nil && bazaarCh.Error) != nil {
		fmt.Println("ERROR: All price requests failed -- are you sure the item name is correct?")
	}
	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}
