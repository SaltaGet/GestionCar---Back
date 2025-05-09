package services

import (
	"errors"

	"github.com/DanielChachagua/GestionCar/models"
	"github.com/DanielChachagua/GestionCar/repositories"
	"gorm.io/gorm"
)

func GetAttendanceByID(id string, workplace string) (*models.AttendanceLaundry, *models.AttendanceWorkshop, error) {
	attendanceLaundry, attendanceWorkshop, err := repositories.Repo.GetAttendanceByID(id, workplace)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil, models.ErrorResponse(404, "Empleado no encontrado", err)
		}
		return nil, nil, models.ErrorResponse(500, "Error al actualizar cliente", err)
	}
	return attendanceLaundry, attendanceWorkshop, nil
}

func GetAllAttendances(workplace string) (*[]models.AttendanceLaundry, *[]models.AttendanceWorkshop, error) {
	attendancesLaundry, attendancesWorkshop, err := repositories.Repo.GetAllAttendances(workplace)
	if err != nil {
		return nil, nil, models.ErrorResponse(500, "Error al actualizar cliente", err)
	}
	return attendancesLaundry, attendancesWorkshop, nil
}

func GetAllAttendancesByDate(date_start string, date_end string, workplace string) (*[]models.AttendanceLaundry, *[]models.AttendanceWorkshop, error) {
	attendancesLaundry, attendancesWorkshop, err := repositories.Repo.GetAttendancesByDate(date_start, date_end, workplace)
	if err != nil {
		return nil, nil, models.ErrorResponse(500, "Error al actualizar cliente", err)
	}
	return attendancesLaundry, attendancesWorkshop, nil
}

func GetAttendanceByEmployeeID(employeeID string, workplace string) (*[]models.AttendanceLaundry, *[]models.AttendanceWorkshop, error) {
	attendancesLaundry, attendancesWorkshop, err := repositories.Repo.GetAttendanceByEmployeeID(employeeID, workplace)
	if err != nil {
		return nil, nil, models.ErrorResponse(500, "Error al actualizar cliente", err)
	}
	return attendancesLaundry, attendancesWorkshop, nil
}

func CreateAttendance(attendance *models.AttendanceCreate, workplace string) (string, error) {
	id, err := repositories.Repo.CreateAttendance(attendance, workplace)
	if err != nil {
		return "", models.ErrorResponse(500, "Error al actualizar cliente", err)
	}
	return id, nil
}

func UpdateAttendance(attendance *models.AttendanceUpdate, workplace string) error {
	err := repositories.Repo.UpdateAttendance(attendance, workplace)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.ErrorResponse(404, "Empleado no encontrado", err)
		}
		return models.ErrorResponse(500, "Error al actualizar cliente", err)
	}
	return nil
}

func DeleteAttendance(id string, workplace string) error {
	err := repositories.Repo.DeleteAttendance(id, workplace)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.ErrorResponse(404, "Empleado no encontrado", err)
		}
		return models.ErrorResponse(500, "Error al actualizar cliente", err)
	}
	return nil
}