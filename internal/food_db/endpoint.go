package food_db

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type foodData struct {
	Page          int    `form:"page,default=1"`
	PerPage       int    `form:"perPage,default=5"`
	SearchKeyword string `form:"searchKeyword"`
}

// FoodData godoc
//	@Summary		Get food data
//	@Description	get food data
//	@Tags			food
//	@Accept			json
//	@Produce		json
//	@Param			page			query		string	false	"page"
//	@Param			perPage			query		string	false	"per page"
//	@Param			searchKeyword	query		string	false	"search keyword"
//	@Success		200				{object}	Response
//	@Failure		400				{string}	string	"Request Error"
//	@Failure		404				{string}	string	"Not Found Error"
//	@Failure		500				{string}	string	"Internal Server Error"
//	@Router			/foodData [get]
func FoodDataEndpoint(ctx *gin.Context) {
	var param foodData
	if err := ctx.ShouldBind(&param); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
	} else {
		foodData := GetData(param.Page, param.PerPage, param.SearchKeyword)
		ctx.JSON(http.StatusOK, foodData)
	}
}
