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
	router.GET("/getusersbytwoparameters/:first_name/:city", Getusersbytwopara)

	router.DELETE("/calld/:id", DeleteTasks)
	router.DELETE("/calld1/:city/:first_name", DeleteTasksbytwopara)
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
	var get []models.Info
	models.DB.Preload("Person").Preload("Address").Find(&get)

	c.IndentedJSON(http.StatusOK, gin.H{"called upon": get})
}

func getUsersByEmail(c *gin.Context) {
	var email []models.Info
	res := models.DB.Preload("Address").Preload("Person", "email =?", c.Param("email")).Find(&email)

	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"dbdata": email})
	}
}

func getusersbyfirstname(c *gin.Context) {
	var firstname []models.Info
	res := models.DB.Preload("Address").Preload("Person", "firstname = ?", c.Param("firstname")).Find(&firstname)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"dbdata": firstname})
	}
}

func getusersbyaddress(c *gin.Context) {
	var address []models.Info
	res := models.DB.Preload("Address").Preload("Person", "address = ?", c.Param("address")).Find(&address)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"dbdata": address})
	}
}

func getusersbylastname(c *gin.Context) {
	var lastname []models.Info
	res := models.DB.Preload("Address").Preload("Person", "lastname = ?", c.Param("last_name")).Find(&lastname)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"dbdata": lastname})
	}
}

func Getusersbytwopara(c *gin.Context) {
	var getbytwo []models.Info
	res := models.DB.Preload("Address").Preload("Person").Raw(`SELECT i.*,p.*,a.* FROM info as i LEFT JOIN Person as p ON i.ClientID = p.ID LEFT JOIN Address as a ON i.ClientID = a.id Where p.firstname = ? and a.city = ? `, c.Param("firstname"), c.Param("city")).Find(&getbytwo)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"dbdata": getbytwo})
	}
}

func PostTasks(c *gin.Context) {
	var post models.Info
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//kw := models.CallingTasks{Data: &todo.Data}
	models.DB.Preload("Address").Preload("Person").Create(&post)
	//c.IndentedJSON(http.StatusOK, gin.H{"message": "ho gya"})
	GetTasks(c)
}

func DeleteTasks(c *gin.Context) {
	var del models.Info
	if err := models.DB.Preload("Address").Preload("Person", "id = ?", c.Param("id")).First(&del).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	models.DB.Delete(&del)

	GetTasks(c)
}

func DeleteTasksbytwopara(c *gin.Context) {
	var Delbytwopara models.Info
	if err := models.DB.Preload("Address").Preload("Person", "city = ? and firstname = ?", c.Param("city"), c.Param("first_name")).Find(&Delbytwopara).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	models.DB.Delete(&Delbytwopara)

	GetTasks(c)
}

func DeleteTasksbyfirstname(c *gin.Context) {
	var delfir models.Info
	if err := models.DB.Preload("Address").Preload("Person", "firstname = ?", c.Param("first_name")).First(&delfir).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	models.DB.Delete(&delfir)

	GetTasks(c)
}

func DeleteTaskslastname(c *gin.Context) {
	var dellas models.Info
	if err := models.DB.Preload("Address").Preload("Person", "lastname = ?", c.Param("last_name")).First(&dellas).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	models.DB.Delete(&dellas)

	GetTasks(c)
}

func DeleteTasksbyemail(c *gin.Context) {
	var delem models.Info
	if err := models.DB.Preload("Address").Preload("Person", "email = ?", c.Param("email")).First(&delem).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	models.DB.Delete(&delem)

	GetTasks(c)
}
func UpdateTasks(c *gin.Context) {

	var change models.Info
	var input models.Info
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.DB.Model(&change).Preload("Address").Preload("Person").Where("id = ?", c.Param("id")).First(&change).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	models.DB.Model(&change).Updates(input)
	GetTasks(c)
}

func UpdateTasksbyaddress(c *gin.Context) {

	var changebyad models.Info
	if err := models.DB.Preload("Address").Preload("Person", "address = ?", c.Param("address")).First(&changebyad).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var inputbyad models.Info
	if err := c.ShouldBindJSON(&inputbyad); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.DB.Model(&changebyad).Updates(inputbyad)

}

func UpdateTasksbyfirstname(c *gin.Context) {

	var changebyf models.Info
	if err := models.DB.Preload("Address").Preload("Person", "first_name = ?", c.Param("first_name")).First(&changebyf).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var inputbyf models.Info
	if err := c.ShouldBindJSON(&inputbyf); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.DB.Model(&changebyf).Updates(inputbyf)

	//GetTasks(c)
}

func UpdateTasksbylastname(c *gin.Context) {

	var changebyl models.Info
	if err := models.DB.Preload("Address").Preload("Person", "last_name = ?", c.Param("last_name")).First(&changebyl).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var inputbyl models.Info
	if err := c.ShouldBindJSON(&inputbyl); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.DB.Model(&changebyl).Updates(inputbyl)

	//GetTasks(c)
	//c.IndentedJSON(http.StatusAccepted,gin.H{"UPDATED BY LAST NAME":"update by last name "})
}

func UpdateTasksbyemail(c *gin.Context) {

	var changebye models.Info
	if err := models.DB.Preload("Address").Preload("Person", "email = ?", c.Param("email")).First(&changebye).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var inpute models.Info
	if err := c.ShouldBindJSON(&inpute); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.DB.Model(&changebye).Updates(inpute)

	//GetTasks(c)
}
