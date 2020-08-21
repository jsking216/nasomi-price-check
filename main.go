package main

import (
	"flag"
	"net/http"
	"fmt"
	"strings"
	"io/ioutil"
)

func main() {
	itemPtr := flag.String("item", "", "ffxi item name")
	flag.Parse()

	ahPayload := strings.NewReader("itemname=" + strings.ReplaceAll(*itemPtr, " ", "+"))
	ahRes, err := http.Post(
		"https://na.nasomi.com/auctionhouse/data/ah-data/searchItemByName.php",
		"text/html; charset=UTF-8",
		ahPayload,
	)
	if err != nil {
		fmt.Println(err)
	}

	bazaarUrl := "https://na.nasomi.com/status/bazaar.php"
	bazaarResp, bazaarErr := http.Get(bazaarUrl)
	if bazaarErr != nil {
		fmt.Println(bazaarErr)
	}
	// use equalFold to find the item if it is in the list

	toCapitalized := strings.Title(*itemPtr)
	withUnderscores := strings.ReplaceAll(toCapitalized, " ", "_")
	vendorUrl := "https://classicffxi.fandom.com/wiki/"  + withUnderscores
	vendorRes, vendorErr := http.Get(vendorUrl)
	if vendorErr != nil {
		fmt.Println(vendorErr)
	}
}