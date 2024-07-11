package main

import (
    "fmt"
    "net/http"
    "github.com/PuerkitoBio/goquery"
)

func main() {
    url := "https://g1.globo.com"

    res, err := http.Get(url)
    if err != nil {
        fmt.Print(err)
    }
    doc, err := goquery.NewDocumentFromReader(res.Body)
    if err != nil {
        fmt.Print(err)
    }
    doc.Find(".feed-post-body").Each(func(index int, item *goquery.Selection) {
        title := item.Text()
        fmt.Printf("Artigo #%d: %s\n", index+1, title)
    })
}
