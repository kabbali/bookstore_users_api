package users

import (
	"github.com/gin-gonic/gin"
	"github.com/kabbali/bookstore_users_api/domain/users"
	"github.com/kabbali/bookstore_users_api/services"
	"github.com/kabbali/bookstore_users_api/utils/errors"
	"net/http"
	"strconv"
)

func CreateUser(c *gin.Context)  {
	user := users.User{}
	//bytes, err := ioutil.ReadAll(c.Request.Body)
	//if err != nil {
	//	//TODO: Handle error
	//	return
	//}
	//if err := json.Unmarshal(bytes, &user); err != nil {
	//	//Handle json error
	//	fmt.Println(err.Error())
	//	return
	//}
	// The above code can be replaced by:
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, savedError := services.CreateUser(user)
	if savedError != nil {
		//create user error
		c.JSON(savedError.Status, savedError)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context)  {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("invalid, user id should be a number")
		c.JSON(err.Status, err)
		return
	}
	user, getError := services.GetUser(userId)
	if getError != nil {
		//create user error
		c.JSON(getError.Status, getError)
		return
	}
	c.JSON(http.StatusOK, user)
	//c.String(http.StatusNotImplemented, "SearchUser not implemented")
}

//func SearchUser(c *gin.Context)  {
//	c.String(http.StatusNotImplemented, "SearchUser not implemented")
//}


