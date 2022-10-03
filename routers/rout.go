package routers

import (
	//"fmt"
	"calling/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes() *gin.Engine {
	router := gin.Default()
	//router.Use(middleware.CORSMiddleware())

	router.POST("/callp", PostTasks)
	router.GET("/callg", GetTasks)
	router.GET("/callg/:email", getUsersByEmail)
	router.GET("/getusersbyfirstname/:firstname", getusersbyfirstname)

	router.DELETE("/callg/:id", DeleteTasks)
	router.PUT("/callpu/:id", UpdateTasks)
	return router
}

/*type work struct {
	Data string `json:"data"`
    FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Address   string `json:"address"`
	Email     string `json:"email"`
}*/
func GetTasks(c *gin.Context) {
	var get []models.CallingTasks
	models.DB.Find(&get)

	c.IndentedJSON(http.StatusOK, gin.H{"called upon": get})
}

func getUsersByEmail(c *gin.Context) {
	var email []models.CallingTasks
	res := models.DB.Where("email =?", c.Param("email")).Find(&email)

	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"dbdata": email})
	}
}

func getusersbyfirstname(c *gin.Context) {
	var first []models.CallingTasks
	res := models.DB.Where("first_name = ?", c.Param("firstname")).Find(&first)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"dbdata": first})
	}
}

func PostTasks(c *gin.Context) {
	var post models.CallingTasks
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//kw := models.CallingTasks{Data: &todo.Data}
	models.DB.Create(&post)
	//c.IndentedJSON(http.StatusOK, gin.H{"message": "ho gya"})
	GetTasks(c)
}

func DeleteTasks(c *gin.Context) {
	var del models.CallingTasks
	if err := models.DB.Where("id = ?", c.Param("id")).First(&del).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	models.DB.Delete(&del)

	GetTasks(c)
}

func UpdateTasks(c *gin.Context) {

	var change models.CallingTasks
	if err := models.DB.Where("id = ?", c.Param("id")).First(&change).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input models.CallingTasks
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.DB.Model(&change).Updates(input)

	GetTasks(c)
}
