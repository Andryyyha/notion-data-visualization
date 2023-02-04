package sender

import (
	"context"
	"fmt"
	"notion-data-visualization/client"
	"time"

	"github.com/jomei/notionapi"
)

// TODO: parent type page_id doesn't work in notion API

func CreatePage(notionClient *notionapi.Client) {
	now := time.Now()
	previousMonth := now.AddDate(0, -1, 0)
	year := previousMonth.Year()
	month := previousMonth.Month().String()
	pageID := notionapi.PageID(client.GetBudgetPageID("/Users/avyazkov/notion-data-visualization.old"))
	p, err := notionClient.Page.Create(context.Background(), &notionapi.PageCreateRequest{
		Parent: notionapi.Parent{
			Type:   notionapi.ParentTypePageID,
			PageID: pageID,
		},
		Properties: notionapi.Properties{
			"Name": notionapi.TitleProperty{
				Title: []notionapi.RichText{
					{Text: &notionapi.Text{Content: fmt.Sprintf("%s %d", month, year)}},
				},
			},
		},
	})
	fmt.Println(p, err)
}
