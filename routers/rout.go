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
	router.GET("/getbyemail/:email", getUsersByEmail)
	router.GET("/getusersbyfirstname/:firstname", getusersbyfirstname)
	router.GET("/getbyaddress/:address", getusersbyaddress)
	router.GET("/callg1/:la", getusersbylastname)

	router.DELETE("/calld/:id", DeleteTasks)
	router.DELETE("/calld1/:address", DeleteTasksbyaddress)
	router.DELETE("/calld2/:first_name", DeleteTasksbyfirstname)
	router.DELETE("/calld3/:last_name", DeleteTaskslastname)
	router.DELETE("/calld4/:email", DeleteTasksbyemail)

	router.PUT("/callpu/:id", UpdateTasks)
	router.PUT("/callpu1/:address", UpdateTasksbyaddress)
	router.PUT("/callpu2/:first_name", UpdateTasksbyfirstname)
	router.PUT("/callpu3/:last_name", UpdateTasksbylastname)
	router.PUT("/callpu4/:address", UpdateTasksbyaddress)

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
	models.DB.Preload("Address").Find(&get)

	c.IndentedJSON(http.StatusOK, gin.H{"called upon": get})
}

func getUsersByEmail(c *gin.Context) {
	var email []models.CallingTasks
	res := models.DB.Preload("Address").Where("email =?", c.Param("email")).Find(&email)

	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"dbdata": email})
	}
}

func getusersbyfirstname(c *gin.Context) {
	var firstname []models.CallingTasks
	res := models.DB.Preload("Address").Where("first_name = ?", c.Param("firstname")).Find(&firstname)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"dbdata": firstname})
	}
}

func getusersbyaddress(c *gin.Context) {
	var address []models.CallingTasks
	res := models.DB.Preload("Address").Where("address = ?", c.Param("address")).Find(&address)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"dbdata": address})
	}
}

func getusersbylastname(c *gin.Context) {
	var lastname []models.CallingTasks
	res := models.DB.Preload("Address").Where("last_name = ?", c.Param("last_name")).Find(&lastname)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"dbdata": lastname})
	}
}

func PostTasks(c *gin.Context) {
	var post models.CallingTasks
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//kw := models.CallingTasks{Data: &todo.Data}
	models.DB.Preload("Address").Create(&post)
	//c.IndentedJSON(http.StatusOK, gin.H{"message": "ho gya"})
	GetTasks(c)
}

func DeleteTasks(c *gin.Context) {
	var del models.CallingTasks
	if err := models.DB.Preload("Address").Where("id = ?", c.Param("id")).First(&del).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	models.DB.Delete(&del)

	GetTasks(c)
}

func DeleteTasksbyaddress(c *gin.Context) {
	var deladdress models.CallingTasks
	if err := models.DB.Preload("Address").Where("state = ?", c.Param("address")).Find(&deladdress).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	models.DB.Delete(&deladdress)

	GetTasks(c)
}

func DeleteTasksbyfirstname(c *gin.Context) {
	var delfir models.CallingTasks
	if err := models.DB.Preload("Address").Where("first_name = ?", c.Param("first_name")).First(&delfir).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	models.DB.Delete(&delfir)

	GetTasks(c)
}

func DeleteTaskslastname(c *gin.Context) {
	var dellas models.CallingTasks
	if err := models.DB.Preload("Address").Where("last_name = ?", c.Param("last_name")).First(&dellas).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	models.DB.Delete(&dellas)

	GetTasks(c)
}

func DeleteTasksbyemail(c *gin.Context) {
	var delem models.CallingTasks
	if err := models.DB.Preload("Address").Where("email = ?", c.Param("email")).First(&delem).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	models.DB.Delete(&delem)

	GetTasks(c)
}
func UpdateTasks(c *gin.Context) {

	var change models.CallingTasks
	var input models.CallingTasks
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.DB.Model(&change).Preload("Address").Where("id = ?", c.Param("id")).First(&change).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	models.DB.Model(&change).Updates(input)
	GetTasks(c)
}

func UpdateTasksbyaddress(c *gin.Context) {

	var changebyad models.CallingTasks
	if err := models.DB.Preload("Address").Where("address = ?", c.Param("address")).First(&changebyad).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var inputbyad models.CallingTasks
	if err := c.ShouldBindJSON(&inputbyad); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.DB.Model(&changebyad).Updates(inputbyad)

}

func UpdateTasksbyfirstname(c *gin.Context) {

	var changebyf models.CallingTasks
	if err := models.DB.Preload("Address").Where("first_name = ?", c.Param("first_name")).First(&changebyf).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var inputbyf models.CallingTasks
	if err := c.ShouldBindJSON(&inputbyf); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.DB.Model(&changebyf).Updates(inputbyf)

	//GetTasks(c)
}

func UpdateTasksbylastname(c *gin.Context) {

	var changebyl models.CallingTasks
	if err := models.DB.Preload("Address").Where("last_name = ?", c.Param("last_name")).First(&changebyl).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var inputbyl models.CallingTasks
	if err := c.ShouldBindJSON(&inputbyl); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.DB.Model(&changebyl).Updates(inputbyl)

	//GetTasks(c)
	//c.IndentedJSON(http.StatusAccepted,gin.H{"UPDATED BY LAST NAME":"update by last name "})
}

func UpdateTasksbyemail(c *gin.Context) {

	var changebye models.CallingTasks
	if err := models.DB.Preload("Address").Where("email = ?", c.Param("email")).First(&changebye).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var inpute models.CallingTasks
	if err := c.ShouldBindJSON(&inpute); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.DB.Model(&changebye).Updates(inpute)

	//GetTasks(c)
}
