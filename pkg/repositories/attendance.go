package repositories

import (
	"github.com/DanielChachagua/GestionCar/pkg/models"
	"github.com/google/uuid"
)

func (r *Repository) GetAttendanceByID(id string) (*models.Attendance, error) {
	var attendance models.Attendance
	if err := r.DB.Where("id = ?", id).First(&attendance).Error; err != nil {
		return nil, err
	}
	return &attendance, nil
}

func (r *Repository) GetAllAttendances() (*[]models.Attendance, error) {

	var attendances []models.Attendance
	if err := r.DB.Find(&attendances).Error; err != nil {
		return nil, err
	}
	return &attendances, nil

}

func (r *Repository) CreateAttendance(attendance *models.AttendanceCreate) (string, error) {
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
		return "", err
	}
	return newId, nil

}

func (r *Repository) UpdateAttendance(attendance *models.AttendanceUpdate) error {
	if err := r.DB.Where("id = ?", attendance.ID).Updates(&models.Attendance{
		Attendance: attendance.Attendance,
		Hours:      attendance.Hours,
		Date:       attendance.Date,
		Amount:     attendance.Amount,
		IsHoliday:  attendance.IsHoliday,
	}).Error; err != nil {
		return err
	}
	return nil

}

func (r *Repository) DeleteAttendance(id string) error {
	var attendance models.Attendance
	if err := r.DB.Where("id = ?", id).Delete(&attendance).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetAttendancesByDate(date_start string, date_end string) (*[]models.Attendance, error) {
	var attendances []models.Attendance
	if err := r.DB.Where("DATE(date) >= ? AND DATE(date) <= ?", date_start, date_end).Find(&attendances).Error; err != nil {
		return nil, err
	}
	return &attendances, nil
}

func (r *Repository) GetAttendanceByEmployeeID(userID string) (*[]models.Attendance, error) {
	var attendances []models.Attendance
	if err := r.DB.Where("employee_id = ?", userID).Find(&attendances).Error; err != nil {
		return nil, err
	}
	return &attendances, nil
}
