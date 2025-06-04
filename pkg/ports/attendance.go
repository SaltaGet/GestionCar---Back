package ports

import "github.com/DanielChachagua/GestionCar/pkg/models"

type AttendanceService interface {
	AttendanceGetByID(id string) (attendance *models.Attendance, err error)
	AttendanceGetAll() (attendances *[]models.Attendance, err error) 
	AttendanceGetByDate(date_start string, date_end string) (attendances *[]models.Attendance, err error)
	AttendanceGetByEmployeeID(employeeID string) (attendancesLaundry *[]models.Attendance, err error)
	AttendanceCreate(attendance *models.AttendanceCreate) (id string, err error)
	AttendanceUpdate(attendance *models.AttendanceUpdate) (err error)
	AttendanceDelete(id string) (err error)
}

type AttendanceRespository interface {
	AttendanceGetByID(id string) (attendance *models.Attendance, err error)
	AttendanceGetAll() (attendances *[]models.Attendance, err error) 
	AttendanceGetByDate(date_start string, date_end string) (attendancesLaundry *[]models.Attendance, err error)
	AttendanceGetByEmployeeID(employeeID string) (attendancesLaundry *[]models.Attendance, err error)
	AttendanceCreate(attendance *models.AttendanceCreate) (id string, err error)
	AttendanceUpdate(attendance *models.AttendanceUpdate) (err error)
	AttendanceDelete(id string) (err error)
}