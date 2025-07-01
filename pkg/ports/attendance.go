package ports

import "github.com/DanielChachagua/GestionCar/pkg/models"

type AttendanceService interface {
	AttendanceGetByID(id string) (attendance *models.AttendanceDTO, err error)
	AttendanceGetAll() (attendances *[]models.AttendanceDTO, err error) 
	AttendanceGetByDate(date_start string, date_end string) (attendances *[]models.AttendanceDTO, err error)
	AttendanceGetByEmployeeID(employeeID string) (attendancesLaundry *[]models.AttendanceDTO, err error)
	AttendanceCreate(attendance *models.AttendanceCreate) (id string, err error)
	AttendanceUpdate(attendance *models.AttendanceUpdate) (err error)
	AttendanceDelete(id string) (err error)
	AttendanceUpdatePay(listIDs []string) (err error)
	AttendancePay(listIDs *[]models.AttendancePay) (err error)
}

type AttendanceRespository interface {
	AttendanceGetByID(id string) (attendance *models.AttendanceDTO, err error)
	AttendanceGetAll() (attendances *[]models.AttendanceDTO, err error) 
	AttendanceGetByDate(date_start string, date_end string) (attendancesLaundry *[]models.AttendanceDTO, err error)
	AttendanceGetByEmployeeID(employeeID string) (attendancesLaundry *[]models.AttendanceDTO, err error)
	AttendanceCreate(attendance *models.AttendanceCreate) (id string, err error)
	AttendanceUpdate(attendance *models.AttendanceUpdate) (err error)
	AttendanceDelete(id string) (err error)
	AttendanceUpdatePay(listIDs []string) (err error)
	AttendancePay(listIDs *[]models.AttendancePay) (err error)
}