package public

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/internal/database"
	"server/internal/models"
	"server/internal/services"
)

type UserHandlersRoutes interface {
	Index(c *gin.Context)
	Login(c *gin.Context)
}

type UserHandlers struct {
	UsersTable database.UsersTable
}

func (h UserHandlers) Index(c *gin.Context) {
	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "Hello! Current API version: 0.1",
	})
}

func (h UserHandlers) Login(c *gin.Context) {
	type login struct {
		Username string `json:"username,omitempty"`
		Password string `json:"password,omitempty"`
	}

	loginParams := login{}
	c.ShouldBindJSON(&loginParams)

	user, err := h.UsersTable.GetOne(loginParams.Username, services.GetMD5Hash(loginParams.Password))

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.SuccessResponse{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SignedResponse{
		Token:   user.Token,
		Message: "Logged in",
	})
}
