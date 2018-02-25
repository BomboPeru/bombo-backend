package api

import (
	"errors"
	"log"

	"github.com/kataras/iris"
	"github.com/satori/go.uuid"
	"../helpers"
	"github.com/k0kubun/pp"
)


// LinkWithUserType ...
func LinkWithUserType(api iris.Party) {
	api.Get("/user/{id:string}", func(c iris.Context) {
		userID := c.Params().Get("id")
		ID, err := uuid.FromString(userID)
		pp.Println(ID)
		if err != nil {
			log.Println("uuid.FromString(), ", err)
			c.StatusCode(iris.StatusInternalServerError)
			c.JSON(Response{
				Data:  nil,
				Error: err.Error(),
			})
			return
		}

		user, err := helpers.GetUserByID(ID)
		if err != nil {
			log.Println("GetUserByID(), ", err)
			c.StatusCode(iris.StatusInternalServerError)
			c.JSON(Response{
				Data:  nil,
				Error: err.Error(),
			})
		} else {
			c.StatusCode(iris.StatusOK)
			c.JSON(Response{
				Data:  user,
				Error: nil,
			})
		}

	})

	api.Get("/user/all", func(c iris.Context) {
		users, err := helpers.GetAllUsers()
		if err != nil {
			log.Println("GetAllUsers(), ", err)
			c.StatusCode(iris.StatusInternalServerError)
			c.JSON(Response{
				Data:  nil,
				Error: err.Error(),
			})
		} else {
			c.StatusCode(iris.StatusOK)
			c.JSON(Response{
				Data:  users,
				Error: nil,
			})
		}

	})

	api.Post("/user/create", func(c iris.Context) {
		user := new(helpers.User)
		err := c.ReadForm(user)
		if err != nil {
			log.Println("c.ReadForm(), ", err)
			c.StatusCode(iris.StatusInternalServerError)
			c.JSON(Response{
				Data:  nil,
				Error: err.Error(),
			})
			return
		}

		// If not have Id
		if uuid.Equal(user.ID, uuid.Nil){
			newID, err := uuid.NewV4()
			if err != nil {
				log.Println("uuid.NewV4(), ", err)
				c.StatusCode(iris.StatusInternalServerError)
				c.JSON(Response{
					Data:  nil,
					Error: err.Error(),
				})
				return
			}
			user.ID = newID
		}

		returnedUser, err := helpers.CreateNewUser(user)
		if err != nil {
			log.Println("CreateNewUser(), ", err)
			c.StatusCode(iris.StatusInternalServerError)
			c.JSON(Response{
				Data:  nil,
				Error: err.Error(),
			})
			return
		}

		c.StatusCode(iris.StatusOK)
		c.JSON(Response{
			Data:  returnedUser,
			Error: nil,
		})
	})

	api.Post("/user/update", func(c iris.Context) {
		user := new(helpers.User)
		err := c.ReadForm(user)
		if err != nil {
			log.Println("c.ReadForm(), ", err)
			c.StatusCode(iris.StatusInternalServerError)
			c.JSON(Response{
				Data:  nil,
				Error: err.Error(),
			})
			return
		}

		// If not have Id
		if user.ID.String() == "" {
			log.Println("Update method needs Id")
			c.StatusCode(iris.StatusInternalServerError)
			c.JSON(Response{
				Data:  nil,
				Error: errors.New("Update method needs Id").Error(),
			})
			return
		}

		updatedUser, err := helpers.UpdateUser(user)
		if err != nil {
			log.Println("UpdateUser(), ", err)
			c.StatusCode(iris.StatusInternalServerError)
			c.JSON(Response{
				Data:  nil,
				Error: err.Error(),
			})
			return
		}

		c.StatusCode(iris.StatusOK)
		c.JSON(Response{
			Data:  updatedUser,
			Error: nil,
		})

	})

	api.Post("/user/delete", func(c iris.Context) {
		type IDForm struct {
			ID uuid.UUID `json:"id"`
		}

		idForm := new(IDForm)
		err := c.ReadForm(idForm)
		if err != nil {
			log.Println("c.ReadForm(), ", err)
			c.StatusCode(iris.StatusInternalServerError)
			c.JSON(Response{
				Data:  nil,
				Error: err.Error(),
			})
			return
		}

		returnedUser, err := helpers.DeleteUserByID(idForm.ID)

		if err != nil {
			c.StatusCode(iris.StatusInternalServerError)
			c.JSON(Response{
				Data:  nil,
				Error: err.Error(),
			})
		} else {
			c.StatusCode(iris.StatusOK)
			c.JSON(Response{
				Data:  returnedUser,
				Error: nil,
			})
		}
	})
}
