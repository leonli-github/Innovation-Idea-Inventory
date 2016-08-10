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
    fmt.Println("creating controllers")
    //router.GET("/",innocontroller.IndexPage)
    router.ServeFiles("/index/*filepath",http.Dir("../template"))
    //Insert new idea into db
    router.POST("/insert",innocontroller.CreateInnoIdea)
    //Search by title
    router.GET("/get/:title",innocontroller.GetInnoIdeaByTitle)
    router.GET("/gettitles",innocontroller.GetTitles)
    router.GET("/getNEXT",innocontroller.GetNEXT)
    router.GET("/getLUCKY",innocontroller.GetLUCKY)
    router.NotFound = http.FileServer(http.Dir("../template"))
    fmt.Println("setting router")
	http.ListenAndServe("192.168.54.163:3000", router)
    fmt.Println("listening")
}



