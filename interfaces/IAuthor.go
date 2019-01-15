package interfaces

import (

	"github.com/siddhiparekh11/GoChallenge/models"
	"context"

)

type IAuthor interface {

	//func CreateAuthor(ctx context.Context, author models.Author) (bool,error)
	GetAuthors(ctx context.Context) ([] *models.Author, error)

}