package api

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/neverhover/Go-000/Week02/core"
	"github.com/neverhover/Go-000/Week02/service"
	"github.com/pkg/errors"
	"net/http"
)

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var res interface{}
		var err error
		// http://localhost:8081/user/1000
		if id == "1000" {
			res, err = service.GetUserError(id)
		} else {
			res, err = service.GetUser(id)
		}
		resp := core.ResponseObject{
			Code: http.StatusOK,
		}
		if errors.Is(err, sql.ErrNoRows) {
			resp.Error = fmt.Sprintf("User %s %s", id, core.ErrorNotFound)
			resp.Code = http.StatusNotFound
			resp.Message = core.ErrorNotFound
			c.JSON(200, resp)
			return
		}
		if err != nil {
			resp.Error = err.Error()
			resp.Code = http.StatusInternalServerError
		}

		resp.Data = res
		c.JSON(200, resp)
	}
}
