package database

import (
	"log"
	"os"

	"github.com/DanielChachagua/GestionCar/models"
	"github.com/DanielChachagua/GestionCar/utils"
	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect(uri string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(uri), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.User{}, &models.Client{}, &models.Vehicle{}, &models.Workplace{}, &models.Role{}, &models.AuditLog{})
	db.AutoMigrate(
		&models.AttendanceLaundry{},
		// &laundry.ClientLaundry{},
		&models.DailyMovementsLaundry{},
		&models.EmployeeLaundry{},
		&models.ExpenseLaundry{},
		&models.IncomeLaundry{},
		&models.MovementTypeLaundry{},
		&models.ProductLaundry{},
		&models.PurchaseOrderLaundry{},
		&models.PurchaseProductLaundry{},
		&models.ServiceLaundry{},
		&models.SupplierLaundry{},
		// &laundry.VehicleLaundry{},
	)
	db.AutoMigrate(
		&models.AttendanceWorkshop{},
		// &workshop.ClientWorkshop{},
		&models.DailyMovementsWorkshop{},
		&models.EmployeeWorkshop{},
		&models.ExpenseWorkshop{},
		&models.IncomeWorkshop{},
		&models.MovementTypeWorkshop{},
		&models.PartWorkshop{},
		&models.PurchaseOrderWorkshop{},
		&models.PurchasePartWorkshop{},
		&models.ServiceWorkshop{},
		&models.SupplierWorkshop{},
		// &workshop.VehicleWorkshop{},
	)

	var email string
	db.Model(&models.User{}).Select("email").Where("email = ?", os.Getenv("ADMIN_EMAIL")).Scan(&email)

	if email != "" {
		log.Println("El admin ya existe")
		return db, nil
	}
	newId := uuid.NewString()

	pass, err := utils.HashPassword(os.Getenv("ADMIN_PASSWORD"))

	if err != nil {
		return nil, err
	}

	db.Create(&models.User{ID: newId, FirstName: os.Getenv("FIRSTNAME_ADMIN"), LastName: os.Getenv("LASTNAME_ADMIN"),Username: os.Getenv("ADMIN_USERNAME"), Email: os.Getenv("ADMIN_EMAIL"), Password: pass, Role: os.Getenv("ROLE_ADMIN")})

	var laundry string
	db.Model(&models.Workplace{}).Select("identifier").Where("identifier = ?", "laundry").Scan(&laundry)

	var workshop string
	db.Model(&models.Workplace{}).Select("identifier").Where("identifier = ?", "workshop").Scan(&workshop)
	if laundry == "" {
		log.Println("Creando lavanderia")
		db.Create(&models.Workplace{ID: uuid.NewString(), Name: "Lavanderia", Address: "Av. Los Olivos", Phone: "123456789", Email: "laundry@example.com",Identifier: "laundry"})
	}
	if workshop == "" {
		log.Println("Creando taller")
		db.Create(&models.Workplace{ID: uuid.NewString(), Name: "Taller", Address: "Av. Los Olivos", Phone: "123456789", Email: "workshop@example.com", Identifier: "workshop"})
	}

	var roles []models.Role
	db.Model(&models.Role{}).Select("name").Where("name IN ?", []string{"super_admin", "admin", "admin_laundry", "admin_workshop", "employee_laundry", "employee_workshop"}).Scan(&roles)
	if len(roles) == 0 {
		log.Println("Creando roles")
		db.Create(&models.Role{ID: uuid.NewString(), Name: "super_admin", Hierarchy: 1, Workplace: "all"})
		db.Create(&models.Role{ID: uuid.NewString(), Name: "admin", Hierarchy: 2, Workplace: "all"})
		db.Create(&models.Role{ID: uuid.NewString(), Name: "admin_laundry", Hierarchy: 3, Workplace: "laundry"})
		db.Create(&models.Role{ID: uuid.NewString(), Name: "admin_workshop", Hierarchy: 3, Workplace: "workshop"})
		db.Create(&models.Role{ID: uuid.NewString(), Name: "employee_laundry", Hierarchy: 4, Workplace: "laundry"})
		db.Create(&models.Role{ID: uuid.NewString(), Name: "employee_workshop", Hierarchy: 4, Workplace: "workshop"})
	}

	return db, nil
}

func CloseDB(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("No se pudo obtener la conexión de bajo nivel:", err)
	}

	if sqlDB != nil {
		if err := sqlDB.Close(); err != nil {
			log.Fatal("Error al cerrar la conexión:", err)
		}
	}
	return nil
}