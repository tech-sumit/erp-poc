package main

import (
	"bitbucket.org/perennialsys/erp_database/connection/model"
	"bitbucket.org/perennialsys/erp_database/connection/sql/postgresql"
	"bitbucket.org/perennialsys/erp_database/sql/stores/store"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"net/http"
	"os"
)
var redisPool *redis.Pool

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
	const maxConnections = 10
	redisPool = &redis.Pool{
		MaxIdle: maxConnections,
		Dial:    func() (redis.Conn, error) { return redis.Dial("tcp", os.Getenv("CACHE_HOST_URL")) },
	}
	r.GET("/cache", func(c *gin.Context) {
		conn := redisPool.Get()
		defer conn.Close()

		counter, err := redis.Int(conn.Do("INCR", "visits"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, "Error incrementing visitor counter")
			return
		}
		c.JSON(200,fmt.Sprintf("Visitor number: %d", counter))
	})
	err = r.Run("0.0.0.0:80")
	if err != nil {
		log.Fatal(err)
	}
}


