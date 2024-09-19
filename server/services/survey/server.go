package main

import (
	"io"
	"log"
	"os"
	"survey/graph"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func initLogger() logger.Interface {
	logLevel := logger.Info
	f, _ := os.Create("gorm.log")
	newLogger := logger.New(
		log.New(
			io.MultiWriter(f, os.Stdout), "\r\n", log.LstdFlags), logger.Config{
			Colorful:                  true,
			LogLevel:                  logLevel,
			SlowThreshold:             time.Second,
			IgnoreRecordNotFoundError: true,
		})

	return newLogger
}

func main() {
	dsn := "host=localhost user=postgres password=1234 dbname=SurveyDB port=5432 sslmode=prefer TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: initLogger()})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	resolver := &graph.Resolver{Db: db}
	schema := graph.NewExecutableSchema(graph.Config{Resolvers: resolver})
	h := handler.NewDefaultServer(schema)
	h.Use(extension.FixedComplexityLimit(20))
	router := gin.Default()
	//router.Use(dataloader.Middleware())
	router.POST("/query", func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	})

	router.GET("/", func(c *gin.Context) {
		playground.Handler("GraphQL", "/query").ServeHTTP(c.Writer, c.Request)
	})

	log.Printf("connect to http://localhost:8080/ for GraphQL playground")
	log.Fatal(router.Run(":8080"))
}
