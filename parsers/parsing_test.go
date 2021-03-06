package parsers_test

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jsking216/nasomi-price-check/parsers"
)

func Test_GetAllBazaarRecordsForItem(t *testing.T) {
	exampleBazaarResponse := "Record Count: 558.<br><br><table border=1><tr><td><strong>Item</td><td><strong>Location</td><td><strong>Seller</td><td><strong>Price</td><td>Quantity</td></tr><tr><td> ginger cookie   </td><td>    West_Ronfaure </td><td>Tonkatough</td><td>50</td><td>2</td></tr><tr><td> ginger cookie   </td><td>    Crawlers_Nest </td><td>Littlearms</td><td>75</td><td>1</td></tr><tr><td> gold key   </td><td>    Bastok_Mines </td><td>Typical</td><td>39999</td><td>1</td></tr><tr><td> gold key   </td><td>    Bastok_Mines </td><td>Thunderingtapir</td><td>40000</td><td>1</td></tr><tr><td> golden Hakutaku eye   </td><td>    Southern_San_dOria </td><td>Alawn</td><td>15000</td><td>1</td></tr><tr><td> grape daifuku   </td><td>    Valkurm_Dunes </td><td>Ardbegislay</td><td>3000</td><td>1</td></tr><tr><td> grape daifuku +1   </td><td>    Valkurm_Dunes </td><td>Ardbegislay</td><td>10000</td><td>1</td></tr><tr><td> great bow +1   </td><td>    Lower_Jeuno </td><td>Veticjeuno</td><td>23000</td><td>1</td></tr><tr><td> guespiere   </td><td>    Lower_Jeuno </td><td>Yaasha</td><td>15900</td><td>1</td></tr><tr><td> Hakutaku eye cluster   </td><td>    Lower_Jeuno </d>Ironballs</td><td>4000</td><td>3</td></tr><tr><td> marinara pizza   </td><td>    West_Ronfaure </td><td>Tonkatough</td><td>4000</td><td>3</td></tr><tr><td> meat mithkabob   </td><td>    Valkurm_Dunes </td><td>Ironballs</td><td>750</td><td>3</td></tr><tr><td> meat mithkabob   </td><td>    West_Ronfaure </td><td>Tonkatough</td><td>750</td><td>4</td></tr><tr><td> meat mithkabob   </td><td>    Residential_Area </td><td>Jrkillah</td><td>1000</td><td>1</td></tr><tr><td> melon pie   </td><td>    Qufim_Island </td><td>Zadrake</td><td>250</td><td>2</td></tr><tr><td> melon pie +1   </td><td>    Valkurm_Dunes </td><td>Scarednewbie</td><td>3000</td><td>1</td></tr><tr><td> mezraq   </td><td>    Lower_Jeuno </td><td>Veticjeuno</td><td>545000</td><td>1</td></tr><tr><td> mezraq   </td><td>    Lower_Jeuno </td><td>Veticjeuno</td><td>600000</td><td>1</td></tr><tr><td> Minerva's ring   </td><td>    Lower_Jeuno </td><td>Moneyplz</td><td>3000000</td><td>1</td></tr><tr><td> Minerva's ring   </td><td>    Dynamis-Valkurm </td><td>Myelrah</td><td>7500000</td><td>1</td></tr><tr><td> Minerva's ring   </td><td>    Lower_Jeuno </td><td>Antzatemywife</td><td>8000000</td><td>1</td></tr><tr><td> minstrel's ring   </td><td>    Crawlers_Nest </td><td>Littlearms</td><td>5200000</td><td>1</td></td><td>1</td></tr><tr><td> piece of akamochi +1   </td><td>    Valkurm_Dunes </td><td>Ardbegislay</td><td>5500</td><td>1</td></tr><tr><td> piece of bubble chocolate   </td><td>    Valkurm_Dunes </td><td>Ironballs</td><td>75</td><td>1</td></tr><tr><td> piece of dogwood lumber   </td><td>    Bastok_Markets </td><td>Gilseller</td><td>20000</td><td>1</td></tr><tr><td> piece of kusamochi   </td><td>    Valkurm_Dunes </td><td>Ardbegislay</td><td>10000</td><td>1</td></tr><tr><td> piece of magnolia lumber   </td><td>    Dynamis-Windurst </td><td>Malag</td><td>50000</td><td>1</td></tr><tr><td> piece of magnolia lumber   </td><td>    Dynamis-Windurst </td><td>Blackhammer</td><td>69999</td><td>1</td></tr><tr><td> piece of oxblood   </td><td>    Lower_Jeuno </td><td>Alphaq</td><td>45000</td><td>1</td></tr><tr><td> pinch of bomb queen ash   </td><td>    Lower_Jeuno </td><td>Almond</td><td>2999</td><td>3</td></tr><tr><td> pinch of bomb queen ash   </td><td>    Port_Windurst </td><td>Ooopsie</td><td>4500</td><td>3</td></tr><tr><td> pinch of bomb queen ash   </td><td>    Valkurm_Dunes </td><td>Ibebe</td><td>5000</td><td>1</td></tr><tr><td> pinch of bomb queen ash   </td><td>    Lower_Jeuno </td><td>Modelo</td><td>5000</td><td>2</td></tr><tr><td> pinch of Valkurm sunsand   </td><td>    Lower_Jeuno </td><td>Yaasha</td><td>900</td><td>1</td></tr><tr><td> plate of crab sushi   </td><td>    Qufim_Island </td><td>Cringyedgelord</td><td>500</td><td>1</td></tr><tr><td> plate of crab sushi   </td><td>    Valkurm_Dunes </td><td>Ironballs</td><td>750</td><td>2</td></tr><tr><td> plate of crab sushi   </td><td>    West_Ronfaure </td><td>Tonkatough</td><td>750</td><td>1</td></tr><tr><td> plate of crab sushi +1   </td><td>    Valkurm_Dunes </td><td>Scarednewbie</td><td>800</td><td>1</td></tr><tr><td> plate of crab sushi +1   </td><td>    Valkurm_Dunes </td><td>Ardbegislay</td><td>950</td><td>1</td></tr><tr><td> plate of crab sushi +1   </td><td>    Valkurm_Dunes </td><td>Ironballs</td><td>1200</td><td>1</td></tr><tr><td> plate of sole sushi   </td><td>    Valkurm_Dunes </td><td>Scarednewbie</td><td>1700</td><td>3</td></tr><tr><td> plate of sole sushi   </td><td>    Qufim_Island </td><td>Cringyedgelord</td><td>1700</td><td>3</td></tr><tr><td> plate of sole sushi   </td><td>    Valkurm_Dunes </td><td>Ardbegislay</td><td>1900</td><td>3</td></tr></table>"
	expectedResult := parsers.BazaarResult{
		BazaarList: []parsers.BazaarItem{
			{
				Item:     "PLATE OF SOLE SUSHI",
				Zone:     "VALKURM_DUNES",
				Player:   "SCAREDNEWBIE",
				Price:    "1700",
				Quantity: "3",
			},
			{
				Item:     "PLATE OF SOLE SUSHI",
				Zone:     "QUFIM_ISLAND",
				Player:   "CRINGYEDGELORD",
				Price:    "1700",
				Quantity: "3",
			},
			{
				Item:     "PLATE OF SOLE SUSHI",
				Zone:     "VALKURM_DUNES",
				Player:   "ARDBEGISLAY",
				Price:    "1900",
				Quantity: "3",
			},
		},
	}

	bazaarResult, parseErr := parsers.GetAllBazaarRecordsForItem(strings.ToUpper("plate of sole sushi"), strings.ToUpper(exampleBazaarResponse))

	if parseErr != nil {
		t.Error("Receved parsing error:", parseErr)
	}
	if !cmp.Equal(bazaarResult, expectedResult) {
		t.Error("Results do not match expectation")
	}
}

func Test_GetCheapestBazaarItemFirst(t *testing.T) {
	exampleBazaar := parsers.BazaarResult{
		BazaarList: []parsers.BazaarItem{
			{
				Item:     "testitem1",
				Zone:     "Big Towm",
				Player:   "Deeznuts",
				Price:    "999",
				Quantity: "1",
			},
			{
				Item:     "testitem1",
				Zone:     "Small Town",
				Player:   "Othernuts",
				Price:    "1000",
				Quantity: "1",
			},
			{
				Item:     "testitem1",
				Zone:     "Small Town",
				Player:   "zzzz",
				Price:    "345345",
				Quantity: "1",
			},
		},
	}

	expectedMinimum := parsers.BazaarItem{
		Item:     "testitem1",
		Zone:     "Big Towm",
		Player:   "Deeznuts",
		Price:    "999",
		Quantity: "1",
	}

	bazaarResult := parsers.GetCheapestBazaarItem(exampleBazaar)
	if !cmp.Equal(bazaarResult, expectedMinimum) {
		t.Error("Results do not match expectation")
	}
}

func Test_GetCheapestBazaarItemLast(t *testing.T) {
	exampleBazaar := parsers.BazaarResult{
		BazaarList: []parsers.BazaarItem{
			{
				Item:     "testitem1",
				Zone:     "Big Towm",
				Player:   "Deeznuts",
				Price:    "999",
				Quantity: "1",
			},
			{
				Item:     "testitem1",
				Zone:     "Small Town",
				Player:   "Othernuts",
				Price:    "1000",
				Quantity: "1",
			},
			{
				Item:     "testitem1",
				Zone:     "Small Town",
				Player:   "zzzz",
				Price:    "111",
				Quantity: "1",
			},
		},
	}

	expectedMinimum := parsers.BazaarItem{
		Item:     "testitem1",
		Zone:     "Small Town",
		Player:   "zzzz",
		Price:    "111",
		Quantity: "1",
	}

	bazaarResult := parsers.GetCheapestBazaarItem(exampleBazaar)
	if !cmp.Equal(bazaarResult, expectedMinimum) {
		t.Error("Results do not match expectation")
	}
}

func Test_GetAllBazaarRecordsPlusOne(t *testing.T) {
	exampleBazaarResponse := "Record Count: 558.<br><br><table border=1><tr><td><strong>Item</td><td><strong>Location</td><td><strong>Seller</td><td><strong>Price</td><td>Quantity</td></tr><tr><td> ginger cookie   </td><td>    West_Ronfaure </td><td>Tonkatough</td><td>50</td><td>2</td></tr><tr><td> ginger cookie   </td><td>    Crawlers_Nest </td><td>Littlearms</td><td>75</td><td>1</td></tr><tr><td> gold key   </td><td>    Bastok_Mines </td><td>Typical</td><td>39999</td><td>1</td></tr><tr><td> gold key   </td><td>    Bastok_Mines </td><td>Thunderingtapir</td><td>40000</td><td>1</td></tr><tr><td> golden Hakutaku eye   </td><td>    Southern_San_dOria </td><td>Alawn</td><td>15000</td><td>1</td></tr><tr><td> grape daifuku   </td><td>    Valkurm_Dunes </td><td>Ardbegislay</td><td>3000</td><td>1</td></tr><tr><td> grape daifuku +1   </td><td>    Valkurm_Dunes </td><td>Ardbegislay</td><td>10000</td><td>1</td></tr><tr><td> great bow +1   </td><td>    Lower_Jeuno </td><td>Veticjeuno</td><td>23000</td><td>1</td></tr><tr><td> guespiere   </td><td>    Lower_Jeuno </td><td>Yaasha</td><td>15900</td><td>1</td></tr><tr><td> Hakutaku eye cluster   </td><td>    Lower_Jeuno </d>Ironballs</td><td>4000</td><td>3</td></tr><tr><td> marinara pizza   </td><td>    West_Ronfaure </td><td>Tonkatough</td><td>4000</td><td>3</td></tr><tr><td> meat mithkabob   </td><td>    Valkurm_Dunes </td><td>Ironballs</td><td>750</td><td>3</td></tr><tr><td> meat mithkabob   </td><td>    West_Ronfaure </td><td>Tonkatough</td><td>750</td><td>4</td></tr><tr><td> meat mithkabob   </td><td>    Residential_Area </td><td>Jrkillah</td><td>1000</td><td>1</td></tr><tr><td> melon pie   </td><td>    Qufim_Island </td><td>Zadrake</td><td>250</td><td>2</td></tr><tr><td> melon pie +1   </td><td>    Valkurm_Dunes </td><td>Scarednewbie</td><td>3000</td><td>1</td></tr><tr><td> mezraq   </td><td>    Lower_Jeuno </td><td>Veticjeuno</td><td>545000</td><td>1</td></tr><tr><td> mezraq   </td><td>    Lower_Jeuno </td><td>Veticjeuno</td><td>600000</td><td>1</td></tr><tr><td> Minerva's ring   </td><td>    Lower_Jeuno </td><td>Moneyplz</td><td>3000000</td><td>1</td></tr><tr><td> Minerva's ring   </td><td>    Dynamis-Valkurm </td><td>Myelrah</td><td>7500000</td><td>1</td></tr><tr><td> Minerva's ring   </td><td>    Lower_Jeuno </td><td>Antzatemywife</td><td>8000000</td><td>1</td></tr><tr><td> minstrel's ring   </td><td>    Crawlers_Nest </td><td>Littlearms</td><td>5200000</td><td>1</td></td><td>1</td></tr><tr><td> piece of akamochi +1   </td><td>    Valkurm_Dunes </td><td>Ardbegislay</td><td>5500</td><td>1</td></tr><tr><td> piece of bubble chocolate   </td><td>    Valkurm_Dunes </td><td>Ironballs</td><td>75</td><td>1</td></tr><tr><td> piece of dogwood lumber   </td><td>    Bastok_Markets </td><td>Gilseller</td><td>20000</td><td>1</td></tr><tr><td> piece of kusamochi   </td><td>    Valkurm_Dunes </td><td>Ardbegislay</td><td>10000</td><td>1</td></tr><tr><td> piece of magnolia lumber   </td><td>    Dynamis-Windurst </td><td>Malag</td><td>50000</td><td>1</td></tr><tr><td> piece of magnolia lumber   </td><td>    Dynamis-Windurst </td><td>Blackhammer</td><td>69999</td><td>1</td></tr><tr><td> piece of oxblood   </td><td>    Lower_Jeuno </td><td>Alphaq</td><td>45000</td><td>1</td></tr><tr><td> pinch of bomb queen ash   </td><td>    Lower_Jeuno </td><td>Almond</td><td>2999</td><td>3</td></tr><tr><td> pinch of bomb queen ash   </td><td>    Port_Windurst </td><td>Ooopsie</td><td>4500</td><td>3</td></tr><tr><td> pinch of bomb queen ash   </td><td>    Valkurm_Dunes </td><td>Ibebe</td><td>5000</td><td>1</td></tr><tr><td> pinch of bomb queen ash   </td><td>    Lower_Jeuno </td><td>Modelo</td><td>5000</td><td>2</td></tr><tr><td> pinch of Valkurm sunsand   </td><td>    Lower_Jeuno </td><td>Yaasha</td><td>900</td><td>1</td></tr><tr><td> plate of crab sushi   </td><td>    Qufim_Island </td><td>Cringyedgelord</td><td>500</td><td>1</td></tr><tr><td> plate of crab sushi   </td><td>    Valkurm_Dunes </td><td>Ironballs</td><td>750</td><td>2</td></tr><tr><td> plate of crab sushi   </td><td>    West_Ronfaure </td><td>Tonkatough</td><td>750</td><td>1</td></tr><tr><td> plate of crab sushi +1   </td><td>    Valkurm_Dunes </td><td>Scarednewbie</td><td>800</td><td>1</td></tr><tr><td> plate of crab sushi +1   </td><td>    Valkurm_Dunes </td><td>Ardbegislay</td><td>950</td><td>1</td></tr><tr><td> plate of crab sushi +1   </td><td>    Valkurm_Dunes </td><td>Ironballs</td><td>1200</td><td>1</td></tr><tr><td> plate of sole sushi   </td><td>    Valkurm_Dunes </td><td>Scarednewbie</td><td>1700</td><td>3</td></tr><tr><td> plate of sole sushi   </td><td>    Qufim_Island </td><td>Cringyedgelord</td><td>1700</td><td>3</td></tr><tr><td> plate of sole sushi   </td><td>    Valkurm_Dunes </td><td>Ardbegislay</td><td>1900</td><td>3</td></tr><tr><td> plate of sole sushi +1</td><td>    Valkurm_Dunes </td><td>DO_NOT_INCLUDE</td><td>666</td><td>6</td></tr></table>"
	expectedResult := parsers.BazaarResult{
		BazaarList: []parsers.BazaarItem{
			{
				Item:     "PLATE OF SOLE SUSHI",
				Zone:     "VALKURM_DUNES",
				Player:   "SCAREDNEWBIE",
				Price:    "1700",
				Quantity: "3",
			},
			{
				Item:     "PLATE OF SOLE SUSHI",
				Zone:     "QUFIM_ISLAND",
				Player:   "CRINGYEDGELORD",
				Price:    "1700",
				Quantity: "3",
			},
			{
				Item:     "PLATE OF SOLE SUSHI",
				Zone:     "VALKURM_DUNES",
				Player:   "ARDBEGISLAY",
				Price:    "1900",
				Quantity: "3",
			},
		},
	}

	bazaarResult, parseErr := parsers.GetAllBazaarRecordsForItem(strings.ToUpper("plate of sole sushi"), strings.ToUpper(exampleBazaarResponse))

	if parseErr != nil {
		t.Error("Receved parsing error:", parseErr)
	}
	if !cmp.Equal(bazaarResult, expectedResult) {
		t.Error("Results do not match expectation")
	}
}
