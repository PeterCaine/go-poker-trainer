package main

import (
	"context"
	"fmt"
	"net/http"
	"github.com/PeterCaine/go-poker-trainer/pkg/poker"
	"github.com/PeterCaine/go-poker-trainer/web/templates"
	"github.com/PeterCaine/go-poker-trainer/web/static"
)


func deckHandler(w http.ResponseWriter, r *http.Request){
    deck := poker.CreateDeck()
    deck.ShuffleDeck()
    communityCards := deck.Deal(3)
    component := templates.TableComponent(communityCards)
    component.Render(context.Background(), w)
}

func main(){
    // Serve static files
    fileserver := http.FileServer(http.FS(static.Files))
    http.Handle("/static/", http.StripPrefix("/static/", fileserver))
    http.HandleFunc("/", deckHandler)
    fmt.Println("Server running at http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}
