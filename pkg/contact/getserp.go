package contact
import (
	"context"
	"time"
	"github.com/rocketlaunchr/google-search"
	"golang.org/x/time/rate"
	"log"
	"fmt"
)

func GetSerp(query, countrycode, langcode string, maxLenght int) []string {
		//Get Google SERP urls
		fmt.Println("Getting Google SERP urls...")
		start, limit := 0, 10
		sitesSlice := []string{}
		for start < maxLenght {
			RateLimit := rate.NewLimiter(rate.Every(time.Second), 1)
			googlesearch.RateLimit = RateLimit
			options := googlesearch.SearchOptions{
				CountryCode:countrycode,LanguageCode:langcode,FollowNextPage:true,Limit: limit,Start: start,
				UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36",
			}
			sitesserp, err := googlesearch.Search(context.Background(), query, options)
			
			if err != nil {
				log.Fatal(err)
			}
	
			for _, result := range sitesserp {
				sitesSlice = append(sitesSlice, result.URL)
			}
			if len(sitesserp) > 0 {
				end := len(sitesserp)-1
				last := sitesserp[end]
				limit = last.Rank + 1
			}
	
			start += limit
			time.Sleep(1*time.Second)
		}
		
	return sitesSlice
}