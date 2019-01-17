package repository_test

import (
	repo "github.com/siddhiparekh11/GoChallenge/repository"
	"github.com/siddhiparekh11/GoChallenge/models"
	"testing"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"github.com/stretchr/testify/assert"
	"context"
)


func TestGetAuthors (t *testing.T) {

	db,mock,err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mockAuthors := []models.Author{

			models.Author {
				ID: "1",
				Firstname: "Siddhi",
				Lastname: "Parekh",
			},
			models.Author {
				ID: "2",
				Firstname: "Ishani",
				Lastname: "Parekh",
			},
	}

	rows := sqlmock.NewRows([]string{"idAuthor","firstName","lastName"}).
			AddRow(mockAuthors[0].ID,mockAuthors[0].Firstname,mockAuthors[0].Lastname).
			AddRow(mockAuthors[1].ID,mockAuthors[1].Firstname,mockAuthors[1].Lastname)

	query := "SELECT idAuthor,firstName,lastName from Authors"
	
	mock.ExpectQuery(query).WillReturnRows(rows)

	a := repo.NewAuthorRepository(db)

	authors,err := a.GetAuthors(context.TODO())

	assert.NoError(t, err)
	assert.NotNil(t, authors)		

}