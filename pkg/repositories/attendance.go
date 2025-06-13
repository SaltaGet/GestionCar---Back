package repositories

import (
	"errors"

	"github.com/DanielChachagua/GestionCar/pkg/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (r *AttendanceRepository) AttendanceGetByID(id string) (*models.Attendance, error) {
	var attendance models.Attendance
	if err := r.DB.Where("id = ?", id).First(&attendance).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, models.ErrorResponse(404, "Asistencia no encontrada", err)
		}
		return nil, models.ErrorResponse(500, "Error interno al buscar la asistencia", err)
	}
	return &attendance, nil
}

func (r *AttendanceRepository) AttendanceGetAll() (*[]models.Attendance, error) {
	var attendances []models.Attendance
	if err := r.DB.Find(&attendances).Error; err != nil {
		return nil, models.ErrorResponse(500, "Error al buscar las asistencias", err)
	}
	return &attendances, nil

}

func (r *AttendanceRepository) AttendanceGetByEmployeeID(userID string) (*[]models.Attendance, error) {
	var attendances []models.Attendance
	if err := r.DB.Where("employee_id = ?", userID).Find(&attendances).Error; err != nil {
		return nil, models.ErrorResponse(500, "Error interno al buscar asistencias", err)
	}
	return &attendances, nil
}

func (r *AttendanceRepository) AttendanceCreate(attendance *models.AttendanceCreate) (string, error) {
	newId := uuid.NewString()
	if err := r.DB.Create(&models.Attendance{
		ID:         newId,
		EmployeeID: attendance.EmployeeID,
		Attendance: attendance.Attendance,
		Hours:      attendance.Hours,
		Date:       attendance.Date,
		Amount:     attendance.Amount,
		IsHoliday:  attendance.IsHoliday,
	}).Error; err != nil {
		return "", models.ErrorResponse(500, "Error interno al crear la asistencia", err)
	}

	return newId, nil
}

func (r *AttendanceRepository) AttendanceUpdate(attendance *models.AttendanceUpdate) error {
	if err := r.DB.Where("id = ?", attendance.ID).Updates(&models.Attendance{
		Attendance: attendance.Attendance,
		Hours:      attendance.Hours,
		Date:       attendance.Date,
		Amount:     attendance.Amount,
		IsHoliday:  attendance.IsHoliday,
	}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.ErrorResponse(404, "Asistencia no encontrada", err)
		}
		return models.ErrorResponse(500, "Error interno al actualizar la asistencia", err)
	}
	return nil

}

func (r *AttendanceRepository) AttendanceDelete(id string) error {
	var attendance models.Attendance
	if err := r.DB.Where("id = ?", id).Delete(&attendance).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.ErrorResponse(404, "Asistencia no encontrada", err)
		}
		return models.ErrorResponse(500, "Error interno al eliminar la asistencia", err)
	}
	return nil
}

func (r *AttendanceRepository) AttendanceGetByDate(date_start string, date_end string) (*[]models.Attendance, error) {
	var attendances []models.Attendance
	if err := r.DB.Where("DATE(date) >= ? AND DATE(date) <= ?", date_start, date_end).Find(&attendances).Error; err != nil {
		return nil, models.ErrorResponse(500, "Error interno al buscar las asistencias", err)
	}
	return &attendances, nil
}

