package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/visket-lang/playground/domain"
	"github.com/visket-lang/playground/model"
	"net/http"
)

func PostCode(c *gin.Context) {
	var job model.Job
	err := c.Bind(&job)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Result{
			Error: err.Error(),
		})
		return
	}

	result := domain.CompileProgram(job)
	if result.Error != "" {
		c.JSON(http.StatusBadRequest, result)
		return
	}

	c.JSON(http.StatusOK, result)
}
