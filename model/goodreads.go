package model

import "encoding/xml"

type NameLookupResponse struct {
	XMLName xml.Name      `xml:"GoodreadsResponse"`
	Request *Request      `xml:"Request"`
	Author  *AuthorByName `xml:"author"`
}

type IdLookupResponse struct {
	XMLName xml.Name   `xml:"GoodreadsResponse"`
	Request *Request   `xml:"Request"`
	Author  AuthorById `xml:"author"`
}

type Request struct {
	XMLName        xml.Name `xml:"Request"`
	Authentication bool     `xml:"authentication"`
	Key            string   `xml:"key"`
	Method         string   `xml:"method"`
}

type AuthorByName struct {
	XMLName xml.Name `xml:"author"`
	Id      string   `xml:"id,attr"`
	Name    string   `xml:"name"`
	Link    string   `xml:"link"`
}

type AuthorById struct {
	XMLName         xml.Name `xml:"author"`
	Id              string   `xml:"id"`
	Name            string   `xml:"name"`
	Link            string   `xml:"link"`
	FansCount       int      `xml:"fans_count"`
	AuthorFollowers int      `xml:"author_followers_count"`
	LargeImageUrl   string   `xml:"large_image_url"`
	ImageUrl        string   `xml:"image_url"`
	SmallImageUrl   string   `xml:"small_image_url"`
	About           string   `xml:"about"`
	Influences      string   `xml:"influences"`
	WorksCount      int      `xml:"works_count"`
	Gender          string   `xml:"gender"`
	Hometown        string   `xml:"hometown"`
	BornAt          string   `xml:"born_at"`
	DiedAt          string   `xml:"died_at"`
	GoodreadsAuthor string   `xml:"goodreads_author"`
}

type Book struct {
	XMLName            xml.Name `xml:"book"`
	Id                 int      `xml:"id"`
	Isbn               string   `xml:"isbn"`
	Isbn13             string   `xml:"isbn13"`
	TextReviewsCount   int      `xml:"text_reviews_count"`
	AmazonURI          string   `xml:"uri"`
	Title              string   `xml:"title"`
	TitleWithoutSeries string   `xml:"title_without_series"`
	ImageUrl           string   `xml:"image_url"`
	SmallImageUrl      string   `xml:"small_image_url"`
	LargeImageUrl      string   `xml:"large_image_url"`
	Link               string   `xml:"link"`
	NumPages           string   `xml:"num_pages"`
	Format             string   `xml:"format"`
	EditionInfo        string   `xml:"edition_information"`
	Publisher          string   `xml:"publisher"`
	PublicationDay     string   `xml:"publication_day"`
	PublicationYear    string   `xml:"publication_year"`
	PublicationMonth   string   `xml:"publication_month"`
	AverageRating      float32  `xml:"average_rating"`
	RatingsCount       int      `xml:"ratings_count"`
	Description        string   `xml:"description"`
	Published          string   `xml:"published"`
	Work               Work     `xml:"work"`
}

type Work struct {
	XMLName xml.Name `xml:"work"`
	Id      string   `xml:"id"`
	Uri     string   `xml:"uri"`
}
