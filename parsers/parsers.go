package parsers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

// VendorParse takes a pointer to the item name and retrieves the minumum vendor price and the itemID
func VendorParse(itemPtr *string) (string, string, error) {
	toCapitalized := strings.Title(*itemPtr)
	withUnderscores := strings.ReplaceAll(toCapitalized, " ", "_")
	vendorURL := "https://classicffxi.fandom.com/wiki/" + withUnderscores
	vendorRes, vendorErr := http.Get(vendorURL)
	if vendorErr != nil {
		fmt.Println(vendorErr)
	}
	vendorData, _ := ioutil.ReadAll(vendorRes.Body)
	defer vendorRes.Body.Close()
	vendorString := string(vendorData)

	itemIDStringMarker := "http://www.edenxi.com/db/items/"
	itemIDStartIndex := strings.Index(vendorString, itemIDStringMarker) + len(itemIDStringMarker)
	forParsingItemID := vendorString[itemIDStartIndex : itemIDStartIndex+20]
	itemID := forParsingItemID[0:strings.Index(forParsingItemID, "\"")]

	priceSubStart := strings.Index(vendorString, "Price:")
	priceSubEnd := strings.Index(vendorString, "Gil")
	if priceSubStart < 0 || priceSubEnd < 0 {
		return itemID, "", errors.New(string(*itemPtr) + " is not sold by a vendor.")
	}
	priceInfo := vendorString[priceSubStart:priceSubEnd]

	// parse price
	price := priceInfo[strings.Index(priceInfo, ">")+2 : strings.Index(priceInfo, "-")]
	return itemID, price, nil
}

// AuctionParse takes an itemID and returns the most recent AH price.
func AuctionParse(itemID string) (string, error) {
	type RecentSummary struct {
		OnStock    string
		Sold15Days string
	}

	type Sale struct {
		ID            string
		Name          string
		ItemID        string
		Price         string
		Stack         string
		Date          string
		Sale          string
		Sell_Date     string
		Time          string
		Buyer         string
		Item_Name     string
		Item_Desc     string
		Name_Singular string
		Name_Plural   string
		StackSize     string
	}

	type AHResponse struct {
		Sale_List []Sale
		Sales     RecentSummary
		Price     string
	}

	ahPayload := strings.NewReader("itemid=" + itemID)
	fmt.Println("Searching AH for ItemID: " + itemID)
	ahRes, err := http.Post(
		"https://na.nasomi.com/auctionhouse/data/ah-data/searchItem.php",
		"application/x-www-form-urlencoded; charset=UTF-8",
		ahPayload,
	)
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(ahRes.Body)
	newStr := buf.String()

	jsonifiedResp := AHResponse{}
	jsonErr := json.Unmarshal([]byte(newStr), &jsonifiedResp)
	if jsonErr != nil {
		return "", jsonErr
	}
	defer ahRes.Body.Close()

	return jsonifiedResp.Sale_List[0].Price, nil
}

// BazaarParse takes a string pointer to the item name and pulls the lowest bazaar price
func BazaarParse(itemPtr *string) (string, error) {
	bazaarURL := "https://na.nasomi.com/status/bazaar.php"
	bazaarRes, bazaarErr := http.Get(bazaarURL)
	if bazaarErr != nil {
		fmt.Println(bazaarErr)
	}
	bazaarData, _ := ioutil.ReadAll(bazaarRes.Body)
	defer bazaarRes.Body.Close()
	bazaarString := string(bazaarData)
	dataUpper, subUpper := strings.ToUpper(bazaarString), strings.ToUpper(*itemPtr)
	if strings.Contains(dataUpper, subUpper) {
		// parse price

		return "999", nil
	}

	return "", errors.New("No bazaar listings found for: " + string(*itemPtr))
}

type BazaarResult struct {
	BazaarList []BazaarItem
}

type BazaarItem struct {
	Item     string
	Zone     string
	Player   string
	Price    string
	Quantity string
}

// GetCheapestBazaarItem returns the lowest priced item from a BazaarResult
func GetCheapestBazaarItem(bazaarItems BazaarResult) BazaarItem {
	itemList := bazaarItems.BazaarList
	minIndex := 0
	for ind, item := range itemList {
		i, _ := strconv.Atoi(item.Price)
		itemListMin, _ := strconv.Atoi(itemList[minIndex].Price)
		if i < itemListMin {
			minIndex = ind
		}
	}
	return itemList[minIndex]
}

// GetAllBazaarRecordsForItem does the string cleaning and returns only the rows which contain the desired item
func GetAllBazaarRecordsForItem(name string, bazaarData string) (BazaarResult, error) {
	var numberOfCellsPerRecord = 5
	allCells := strings.Split(bazaarData, "<td>")
	headersRemoved := allCells[6:]
	var trimmedList []string
	var stringsToRemove = []string{
		"<td>",
		"</td>",
		"<tr>",
		"</tr>",
		"</table>",
	}

	var currentRecordWindow = 999
	for _, ch := range headersRemoved {
		if strings.Contains(ch, name) {
			currentRecordWindow = 0
			trimmedList = append(trimmedList, strings.TrimSpace(removeStringFromString(ch, stringsToRemove)))
			currentRecordWindow++
			continue
		}

		if currentRecordWindow < numberOfCellsPerRecord {
			trimmedList = append(trimmedList, strings.TrimSpace(removeStringFromString(ch, stringsToRemove)))
		}
		currentRecordWindow++
	}

	if len(trimmedList) < 1 {
		return BazaarResult{}, errors.New("Item not found in bazaar data")
	}

	var bazaarList []BazaarItem
	fmt.Println(len(trimmedList) / numberOfCellsPerRecord)
	for i := 0; i < len(trimmedList); i += 5 {
		bazaarList = append(bazaarList, BazaarItem{
			Item:     trimmedList[i],
			Zone:     trimmedList[i+1],
			Player:   trimmedList[i+2],
			Price:    trimmedList[i+3],
			Quantity: trimmedList[i+4],
		})
	}
	response := BazaarResult{
		BazaarList: bazaarList,
	}
	return response, nil
}

func removeStringFromString(source string, toRemove []string) string {
	var finalString = source
	for _, remove := range toRemove {
		finalString = strings.ReplaceAll(finalString, remove, "")
	}
	return finalString
}
