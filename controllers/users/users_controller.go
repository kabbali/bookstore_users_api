package users

import (
	"github.com/gin-gonic/gin"
	"github.com/kabbali/bookstore_users_api/domain/users"
	"github.com/kabbali/bookstore_users_api/services"
	"github.com/kabbali/bookstore_users_api/utils/errors"
	"net/http"
	"strconv"
)

func getUserId(userIdParam string) (int64, *errors.RestErr) {
	userId, idErr := strconv.ParseInt(userIdParam, 10, 64)
	if idErr != nil {
		return 0, errors.NewBadRequestError("invalid, user id should be a number")
	}
	return userId, nil
}

func Create(c *gin.Context)  {
	user := users.User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, savedError := services.CreateUser(user)
	if savedError != nil {
		c.JSON(savedError.Status, savedError)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func Get(c *gin.Context)  {
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	user, getError := services.GetUser(userId)
	if getError != nil {
		c.JSON(getError.Status, getError)
		return
	}
	c.JSON(http.StatusOK, user)
}

func Update(c *gin.Context) {
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	user := users.User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	user.Id = userId

	isPartial := c.Request.Method == http.MethodPatch

	result, updateError := services.UpdateUser(isPartial, user)
	if updateError != nil {
		//create user error
		c.JSON(updateError.Status, updateError)
		return
	}
	c.JSON(http.StatusOK, result)
}

func Delete(c *gin.Context)  {
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	if err := services.DeleteUser(userId); err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

func Search(c *gin.Context)  {
	status := c.Query("status")
	result, err := services.Search(status)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, result)
}


