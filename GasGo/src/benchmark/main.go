package main

import (
  "github.com/go-gas/gas"
  "net/http"
  //"fmt"
  "log"
  "strconv"
  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
)

type Person struct {
  Id bson.ObjectId  
  Name string
  Phone string 
}

type PersonAll interface {}

type PersonExtended struct {
  _Id bson.ObjectId
  Name string
  Phone string 
  Index float32
  isActive bool
  Guid string
  Balance string
  Picture string
  Age float32
  Eyecolor string
  Gender string

  Company string
  Email string 
  Address string
  About string
  Registered string
  Latitude float32
  Longitude float32
  Tags []string
  Friends []interface{}
  Greeting string
  FavoriteFruit string
}

func main() {
  g := gas.Default()
  g.LoadConfig("config.yaml")

  g.Router.Get("/", Index)
  g.Router.Get("/user", GetUser)
  g.Router.Post("/person", CreatePerson)
  g.Router.Get("/test", GetTestData)

  g.Run("localhost:8181")
}

func GetDBSession() *mgo.Session {
  session, err := mgo.Dial("localhost:27017")
  if err != nil {
    panic(err)
  }  

  // Optional. Switch the session to a monotonic behavior.
  session.SetMode(mgo.Monotonic, true)  

  return session
}

func Index(ctx *gas.Context) error {
  return ctx.HTML(http.StatusOK, "Micro service! <br> <a href=\"/user\">json response example</a>")
}

func GetUser(ctx *gas.Context) error {
  return ctx.JSON(http.StatusOK, gas.H{
    "name": "John",
    "age":  32,
  })
}

func GetTestData(ctx *gas.Context) error {

  session := GetDBSession()
  defer session.Close()

  c := session.DB("local").C("testData")

  limit := 100

  if i, err := strconv.Atoi(ctx.GetParam("limit")); err == nil {
      //fmt.Printf("i=%d, type: %T\n", i, i)
      limit = i
  }  

  result := []PersonAll{}
  //var interfaceSlice []interface{} = make([]interface{}, 100)
  err := c.Find(bson.M{}).Limit(limit).All(&result)
  if err != nil {
    log.Fatal(err)
  }

  //fmt.Println("result[0]:", result[0])  

  return ctx.JSON(http.StatusOK, result)
}


func CreatePerson(ctx *gas.Context) error {

  session := GetDBSession()
  defer session.Close()

  c := session.DB("test").C("people")

  err := c.Insert(&Person{bson.NewObjectId(), "Ale", "+55 53 8116 9639"},
           &Person{bson.NewObjectId(), "Cla", "+55 53 8402 8510"})
  if err != nil {
    log.Fatal(err)
  }  

  return ctx.JSON(http.StatusOK, err)
}


