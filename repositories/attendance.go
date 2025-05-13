package repositories

import (
	"fmt"

	"github.com/DanielChachagua/GestionCar/models"
	"github.com/google/uuid"
)

func (r *Repository) GetAttendanceByID(id string, workplace string) (*models.AttendanceLaundry, *models.AttendanceWorkshop, error) {
	if workplace == "laundry" {
		var attendance models.AttendanceLaundry
		if err := r.DB.Where("id = ?", id).First(&attendance).Error; err != nil {
			return nil, nil, err
		}
		return &attendance, nil, nil
	} else if workplace == "workshop" {
		var attendance models.AttendanceWorkshop
		if err := r.DB.Where("id = ?", id).First(&attendance).Error; err != nil {
			return nil, nil, err
		}
		return nil, &attendance, nil
	}
	return nil, nil, nil
}

func (r *Repository) GetAllAttendances(workplace string) (*[]models.AttendanceLaundry, *[]models.AttendanceWorkshop, error) {
	if workplace == "laundry" {
		var attendances []models.AttendanceLaundry
		if err := r.DB.Find(&attendances).Error; err != nil {
			return nil, nil, err
		}
		return &attendances, nil, nil
	} else if workplace == "workshop" {
		var attendances []models.AttendanceWorkshop
		if err := r.DB.Find(&attendances).Error; err != nil {
			return nil, nil, err
		}
		return nil, &attendances, nil
	}
	return nil, nil, nil
}

func (r *Repository) CreateAttendance(attendance *models.AttendanceCreate, workplace string) (string, error) {
	newId := uuid.NewString()
	switch workplace {
		case "laundry":
			if err := r.DB.Create(&models.AttendanceLaundry{
				ID: newId,
				EmployeeID: attendance.EmployeeID,
				Attendance: attendance.Attendance,
				Hours: attendance.Hours,
				Date: attendance.Date,
				Amount: attendance.Amount,
				IsHoliday: attendance.IsHoliday,
			}).Error; err != nil {
				return "", err
			}
			return newId, nil
		case "workshop":
			if err := r.DB.Create(&models.AttendanceWorkshop{
				ID: newId,
				EmployeeID: attendance.EmployeeID,
				Attendance: attendance.Attendance,
				Hours: attendance.Hours,
				Date: attendance.Date,
				Amount: attendance.Amount,
				IsHoliday: attendance.IsHoliday,
			}).Error; err != nil {
				return "", err
			}
			return newId, nil
		default:
			return "", fmt.Errorf("tipo de asistencia no soportado: %T", attendance)
	}
}

func (r *Repository) UpdateAttendance(attendance *models.AttendanceUpdate, workplace string) error {
	switch workplace {
		case "laundry":
			if err := r.DB.Where("id = ?", attendance.ID).Updates(&models.AttendanceLaundry{
				EmployeeID: attendance.EmployeeID,
				Attendance: attendance.Attendance,
				Hours: attendance.Hours,
				Date: attendance.Date,
				Amount: attendance.Amount,
				IsHoliday: attendance.IsHoliday,
			}).Error; err != nil {
				return err
			}
			return nil
		case "workshop":
			if err := r.DB.Where("id = ?", attendance.ID).Updates(&models.AttendanceWorkshop{
				EmployeeID: attendance.EmployeeID,
				Attendance: attendance.Attendance,
				Hours: attendance.Hours,
				Date: attendance.Date,
				Amount: attendance.Amount,
				IsHoliday: attendance.IsHoliday,
			}).Error; err != nil {
				return err
			}
			return nil
		default:
			return fmt.Errorf("tipo de asistencia no soportado: %T", attendance)
	}
}

func (r *Repository) DeleteAttendance(id string, workplace string) error {
	if workplace == "laundry" {
		var attendance models.AttendanceLaundry
		if err := r.DB.Where("id = ?", id).Delete(&attendance).Error; err != nil {
			return err
		}
	} else if workplace == "workshop" {
		var attendance models.AttendanceWorkshop
		if err := r.DB.Where("id = ?", id).Delete(&attendance).Error; err != nil {
			return err
		}
	} else {
		return fmt.Errorf("tipo de asistencia no soportado: %s", workplace)
	}
	return nil
}

func (r *Repository) GetAttendancesByDate(date_start string, date_end string, workplace string) (*[]models.AttendanceLaundry, *[]models.AttendanceWorkshop, error) {
	if workplace == "laundry" {
		var attendances []models.AttendanceLaundry
		if err := r.DB.Where("DATE(created_at) >= ? AND DATE(created_at) <= ?", date_start, date_end).Find(&attendances).Error; err != nil {
			return nil, nil, err
		}
		return &attendances, nil, nil
	} else if workplace == "workshop" {
		var attendances []models.AttendanceWorkshop
		if err := r.DB.Where("DATE(created_at) >= ? AND DATE(created_at) <= ?", date_start, date_end).Find(&attendances).Error; err != nil {
			return nil, nil, err
		}
		return nil, &attendances, nil
	}
	return nil, nil, nil
}

func (r *Repository) GetAttendanceByEmployeeID(userID string, workplace string) (*[]models.AttendanceLaundry, *[]models.AttendanceWorkshop, error) {
	if workplace == "laundry" {
		var attendances []models.AttendanceLaundry
		if err := r.DB.Where("employee_id = ?", userID).Find(&attendances).Error; err != nil {
			return nil, nil, err
		}
		return &attendances, nil, nil
	} else if workplace == "workshop" {
		var attendances []models.AttendanceWorkshop
		if err := r.DB.Where("employee_id = ?", userID).Find(&attendances).Error; err != nil {
			return nil, nil, err
		}
		return nil, &attendances, nil
	}
	return nil, nil, nil
}