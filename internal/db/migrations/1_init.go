package migrations

import (
	"github.com/sahandPgr/car-selling-service/constatns"
	"github.com/sahandPgr/car-selling-service/internal/db"
	"github.com/sahandPgr/car-selling-service/internal/models"
	"github.com/sahandPgr/car-selling-service/pkg/logger"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const countStarExp = "count(*)"

// Up_1 is a migration function for creating tables
func Up_1(log logger.Logger) {
	var database = db.GetDB()

	createTables(database, log)
	createDefault(database)
	createCountry(database)
	createPropertyCategory(database)
	createCarType(database)
	createGearbox(database)

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
		Select(countStarExp).
		Find(&count)
	if count == 0 {
		database.Create(&models.Country{Name: "Iran", Cities: []models.City{
			{Name: "Tehran"},
			{Name: "Isfahan"},
			{Name: "Shiraz"},
			{Name: "Chalus"},
			{Name: "Ahwaz"},
		}, Companies: []models.Company{
			{Name: "Saipa"},
			{Name: "Iran khodro"},
		}})
		database.Create(&models.Country{Name: "USA", Cities: []models.City{
			{Name: "New York"},
			{Name: "Washington"},
		}, Companies: []models.Company{
			{Name: "Tesla"},
			{Name: "Jeep"},
		}})
		database.Create(&models.Country{Name: "Germany", Cities: []models.City{
			{Name: "Berlin"},
			{Name: "Munich"},
		}, Companies: []models.Company{
			{Name: "Opel"},
			{Name: "Benz"},
		}})
		database.Create(&models.Country{Name: "China", Cities: []models.City{
			{Name: "Beijing"},
			{Name: "Shanghai"},
		}, Companies: []models.Company{
			{Name: "Chery"},
			{Name: "Geely"},
		}})
		database.Create(&models.Country{Name: "Italy", Cities: []models.City{
			{Name: "Roma"},
			{Name: "Turin"},
		}, Companies: []models.Company{
			{Name: "Ferrari"},
			{Name: "Fiat"},
		}})
		database.Create(&models.Country{Name: "France", Cities: []models.City{
			{Name: "Paris"},
			{Name: "Lyon"},
		}, Companies: []models.Company{
			{Name: "Renault"},
			{Name: "Bugatti"},
		}})
		database.Create(&models.Country{Name: "Japan", Cities: []models.City{
			{Name: "Tokyo"},
			{Name: "Kyoto"},
		}, Companies: []models.Company{
			{Name: "Toyota"},
			{Name: "Honda"},
		}})
		database.Create(&models.Country{Name: "South Korea", Cities: []models.City{
			{Name: "Seoul"},
			{Name: "Ulsan"},
		}, Companies: []models.Company{
			{Name: "Kia"},
			{Name: "Hyundai"},
		}})
	}

}

func createPropertyCategory(database *gorm.DB) {
	count := 0
	database.Model(&models.PropertyCategory{}).
		Select(countStarExp).
		Find(&count)
	if count == 0 {
		database.Create(&models.PropertyCategory{Name: "Body"})
		database.Create(&models.PropertyCategory{Name: "Engine"})
		database.Create(&models.PropertyCategory{Name: "Drivetrain"})
		database.Create(&models.PropertyCategory{Name: "Suspension"})
		database.Create(&models.PropertyCategory{Name: "Equipment"})
		database.Create(&models.PropertyCategory{Name: "Driver support systems"})
		database.Create(&models.PropertyCategory{Name: "Lights"})
		database.Create(&models.PropertyCategory{Name: "Multimedia"})
		database.Create(&models.PropertyCategory{Name: "Safety equipment"})
		database.Create(&models.PropertyCategory{Name: "Seats and steering wheel"})
		database.Create(&models.PropertyCategory{Name: "Windows and mirrors"})
	}

	var categories []string
	database.Model(&models.PropertyCategory{}).
		Select("name").
		Find(&categories)
	for _, cat := range categories {
		createProperty(database, cat)
	}
}

func createProperty(database *gorm.DB, cat string) {
	count := 0
	catModel := new(models.PropertyCategory)

	database.Model(models.PropertyCategory{}).
		Where("name = ?", cat).
		Find(catModel)

	database.Model(&models.Property{}).
		Select(countStarExp).
		Where("category_id = ?", catModel.ID).
		Find(&count)
	if count > 0 || catModel.ID == 0 {
		return
	}
	var properties *[]models.Property
	switch cat {
	case "Body":
		properties = getBodyProperties(catModel.ID)
	case "Engine":
		properties = getEngineProperties(catModel.ID)
	case "Drivetrain":
		properties = getDrivetrainProperties(catModel.ID)
	case "Suspension":
		properties = getSuspensionProperties(catModel.ID)
	case "Comfort":
		properties = getComfortProperties(catModel.ID)
	case "Driver support systems":
		properties = getDriverSupportSystemProperties(catModel.ID)
	case "Lights":
		properties = getLightsProperties(catModel.ID)
	case "Multimedia":
		properties = getMultimediaProperties(catModel.ID)

	case "Safety equipment":
		properties = getSafetyEquipmentProperties(catModel.ID)

	case "Seats and steering wheel":
		properties = getSeatsProperties(catModel.ID)

	case "Windows and mirrors":
		properties = getWindowsProperties(catModel.ID)

	default:
		properties = &([]models.Property{})
	}

	for _, prop := range *properties {
		database.Create(&prop)
	}

}

func createCarType(database *gorm.DB) {
	count := 0
	database.
		Model(&models.CarType{}).
		Select(countStarExp).
		Find(&count)
	if count == 0 {
		database.Create(&models.CarType{Name: "Crossover"})
		database.Create(&models.CarType{Name: "Sedan"})
		database.Create(&models.CarType{Name: "Sports"})
		database.Create(&models.CarType{Name: "Coupe"})
		database.Create(&models.CarType{Name: "Hatchback"})
	}
}

func createGearbox(database *gorm.DB) {
	count := 0
	database.
		Model(&models.Gearbox{}).
		Select(countStarExp).
		Find(&count)
	if count == 0 {
		database.Create(&models.Gearbox{Name: "Manual"})
		database.Create(&models.Gearbox{Name: "Automatic"})
	}
}
