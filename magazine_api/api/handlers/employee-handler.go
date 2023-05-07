package handlers

import (
	"magazine_api/lib"
	"magazine_api/models"
	"magazine_api/orchestrators"
	"magazine_api/services"
	"time"

	"github.com/danhper/structomap"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type EmployeeHandler struct {
	logger            lib.Logger
	service           services.UserService
	user_orchestrator orchestrators.UserOrchestrator
}

func NewEmployeeHandler(logger lib.Logger, service services.UserService, u orchestrators.UserOrchestrator) EmployeeHandler {
	return EmployeeHandler{logger: logger, service: service, user_orchestrator: u}
}

// ListUsers godoc
// @Summary      Lists users
// @Description  List users
// @Tags         Employee
// @Produce      json
// @Success      200   {object}  object{data=[]models.User}
// @Router       /employee [get]
//
// List Employees controller
func (u EmployeeHandler) ListEmployees(c *gin.Context) {
	users, err := u.service.ListEmployees(c)
	if err != nil {
		handleError(u.logger, c, err)
		return
	}

	c.JSON(200, users)
}

// ListDeletedUsers godoc
// @Summary      Lists Deleted users
// @Description  Lists Deleted users
// @Tags         Employee
// @Produce      json
// @Success      200  {object}  object{data=[]models.User}
// @Router       /employee/deleted [get]
//
// List Deleted Employee controller
func (u EmployeeHandler) ListDeletedEmployee(c *gin.Context) {
	users, err := u.service.ListsDeletedUsers()
	if err != nil {
		handleError(u.logger, c, err)
		return
	}

	c.JSON(200, gin.H{"data": users})
}

// GetProfielByEmployeeID godoc
// @Summary      Gets Employee profile by ID
// @Description  Get employee details by employee profile id.
// @Tags         Employee
// @Produce      json
// @Param        id   path      string  true  "ID"
// @Success      200  {object}  object{data=responses.EmployeeAll}
// @Router       /employee/profile/{id} [get]
// Gets Employee By ID controller
func (u EmployeeHandler) GetProfileByID(c *gin.Context) {
	id := c.Param("id")

	user, err := u.service.GetProfileByID(uuid.MustParse(id))
	if err != nil {
		handleError(u.logger, c, err)
		return
	}

	c.JSON(200, gin.H{"data": user})
}

// GetOneUserByEmail godoc
// @Summary      Gets One User by email
// @Description  Gets One user by email
// @Tags         Employee
// @Produce      json
// @Param        email  path      string  true  "Email"  Format(email)
// @Success      200    {object}  object{data=models.User}
// @Router       /employee/email/{email} [get]
// Gets Users By Email controller
func (u EmployeeHandler) GetEmployeeByEmail(c *gin.Context) {
	email := c.Param("email")

	user, err := u.service.GetUserByEmail(email)
	if err != nil {
		handleError(u.logger, c, err)
		return
	}

	c.JSON(200, gin.H{"data": user})
}

// ListUsersByType godoc
// @Summary      Lists users by type
// @Description  List users by type
// @Tags         Employee
// @Produce      json
// @Param        type  path      string  true  "Type"
// @Success      200  {object}  object{data=[]models.User}
// @Router       /employee/type/{type} [get]
//
// List Users By Type controller
func (u EmployeeHandler) ListEmployeesByType(c *gin.Context) {
	userType := c.Param("type")
	users, err := u.service.ListEmployeesByType(c, userType)
	if err != nil {
		handleError(u.logger, c, err)
		return
	}

	c.JSON(200, gin.H{"data": users})
}

// GetOneUserByContactNumber godoc
// @Summary      Gets One User by Contact Number
// @Description  Gets One user by Contact Number
// @Tags         Employee
// @Produce      json
// @Param        contact_number  path      string  true  "Contact Number"
// @Success      200             {object}  object{data=models.User}
// @Router       /employee/contact/{contact} [get]
// Gets Users By Contact Number controller
func (u EmployeeHandler) GetEmployeeByContactNumber(c *gin.Context) {
	contact := c.Param("contact")

	user, err := u.service.GetUserByContactNumber(contact)
	if err != nil {
		handleError(u.logger, c, err)
		return
	}

	c.JSON(200, gin.H{"data": user})
}

// GetOneEmployeeByID godoc
// @Summary      Gets Employee By ID
// @Description  Get employee details by employee profile id.
// @Tags         Employee
// @Produce      json
// @Param        id   path      string  true  "ID"
// @Success      200  {object}  object{data=responses.EmployeeAll}
// @Router       /employee/id/{id} [get]
// Gets Employee By ID controller
func (u EmployeeHandler) GetEmployeeByID(c *gin.Context) {
	id := c.Param("id")

	user, err := u.service.GetUserByID(uuid.MustParse(id))
	if err != nil {
		handleError(u.logger, c, err)
		return
	}

	c.JSON(200, gin.H{"data": user})
}

// DeleteUser godoc
// @Summary      Soft Delete an user
// @Description  Delete by user ID
// @Tags         Employee
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      204  {object}  object{data=string}
// @Router       /employee/{id} [delete]
// Delete Users By ID controller
func (u EmployeeHandler) DeleteEmployeeByID(c *gin.Context) {
	id := c.Param("id")

	err := u.service.DeleteUser(uuid.MustParse(id))
	if err != nil {
		handleError(u.logger, c, err)
		return
	}

	c.JSON(200, gin.H{"data": "successfully deleted"})
}

// UpdateUser godoc
// @Summary      Update an user
// @Description  Updates user attribute
// @Tags         Employee
// @Accept       json
// @Produce      json
// @Param        id    path      string           true  "User ID"
// @Param        user  body      models.UserBase  true  "Update user"
// @Success      200   {object}  object{data=models.User}
// @Router       /employee/{id} [patch]
// Patch Users By ID controller
func (u EmployeeHandler) PatchEmployee(c *gin.Context) {
	id := c.Param("id")

	user, err := u.service.GetUserByID(uuid.MustParse(id))
	if err != nil {
		handleError(u.logger, c, err)
		return
	}

	var newUser models.UserBase
	if err := c.ShouldBindJSON(&newUser); err != nil {
		handleError(u.logger, c, err)
		return
	}

	userMap := structomap.New().UseSnakeCase().PickAll().
		Omit("Password").
		OmitIf(
			func(u interface{}) bool {
				return newUser.Name == nil
			}, "Name",
		).
		OmitIf(
			func(u interface{}) bool {
				return newUser.Email == nil
			}, "Email",
		).
		OmitIf(
			func(u interface{}) bool {
				return newUser.ContactNumber == nil
			}, "ContactNumber",
		).
		OmitIf(
			func(u interface{}) bool {
				return newUser.Role == nil
			}, "Role",
		).
		Transform(newUser)

	if len(userMap) > 0 {
		userMap["updated_on"] = time.Now()
		userMap["id"] = user.ID

		err := u.service.UpdateUser(user.ID, &userMap)
		if err != nil {
			handleError(u.logger, c, err)
			return
		}

		c.JSON(200, gin.H{"data": userMap})
		return
	}

	c.JSON(200, gin.H{"data": "nothing to update"})
}
