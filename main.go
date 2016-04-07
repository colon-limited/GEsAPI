package main

import (
  "gopkg.in/olivere/elastic.v3"
  "github.com/BurntSushi/toml"
  "github.com/gin-gonic/gin"
)

func main {
  route := gin.Default()
  v1 := router.Group("/v1")
  {
    v1.POST("/list/reports", ListReports)
    v1.POST("/get/report/:id", GetReport)
  }
}

func ListReports(c *gin.Context) {}
func GetReport(c *gin.Context) {}
