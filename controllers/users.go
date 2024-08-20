package controllers

import (
	"backend/lib"
	"backend/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ListAllUsers(ctx *gin.Context){
	results := models.FindAllUsers()
	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "List all users",
		Results:results,
	})
}

func DetailUsers(ctx *gin.Context){
	id, _ := strconv.Atoi(ctx.Param("id"))
	data := models.FindOneUsers(id)

	if data.Id != 0 {
		ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "Users OK",
		Results: data,
	})
	}else{
		ctx.JSON(http.StatusNotFound, lib.Response{
		Success: false,
		Message: "Users Not Found",
		Results:data,
	})
	}
	
}

func CreateUsers(ctx *gin.Context) {
     newUser := models.Users{}

    if err := ctx.ShouldBind(&newUser); err != nil {
        ctx.JSON(http.StatusBadRequest, lib.Response{
            Success: false,
            Message: "Invalid input data",
        })
        return
    }

    err := models.CreateUser(newUser)
    if err != nil {
		ctx.JSON(http.StatusInternalServerError, lib.Response{
			Success: false,
			Message: "Failed to create user",
		})
		return
    }
	ctx.JSON(http.StatusOK, lib.Response{
   Success: true,
   Message: "User created successfully",
   Results: newUser,
})
	
   
}

func DeleteUsers(ctx *gin.Context){
    id, err := strconv.Atoi(ctx.Param("id"))
    dataUser := models.FindOneUsers(id)

    if err != nil {
        ctx.JSON(http.StatusBadRequest, lib.Response{
            Success: false,
            Message: "Invalid user ID",
        })
        return
    }

    err = models.DeleteUsers(id)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, lib.Response{
            Success: false,
            Message: "Id Not Found",
        })
        return
    }

    ctx.JSON(http.StatusOK, lib.Response{
        Success: true,
        Message: "User deleted successfully",
        Results: dataUser,
    })
}

func UpdateUser(c *gin.Context) {
    param := c.Param("id")
    id, _ := strconv.Atoi(param)
    data := models.FindAllUsers()

    user := models.Users{}
    err := c.Bind(&user)
    if err != nil {
        fmt.Println(err)
        return
    }

    result := models.Users{}
    for _, v := range data {
        if v.Id == id {
            result = v
        }
    }

    if result.Id == 0 {
        c.JSON(http.StatusNotFound, lib.Response{
            Success: false,
            Message: "user whit id " + param + " not found",
        })
        return
    }

    models.EditUser(user.Email, user.Username, user.Password, param)

    c.JSON(http.StatusOK, lib.Response{
        Success: true,
        Message: "user whit id " + param + " Edit Success",
        Results: user,
    })
}
