package main

import(
	"fmt"
	"net/http"
	"github.com/julienschmidt/httprouter" //go get github.com/julienschmidt/httprouter
    "controllers"
    "db"
)

func main() {
	//create http router
	router := httprouter.New()
    fmt.Println("starting server")

    innocontroller := controllers.NewController(db.GetSession())

    //router.GET("/",innocontroller.IndexPage)
    router.ServeFiles("/index/*filepath",http.Dir("../template"))
    //Insert new idea into db
    router.POST("/insert",innocontroller.CreateInnoIdea)
    //Search by title
    router.GET("/get/:title",innocontroller.GetInnoIdeaByTitle)
    router.GET("/gettitles/",innocontroller.GetTitles)
    router.NotFound = http.FileServer(http.Dir("../template"))
	
	http.ListenAndServe("localhost:3000", router) 
}



