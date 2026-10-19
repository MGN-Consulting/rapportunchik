package main

import (
	"fmt"
)

type RapportService struct {
	pdf *PDFService
}

func NewRapportService() *RapportService {
	return &RapportService{
		pdf: NewPDFService(),
	}
}

func (s *RapportService) GetProfile() (*Profile, error) {
	if DB == nil {
		return nil, fmt.Errorf("database not initialized")
	}
	return DB.GetProfile()
}

func (s *RapportService) SaveProfile(p Profile) error {
	if DB == nil {
		return fmt.Errorf("database not initialized")
	}
	return DB.SaveProfile(p)
}

func (s *RapportService) GetAllProfiles() ([]Profile, error) {
	if DB == nil {
		return nil, fmt.Errorf("database not initialized")
	}
	return DB.GetAllProfiles()
}

func (s *RapportService) GetDayLog(dateStr string) (*WorkDayLog, error) {
	if DB == nil {
		return nil, fmt.Errorf("database not initialized")
	}
	return DB.GetDayLog(dateStr)
}

func (s *RapportService) SaveDayLog(logEntry WorkDayLog) error {
	if DB == nil {
		return fmt.Errorf("database not initialized")
	}
	return DB.SaveDayLog(logEntry)
}

func (s *RapportService) GetMonthSummary(year int, month int) (*MonthSummary, error) {
	if DB == nil {
		return nil, fmt.Errorf("database not initialized")
	}
	return DB.GetMonthSummary(year, month)
}

func (s *RapportService) GetOverallStats() (*OverallStats, error) {
	if DB == nil {
		return nil, fmt.Errorf("database not initialized")
	}
	return DB.GetOverallStats()
}

func (s *RapportService) GetAvailableMonths() ([]MonthOption, error) {
	if DB == nil {
		return nil, fmt.Errorf("database not initialized")
	}
	return DB.GetAvailableMonths()
}

func (s *RapportService) PrintRapportino(year int, month int) (string, error) {
	if DB == nil {
		return "", fmt.Errorf("database not initialized")
	}
	profile, err := DB.GetProfile()
	if err != nil {
		return "", fmt.Errorf("could not get profile: %w", err)
	}

	summary, err := DB.GetMonthSummary(year, month)
	if err != nil {
		return "", fmt.Errorf("could not get month summary: %w", err)
	}

	pdfPath, err := s.pdf.GenerateRapportinoPDF(profile, summary)
	if err != nil {
		return "", fmt.Errorf("could not generate PDF: %w", err)
	}

	// Open the generated PDF in macOS default preview
	_ = s.pdf.OpenFile(pdfPath)

	return pdfPath, nil
}

func (s *RapportService) OpenFile(path string) error {
	return s.pdf.OpenFile(path)
}
