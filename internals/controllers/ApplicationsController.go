package controllers

import (
	models "bricks-core/internals/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ReadApplications(c *gin.Context) {
	var applications []models.Application

	//sql, _, _ := goqu.From("applications").ToSQL()
	//
	//query, err := pg.Conn.Query(context.Background(), sql)
	//
	//for query.Next() {
	//	var id int
	//	var name string
	//	err := query.Scan(&id, &name)
	//
	//	if err != nil {
	//		c.JSON(http.StatusBadRequest, err)
	//		return
	//	}
	//
	//	applications = append(applications, models.Application{Name: name})
	//}
	//
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, err)
	//	return
	//}

	applications = append(applications, models.Application{Name: "foo"})

	c.JSON(http.StatusOK, applications)
}

func CreateApplication(c *gin.Context) {
	//var applications []models.Application
	application := models.Application{
		Name: "foo",
	}
	//c.ShouldBindJSON(&application)
	//
	//log.Print(application)
	//
	//tx, _ := pg.Conn.Begin(context.Background())
	//vals := goqu.Insert("applications").
	//	Cols("name").
	//	Vals(goqu.Vals{application.Name})
	//sql, params, _ := vals.ToSQL()
	//tx.Exec(context.Background(), fmt.Sprint(sql, params))
	//_ = tx.Commit(context.Background())

	c.JSON(http.StatusOK, application)
}
