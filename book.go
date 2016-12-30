package main

import (
	"log"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
	"github.com/labstack/echo"
)

type (
	Book struct {
		Name        string `json:"name"`
		Image       string `json:"image"`
		Link        string `json:"link"`
		Free        bool   `json:"free"`
		Description string `json:"description"`
		Comment     string `json:"comment"`
	}
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		book := scrape(c.QueryParam("item"))
		return c.JSON(http.StatusOK, book)
	})
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}

func scrape(uri string) Book {
	doc, err := goquery.NewDocument(uri)
	if err != nil {
		log.Fatal(err)
	}

	book := Book{
		Name:        title(doc),
		Image:       image(doc),
		Link:        link(uri),
		Free:        free(doc),
		Description: "",
		Comment:     "",
	}

	return book
}

func title(doc *goquery.Document) string {
	title := doc.Find("#ebooksProductTitle").Text()
	return title
}

func image(doc *goquery.Document) string {
	image, _ := doc.Find("#ebooksImgBlkFront").Attr("src")
	return image
}

func link(uri string) string {
	_, err := url.Parse(uri)
	if err != nil {
		log.Fatal(err)
	}

	return uri
}

func free(doc *goquery.Document) bool {
	free := doc.Find("#buybox #ku-logo").Length()
	return free > 0
}
