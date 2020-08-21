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

	toCapitalized := strings.Title(*itemPtr)
	withUnderscores := strings.ReplaceAll(toCapitalized, " ", "_")
	vendorUrl := "https://classicffxi.fandom.com/wiki/"  + withUnderscores
	vendorRes, vendorErr := http.Get(vendorUrl)
	if vendorErr != nil {
		fmt.Println(vendorErr)
	}
	vendorData, _ := ioutil.ReadAll(vendorRes.Body)
	defer vendorRes.Body.Close()
	vendorString := string(vendorData)
	priceInfo := vendorString[strings.Index(vendorString, "Price:"):strings.Index(vendorString, "Gil")]
	itemIDStringMarker := "http://www.edenxi.com/db/items/"
	itemIDStartIndex := strings.Index(vendorString, itemIDStringMarker) + len(itemIDStringMarker)
	forParsingItemID := vendorString[itemIDStartIndex:itemIDStartIndex+20]
	fmt.Println(priceInfo)
	fmt.Println(forParsingItemID)


	ahPayload := strings.NewReader("itemid=17396")
	ahRes, err := http.Post(
		"https://na.nasomi.com/auctionhouse/data/ah-data/searchItem.php",
		"text/html; charset=UTF-8",
		ahPayload,
	)
	if err != nil {
		fmt.Println(err)
	}
	ahData, _ := ioutil.ReadAll(ahRes.Body)
	defer ahRes.Body.Close()

	fmt.Printf("%s\n", ahData)

	bazaarUrl := "https://na.nasomi.com/status/bazaar.php"
	bazaarRes, bazaarErr := http.Get(bazaarUrl)
	if bazaarErr != nil {
		fmt.Println(bazaarErr)
	}
	bazaarData, _ := ioutil.ReadAll(bazaarRes.Body)
	defer bazaarRes.Body.Close()
	bazaarString := string(bazaarData)
	dataUpper, subUpper := strings.ToUpper(bazaarString), strings.ToUpper(*itemPtr)
	if (strings.Contains(dataUpper, subUpper)) {
		// parse price
	} else {
		fmt.Println("No bazaar listings found for: " + string(*itemPtr))
	}
}