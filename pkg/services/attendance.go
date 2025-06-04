package services

import (
	"errors"

	"github.com/DanielChachagua/GestionCar/pkg/models"
	"gorm.io/gorm"
)

func (s *AttendanceService) AttendanceGetByID(id string) (*models.Attendance, error) {
	attendance, err := s.AttendanceRepository.AttendanceGetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, models.ErrorResponse(404, "Empleado no encontrado", err)
		}
		return nil, models.ErrorResponse(500, "Error al actualizar cliente", err)
	}
	return attendance, nil
}

func (s *AttendanceService) AttendanceGetAll() (*[]models.Attendance, error) {
	attendances, err := s.AttendanceRepository.AttendanceGetAll()
	if err != nil {
		return nil, models.ErrorResponse(500, "Error al actualizar cliente", err)
	}
	return attendances, nil
}

func (s *AttendanceService) AttendanceGetByDate(date_start string, date_end string) (*[]models.Attendance, error) {
	attendances, err := s.AttendanceRepository.AttendanceGetByDate(date_start, date_end)
	if err != nil {
		return nil, models.ErrorResponse(500, "Error al actualizar cliente", err)
	}
	return attendances, nil
}

func (s *AttendanceService) AttendanceGetByEmployeeID(employeeID string) (*[]models.Attendance, error) {
	attendances, err := s.AttendanceRepository.AttendanceGetByEmployeeID(employeeID)
	if err != nil {
		return nil, models.ErrorResponse(500, "Error al actualizar cliente", err)
	}
	return attendances, nil
}

func (s *AttendanceService) AttendanceCreate(attendance *models.AttendanceCreate) (string, error) {
	id, err := s.AttendanceRepository.AttendanceCreate(attendance)
	if err != nil {
		return "", models.ErrorResponse(500, "Error al actualizar cliente", err)
	}
	return id, nil
}

func (s *AttendanceService) AttendanceUpdate(attendance *models.AttendanceUpdate) error {
	err := s.AttendanceRepository.AttendanceUpdate(attendance)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.ErrorResponse(404, "Empleado no encontrado", err)
		}
		return models.ErrorResponse(500, "Error al actualizar cliente", err)
	}
	return nil
}

func (s *AttendanceService) AttendanceDelete(id string) error {
	err := s.AttendanceRepository.AttendanceDelete(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.ErrorResponse(404, "Empleado no encontrado", err)
		}
		return models.ErrorResponse(500, "Error al actualizar cliente", err)
	}
	return nil
}