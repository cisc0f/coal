package main

import (
	"io"
	"net/http"
	"time"

	fire "github.com/cisc0f/coal/internal/db"
	"github.com/cisc0f/coal/internal/handler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// GET: get server status
func getStatus() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.IndentedJSON(http.StatusOK, "Server up and running")
	}
}

// GET: get all songs
func getSongs(db *fire.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		songs, err := handler.GetAllSongs(db)
		if err != nil {
			ctx.IndentedJSON(http.StatusNotFound, "Something went wrong")
		}

		ctx.IndentedJSON(http.StatusOK, songs)
	}
}

// GET: Delete all songs on firebase (KEEP COMMENTED IF NOT ADMIN)
// func deleteAllSongs(db *DB) gin.HandlerFunc {
// 	return func(ctx *gin.Context) {

// 		ref := db.Client.NewRef("songs")

// 		if err := ref.Delete(context.TODO()); err != nil {
// 			fmt.Print("error in deleting ref: ", err)
// 		}

// 		fmt.Println("all songs deleted successfully")
// 	}
// }

// POST: compare song using Coal
func postCompare(db *fire.DB, s *fire.Storage) gin.HandlerFunc {
	return func(c *gin.Context) {

		err := c.Request.ParseMultipartForm(10 << 20) // 10 MB limit
		if err != nil {
			c.String(http.StatusBadRequest, "Unable to parse form")
			return
		}

		// Get the file from the form data
		file, _, err := c.Request.FormFile("midiFile")
		if err != nil {
			c.String(http.StatusBadRequest, "Unable to get file from form")
			return
		}
		defer file.Close()

		var fileReader io.Reader = file

		res, err := handler.PostCompare(db, s, &fileReader)
		if err != nil {
			c.JSON(400, gin.H{"error": err})
			return
		}
		c.JSON(http.StatusOK, gin.H{"id": res.Id, "title": res.Title, "author": res.Author, "matchingRate": res.MatchingRate})
	}
}

// POST: create new song
func postSong(db *fire.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

// Main entrypotn
func main() {

	db := &fire.DB{}
	if err := db.SetupDB(); err != nil {
		panic("failed to setup db: " + err.Error())
	}

	s := &fire.Storage{}
	if err := s.SetupStorage(); err != nil {
		panic("failed to setup storage: " + err.Error())
	}

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"POST", "GET"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	router.GET("/status", getStatus())
	router.GET("/songs", getSongs(db))
	router.POST("/compare", postCompare(db, s))
	router.POST("/song", postSong(db))
	// router.GET("/deleteSongs", deleteAllSongs(db))

	router.Run("localhost:8888")
}
