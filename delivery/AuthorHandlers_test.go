package delivery_test


import (
	
	"testing"
	"net/http"
	"net/http/httptest"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/siddhiparekh11/GoChallenge/delivery"
	"github.com/siddhiparekh11/GoChallenge/models"
	"io/ioutil"
	"github.com/gorilla/mux"
	"context"

)

type fakeContr struct {}


func (f fakeContr)GetAuthors(ctx context.Context) ([] *models.Author,error) {

	authors := make([] *models.Author,0)
	authors = append(authors,&models.Author{"1","Siddhi","Parekh"})
	authors = append(authors,&models.Author{"2","Ishani","Parekh"})

	return authors,nil

}



func TestGetAuthors(t *testing.T) {

	req, err := http.NewRequest("GET","/api/authors",nil)
	if err != nil {
				t.Fatalf("could not create request: %v", err)
	}
	rec := httptest.NewRecorder()

	rou := delivery.NewAuthorHandler(mux.NewRouter(),nil,fakeContr{})

	rou.Mux.ServeHTTP(rec,req)

	assert.Equal(t,200,rec.Code,"Ok response is expected")

	res := rec.Result()
	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
				t.Fatalf("could not read response: %v", err)
    }

    fmt.Println(b)

}




