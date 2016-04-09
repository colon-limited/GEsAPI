package main

import (
  //"gopkg.in/olivere/elastic.v3"
  "github.com/BurntSushi/toml"
  "github.com/gin-gonic/gin"
  // use for hash from scrypt
  "encoding/base64"
  "strings"
)

// Config
type Config struct {
  Base Base
  Auth Auth
  ES ES
}
type Base struct {
  Port string `toml:"port"`
}
type ES struct {
  Url string `toml:"url"`
  Port string `toml:"port"`
}
type Auth struct {
  Salt string `toml:"salt"`
  Pass string `toml:"seed"`
}

// Request Paramater
type ReportList struct {
  Auth string `json:"auth"`
}
type ReportDetail struct {
  Auth string `json:"auth"`
  Id int `json:"report"`
}
type SegmentList struct {
  Auth string `json:"auth"`
  Id int `json:"report"`
}
type SegmentUpdate struct {
  Auth string `json:"auth"`
  Id int `json:"report"`
  Segment []SegmentItem `json:"segment"`
}
type SegmentItem struct {
  Name string `json:"name"`
  Cond string `json:"condition"`
}

var con Config

func init() {
  _, err := toml.DecodeFile("config.toml", &con)
  if err != nil {
    panic(err)
  }
}

/*
* auth code is "base64encodedpass" + "," + "base64encodedsalt" strings over base64 encoded.
*/
func auth(st string)(res bool) {
  de, _ := base64.StdEncoding.DecodeString(st)
  spstr := strings.Split(string(de), ",")
  pass_code, _ := base64.StdEncoding.DecodeString(spstr[0])
  salt_code, _ := base64.StdEncoding.DecodeString(spstr[1])
  if string(pass_code) == con.Auth.Pass && string(salt_code) == con.Auth.Salt {
    res = true
  } else {
    res = false
  }
  return res
}

func main() {
  router := gin.Default()
  v1 := router.Group("/v1")
  {
    v1.POST("/report", ListReports)
    v1.POST("/report/detail", GetReport)
    v1.POST("/report/segment", GetSegment)
    v1.POST("/report/segment/edit", EditSegment)
  }
  router.Run()
}

func ListReports(c *gin.Context) {
  var req ReportList
  c.BindJSON(&req)
  if auth(req.Auth) {
    c.JSON(200,gin.H{"status":"200",})
  } else {
    c.JSON(500,gin.H{"status":"500",})
  }
}
func GetReport(c *gin.Context) {
  var req ReportDetail
  c.BindJSON(&req)
  if auth(req.Auth) {
    c.JSON(200,gin.H{"status":"200",})
  } else {
    c.JSON(500,gin.H{"status":"500",})
  }
}
func GetSegment(c *gin.Context) {
  var req SegmentList
  c.BindJSON(&req)
  if auth(req.Auth) {
    c.JSON(200,gin.H{"status":"200",})
  } else {
    c.JSON(500,gin.H{"status":"500",})
  }
}
func EditSegment(c *gin.Context) {
  var req SegmentUpdate
  c.BindJSON(&req)
  if auth(req.Auth) {
    c.JSON(200,gin.H{"status":"200",})
  } else {
    c.JSON(500,gin.H{"status":"500",})
  }
}
