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
	ahData, _ := ioutil.ReadAll(ahRes.Body)
	ahRes.Body.Close()

	fmt.Println(ahData)

	bazaarUrl := "https://na.nasomi.com/status/bazaar.php"
	bazaarRes, bazaarErr := http.Get(bazaarUrl)
	if bazaarErr != nil {
		fmt.Println(bazaarErr)
	}
	bazaarData, _ := ioutil.ReadAll(bazaarRes.Body)
	bazaarRes.Body.Close()
	fmt.Println(bazaarData)
	// use equalFold to find the item if it is in the list

	toCapitalized := strings.Title(*itemPtr)
	withUnderscores := strings.ReplaceAll(toCapitalized, " ", "_")
	vendorUrl := "https://classicffxi.fandom.com/wiki/"  + withUnderscores
	vendorRes, vendorErr := http.Get(vendorUrl)
	if vendorErr != nil {
		fmt.Println(vendorErr)
	}
	vendorData, _ := ioutil.ReadAll(vendorRes.Body)
	vendorRes.Body.Close()
	fmt.Println(vendorData)
}