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

var books = []book{
	{ID: "1", Title: "Book 1", Author: "Author 1", Quantity: 10},
	{ID: "2", Title: "Book 2", Author: "Author 2", Quantity: 5},
	{ID: "3", Title: "Book 3", Author: "Author 3", Quantity: 20},


}

func getbooks(c *gin.Context){
	c.IndentedJSON(http.StatusOK, books)  
}



func bookById (c *gin.Context) {
	id := c.Param("id")

	book, err := getBookById(id)
	if err != nil {	
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, book)

}

func getBookById(id string) (*book, error){
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("book not found")
}

func checkoutBook (c *gin.Context){
	id, ok := c.GetQuery(("id"))
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter"})
		return
	}
	book, err := getBookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message":"Book not found"})
		return
	}

	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusConflict, gin.H{"message": "Book out of stock"})
		return
	}
	book.Quantity -= 1
	c.IndentedJSON(http.StatusOK, book)
}

func returnbook(c *gin.Context){
	id, ok := c.GetQuery(("id"))
	if !ok {	
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter"})
		return
	}
	book, err := getBookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message":"Book not found"})
		return
	}

	book.Quantity += 1
	c.IndentedJSON(http.StatusOK, book)
}

func createBook(c *gin.Context){
	var newBook book
	if err := c.BindJSON(&newBook); err != nil {
		return
	}
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)

}

func main(){
	router := gin.Default()
	router.GET("/books", getbooks)
	router.GET("/books/:id", bookById)
	router.POST("/books", createBook)
	router.PATCH("/checkout", checkoutBook)
	router.PATCH("/return", returnbook)
	router.Run("localhost:8080")
}