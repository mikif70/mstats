package main

import "github.com/gin-gonic/gin"

func dbstats (c *gin.Context) {
	c.JSON(200, gin.H{ 
		"objects" : 11648137,
        "avgObjSize" : 135.86573715607912,
        "dataSize" : 1588168578,
        "storageSize" : 430227456,
        "indexes" : 42,
        "indexSize" : 1209020416,
	})
}

func main() {
	r := gin.Default()
	r.GET("/dbstats", dbstats)	
	r.Static("/dash", "./client")
    r.Run(":8080") // listen and serve on 0.0.0.0:8080
}