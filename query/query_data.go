package query

import (
	"context"
	"time"

	"github.com/jomei/notionapi"
)

// GetLastMonth TODO: update signature add Filter as parameter
func GetLastMonth(notionClient *notionapi.Client, databaseID string) []notionapi.Page {
	location := time.Now().Location()
	curYear, curMonth, _ := time.Now().Date()
	beginingOfCurMonth := time.Date(curYear, curMonth, 1, 0, 0, 0, 0, location)
	// endOfLastMonth := beginingOfCurMonth.AddDate(0, 0, -1)
	// endOfLastMonthObj := notionapi.Date(endOfLastMonth)
	beginingOfLastMonth := beginingOfCurMonth.AddDate(0, -1, 0)
	beginingOfLastMonthObj := notionapi.Date(beginingOfLastMonth)
	beginingOfCurMonthObj := notionapi.Date(beginingOfCurMonth)
	result, err := notionClient.Database.Query(context.Background(), notionapi.DatabaseID(databaseID), &notionapi.DatabaseQueryRequest{
		Filter: notionapi.AndCompoundFilter{
			// notionapi.PropertyFilter{
			// 	Property: "Date",
			// 	Date: &notionapi.DateFilterCondition{
			// 		PastMonth: &struct{}{},
			// 	},
			// },
			notionapi.PropertyFilter{
				Property: "Date",
				Date: &notionapi.DateFilterCondition{
					OnOrAfter: &beginingOfLastMonthObj,
				},
			},
			notionapi.PropertyFilter{
				Property: "Date",
				Date: &notionapi.DateFilterCondition{
					Before: &beginingOfCurMonthObj,
				},
			},
		},
	})
	if err != nil {
		panic("Error executing query")
	}
	return result.Results
}
