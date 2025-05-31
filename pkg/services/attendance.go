package services

import (
	"errors"

	"github.com/DanielChachagua/GestionCar/pkg/models"
	"gorm.io/gorm"
)

func (s *AttendanceService) GetAttendanceByID(id string) (*models.Attendance, error) {
	attendance, err := s.AttendanceRepository.GetAttendanceByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, models.ErrorResponse(404, "Empleado no encontrado", err)
		}
		return nil, models.ErrorResponse(500, "Error al actualizar cliente", err)
	}
	return attendance, nil
}

func (s *AttendanceService) GetAllAttendances() (*[]models.Attendance, error) {
	attendances, err := s.AttendanceRepository.GetAllAttendances()
	if err != nil {
		return nil, models.ErrorResponse(500, "Error al actualizar cliente", err)
	}
	return attendances, nil
}

func (s *AttendanceService) GetAllAttendancesByDate(date_start string, date_end string) (*[]models.Attendance, error) {
	attendances, err := s.AttendanceRepository.GetAttendancesByDate(date_start, date_end)
	if err != nil {
		return nil, models.ErrorResponse(500, "Error al actualizar cliente", err)
	}
	return attendances, nil
}

func (s *AttendanceService) GetAttendanceByEmployeeID(employeeID string) (*[]models.Attendance, error) {
	attendances, err := s.AttendanceRepository.GetAttendanceByEmployeeID(employeeID)
	if err != nil {
		return nil, models.ErrorResponse(500, "Error al actualizar cliente", err)
	}
	return attendances, nil
}

func (s *AttendanceService) CreateAttendance(attendance *models.AttendanceCreate) (string, error) {
	id, err := s.AttendanceRepository.CreateAttendance(attendance)
	if err != nil {
		return "", models.ErrorResponse(500, "Error al actualizar cliente", err)
	}
	return id, nil
}

func (s *AttendanceService) UpdateAttendance(attendance *models.AttendanceUpdate) error {
	err := s.AttendanceRepository.UpdateAttendance(attendance)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.ErrorResponse(404, "Empleado no encontrado", err)
		}
		return models.ErrorResponse(500, "Error al actualizar cliente", err)
	}
	return nil
}

func (s *AttendanceService) DeleteAttendance(id string) error {
	err := s.AttendanceRepository.DeleteAttendance(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.ErrorResponse(404, "Empleado no encontrado", err)
		}
		return models.ErrorResponse(500, "Error al actualizar cliente", err)
	}
	return nil
}