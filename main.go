package main

import (
	"bwastartup/auth"
	"bwastartup/handler"
	"bwastartup/user"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Connection to database is good")

	// call func NewRepository dari /user/repository
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	authService := auth.NewService()

	// fmt.Println(authService.GenerateToken(1000))

	// userService.SaveAvatar(1, "images/1-profile.png")

	// test service
	// input := user.LoginInput{
	// 	Email:    "opick@email.com",
	// 	Password: "password",
	// }
	// user, err := userService.Login(input)
	// if err != nil {
	// 	fmt.Println("Terjadi kesalahan")
	// 	fmt.Println(err.Error())
	// }

	// fmt.Println(user.Email)
	// fmt.Println(user.Name)

	// test repository
	// userByEmail, err := userRepository.FindByEmail("opick@email.com")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// // fmt.Println(userByEmail.Name)
	// if userByEmail.ID == 0 {
	// 	fmt.Println("User tidak ditemukan")
	// } else {
	// 	fmt.Println(userByEmail.Name)
	// }

	userHandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatars", userHandler.UploadAvatar)

	router.Run()

	// userInput := user.RegisterUserInput{}
	// userInput.Name = "Tes simpan dari service"
	// userInput.Email = "test@email.com"
	// userInput.Occupation = "Desaigner"
	// userInput.Password = "password"

	// userService.RegisterUser(userInput)

	// user := user.User{
	// 	Name: "Test simpan",
	// }

	// userRepository.Save(user)

	// input dari user
	// handler -> mapping input dari user -> struct input
	// service -> mapping dari struct input ke struct user
	// repository -> transaksi ke db

	// var users []user.User // variable users dengan tipe array of entity struct User mewakili table users
	// length := len(users)
	// fmt.Println(length) // hasilnya 0

	// find data users
	// db.Find(&users)
	// length = len(users)
	// fmt.Println(length) // hasilnya 2 (sesuai data di db)

	// looping hasil find data users
	// for _, user := range users {
	// 	fmt.Println(user.Name)
	// 	fmt.Println(user.Email)
	// 	fmt.Println("=============")
	// }

	// router := gin.Default()
	// router.GET("/handler", handler) // mengakses /handler akan memanggil func handler
	// router.Run()
}

// func handler(c *gin.Context) {
// 	dsn := "root:@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}

// 	var users []user.User
// 	db.Find(&users) // &users -> menandakan pointer dari variable users

// 	c.JSON(http.StatusOK, users)

// 	/*
// 		flow yg diakses
// 		- input data
// 		- handler -> menangkap inputan, kemudian dimapping datanya ke struct
// 		- service -> dimapping ke struct User
// 		- repository -> save struct User ke db
// 	*/
// }
