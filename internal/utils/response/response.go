package response

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ErrMsg struct {
	Message any `json:"message"`
}

type ResponseList struct {
	Pages    Pages `json:"pages"`
	ListData any   `json:"data"`
}

func OK(data any, c *gin.Context) {
	c.JSON(http.StatusOK, data)
	c.Next()
}

func Error(msg any, code int, c *gin.Context) {
	c.JSON(code, &ErrMsg{
		Message: msg,
	})
	c.Abort()
}

func ValidMsg(err error, _model any) string {
	validModel := reflect.TypeOf(_model)
	if errs, ok := err.(validator.ValidationErrors); ok {
		for _, e := range errs {
			if f, exist := validModel.Elem().FieldByName(e.Field()); exist {
				return f.Tag.Get("msg")
			}
		}
	}
	return err.Error()
}
