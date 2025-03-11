package main

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"github.com/PeterCaine/go-poker-trainer/pkg/poker"
	"github.com/PeterCaine/go-poker-trainer/web/static"
	"github.com/PeterCaine/go-poker-trainer/web/templates"
)
var (
    game *poker.Game
    mutex sync.Mutex
)

var rangeGrid *poker.Range = poker.NewRange()

func dealHandler(w http.ResponseWriter, r *http.Request){
    mutex.Lock()
    defer mutex.Unlock()
    
    if game == nil || game.CurrentPhase == "showdown"{
        game = poker.NewGame()
    }
    
    game.DealNextPhase()
    component := templates.TableComponent(game.CommunityCards, game.PlayerHand, game.CurrentPhase, rangeGrid)
    component.Render(context.Background(), w)
}

func deckHandler(w http.ResponseWriter, r *http.Request){
    mutex.Lock()
    defer mutex.Unlock()
    
    if game == nil || game.CurrentPhase == "showdown"{
        game = poker.NewGame()
    }
    
    component := templates.TableComponent(game.CommunityCards, game.PlayerHand, game.CurrentPhase, rangeGrid)
    component.Render(context.Background(), w)
}

func rangeToggleHandler(w http.ResponseWriter, r *http.Request) {
    mutex.Lock()
    defer mutex.Unlock()
    path := strings.TrimPrefix(r.URL.Path, "/toggle-range/")
    parts := strings.Split(path, "/")

    if len(parts) != 2 {
        http.Error(w, "incorrect number of url digits passed", http.StatusBadRequest)
        return
    }

    row, err1 := strconv.Atoi(parts[0])
    col, err2 := strconv.Atoi(parts[1])

    if err1 != nil || err2 != nil{
        http.Error(w, "incorrect path suffix passed - expected to be able to cast to integers", http.StatusBadRequest)
        return
    }

    rangeGrid.Grid[row][col].Selected =  !rangeGrid.Grid[row][col].Selected // Return just the updated button HTML

    cell := &rangeGrid.Grid[row][col]

    handKey := cell.String()
    if combo, exists := rangeGrid.Combos[handKey]; exists {
        combo.Selected = cell.Selected
        rangeGrid.Combos[handKey] = combo
    }

    classAttr := "grid-cell"
    if cell.Selected {
        classAttr = "grid-cell selected"
    }
    // Return only the button HTML
    fmt.Fprintf(w, `<button 
        class="%s"
        hx-post="/toggle-range/%d/%d"
        hx-swap="outerHTML"
    >
        %s
    </button>`, classAttr, row, col, cell.String())

}

func statsHandler(w http.ResponseWriter, r * http.Request){
    mutex.Lock()
    defer mutex.Unlock()
    templates.RangeStats(rangeGrid).Render(r.Context(), w)
}

func main(){
    // Serve static files
    fileserver := http.FileServer(http.FS(static.Files))
    http.Handle("/static/", http.StripPrefix("/static/", fileserver))
    http.HandleFunc("/", deckHandler)
    http.HandleFunc("/deal", dealHandler)    // Register handlers with HTTP method constraints
    http.HandleFunc("/toggle-range/", rangeToggleHandler)
    http.HandleFunc("/range-stats", statsHandler)

    fmt.Println("Server running at http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}
