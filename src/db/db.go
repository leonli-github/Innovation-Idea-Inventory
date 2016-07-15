package db
/*
import(
   "encoding/json"
   "fmt"
   "model"
   "log"
   "gopkg.in/mgo.v2"
   "gopkg.in/mgo.v2/bson"
)
*/
import(
    "gopkg.in/mgo.v2"
    //"gopkg.in/mgo.v2/bson"
)

func GetSession() *mgo.Session{  
    // Connect to our local mongo
    s, err := mgo.Dial("mongodb://localhost")

    // Check if connection error, is mongo running?
    if err != nil {
        panic(err)
    }
    return s
}



 