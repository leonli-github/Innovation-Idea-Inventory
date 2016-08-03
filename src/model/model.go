package model

import "gopkg.in/mgo.v2/bson"
//curl http://localhost:3000/title/facial

type Idea struct{
    Id bson.ObjectId `json:"id" bson:"id"`
	ProjectTitle string `json:"projecttitle" bson:"projecttitle"`
	ApplicationDate string `json:"applicationdate" bson:"applicationdate"`
	Applicant string `json:"applicant" bson:"applicant"`
	ContactNumber string `json:"contactnumber" bson:"contactnumber"`
	Description string `json:"description" bson:"description"`
	PotentialInterest string `json:"potentialinterest" bson:"potentialinterest"`
	RelatedTechnique string `json:"relatedtechnique" bson:"relatedtechnique"`
	RelatedBusiness string `json:"relatedbusiness" bson:"relatedbusiness"`
	status string `json:"status" bson:"status"`
	submittedby string `json:"submittedby" bson:"submittedby"` 
}

//test insert:  curl -XPOST -H 'Content-Type: application/json' -d '{"title": "facial", "staffid": "46463", "status": "In Progress", "submittime": "20160707","description":""}' http://localhost:3000/insert
//test get: curl http://localhost:3000/gettitles