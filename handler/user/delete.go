package user

import (
	"api-go/handler"
	"api-go/model"
	"api-go/pkg/errno"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"strconv"
)

// Delete delete an user by the user identifier.
func Delete(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))
	log.Infof("Delete func called , userId:%v", userId)
	if err := model.DeleteUesr(uint64(userId)); err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	handler.SendResponse(c, nil, nil)
}
