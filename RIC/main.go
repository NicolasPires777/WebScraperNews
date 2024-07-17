package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	// "time"

	"github.com/PuerkitoBio/goquery"
	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Artigo struct {
	Noticia   string `bson:"news"`
	Data      string `bson:"date"`
	Categoria string `bson:"categoria"`
}

func main() {

	mongodbconf := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), mongodbconf)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Conectado ao Banco\n\n")

	colecao := client.Database("Nicolas").Collection("news")
	development := true

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
		var article Artigo
		article.Noticia = item.Find("p").First().Text()
		article.Data = item.Find(".feed-post-datetime").First().Text()
		article.Categoria = item.Find(".feed-post-metadata-section").Text()

		if development != true {
			_, err := colecao.InsertOne(context.TODO(), article)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			fmt.Println(article)
		}
	})

	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\n\nDesconectado do MongoDB!")
}
