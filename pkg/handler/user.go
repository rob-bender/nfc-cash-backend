package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary		GetUserProfile
// @Tags			user
// @Description	get user profile
// @ID				get-user-profile
// @Accept			json
// @Produce		json
// @Success		200		{integer}	integer				1
// @Failure		400,404	{object}	error
// @Failure		500		{object}	error
// @Failure		default	{object}	error
// @Router			/user/get-user-profile [get]
func (h *Handler) getUserProfile(c *gin.Context) { // получение профиля пользователя
	userId, err := getUserId(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"status":  http.StatusUnauthorized,
			"message": err.Error(),
		})
		return
	}
	if userId == 0 {
		c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"status":  http.StatusUnauthorized,
			"message": "идентификатор пользователя имеет недопустимый тип",
		})
		return
	}
	if userId > 0 {
		resGetUserProfile, statusCode, err := h.services.GetUserProfile(userId)
		if err != nil {
			c.JSON(statusCode, map[string]interface{}{
				"status":  statusCode,
				"message": err.Error(),
			})
			return
		}
		if len(resGetUserProfile) > 0 {
			c.JSON(http.StatusOK, map[string]interface{}{
				"status":  http.StatusOK,
				"message": "успешное получение профиля пользователя",
				"result":  resGetUserProfile,
			})
		} else {
			c.JSON(http.StatusOK, map[string]interface{}{
				"status":  http.StatusOK,
				"message": "успешное получение профиля пользователя",
				"result":  resGetUserProfile,
			})
		}
	}
}

// @Summary		CheckIsAdmin
// @Tags			user
// @Description	check is admin
// @ID				check-is-admin
// @Accept			json
// @Produce		json
// @Success		200		{integer}	integer				1
// @Failure		400,404	{object}	error
// @Failure		500		{object}	error
// @Failure		default	{object}	error
// @Router			/user/check-is-admin [get]
func (h *Handler) checkIsAdmin(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"status":  http.StatusUnauthorized,
			"message": err.Error(),
		})
		return
	}
	if userId == 0 {
		c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"status":  http.StatusUnauthorized,
			"message": "идентификатор пользователя имеет недопустимый тип",
		})
		return
	}
	if userId > 0 {
		resCheckIsAdmin, statusCode, err := h.services.CheckIsAdmin(userId)
		if err != nil {
			c.JSON(statusCode, map[string]interface{}{
				"status":  statusCode,
				"message": err.Error(),
				"result":  resCheckIsAdmin,
			})
			return
		}
		if resCheckIsAdmin {
			c.JSON(http.StatusOK, map[string]interface{}{
				"status":  http.StatusOK,
				"message": "пользователь является администратором",
				"result":  resCheckIsAdmin,
			})
		}
		if !resCheckIsAdmin {
			c.JSON(http.StatusOK, map[string]interface{}{
				"status":  http.StatusOK,
				"message": "пользователь не является администратором",
				"result":  resCheckIsAdmin,
			})
		}
	}
}
