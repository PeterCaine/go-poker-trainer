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
    mutex sync.Mutex
)

func dealHandler(w http.ResponseWriter, r *http.Request){
    mutex.Lock()
    defer mutex.Unlock()
    
    if game == nil || game.CurrentPhase == "showdown"{
        game = poker.NewGame()
    }
    
    game.DealNextPhase()
    component := templates.TableComponent(game.CommunityCards, game.PlayerHand, game.CurrentPhase)
    component.Render(context.Background(), w)
}

func deckHandler(w http.ResponseWriter, r *http.Request){
    mutex.Lock()
    defer mutex.Unlock()
    
    if game == nil || game.CurrentPhase == "showdown"{
        game = poker.NewGame()
    }
    
    component := templates.TableComponent(game.CommunityCards, game.PlayerHand, game.CurrentPhase)
    component.Render(context.Background(), w)
}

func main(){
    // Serve static files
    fileserver := http.FileServer(http.FS(static.Files))
    http.Handle("/static/", http.StripPrefix("/static/", fileserver))
    http.HandleFunc("/", deckHandler)
    http.HandleFunc("/deal", dealHandler)
    fmt.Println("Server running at http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}
