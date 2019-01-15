package controller


import (
	
	"context"
	//"github.com/siddhiparekh11/GoChallenge/repository"
	"github.com/siddhiparekh11/GoChallenge/interfaces"
	"github.com/siddhiparekh11/GoChallenge/models"
	"log"
)

// fan-out, fan-in, pipes, channels 

type AuthorController struct {

	ARepo interfaces.IAuthor
}


type AuthorChannel struct {

	Authors []*models.Author
	Error error
}


func NewAuthorController(aRepo interfaces.IAuthor) interfaces.IAuthor {
	return &AuthorController { ARepo: aRepo }
}


func (authContr *AuthorController) GetAuthors(ctx context.Context) ([] *models.Author, error) {

	authorsChan := make(chan AuthorChannel)

	go func() {

		authrs,err := authContr.ARepo.GetAuthors(ctx)

		chanObj := new(AuthorChannel)
		chanObj.Authors = authrs
		chanObj.Error = err

		authorsChan <- *chanObj

	}()

	

	authors := make([] *models.Author,0)

	for c := range authorsChan {	
			log.Println("I am called")	
			authors = c.Authors
			break
	}
	log.Println("controller")
	close(authorsChan)
	
	log.Println(authors[0])

	/*for author := range c.Authors {
					if author!=nil {
						authors=append(authors,author)
					}			
			}*/

	return authors, nil



}

/*func (authContr *AuthorController) CreateAuthor(ctx context.Context,author models.Author) (bool, error) {


}*/