package quote

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/c0d33ngr/quote-generator/models"
)

// Get random quote
func FetchRandomQuote(tags string) (string, error) {
	baseUrl := "https://api.quotable.io/"
	URL := baseUrl + "random?tags=" + tags

	resp, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
	}

	bd, ee := ioutil.ReadAll(resp.Body)
	if ee != nil {
		log.Fatal(ee)
	}

	defer resp.Body.Close()

	var quote models.QuoteResponse

	err = json.Unmarshal(bd, &quote)
	if err != nil {
		fmt.Println(resp.StatusCode)
		log.Fatal(err)
	}

	return quote.TextOutput(), nil
}
