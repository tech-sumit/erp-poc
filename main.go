package main

import (
	"bitbucket.org/perennialsys/erp_database/connection/model"
	"bitbucket.org/perennialsys/erp_database/connection/sql/postgresql"
	"bitbucket.org/perennialsys/erp_database/sql/stores/store"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"os"
)

func main() {
	var err error
	r := gin.Default()
	sqlConfig := &model.ConnectionMeta{
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Database: os.Getenv("DB_NAME"),
	}

	connection, err := postgresql.GetConnection(sqlConfig)
	if err != nil {
		log.Fatal(err)
		return
	}
	storeDB := store.NewConnection(connection)
	//marketplaceCache, err := marketplace.NewConnection(&cacheModel.ConnectionMeta{
	//	Host: os.Getenv("CACHE_HOST_URL"),
	//})
	//if err != nil {
	//	log.Fatal(err)
	//	return
	//}

	r.GET("/store", func(c *gin.Context) {
		data, err := storeDB.GetStores(&map[string]interface{}{})
		if err != nil {
			c.JSON(400, gin.H{
				"status": "failed",
				"message": err.Error(),
			})
			return
		}
		if len(data) > 0 {
			c.JSON(200, data)
		} else {
			c.JSON(204, gin.H{
				"status": "No Content",
			})
		}
	})
	//r.GET("/categories", func(c *gin.Context) {
	//	rawCategories, err := marketplaceCache.GetCategories(c.Request.Context(), "bukalapak")
	//	if err != nil {
	//		c.JSON(400, gin.H{
	//			"status": "failed",
	//			"message": err.Error(),
	//		})
	//		return
	//	}
	//	c.JSON(200, rawCategories)
	//})
	err = r.Run("0.0.0.0:80")
	if err != nil {
		log.Fatal(err)
	}
}
