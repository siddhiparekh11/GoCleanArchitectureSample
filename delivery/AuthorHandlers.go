package delivery


import (

	"context"
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
	//"github.com/siddhiparekh11/GoChallenge/controllers"
	"github.com/siddhiparekh11/GoChallenge/interfaces"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
)


type IAuthorHandlers interface {
	GetAuthors(w http.ResponseWriter, r *http.Request) 
}


type AuthorHandler struct {

	mux *mux.Router
	conn *sql.DB
	IAuthorHandlers
	aContr interfaces.IAuthor

}


func NewAuthorHandler(m *mux.Router,con *sql.DB, aCon interfaces.IAuthor) {

	authorHandler := &AuthorHandler {
		mux : m,
		conn: con,
		aContr: aCon,
	}

	authorHandler.mux.HandleFunc("/api/authors",authorHandler.GetAuthors).Methods("GET")



}


func (authorHandler *AuthorHandler) GetAuthors(w http.ResponseWriter, r *http.Request)  {

	w.Header().Set("Content-Type","application/json")
	authors, err := authorHandler.aContr.GetAuthors(context.Background())
	if err != nil {
		
	}

	json.NewEncoder(w).Encode(authors)

}

