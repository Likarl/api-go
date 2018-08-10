package user

import (
	"api-go/handler"
	"api-go/model"
	"api-go/pkg/errno"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"strconv"
)

// @Summary Delete an user by the user identifier
// @Description Delete user by ID
// @Tags user
// @Accept  json
// @Produce  json
// @Param id path uint64 true "The user's database id index num"
// @Success 200 {object} handler.Response "{"code":0,"message":"OK","data":null}"
// @Router /user/{id} [delete]
func Delete(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))
	log.Infof("Delete func called , userId:%v", userId)
	if err := model.DeleteUesr(uint64(userId)); err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	handler.SendResponse(c, nil, nil)
}
