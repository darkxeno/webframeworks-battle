package main

import (
  "github.com/go-gas/gas"
  "net/http"
  "fmt"
  "log"
  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
)

type Person struct {
  Name string
  Phone string
}

func main() {
  g := gas.Default("config.yaml")

  g.Router.Get("/", Index)
  g.Router.Get("/user", GetUser)
  g.Router.Post("/person", CreatePerson)
  g.Router.Get("/person", GetPersons)

  g.Run()
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

func GetPersons(ctx *gas.Context) error {

  session := GetDBSession()
  defer session.Close()

  c := session.DB("test").C("people")

  result := []Person{}
  err := c.Find(bson.M{}).All(&result)
  if err != nil {
          log.Fatal(err)
  }

  fmt.Println("Phone:", result[0])  

  return ctx.JSON(http.StatusOK, result)
}


func CreatePerson(ctx *gas.Context) error {

  session := GetDBSession()
  defer session.Close()

  c := session.DB("test").C("people")

  err := c.Insert(&Person{"Ale", "+55 53 8116 9639"},
           &Person{"Cla", "+55 53 8402 8510"})
  if err != nil {
          log.Fatal(err)
  }  

  return ctx.JSON(http.StatusOK, err)
}


