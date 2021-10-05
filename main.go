package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

var fileWriter *csv.Writer

func init() {
	f, err := os.OpenFile("./products.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()
	if err != nil {
		log.Fatalln("failed to open file", err)
	}
	fileWriter = csv.NewWriter(f)
	header := []string{"date", "time", "platform", "pin", "manufacturer", "name", "numberOfskus", "categoryname", "price_MRP", "price_SALE", "save_price", "in stock", "volume", "skuUniqueID", "sku_name"}

	err = fileWriter.Write(header)
	if err != nil {
		fmt.Println("Error while writing to the file: ", err)
		return
	}
	fileWriter.Flush()
}

func writeToFile(data [][]string) {
	f, err := os.OpenFile("./products.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()
	if err != nil {
		log.Fatalln("failed to open file", err)
	}
	fileWriter = csv.NewWriter(f)
	err = fileWriter.WriteAll(data)
	if err != nil {
		fmt.Println("Error while writing to the file: ", err)
		return
	}
	fileWriter.Flush()
}
func main() {
	fmt.Println("Starting scraping...")
	start := time.Now()
	categories := getCategories()
	for _, cat := range categories.Catarray {
		limit := 10
		products := getProducts(cat.Seotoken, limit)
		limit = products.Totalrecords
		csvRows := [][]string{}
		products = getProducts(cat.Seotoken, limit)
		for _, prod := range products.Suggestionview {
			for _, sku := range prod.Skus {
				volume := ""
				if len(sku.Defining) > 0 {
					volume = sku.Defining[0].Volume
				}
				available := ""
				if sku.Availabilitytype == "A" {
					available = "In Stock"
				} else {
					available = "Out of Stock"
				}
				pRow := []string{
					fmt.Sprintf("%d-%02d-%02d", start.Year(), start.Month(), start.Day()),
					fmt.Sprintf("%02d:%02d:%02d", start.Hour(), start.Minute(), start.Second()),
					"DMART",
					fmt.Sprint(categories.Defaultzip),
					prod.Manufacturer,
					prod.Name,
					fmt.Sprint(prod.Numberofskus),
					prod.Categoryname,
					sku.PriceMrp,
					sku.PriceSale,
					sku.SavePrice,
					available,
					volume,
					sku.Skuuniqueid,
					sku.Name,
				}
				csvRows = append(csvRows, pRow)
			}
			writeToFile(csvRows)
		}
	}
	fmt.Println("Scrapping completed in ", time.Since(start), "!!!")
}

func getProducts(category string, limit int) (products Products) {

	url := "https://digital.dmart.in/api/v1/plp/" + category + "?page=1&size=" + fmt.Sprint(limit)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "text/plain")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	products = Products{}
	err = json.Unmarshal(body, &products)
	return
}

func getCategories() (cats Categories) {

	url := "https://digital.dmart.in/api/v1/categories/@top?profile=details"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	resp := Categories{}
	err = json.Unmarshal(body, &resp)
	return resp
}
