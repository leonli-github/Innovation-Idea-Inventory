package main

import (
        "fmt"
	"log"
        "gopkg.in/mgo.v2"
        "gopkg.in/mgo.v2/bson"
)


type Innovation struct {
        Id bson.ObjectId `json:"id" bson:"id"`
        Title string `json:"title" bson:"title"`
        StaffID string `json:"staffid" bson:"staffid"`
        Status string `json:"status" bson:"status"`
        SubmitTime string  `json:"submittime" bson:"submittime"`
        Description string `json:"description" bson"description"`
    }


func main() {
        session, err := mgo.Dial("mongodb://localhost")
        if err != nil {
                panic(err)
        }
        defer session.Close()

        // Optional. Switch the session to a monotonic behavior.
        session.SetMode(mgo.Monotonic, true)

        c := session.DB("innodb").C("innoidea")
        //err = c.Insert(&Person{"Ale", "+55 53 8116 9639"},
	               //&Person{"Cla", "+55 53 8402 8510"})
        //if err != nil {
                //log.Fatal(err)
        //}

        result := []Innovation{}
        //result []struct{ Title string `bson:"title"` }
        err = c.Find(nil).Select(bson.M{"description":1}).All(&result)
        if err != nil {
                log.Fatal(err)
        }

        //fmt.Println("title:", result[1].Title)
        for _, v := range result {
        fmt.Println(v.Description)
}
}