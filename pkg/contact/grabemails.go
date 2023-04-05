package contact

import (
	"encoding/csv"
	"log"
	"net/url"
	"os"
	"regexp"
	"path/filepath"
	"strings"
	"fmt"

	"github.com/gocolly/colly"
	"github.com/mcnijman/go-emailaddress"
	"github.com/wizzybenson/krawler/pkg/contact/types"
)

func GrabEmails(filename, countryCode, lang, query string, maxLength int) {
	sites := GetSerp(query, countryCode, lang, maxLength)
	GrabContacts(sites, filename)
}

func GrabContacts(sites []string, filename string) {
	scraper := colly.NewCollector()

	scraper.DisallowedURLFilters = []*regexp.Regexp{regexp.MustCompile(`facebook|instagram|youtube|twitter|wiki|linkedin|tiktok|tripadvisor`)}
	scraper.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36"
	domainMap := make(map[string]*types.EmailSet)
	bodyString := []byte{}
	validateHost := false
	home := false
	contactPages := []string{}

	scraper.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL)
	})

	scraper.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong: ", err)
	})

	scraper.OnResponse(func(r *colly.Response) {
		bodyString = r.Body
	})

	scraper.OnHTML("body", func(e *colly.HTMLElement) {
		emails := []string{}

		foundEmails := emailaddress.FindWithIcannSuffix(bodyString, validateHost)
		for _, email := range foundEmails {
			emails = append(emails, email.String())
		}
		if domainMap[e.Request.URL.Host] == nil {
			domainMap[e.Request.URL.Host] = &types.EmailSet{}
		}
		domainMap[e.Request.URL.Host].Add(emails)

		if home {
			contactpageRegex := regexp.MustCompile(`about-us|contact|contact-us|nous-contacter|contacter`)
			e.ForEach("a", func(_ int, elem *colly.HTMLElement) {
				if contactpageRegex.MatchString(elem.Attr("href")) {
					base, err := url.Parse(e.Request.URL.Host)
					if err != nil {
						rel, err := base.Parse(elem.Attr("href"))
						if err != nil {
							contactPages = append(contactPages, rel.String())
						}
					}
				}
			})
		}
	})

	scraper.OnScraped(func(r *colly.Response) {
		if len(contactPages) > 0 {
			home = false
			next := contactPages[0]
			contactPages = contactPages[1:]
			scraper.Visit(next)
		}
	})

	for _, site := range sites {
		home = true
		contactPages = []string{}
		scraper.Visit(site)
	}

	saveToCSV(domainMap, filename)
}

func saveToCSV(domainMap map[string]*types.EmailSet, filename string) {
	filename = strings.Trim(filename," ")
	if !strings.HasSuffix(filename, ".csv") {
		filename = filename + ".csv"
	}
	
	filePath, err := os.Executable()
	if err != nil {
	log.Fatal("unable to get the current filename")
}
	dirname := filepath.Dir(filePath)
	path := filepath.Join(dirname, filename)
	file, err := os.Create(path)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)

	// writing the CSV headers
	headers := []string{
		"website",
		"emails",
	}
	writer.Write(headers)

	fmt.Println("Saving emails to ,", file.Name())
	// writing each website contact as a CSV row
	for domain, emailSet := range domainMap {
		// converting a PokemonProduct to an array of strings
		record := []string{
			domain,
			emailSet.ToString(),
		}

		// adding a CSV record to the output file
		writer.Write(record)
	}
	defer writer.Flush()
}