package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		dbUrl = "postgresql://postgres:postgres@localhost:5432/postgres"
	}

	server := newServer(dbUrl)
	defer server.db.Close()

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"hello": "world",
		})
	})
	r.GET("/bundle/:name", server.getBundles)
	r.POST("/bundle/:name", server.createBundle)
	r.Run()
}

func createSchema(db *sqlx.DB) error {
	// assert table exists
	_, err := db.Query(`SELECT id FROM bundles;`)
	if err != nil {
		log.Println("creating table bundles")
		_, err = db.Exec(`CREATE TABLE bundles (
				id serial NOT NULL PRIMARY KEY,
				name varchar(125) NOT NULL,
				data json NOT NULL
			);`)
	}
	return err
}

type Server struct {
	db *sqlx.DB
}

func newServer(dbUrl string) *Server {
	db, err := sqlx.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}
	log.Println("Connected to " + dbUrl)

	err = createSchema(db)
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}

	return &Server{
		db,
	}
}

func (s *Server) getBundles(c *gin.Context) {

}

func (s *Server) createBundle(c *gin.Context) {
	name := c.Param("name")
	data := make(map[string]interface{})
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{
		"name": name,
		"data": data,
	})
}
