package main

import (
	"context"
	"os"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestBackendFlow(t *testing.T) {
	// 1. Test DB Init
	db, err := InitDB()
	if err != nil {
		t.Fatalf("Failed to initialize database: %v", err)
	}
	_ = db.SetActiveProfile("default")

	// 2. Test GetProfile and verify Wednesday schedule & default settings
	profile, err := db.GetProfile()
	if err != nil {
		t.Fatalf("Failed to get profile: %v", err)
	}
	if profile.DailyRate != 160.0 {
		t.Errorf("Expected daily rate 160.0, got %f", profile.DailyRate)
	}
	if profile.RevenueTarget != 20000.0 {
		t.Errorf("Expected revenue target 20000.0, got %f", profile.RevenueTarget)
	}

	wedPreset, exists := profile.WeekdayPresets["3"]
	if !exists {
		t.Fatalf("Wednesday preset not found")
	}
	if wedPreset.AfternoonEnd != "17:30" || wedPreset.TotalHours != 8.0 {
		t.Errorf("Wednesday preset expected 17:30 and 8.0 hours, got %s and %f", wedPreset.AfternoonEnd, wedPreset.TotalHours)
	}

	// 2b. Test Multi-Profile: Save secondary profile and switch active profile
	testCtx, testCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer testCancel()
	_, _ = db.workLogs.DeleteMany(testCtx, bson.M{"profileId": "profile_test_client_b"})
	_, _ = db.profiles.DeleteOne(testCtx, bson.M{"_id": "profile_test_client_b"})

	profileB := DefaultProfile()
	profileB.ID = "profile_test_client_b"
	profileB.ClientCompanyName = "Client Bravo SRL"
	profileB.DailyRate = 220.0

	err = db.SaveProfile(profileB)
	if err != nil {
		t.Fatalf("Failed to save profile B: %v", err)
	}

	allProfiles, err := db.GetAllProfiles()
	if err != nil {
		t.Fatalf("Failed to get all profiles: %v", err)
	}
	if len(allProfiles) < 2 {
		t.Errorf("Expected at least 2 profiles, got %d", len(allProfiles))
	}

	activeP, err := db.GetProfile()
	if err != nil {
		t.Fatalf("Failed to get active profile: %v", err)
	}
	if activeP.ID != "profile_test_client_b" || activeP.DailyRate != 220.0 {
		t.Errorf("Expected active profile to be Client Bravo, got %+v", activeP)
	}

	// Switch back to default
	err = db.SetActiveProfile("default")
	if err != nil {
		t.Fatalf("Failed to set active profile back to default: %v", err)
	}
	activeDefault, err := db.GetProfile()
	if err != nil {
		t.Fatalf("Failed to get active profile after switch: %v", err)
	}
	if activeDefault.ID != "default" {
		t.Errorf("Expected active profile to be default, got %s", activeDefault.ID)
	}

	// 3. Test SaveDayLog and GetDayLog
	testDate := "2026-10-19"
	testLog := WorkDayLog{
		Date:           testDate,
		Year:           2026,
		Month:          10,
		Day:            19,
		IsWorkingDay:   true,
		WorkLocation:   "In Sede",
		MorningStart:   "09:10",
		MorningEnd:     "14:00",
		AfternoonStart: "14:30",
		AfternoonEnd:   "17:10",
		TotalHours:     7.5,
		Objective:      "TEST-101",
		Description:    "Test activity description",
		IsWorked:       true,
		UpdatedAt:      time.Now().Format(time.RFC3339),
	}

	err = db.SaveDayLog(testLog)
	if err != nil {
		t.Fatalf("Failed to save day log: %v", err)
	}

	retrieved, err := db.GetDayLog(testDate)
	if err != nil {
		t.Fatalf("Failed to get day log: %v", err)
	}
	if !retrieved.IsWorked || retrieved.Objective != "TEST-101" {
		t.Errorf("Retrieved log mismatch: %+v", retrieved)
	}

	// 4. Test GetMonthSummary
	summary, err := db.GetMonthSummary(2026, 10)
	if err != nil {
		t.Fatalf("Failed to get month summary: %v", err)
	}
	if summary.WorkedDays < 1 {
		t.Errorf("Expected at least 1 worked day, got %d", summary.WorkedDays)
	}
	if summary.CurrentMonthlyRevenue < 160.0 {
		t.Errorf("Expected current monthly revenue >= 160, got %f", summary.CurrentMonthlyRevenue)
	}

	// 5. Test GetOverallStats
	stats, err := db.GetOverallStats()
	if err != nil {
		t.Fatalf("Failed to get overall stats: %v", err)
	}
	if stats.TotalDaysWorked < 1 {
		t.Errorf("Expected total days worked >= 1, got %d", stats.TotalDaysWorked)
	}
	if stats.TotalMoneyEarned < 160.0 {
		t.Errorf("Expected total money earned >= 160, got %f", stats.TotalMoneyEarned)
	}

	// 5b. Test Multi-Profile Data Isolation
	// Switch to profile_test_client_b
	err = db.SetActiveProfile("profile_test_client_b")
	if err != nil {
		t.Fatalf("Failed to switch to profile B: %v", err)
	}

	// Initially, profile B should have 0 worked days and an unworked default log for testDate
	bLog, err := db.GetDayLog(testDate)
	if err != nil {
		t.Fatalf("Failed to get day log for profile B: %v", err)
	}
	if bLog.IsWorked {
		t.Errorf("Expected profile B day log to be unworked initially, got isWorked=true")
	}

	bSummary, err := db.GetMonthSummary(2026, 10)
	if err != nil {
		t.Fatalf("Failed to get month summary for profile B: %v", err)
	}
	if bSummary.WorkedDays != 0 {
		t.Errorf("Expected profile B to have 0 worked days, got %d", bSummary.WorkedDays)
	}

	bStats, err := db.GetOverallStats()
	if err != nil {
		t.Fatalf("Failed to get overall stats for profile B: %v", err)
	}
	if bStats.TotalDaysWorked != 0 {
		t.Errorf("Expected profile B to have 0 total days worked, got %d", bStats.TotalDaysWorked)
	}

	// Now log work for profile B on the same date
	bLogEntry := WorkDayLog{
		Date:         testDate,
		IsWorkingDay: true,
		IsWorked:     true,
		TotalHours:   5.0,
		Objective:    "BRAVO-101",
	}
	err = db.SaveDayLog(bLogEntry)
	if err != nil {
		t.Fatalf("Failed to save day log for profile B: %v", err)
	}

	bSummaryAfter, err := db.GetMonthSummary(2026, 10)
	if err != nil {
		t.Fatalf("Failed to get month summary for profile B after save: %v", err)
	}
	if bSummaryAfter.WorkedDays != 1 {
		t.Errorf("Expected profile B to have 1 worked day, got %d", bSummaryAfter.WorkedDays)
	}
	if bSummaryAfter.CurrentMonthlyRevenue != 220.0 {
		t.Errorf("Expected profile B revenue to be 220.0, got %f", bSummaryAfter.CurrentMonthlyRevenue)
	}

	// Switch back to default and verify default log is completely unchanged
	err = db.SetActiveProfile("default")
	if err != nil {
		t.Fatalf("Failed to switch back to default profile: %v", err)
	}
	defaultLogAfter, err := db.GetDayLog(testDate)
	if err != nil {
		t.Fatalf("Failed to get default day log after switch: %v", err)
	}
	if defaultLogAfter.Objective != "TEST-101" || defaultLogAfter.TotalHours != 7.5 {
		t.Errorf("Default log was corrupted by profile B: %+v", defaultLogAfter)
	}

	// 6. Test PDF Generation (Italian, English, and Russian)
	pdfService := NewPDFService()
	pdfPath, err := pdfService.GenerateRapportinoPDF(profile, summary)
	if err != nil {
		t.Fatalf("Failed to generate Italian PDF: %v", err)
	}

	fileInfo, err := os.Stat(pdfPath)
	if err != nil {
		t.Fatalf("Generated PDF file does not exist at %s: %v", pdfPath, err)
	}
	if fileInfo.Size() < 500 {
		t.Errorf("Generated PDF file is unusually small (%d bytes)", fileInfo.Size())
	}
	t.Logf("Italian PDF successfully generated at %s (%d bytes)", pdfPath, fileInfo.Size())

	// Test English PDF
	profile.Language = "en"
	pdfPathEn, err := pdfService.GenerateRapportinoPDF(profile, summary)
	if err != nil {
		t.Fatalf("Failed to generate English PDF: %v", err)
	}
	if _, err := os.Stat(pdfPathEn); err != nil {
		t.Fatalf("English PDF missing: %v", err)
	}
	t.Logf("English PDF successfully generated at %s", pdfPathEn)

	// Test Russian PDF with Cyrillic characters in descriptions
	profile.Language = "ru"
	ruSummary := *summary
	if len(ruSummary.Days) > 0 {
		ruSummary.Days[0].Objective = "JIRA-RU-104"
		ruSummary.Days[0].Description = "Разработка и тестирование модуля локализации на русском языке."
	}
	pdfPathRu, err := pdfService.GenerateRapportinoPDF(profile, &ruSummary)
	if err != nil {
		t.Fatalf("Failed to generate Russian PDF: %v", err)
	}
	if _, err := os.Stat(pdfPathRu); err != nil {
		t.Fatalf("Russian PDF missing: %v", err)
	}
	t.Logf("Russian PDF successfully generated at %s", pdfPathRu)
}
