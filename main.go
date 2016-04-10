package main

import (
  "fmt"
  "log"
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

func setAccessHeader(c *gin.Context) {
  c.Header("Access-Control-Allow-Origin", "http://localhost")
  c.Header("Access-Control-Allow-Credentials", "true")
  c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
  c.Header("Access-Control-Allow-Methods","GET, POST, PUT, DELETE, OPTIONS")
}

func before(c *gin.Context, key string)(res bool) {
  setAccessHeader(c)
  if auth(key) {
    res = true
  } else {
    res = false
  }
  return res
}

func main() {
  router := gin.Default()
  router.Use(gin.Logger())
  v1 := router.Group("/v1")
  {
    v1.POST("/report", ListReports)
    v1.POST("/report/detail/:all", GetReport)
    v1.POST("/report/segment", GetSegment)
    v1.POST("/report/segment/edit", EditSegment)
  }
  //router.Run(con.Base.Port)
  router.Run()
}

func ListReports(c *gin.Context) {
  var req ReportList
  log.Println(fmt.Sprintf("req is %d", req))
  c.BindJSON(&req)
  log.Println(fmt.Sprintf("req is %d", req))
  bef := before(c, req.Auth)
  log.Println(fmt.Sprintf("before result is %d", bef))
  if bef {
    c.JSON(200,gin.H{"status":"200",})
  } else {
    c.JSON(500,gin.H{"status":"500",})
  }
}
func GetReport(c *gin.Context) {
  var req ReportDetail
  c.BindJSON(&req)
  if before(c, req.Auth) {
    c.JSON(200,gin.H{"status":"200",})
  } else {
    c.JSON(500,gin.H{"status":"500",})
  }
}
func GetSegment(c *gin.Context) {
  var req SegmentList
  c.BindJSON(&req)
  if before(c, req.Auth) {
    c.JSON(200,gin.H{"status":"200",})
  } else {
    c.JSON(500,gin.H{"status":"500",})
  }
}
func EditSegment(c *gin.Context) {
  var req SegmentUpdate
  c.BindJSON(&req)
  if before(c, req.Auth) {
    c.JSON(200,gin.H{"status":"200",})
  } else {
    c.JSON(500,gin.H{"status":"500",})
  }
}
