package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetUUIDFromMiddleware(c *gin.Context, param string) (uuid.UUID, bool, error) {
	idStr, exists := c.Get(param)
	idUUID, err := uuid.Parse(idStr.(string))
	return idUUID, exists, err
}

func GetUUIDFromParam(c *gin.Context, param string) (uuid.UUID, error) {
	idStr := c.Param(param)
	parsedID, err := uuid.Parse(idStr)

	return parsedID, err
}
