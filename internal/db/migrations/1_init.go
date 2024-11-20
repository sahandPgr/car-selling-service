package migrations

import (
	"github.com/sahandPgr/car-selling-service/constatns"
	"github.com/sahandPgr/car-selling-service/internal/db"
	"github.com/sahandPgr/car-selling-service/internal/models"
	"github.com/sahandPgr/car-selling-service/pkg/logger"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Up_1 is a migration function for creating tables
func Up_1(log logger.Logger) {
	var database = db.GetDB()
	country := models.Country{}
	city := models.City{}
	user := models.User{}
	role := models.Role{}
	user_role := models.UserRole{}

	tables := addToTables(database, country, city, user, role, user_role)

	database.Migrator().CreateTable(tables...)
	createDefault(database)
	log.Info(logger.Potgres, logger.Migration, nil, "Migration Up_1 is done.")
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
