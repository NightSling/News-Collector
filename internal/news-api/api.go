package newsapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

/*
	{
		"status": "ok",
	    "totalResults": 124079,
	    "articles": [
	        {
	            "source": {
	                "id": "business-insider",
	                "name": "Business Insider"
	            },
	            "author": "George Glover",
	            "title": "Eli Lilly's Wegovy rival Zepbound will help make weight-loss drug market worth $130 billion by 2030, Goldman Sachs says",
	            "description": "The anti-obesity drug market will be worth $130 billion by 2030, Goldman Sachs analysts said in a research note.",
	            "url": "https://www.businessinsider.com/eli-lilly-zepbound-ozempic-wegovy-novo-nordisk-weight-loss-drugs-2024-6",
	            "urlToImage": "https://i.insider.com/6659b30fcc442a2f676c9c6b?width=1200&format=jpeg",
	            "publishedAt": "2024-06-02T04:02:01Z",
	            "content": "Zepbound is Eli Lilly's new weight-loss drug.Brendan McDermid/Reuters\r\n<ul><li>The anti-obesity drug market will be worth $130 billion by 2030, Goldman Sachs said.</li><li>New products such as Eli Li… [+1764 chars]"
	        }
	    ]
	}
*/
type NewsResponse struct {
	Status       string `json:"status"`
	TotalResults int    `json:"totalResults"`
	Articles     []struct {
		Source struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"source"`
		Author      string `json:"author"`
		Title       string `json:"title"`
		Description string `json:"description"`
		URL         string `json:"url"`
		URLToImage  string `json:"urlToImage"`
		PublishedAt string `json:"publishedAt"`
		Content     string `json:"content"`
	} `json:"articles"`
}

const URL = "https://newsapi.org/v2/everything"

func FetchNews(api_key string, sources []string, page int) (*NewsResponse, error) {
	// This function fetches news from the NewsAPI.org
	req := fmt.Sprintf("%s?apiKey=%s&sources=%s&page=%d", URL, api_key, strings.Join(sources, ","), page)
	resp, err := http.Get(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var news NewsResponse
	if err := json.Unmarshal(body, &news); err != nil {
		return nil, err
	}
	return &news, nil
}
