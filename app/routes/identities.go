package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/l-leniac-l/golang-signing-app/domain/repositories"
	"github.com/l-leniac-l/golang-signing-app/domain/services"
)

type SignRequest struct {
	IdentityId   string `json:"identityId"`
	DocumentHash string `json:"documentHash"`
}

func SignDocumentHandler(c *gin.Context) {
	signRequest := SignRequest{}

	if err := c.ShouldBindBodyWith(&signRequest, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request body"})
		return
	}

	ir := repositories.NewIdentityRepository()

	id := ir.GetById(signRequest.IdentityId)

	if id == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf(
				"cannot found identity with id: %s",
				signRequest.IdentityId,
			),
		})
		return
	}

	si := services.NewSignIdentity(id, []byte(signRequest.DocumentHash))

	_, err := si.Sign()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to sign document",
			"error":   err.Error(),
		})
		return
	}

	// res, _ := json.Marshal(si)

	c.JSON(http.StatusCreated, gin.H{
		"signedIdentity": si,
	})
}
