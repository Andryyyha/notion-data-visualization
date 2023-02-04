package main

import (
	"notion-data-visualization/client"
	"notion-data-visualization/parser"
	"notion-data-visualization/query"
	"notion-data-visualization/renderer"
)

func main() {
	notionClient := client.InitClient("/Users/avyazkov/notion-data-visualization.old")
	queryResult := query.GetLastMonth(notionClient, "28911724604b4fcf81ec1d8caf84023e")
	parsedData := parser.GetPagesIDs(queryResult)
	results := parser.Parse(parsedData, notionClient.Token.String())
	renderer.Render(results)
	// sender.CreatePage(notionClient)
	// fmt.Println(chart)
	//fmt.Println(notionClient)
}
