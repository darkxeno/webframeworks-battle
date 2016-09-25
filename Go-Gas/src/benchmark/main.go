package main

import (
  "github.com/go-gas/gas"
  "net/http"
  //"fmt"
  "os"
  "runtime"
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

type DataStore struct {
    session *mgo.Session
}

var (
  masterDataStore = DataStore{GetDBSession()}
)

func main() {
  g := gas.Default()
  g.LoadConfig("config.yaml")

  runtime.GOMAXPROCS(4)

  g.Router.Get("/", Index)
  g.Router.Post("/person", CreatePerson)
  g.Router.Get("/test", GetTestData)

  g.Run("0.0.0.0:8181")
}

func GetDBSession() *mgo.Session {

  mongoURL := os.Getenv("MONGODB_URL")

  if mongoURL == "" { mongoURL = "localhost:27017" }

  session, err := mgo.Dial(mongoURL)
  if err != nil {
    log.Fatal("error connecting mongodb:",err,mongoURL)    
    panic(err)
  }  

  // Optional. Switch the session to a monotonic behavior.
  //session.SetMode(mgo.Monotonic, true) 
  session.SetMode(mgo.Strong, true) 

  return session
}

func GetDataStore() *DataStore {
  return &DataStore{masterDataStore.session.Copy()}
  //return &GetDBSession()
}

func Index(ctx *gas.Context) error {
  return ctx.HTML(http.StatusOK, "Micro service! <br> <a href=\"/user\">json response example</a>")
}

func GetTestData(ctx *gas.Context) error {

  session := GetDataStore().session
  defer session.Close()

  c := session.DB("local").C("testData")

  limit := 100

  if i, err := strconv.Atoi(ctx.GetParam("limit")); err == nil {
      limit = i
  }  

  result := []PersonAll{}
  err := c.Find(bson.M{}).Limit(limit).All(&result)
  if err != nil {
    log.Fatal("error on mongodb find:",err)
  }

  //fmt.Println("result[0]:", result[0])  

  return ctx.JSON(http.StatusOK, result)
}


func CreatePerson(ctx *gas.Context) error {

  session := GetDataStore().session
  defer session.Close()

  c := session.DB("test").C("people")

  err := c.Insert(&Person{bson.NewObjectId(), "Ale", "+55 53 8116 9639"},
           &Person{bson.NewObjectId(), "Cla", "+55 53 8402 8510"})
  if err != nil {
    log.Fatal("error on mongodb create:",err)
  }  

  return ctx.JSON(http.StatusOK, err)
}


