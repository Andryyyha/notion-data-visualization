package parser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/jomei/notionapi"
)

func GetPagesIDs(pages []notionapi.Page) []string {
	ids := make([]string, 0)
	// t := make([]byte, 0)
	for _, e := range pages {
		ids = append(ids, e.ID.String())
	}

	return ids
}

func Parse(ids []string, notionToken string) map[string]float64 {
	c := &http.Client{}
	results := make(map[string]float64)
	for _, p := range ids {
		req, err := http.NewRequest("GET", fmt.Sprintf("https://api.notion.com/v1/pages/%s", p), nil)
		if err != nil {
			fmt.Println(err.Error())
		}
		req.Header.Add("Notion-Version", "2022-06-28")
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", notionToken))
		response, err := c.Do(req)
		if err != nil {
			fmt.Println(err.Error())
		}
		page := make(map[string]interface{})
		responseBodyArray, _ := ioutil.ReadAll(response.Body)
		_ = json.Unmarshal(responseBodyArray, &page)
		properties := page["properties"].(map[string]interface{})
		agg := properties["Aggregation"].(map[string]interface{})["formula"].(map[string]interface{})["number"].(float64)
		category := properties["Category"].(map[string]interface{})["select"].(map[string]interface{})["name"].(string)
		results[category] = results[category] + agg
	}
	return results
}
