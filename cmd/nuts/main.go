package main

import (
	"context"
	"fmt"
	"net/http"
    "sync"
	"github.com/PeterCaine/go-poker-trainer/pkg/poker"
	"github.com/PeterCaine/go-poker-trainer/web/templates"
	"github.com/PeterCaine/go-poker-trainer/web/static"
)

var (
    game *poker.Game
    gameMutex sync.Mutex
)

func deckHandler(w http.ResponseWriter, r *http.Request){
    gameMutex.Lock()
    defer gameMutex.Unlock()

    if game == nil || game.CurrentPhase == "showdown"{
        game = poker.NewGame()
    }

    game.DealNextPhase()

    component := templates.TableComponent(game.CommunityCards, game.PlayerHand)
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
