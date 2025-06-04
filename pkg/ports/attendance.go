package ports

import "github.com/DanielChachagua/GestionCar/pkg/models"

type AttendanceService interface {
	GetAttendanceByID(id string) (attendance *models.Attendance, err error)
	GetAllAttendances() (attendances *[]models.Attendance, err error) 
	GetAttendancesByDate(date_start string, date_end string) (attendances *[]models.Attendance, err error)
	GetAttendanceByEmployeeID(employeeID string) (attendancesLaundry *[]models.Attendance, err error)
	CreateAttendance(attendance *models.AttendanceCreate) (id string, err error)
	UpdateAttendance(attendance *models.AttendanceUpdate) (err error)
	DeleteAttendance(id string) (err error)
}

type AttendanceRespository interface {
	GetAttendanceByID(id string) (attendance *models.Attendance, err error)
	GetAllAttendances() (attendances *[]models.Attendance, err error) 
	GetAttendancesByDate(date_start string, date_end string) (attendancesLaundry *[]models.Attendance, err error)
	GetAttendanceByEmployeeID(employeeID string) (attendancesLaundry *[]models.Attendance, err error)
	CreateAttendance(attendance *models.AttendanceCreate) (id string, err error)
	UpdateAttendance(attendance *models.AttendanceUpdate) (err error)
	DeleteAttendance(id string) (err error)
}