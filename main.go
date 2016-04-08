package main

import (
  "gopkg.in/olivere/elastic.v3"
  "github.com/BurntSushi/toml"
  "github.com/gin-gonic/gin"
)

type Config struct {
  Auth Auth
  ES ESConf
}

type ESConf struct {
  Url string `toml:"url"`
  Port string `toml:"port"`
}

type Auth struct {
  Salt string `toml:"salt"`
}

var con Config

func init {
  _, err := toml.DecodeFile("config.toml", &con)
  if err != nil {
    panic(err)
  }
}

func main {
  route := gin.Default()
  v1 := router.Group("/v1")
  {
    v1.POST("/report", ListReports)
    v1.POST("/report/detail", GetReport)
    v1.POST("/report/segment", GetSegment)
    v1.POST("/report/segment/edit", EditSegment)
  }
}

func authorization(key string) {}
func ListReports(c *gin.Context) {}
func GetReport(c *gin.Context) {}
func GetSegment(c *gin.Context) {}
func EditSegment(c *gin.Context) {}
