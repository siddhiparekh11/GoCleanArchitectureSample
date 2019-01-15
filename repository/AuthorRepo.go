package repository

import (
	_ "github.com/go-sql-driver/mysql"
	"context"
	"database/sql"
	"github.com/siddhiparekh11/GoChallenge/interfaces"
	"github.com/siddhiparekh11/GoChallenge/models"
	"log"
)


type AuthorRepository struct {
    Conn *sql.DB
}

//creates and returns an object of author repository interface
func NewAuthorRepository(conn *sql.DB) (interfaces.IAuthor) {

		return &AuthorRepository {conn}

}


/*func (authorRepo *AuthorRepository) CreateAuthor(ctx context.Context, author models.Author) (bool,error) {



}*/

func (authorRepo *AuthorRepository) GetAuthors(ctx context.Context) ([] *models.Author, error) {

	log.Println("I am called from repo")
	query := "SELECT idAuthor,firstName,lastName from Authors"
	rows,err := authorRepo.Conn.QueryContext(ctx,query)

	if err!=nil {
		log.Println(err)
		return nil,err
	}

	defer rows.Close()

	author := make([] *models.Author,0)

	for rows.Next() {


		
		a := new(models.Author)
		err = rows.Scan(&a.ID,&a.Firstname,&a.Lastname)
		
		if err!=nil {
		return nil,err
		}

		//log.Println(a)

		author = append(author,a)
	
	}

	return author,nil

}
