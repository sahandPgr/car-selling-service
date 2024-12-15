package migrations

import (
	"github.com/sahandPgr/car-selling-service/constatns"
	"github.com/sahandPgr/car-selling-service/internal/db"
	"github.com/sahandPgr/car-selling-service/internal/models"
	"github.com/sahandPgr/car-selling-service/pkg/logger"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const countStartExp = "count(*)"

// Up_1 is a migration function for creating tables
func Up_1(log logger.Logger) {
	var database = db.GetDB()

	createTables(database, log)
	createDefault(database)
	createCountry(database)

}

// This function create tables if not exists
func addToTables(database *gorm.DB, models ...interface{}) []interface{} {
	tables := []interface{}{}
	for _, m := range models {
		if !database.Migrator().HasTable(m) {
			tables = append(tables, m)
		}
	}
	return tables
}

func createTables(database *gorm.DB, log logger.Logger) {
	tables := addToTables(database, models.Country{}, models.City{}, models.User{},
		models.Role{}, models.UserRole{}, models.File{}, models.PersianYear{}, models.PropertyCategory{},
		models.Property{}, models.Company{}, models.Gearbox{}, models.Color{}, models.CarType{},
		models.CarModel{}, models.CarModelColor{}, models.CarModelYear{}, models.CarModelImage{},
		models.CarModelComment{}, models.CarModelPriceHistory{}, models.CarModelProperty{},
	)
	if err := database.Migrator().CreateTable(tables...); err != nil {
		log.Fatal(logger.Potgres, logger.Migration, nil, err.Error())
	}
	log.Info(logger.Potgres, logger.Migration, nil, "Migration Up_1 is done.")
}

// This function excecute the default data
func createDefault(database *gorm.DB) {
	adminRole := models.Role{Name: constatns.AdminRole}
	createRoleIfNotExists(database, &adminRole)

	difaultRole := models.Role{Name: constatns.DefaultUserRole}
	createRoleIfNotExists(database, &difaultRole)

	adminUser := models.User{Username: constatns.AdminName, FirstName: "test", LastName: "test", Email: "test@gmail.com",
		PhoneMobile: "09142535335"}
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte("12345678"), bcrypt.DefaultCost)
	adminUser.Password = string(hashPassword)
	createAdminUserIfNotExists(database, &adminUser, adminRole.ID)
}

// This function create role if not exists
func createRoleIfNotExists(database *gorm.DB, r *models.Role) {
	exsits := 0
	database.Model(&models.Role{}).
		Select("1").
		Where("name = ?", r.Name).
		First(&exsits)
	if exsits == 0 {
		database.Create(r)
	}
}

// This function create admin user if not exists
func createAdminUserIfNotExists(database *gorm.DB, u *models.User, roleId int) {
	exsits := 0
	database.Model(&models.User{}).
		Select("1").
		Where("username = ?", u.Username).
		First(&exsits)
	if exsits == 0 {
		database.Create(u)
		ur := models.UserRole{UserId: u.ID, RoleId: roleId}
		database.Create(&ur)
	}
}

func createCountry(database *gorm.DB) {
	count := 0
	database.Model(&models.Country{}).
		Select(countStartExp).
		Find(&count)
	if count == 0 {
		database.Create(&models.Country{Name: "Iran", Cities: []models.City{
			{Name: "Tehran"},
			{Name: "Isfahan"},
			{Name: "Shiraz"},
			{Name: "Chalus"},
			{Name: "Ahwaz"},
		}})
		database.Create(&models.Country{Name: "USA", Cities: []models.City{
			{Name: "New York"},
			{Name: "Washington"},
		}})
		database.Create(&models.Country{Name: "Germany", Cities: []models.City{
			{Name: "Berlin"},
			{Name: "Munich"},
		}})
		database.Create(&models.Country{Name: "China", Cities: []models.City{
			{Name: "Beijing"},
			{Name: "Shanghai"},
		}})
		database.Create(&models.Country{Name: "Italy", Cities: []models.City{
			{Name: "Roma"},
			{Name: "Turin"},
		}})
		database.Create(&models.Country{Name: "France", Cities: []models.City{
			{Name: "Paris"},
			{Name: "Lyon"},
		}})
		database.Create(&models.Country{Name: "Japan", Cities: []models.City{
			{Name: "Tokyo"},
			{Name: "Kyoto"},
		}})
		database.Create(&models.Country{Name: "South Korea", Cities: []models.City{
			{Name: "Seoul"},
			{Name: "Ulsan"},
		}})
	}

}
