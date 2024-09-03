package scraper

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/Laurohms/go-scraper/internal/models"
	"github.com/Laurohms/go-scraper/internal/utils"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"github.com/joho/godotenv"
)

const (
	URL         = "URL"
	OUTPUT_FILE = "postings.json"
)

func Start() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println(".env successfully loaded")

	url := os.Getenv(URL)
	if url == "" {
		log.Fatal("URL from .env is empty")
	}
	fmt.Println("URL loaded:", url)

	c := colly.NewCollector()
	c.SetRequestTimeout(60 * time.Second)

	var postings []models.Postings

	c.OnHTML("#form1 > div.principalImpressao > div.conteudoImpressao > div", func(h *colly.HTMLElement) {
		log.Println("Processing HTML element...")

		h.DOM.Find(".bloco").Each(func(_ int, bloco *goquery.Selection) {
			subAccount := bloco.Find("table > thead > tr:nth-child(1) > th").Text()
			log.Println("SubAccount found:", subAccount)

			switch strings.TrimSpace(subAccount) {
			case "PRINCIPAL", "OBRAS", "FDO RESERVA", "CREDITO ESPECIAL", "ESPECIAL", "SALAO DE FESTAS":
				bloco.Find("tbody > tr").Each(func(i int, row *goquery.Selection) {
					dateStr := row.Find("td").Eq(0).Children().Text()
					description := row.Find("td").Eq(1).Children().Text()
					subdivision := row.Find("td").Eq(2).Children().Text()
					valueStr := row.Find("td").Eq(3).Children().Text()
					balanceStr := row.Find("td").Eq(4).Children().Text()

					date, err := time.Parse("02/01/2006", strings.TrimSpace(dateStr))
					if err != nil {
						log.Println("date str to time parse fail:", err)
					}
					value, err := utils.StrToFloat(valueStr)
					if err != nil {
						log.Printf("fail parsing str to float for value: %v", err)
					}
					balance, err := utils.StrToFloat(balanceStr)
					if err != nil {
						log.Printf("fail parsing str to float for balance: %v", err)
					}

					posting := models.Postings{
						Date:        date,
						Description: strings.TrimSpace(description),
						Subdivision: strings.TrimSpace(subdivision),
						Value:       value,
						Balance:     balance,
					}
					postings = append(postings, posting)
				})
			default:
				fmt.Println("SubAccount not recognized:", subAccount)
			}

		})
	})

	c.OnScraped(func(r *colly.Response) {
		log.Println("start creating json file")

		jsonData, err := json.MarshalIndent(postings, "", "    ")
		if err != nil {
			log.Fatalf("error marshaling json: %v", err)
		}

		file, err := os.Create(OUTPUT_FILE)
		if err != nil {
			log.Fatalf("error creating json file: %v", err)
		}
		defer file.Close()

		_, err = file.Write(jsonData)
		if err != nil {
			log.Fatalf("Error writing to file: %v", err)
		}

		fmt.Println("Data saved to", OUTPUT_FILE)

	})

	err = c.Visit(url)
	if err != nil {
		log.Fatal("Failed to visit URL:", err)
	}

}
