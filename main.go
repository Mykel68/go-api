package main



import ("net/http"
"github.com/gin-gonic/gin"
"errors"
)

type book struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	Quantity int `json:"quantity"`
}

var books []book{
	{ID: "1", Title: "Book 1", Author: "Author 1", Quantity: 10},
	{ID: "2", Title: "Book 2", Author: "Author 2", Quantity: 5},
	{ID: "3", Title: "Book 3", Author: "Author 3", Quantity: 20},}