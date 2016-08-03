package controllers

import(
   "encoding/json"
   "fmt"
   "net/http"
   "github.com/julienschmidt/httprouter"
   "model"
   "gopkg.in/mgo.v2"
   "gopkg.in/mgo.v2/bson"
   //"io/ioutil"
   "html/template"
)

type InnoController struct {
	 session *mgo.Session
}

func NewController(session *mgo.Session) *InnoController{
     return &InnoController{session}
}
func (IC InnoController) IndexPage(writer http.ResponseWriter, rquest *http.Request, param httprouter.Params) {
      
       //page,_ := ioutil.ReadFile("../template/index.html")
       t:= template.Must(template.ParseFiles("../template/index.html"))
       t.Execute(writer, nil)
      
      //fmt.Fprintf(writer,page)
}

func(IC InnoController) CreateInnoIdea(writer http.ResponseWriter, rquest *http.Request, param httprouter.Params){
    
      idea := model.Idea{}
       
      json.NewDecoder(rquest.Body).Decode(&idea)
      
      s:= IC.session.Copy()
      defer s.Close()
      //TODO: Check if idea is valid
      //insert into database
      idea.Id = bson.NewObjectId()

      s.DB("innodb").C("innoidea").Insert(idea)

}

func(IC InnoController) GetInnoIdeaByTitle(writer http.ResponseWriter, ruest *http.Request, param httprouter.Params){
	  
      s:= IC.session.Copy()
      defer s.Close()

      idea:= model.Idea{}

      title := param.ByName("title")

      if err := s.DB("innodb").C("innoidea").Find(bson.M{"projecttitle":title}).One(&idea); err!=nil {
        writer.WriteHeader(404  )
        return
      }

    //ideajson,_ := json.Marshal(idea)

    //writer.Header().Set("Content-Type", "application/json")
    //writer.WriteHeader(200)
    //fmt.Fprintf(writer, "%s", ideajson)

      //use template
      t, _ := template.ParseFiles("../template/idea.html")
      t.Execute(writer,idea)



}
func(IC InnoController) GetTitles(writer http.ResponseWriter, ruest *http.Request, param httprouter.Params){
    
      s:= IC.session.Copy()
      defer s.Close()
      
      ideas:= []model.Idea{}

      if err := s.DB("innodb").C("innoidea").Find(nil).Select(bson.M{"projecttitle":1,"description":1}).All(&ideas); err!=nil {
        writer.WriteHeader(404)
        return    
      }
    
    ideajson,_ := json.Marshal(ideas)

    writer.Header().Set("Content-Type", "application/json")
    writer.WriteHeader(200)
    fmt.Fprintf(writer, "%s", ideajson)


}

func(IC InnoController) GetNEXT(writer http.ResponseWriter, ruest *http.Request, param httprouter.Params){
    
  
      t, _ := template.ParseFiles("../template/NEXT.html")
      t.Execute(writer,nil)



}
func(IC InnoController) GetLUCKY(writer http.ResponseWriter, ruest *http.Request, param httprouter.Params){
    
  
      t, _ := template.ParseFiles("../template/lucky.html")
      t.Execute(writer,nil)



}

//curl -XPOST -H 'Content-Type: application/json' -d '{"name": "Bob Smith", "gender": "male", "age": 50}' http://localhost:3000/user
