package controllers

import (
	models "bricks-core/internals/pkg"
	pointer_db "bricks-core/internals/pkg/db"
	"context"
	unitTest "github.com/Valiben/gin_unit_test"
	"github.com/Valiben/gin_unit_test/utils"
	"github.com/doug-martin/goqu/v9"
	"gotest.tools/v3/assert"
	"testing"
)

func TestShouldListApplications(t *testing.T) {
	var resp []models.Application
	mockData := []models.Application{
		{Name: "foo"},
		{Name: "bar"},
	}

	defer t.Cleanup(func() {
		clearDatabase(mockData, t)
	})

	populateDatabase(mockData, t)
	populateDatabase(mockData, t)

	err := unitTest.TestHandlerUnMarshalResp(utils.GET, "/applications", "json", nil, &resp)

	if err != nil {
		t.Errorf("This failed to get back a list of applications %d", err)
	}

	if len(resp) == 0 {
		t.Errorf("Should get back %d but got %d", 1, len(resp))
	}
}

func TestShouldCreateApplication(t *testing.T) {
	application := models.Application{Name: "bar"}

	var resp models.Application
	err := unitTest.TestHandlerUnMarshalResp(utils.POST, "/applications", "json", application, &resp)

	assert.NilError(t, err)
	assert.Equal(t, resp.Name, "bar")
}

// Takes in array of interface to create object in DB
func populateDatabase(testModels []models.Application, t *testing.T) error {
	for _, model := range testModels {
		sql, _, err := goqu.Insert("applications").
			Cols("name").
			Vals(
				goqu.Vals{model.Name},
			).
			ToSQL()

		_, err1 := pointer_db.Conn.Exec(context.Background(), sql)

		if err != nil {
			return err1
		}
	}
	return nil
}

func clearDatabase(testModels []models.Application, t *testing.T) error {
	return nil
}
