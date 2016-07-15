package model

import "gopkg.in/mgo.v2/bson"
//curl http://localhost:3000/title/facial
type Innovation struct {
        Id bson.ObjectId `json:"id" bson:"id"`
        Title string `json:"title" bson:"title"`
        StaffID string `json:"staffid" bson:"staffid"`
        Status string `json:"status" bson:"status"`
        SubmitTime string  `json:"submittime" bson:"submittime"`
        Description string `json:"description" bson"description"`
    }

//test insert:  curl -XPOST -H 'Content-Type: application/json' -d '{"title": "facial", "staffid": "46463", "status": "In Progress", "submittime": "20160707","description":""}' http://localhost:3000/insert
//test get: curl http://localhost:3000/get/facial