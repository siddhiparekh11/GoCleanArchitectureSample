package main


import (
	
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	 "github.com/siddhiparekh11/GoChallenge/repository"
	"github.com/siddhiparekh11/GoChallenge/controllers"
	"github.com/siddhiparekh11/GoChallenge/delivery"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
	"net/http"
	"log"
	"reflect"
	

)

type App struct {

	Router *mux.Router
	Conn *sql.DB

}


func init() {

	viper.SetConfigFile(`Config.json`)
	err := viper.ReadInConfig()

	if err!=nil {

	}
}


var dObj delivery.AuthorHandler

func main() {

	app := App {
		Router : mux.NewRouter(),
	}

	conn , err := dbConnect()

	if err!=nil {
			fmt.Println(err)
	}

	app.Conn = conn

	aRepo := repository.NewAuthorRepository(app.Conn)
	fmt.Println(reflect.TypeOf(aRepo))
	aContr := controller.NewAuthorController(aRepo)
	delivery.NewAuthorHandler(app.Router,app.Conn,aContr)
	log.Fatal(http.ListenAndServe(":8000",app.Router))

}




func dbConnect() (*sql.DB,error) {

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", viper.GetString(`database.user`), viper.GetString(`database.password`),viper.GetString(`database.host`),viper.GetString(`database.port`),viper.GetString(`database.name`))
	log.Println(connectionString)
	db,err := sql.Open("mysql",connectionString)
	if err!=nil {
		return nil,err
	}

	return db,nil

}




