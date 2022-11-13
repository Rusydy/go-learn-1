package main

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})

	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)

	router.POST("/albums", postAlbums)

	router.Run("localhost:8081")
}

// album represents data about a record album.
type album struct {
	ID     int64   `form:"id" binding:"" json:"id"`
	Title  string  `form:"title" binding:"required" json:"title"`
	Artist string  `form:"artist" binding:"required" json:"artist"`
	Price  float64 `form:"price" binding:"required" json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: 1, Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: 2, Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: 3, Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// getAlbumByID locates the album whose ID value matches the id
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// convert string to int32
	idInt, err := strconv.ParseInt(id, 10, 32)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid album id"})
		return
	}

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == idInt {
			c.IndentedJSON(http.StatusOK, gin.H{
				"message": "Album found",
				"data":    a,
			})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": strings.Split(err.Error(), "\n"),
		})
		return
	}

	lastID := albums[len(albums)-1].ID
	newAlbum.ID = lastID + 1
	albums = append(albums, newAlbum)

	c.IndentedJSON(http.StatusCreated, newAlbum)

}
