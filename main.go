package main

import (
	"crud_api/modules/account"
	"crud_api/modules/customers"
	"crud_api/utility"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type User struct {
	gorm.Model
	//kalau mau langsung delete ke database, pakai property sendiri, jangan pakai go model
	//ID        uint `gorm:"primarykey"`
	//CreatedAt time.Time
	//UpdatedAt time.Time
	Name string `json:"name"`
}

var db *gorm.DB
var err error

func initDB() (*gorm.DB, error) {
	//default
	dsn := "root@tcp(localhost:3306)/db_project?parseTime=true"
	//dsn := "root:password@tcp(db:3306)/app?parseTime=true"
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

func createUser(c *gin.Context) {
	var user User

	// Baca data JSON dari body permintaan
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Simpan data ke database
	err := db.Create(&user).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Tampilkan respons berhasil
	c.JSON(http.StatusOK, gin.H{"message": "User created successfully", "user": user})
}
func getUsers(c *gin.Context) {
	var users []User

	// Dapatkan semua data user dari database
	if err := db.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Tampilkan data user
	c.JSON(http.StatusOK, gin.H{"users": users})
}

func getUserById(c *gin.Context) {
	var user User
	userID := c.Param("id")

	// Dapatkan data user dari database berdasarkan ID
	if err := db.First(&user, userID).Error; err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		//return
	}

	// Tampilkan data user
	c.JSON(http.StatusOK, gin.H{"user": user})
}
func updateUser(c *gin.Context) {
	var user User
	userID := c.Param("id")

	// Dapatkan data user dari database berdasarkan ID
	if err := db.First(&user, userID).Error; err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Baca data JSON dari body permintaan
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Simpan perubahan ke database
	if err := db.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Tampilkan respons berhasil
	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully", "user": user})
}
func deleteUser(c *gin.Context) {
	var user User
	userID := c.Param("id")

	// Dapatkan data user dari database berdasarkan ID
	if err := db.First(&user, userID).Error; err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Hapus data user dari database
	if err := db.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Tampilkan respons berhasil
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/users", createUser)
	r.GET("/users", getUsers)
	r.GET("/users/:id", getUserById)
	r.PUT("/users/:id", updateUser)
	r.DELETE("/users/:id", deleteUser)

	return r
}

func main() {
	db, err := initDB()
	if err != nil {
		log.Fatalln("initDB:", err)
	}

	r := gin.Default()
	customersHandler := customers.DefaultRequestHandler(db)
	adminsHandler := account.DefaultRequestHandler(db)
	//Customer Route
	CustomerPath := "/customers"
	CustomerPathGroup := r.Group(CustomerPath, utility.AdminAuth)
	CustomerPathGroup.GET("/", customersHandler.Read)
	CustomerPathGroup.GET("/:id", customersHandler.ReadID)
	CustomerPathGroup.POST("/", customersHandler.Create)
	CustomerPathGroup.PUT("/:id", customersHandler.Update)
	CustomerPathGroup.DELETE("/:id", customersHandler.Delete)
	//Admin
	AdminPath := "/admin"
	AdminPathGroup := r.Group(AdminPath, utility.SuperAdminAuth)
	AdminPathGroup.GET("/", adminsHandler.Read)
	AdminPathGroup.GET("/:id", adminsHandler.ReadID)
	AdminPathGroup.POST("", adminsHandler.Create)
	AdminPathGroup.PUT("/:id", adminsHandler.Update)
	AdminPathGroup.PUT("/approval/:id", adminsHandler.Approval)
	AdminPathGroup.PUT("/active/:id", adminsHandler.Activate)
	AdminPathGroup.DELETE("/:id", adminsHandler.Delete)
	r.POST("/admins/login", adminsHandler.Login)

	//request-handler : menerima request, mengirim response
	//controller : validasi dan transformasi data
	//usecase : pemrosesan data
	//repository : presistensi data

	err = r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
