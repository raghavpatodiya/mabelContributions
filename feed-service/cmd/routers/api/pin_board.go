package api

import (
	"net/http"
	pin_board_service "super-feed-service/cmd/service/pin_board"
	"super-feed-service/cmd/utils"

	"github.com/gin-gonic/gin"
)

// CreatePinBoard ... Create PinBoard
// @Summary     Create new pin board for the user
// @Description Create new pin board
// @Tags        Pin Boards
// @Accept      json
// @Param       pinboard body models.PinBoard true "Pin Board Data"
// @Success     200
// @Failure     404
// @Failure     500
// @Router      /pinboards/ [post]
// @Security    BearerAuth
func CreatePinBoard(c *gin.Context) {
	res, err := pin_board_service.CreatePinBoardService(c)
	statusCode, data := utils.FormatResponseMessage(res, err, http.StatusCreated)
	c.JSON(statusCode, data)
}

func GetPinBoard(c *gin.Context) {
	res, err := pin_board_service.GetPinBoardService(c)
	statusCode, data := utils.FormatResponseMessage(res, err, http.StatusOK)
	c.JSON(statusCode, data)
}

func DeletePinBoard(c *gin.Context) {
	res, err := pin_board_service.DeletePinBoardService(c)
	statusCode, data := utils.FormatResponseMessage(res, err, http.StatusNoContent)
	c.JSON(statusCode, data)
}

func UpdatePinBoard(c *gin.Context) {
	res, err := pin_board_service.UpdatePinBoardService(c)
	statusCode, data := utils.FormatResponseMessage(res, err, http.StatusNoContent)
	c.JSON(statusCode, data)
}

func PinsByBoard(c *gin.Context) {
	res, err := pin_board_service.GetPinsByBoardService(c)
	statusCode, data := utils.FormatResponseMessage(res, err, http.StatusOK)
	c.JSON(statusCode, data)
}

func GetPinBoardsByUser(c *gin.Context) {
	res, err := pin_board_service.GetPinBoardsByUserService(c)
	statusCode, data := utils.FormatResponseMessage(res, err, http.StatusOK)
	c.JSON(statusCode, data)
}

func PinFeedContent(c *gin.Context) {
	res, err := pin_board_service.PinFeedContentService(c)
	statusCode, data := utils.FormatResponseMessage(res, err, http.StatusCreated)
	c.JSON(statusCode, data)
}

func UnpinFeedContent(c *gin.Context) {
	res, err := pin_board_service.UnpinFeedContentService(c)
	statusCode, data := utils.FormatResponseMessage(res, err, http.StatusNoContent)
	c.JSON(statusCode, data)
}

func SharePinBoard(c *gin.Context) {
	res, err := pin_board_service.SharePinBoardService(c)
	statusCode, data := utils.FormatResponseMessage(res, err, http.StatusNoContent)
	c.JSON(statusCode, data)
}

func LikePinBoard(c *gin.Context) {
	res, err := pin_board_service.LikePinBoardService(c)
	statusCode, data := utils.FormatResponseMessage(res, err, http.StatusNoContent)
	c.JSON(statusCode, data)
}

func UnlikePinBoard(c *gin.Context) {
	res, err := pin_board_service.UnLikePinBoardService(c)
	statusCode, data := utils.FormatResponseMessage(res, err, http.StatusNoContent)
	c.JSON(statusCode, data)
}
