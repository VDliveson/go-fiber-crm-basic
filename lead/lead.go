package lead

import (
	"github.com/VDliveson/go-fiber-crm-basic/database"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Lead struct {
	gorm.Model
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   int    `json:"phone"`
}

func GetLeads(c *fiber.Ctx) error {
	db := database.DBconn
	var leads []Lead
	db.Find(&leads)
	c.JSON(leads)
	return nil
}

func GetLead(c *fiber.Ctx) error {
	id := c.Params(":id")
	db := database.DBconn
	var lead Lead
	db.Find(&lead, id)
	c.JSON(lead)
	return nil
}

func NewLead(c *fiber.Ctx) error{
	db := database.DBconn
	lead := new(Lead)
	if err := c.BodyParser(lead); err != nil {
		c.Status(503).Send([]byte(err.Error()))
		return nil
	}
	db.Create(&lead)
	c.JSON(lead)
	return nil
}

func DeleteLead(c *fiber.Ctx) error{
	id := c.Params(":id")
	db := database.DBconn
	var lead Lead
	db.First(&lead, id)
	if lead.Name == "" {
		c.Status(500).SendString("No lead found with ID")
		return nil
	}
	db.Delete(&lead)
	c.SendString("Lead successfully deleted")
	return nil
}

// func UpdateLead(c *fiber.Ctx) error{
// 	id:= c.Params(":id")
// 	db := database.DBconn
// 	var lead Lead

// 	db.First(&lead, id)
// 	if lead.Name == "" {
// 		c.Status(500).SendString("No lead found with ID")
// 		return nil
// 	}

// 	newlead := new(Lead)
// 	if err := c.BodyParser(newlead); err != nil {
// 		c.Status(503).Send([]byte(err.Error()))
// 		return nil
// 	}

// 	db.Update(&newlead, lead)
// }