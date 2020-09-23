// main.go
package main

import (
    "encoding/json"
    "fmt"
    "log"
    "io/ioutil"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/rs/cors"
)

// Article - Our struct for all articles
type Article struct {
    Id      string `json:"Id"`
    Title   string `json:"Title"`
    Desc    string `json:"desc"`
    Content string `json:"content"`
}

var Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Endpoint Hit: returnAllArticles")
    json.NewEncoder(w).Encode(Articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    key := vars["id"]

    for _, article := range Articles {
        if article.Id == key {
            json.NewEncoder(w).Encode(article)
        }
    }
}


func createNewArticle(w http.ResponseWriter, r *http.Request) {
    // get the body of our POST request
    // unmarshal this into a new Article struct
    // append this to our Articles array.    
    reqBody, _ := ioutil.ReadAll(r.Body)
    var article Article
    json.Unmarshal(reqBody, &article)
    // update our global Articles array to include
    // our new Article
    Articles = append(Articles, article)

    json.NewEncoder(w).Encode(article)
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    for index, article := range Articles {
        if article.Id == id {
            Articles = append(Articles[:index], Articles[index+1:]...)
        }
    }

}

func handleRequests() {
    c := cors.New(cors.Options{
        AllowedOrigins: []string{"*"}, // All origins
        AllowedMethods: []string{"GET"}, // Allowing only get, just an example
    })


    myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/", homePage)
    myRouter.HandleFunc("/articles", returnAllArticles)
    myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
    myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
    myRouter.HandleFunc("/article/{id}", returnSingleArticle)
    log.Fatal(http.ListenAndServe(":10000", c.Handler(myRouter)))
}

func main() {

    Articles = []Article{
        Article{Id: "1", Title: "Hello 1", Desc: "Article Description 1", Content: "Article Content 1"},
        Article{Id: "2", Title: "Hello 2", Desc: "Article Description 2", Content: "Article Content 2"},
        Article{Id: "3", Title: "Hello 3", Desc: "Article Description 3", Content: "Article Content 3"},
        Article{Id: "4", Title: "Hello 4", Desc: "Article Description 4", Content: "Article Content 4"},
        Article{Id: "5", Title: "Hello 5", Desc: "Article Description 5", Content: "Article Content 5"},
        Article{Id: "6", Title: "Hello 6", Desc: "Article Description 6", Content: "Article Content 6"},
    }
    handleRequests()
}
