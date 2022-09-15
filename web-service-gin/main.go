package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	router := gin.Default() // Initialize a Gin router using Default.
	// this sets up an association in which getAlbums handles requests to the /albums endpoint path.
	// in other words, use the GET function to associate the GET HTTP method and /albums path with a handler function.
	router.GET("/albums", getAlbums)
	// Associate the POST method at the /albums path with the postAlbums function.
	router.POST("/albums", postAlbums)
	// attach the router to an http.Server and start the server.
	router.Run("localhost:8080")
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums) //  serialize the struct into JSON and add it to the response.
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON (request body) to newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	// Add a 201 status code to the response, along with JSON representing the album you added.
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
