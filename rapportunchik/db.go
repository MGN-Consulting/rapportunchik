package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"sort"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Database struct {
	client   *mongo.Client
	db       *mongo.Database
	profiles *mongo.Collection
	workLogs *mongo.Collection
}

var DB *Database

func InitDB() (*Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOpts := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(clientOpts)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to mongodb: %w", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("mongodb ping failed: %w", err)
	}

	db := client.Database("rapportunchik")
	d := &Database{
		client:   client,
		db:       db,
		profiles: db.Collection("profile"),
		workLogs: db.Collection("work_logs"),
	}
	DB = d

	// Ensure default profile exists
	_, err = d.GetProfile()
	if err != nil {
		log.Printf("Error getting/seeding profile: %v", err)
	}

	return d, nil
}

func DefaultProfile() Profile {
	presets := make(map[string]DayPreset)
	// 0=Sunday
	presets["0"] = DayPreset{
		IsWorkingDay:   false,
		WorkLocation:   "Smart-working",
		MorningStart:   "09:00",
		MorningEnd:     "14:00",
		AfternoonStart: "14:30",
		AfternoonEnd:   "17:00",
		TotalHours:     0.0,
	}
	// 1=Monday: Office 09:10-14:00, 14:30-17:10 (7h 30m = 7.5h)
	presets["1"] = DayPreset{
		IsWorkingDay:   true,
		WorkLocation:   "In Sede",
		MorningStart:   "09:10",
		MorningEnd:     "14:00",
		AfternoonStart: "14:30",
		AfternoonEnd:   "17:10",
		TotalHours:     7.5,
	}
	// 2=Tuesday: Smart-working 09:00-14:00, 14:30-17:00 (7h 30m = 7.5h)
	presets["2"] = DayPreset{
		IsWorkingDay:   true,
		WorkLocation:   "Smart-working",
		MorningStart:   "09:00",
		MorningEnd:     "14:00",
		AfternoonStart: "14:30",
		AfternoonEnd:   "17:00",
		TotalHours:     7.5,
	}
	// 3=Wednesday: Smart-working 09:00-14:00, 14:30-17:30 (8h)
	presets["3"] = DayPreset{
		IsWorkingDay:   true,
		WorkLocation:   "Smart-working",
		MorningStart:   "09:00",
		MorningEnd:     "14:00",
		AfternoonStart: "14:30",
		AfternoonEnd:   "17:30",
		TotalHours:     8.0,
	}
	// 4=Thursday: Office 09:10-14:00, 14:30-17:10 (7h 30m = 7.5h)
	presets["4"] = DayPreset{
		IsWorkingDay:   true,
		WorkLocation:   "In Sede",
		MorningStart:   "09:10",
		MorningEnd:     "14:00",
		AfternoonStart: "14:30",
		AfternoonEnd:   "17:10",
		TotalHours:     7.5,
	}
	// 5=Friday: Smart-working 09:00-14:00, 14:30-17:00 (7h 30m = 7.5h)
	presets["5"] = DayPreset{
		IsWorkingDay:   true,
		WorkLocation:   "Smart-working",
		MorningStart:   "09:00",
		MorningEnd:     "14:00",
		AfternoonStart: "14:30",
		AfternoonEnd:   "17:00",
		TotalHours:     7.5,
	}
	// 6=Saturday
	presets["6"] = DayPreset{
		IsWorkingDay:   false,
		WorkLocation:   "Smart-working",
		MorningStart:   "09:00",
		MorningEnd:     "14:00",
		AfternoonStart: "14:30",
		AfternoonEnd:   "17:00",
		TotalHours:     0.0,
	}

	return Profile{
		ID:                "default",
		ClientCompanyName: "MGN Consulting",
		ClientAddress:     "Milano (MI), Italia",
		ClientPIVA:        "IT12345678901",
		ClientContact:     "amministrazione@mgnconsulting.it",
		ContractorName:    "Collaboratore",
		ContractorPIVA:    "",
		ContractorCF:      "",
		ContractorAddress: "",
		ContractorIBAN:    "",
		FirstDayOfWork:    "2026-10-16",
		DailyRate:         160.0,
		RevenueTarget:     20000.0,
		WeekdayPresets:    presets,
	}
}

func (d *Database) GetProfile() (*Profile, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var profile Profile
	err := d.profiles.FindOne(ctx, bson.M{"_id": "default"}).Decode(&profile)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			def := DefaultProfile()
			_, insertErr := d.profiles.InsertOne(ctx, def)
			if insertErr != nil {
				return nil, fmt.Errorf("failed to insert default profile: %w", insertErr)
			}
			return &def, nil
		}
		return nil, fmt.Errorf("failed to fetch profile: %w", err)
	}

	// Ensure all weekday presets are populated
	if profile.WeekdayPresets == nil || len(profile.WeekdayPresets) < 7 {
		def := DefaultProfile()
		if profile.WeekdayPresets == nil {
			profile.WeekdayPresets = make(map[string]DayPreset)
		}
		for k, v := range def.WeekdayPresets {
			if _, exists := profile.WeekdayPresets[k]; !exists {
				profile.WeekdayPresets[k] = v
			}
		}
		_ = d.SaveProfile(profile)
	}

	return &profile, nil
}

func (d *Database) SaveProfile(p Profile) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	p.ID = "default"
	_, err := d.profiles.ReplaceOne(ctx, bson.M{"_id": "default"}, p, options.Replace().SetUpsert(true))
	if err != nil {
		return fmt.Errorf("failed to save profile: %w", err)
	}
	return nil
}

func (d *Database) GetDayLog(dateStr string) (*WorkDayLog, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	t, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return nil, fmt.Errorf("invalid date format %s: %w", dateStr, err)
	}

	var logEntry WorkDayLog
	err = d.workLogs.FindOne(ctx, bson.M{"_id": dateStr}).Decode(&logEntry)
	if err == nil {
		return &logEntry, nil
	}

	if err != mongo.ErrNoDocuments {
		return nil, fmt.Errorf("error querying day log: %w", err)
	}

	// Pre-compile from profile defaults
	profile, err := d.GetProfile()
	if err != nil {
		return nil, err
	}

	weekdayKey := fmt.Sprint(int(t.Weekday()))
	preset, ok := profile.WeekdayPresets[weekdayKey]
	if !ok {
		preset = DefaultProfile().WeekdayPresets[weekdayKey]
	}

	precompiled := WorkDayLog{
		Date:           dateStr,
		Year:           t.Year(),
		Month:          int(t.Month()),
		Day:            t.Day(),
		IsWorkingDay:   preset.IsWorkingDay,
		WorkLocation:   preset.WorkLocation,
		MorningStart:   preset.MorningStart,
		MorningEnd:     preset.MorningEnd,
		AfternoonStart: preset.AfternoonStart,
		AfternoonEnd:   preset.AfternoonEnd,
		TotalHours:     preset.TotalHours,
		Objective:      "",
		Description:    "",
		IsWorked:       false,
		UpdatedAt:      time.Now().Format(time.RFC3339),
	}

	return &precompiled, nil
}

func (d *Database) SaveDayLog(logEntry WorkDayLog) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	t, err := time.Parse("2006-01-02", logEntry.Date)
	if err == nil {
		logEntry.Year = t.Year()
		logEntry.Month = int(t.Month())
		logEntry.Day = t.Day()
	}
	logEntry.UpdatedAt = time.Now().Format(time.RFC3339)

	_, err = d.workLogs.ReplaceOne(ctx, bson.M{"_id": logEntry.Date}, logEntry, options.Replace().SetUpsert(true))
	if err != nil {
		return fmt.Errorf("failed to save day log: %w", err)
	}
	return nil
}

func (d *Database) GetMonthLogs(year int, month int) ([]WorkDayLog, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{
		"year":  year,
		"month": month,
	}

	cursor, err := d.workLogs.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to query month logs: %w", err)
	}
	defer cursor.Close(ctx)

	var savedLogs []WorkDayLog
	if err := cursor.All(ctx, &savedLogs); err != nil {
		return nil, fmt.Errorf("failed to decode month logs: %w", err)
	}

	return savedLogs, nil
}

func italianMonthName(m int) string {
	names := []string{
		"", "Gennaio", "Febbraio", "Marzo", "Aprile", "Maggio", "Giugno",
		"Luglio", "Agosto", "Settembre", "Ottobre", "Novembre", "Dicembre",
	}
	if m >= 1 && m <= 12 {
		return names[m]
	}
	return ""
}

func (d *Database) GetMonthSummary(year int, month int) (*MonthSummary, error) {
	profile, err := d.GetProfile()
	if err != nil {
		return nil, err
	}

	// Days in requested month
	firstDay := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Local)
	daysInMonth := time.Date(year, time.Month(month)+1, 0, 0, 0, 0, 0, time.Local).Day()

	savedList, err := d.GetMonthLogs(year, month)
	if err != nil {
		return nil, err
	}

	savedMap := make(map[string]WorkDayLog)
	for _, l := range savedList {
		savedMap[l.Date] = l
	}

	var allDays []WorkDayLog
	expectedWorkingDays := 0
	workedDays := 0
	totalHours := 0.0

	for day := 1; day <= daysInMonth; day++ {
		dateStr := fmt.Sprintf("%04d-%02d-%02d", year, month, day)
		currentDate := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)

		entry, exists := savedMap[dateStr]
		if !exists {
			weekdayKey := fmt.Sprint(int(currentDate.Weekday()))
			preset, ok := profile.WeekdayPresets[weekdayKey]
			if !ok {
				preset = DefaultProfile().WeekdayPresets[weekdayKey]
			}
			entry = WorkDayLog{
				Date:           dateStr,
				Year:           year,
				Month:          month,
				Day:            day,
				IsWorkingDay:   preset.IsWorkingDay,
				WorkLocation:   preset.WorkLocation,
				MorningStart:   preset.MorningStart,
				MorningEnd:     preset.MorningEnd,
				AfternoonStart: preset.AfternoonStart,
				AfternoonEnd:   preset.AfternoonEnd,
				TotalHours:     preset.TotalHours,
				Objective:      "",
				Description:    "",
				IsWorked:       false,
				UpdatedAt:      "",
			}
		}

		if entry.IsWorkingDay {
			expectedWorkingDays++
		}
		if entry.IsWorked {
			workedDays++
			totalHours += entry.TotalHours
		}

		allDays = append(allDays, entry)
	}

	// Calculate MonthIndex relative to first day of work
	startYear := 2026
	startMonth := 10
	if parsedStart, err := time.Parse("2006-01-02", profile.FirstDayOfWork); err == nil {
		startYear = parsedStart.Year()
		startMonth = int(parsedStart.Month())
	}
	monthDiff := (year-startYear)*12 + (month - startMonth) + 1
	if monthDiff < 1 {
		monthDiff = 1
	}

	summary := &MonthSummary{
		Year:                   year,
		Month:                  month,
		MonthName:              fmt.Sprintf("%s %d", italianMonthName(month), year),
		MonthIndex:             monthDiff,
		TotalCalendarDays:      daysInMonth,
		ExpectedWorkingDays:    expectedWorkingDays,
		WorkedDays:             workedDays,
		ExpectedMonthlyRevenue: float64(expectedWorkingDays) * profile.DailyRate,
		CurrentMonthlyRevenue:  float64(workedDays) * profile.DailyRate,
		TotalHours:             math.Round(totalHours*10) / 10,
		Days:                   allDays,
	}

	_ = firstDay
	return summary, nil
}

func (d *Database) GetOverallStats() (*OverallStats, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	profile, err := d.GetProfile()
	if err != nil {
		return nil, err
	}

	filter := bson.M{
		"isWorked": true,
	}
	cursor, err := d.workLogs.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to query worked logs: %w", err)
	}
	defer cursor.Close(ctx)

	var workedLogs []WorkDayLog
	if err := cursor.All(ctx, &workedLogs); err != nil {
		return nil, fmt.Errorf("failed to decode worked logs: %w", err)
	}

	totalDays := len(workedLogs)
	totalHours := 0.0
	for _, l := range workedLogs {
		totalHours += l.TotalHours
	}

	totalMoney := float64(totalDays) * profile.DailyRate
	target := profile.RevenueTarget
	if target <= 0 {
		target = 20000.0
	}

	liquidPercent := (totalMoney / target) * 100.0
	if liquidPercent > 100.0 {
		liquidPercent = 100.0
	}

	// Calculate months worked count from first day of work to today
	now := time.Now()
	startYear := 2026
	startMonth := 10
	if parsedStart, err := time.Parse("2006-01-02", profile.FirstDayOfWork); err == nil {
		startYear = parsedStart.Year()
		startMonth = int(parsedStart.Month())
	}
	monthsElapsed := (now.Year()-startYear)*12 + (int(now.Month()) - startMonth) + 1
	if monthsElapsed < 1 {
		monthsElapsed = 1
	}

	return &OverallStats{
		MonthsWorkedCount: monthsElapsed,
		TotalHoursWorked:  math.Round(totalHours*10) / 10,
		TotalDaysWorked:   totalDays,
		TotalMoneyEarned:  totalMoney,
		RevenueTarget:     target,
		LiquidFillPercent: math.Round(liquidPercent*10) / 10,
	}, nil
}

func (d *Database) GetAvailableMonths() ([]MonthOption, error) {
	profile, err := d.GetProfile()
	if err != nil {
		return nil, err
	}

	startYear := 2026
	startMonth := 10
	if parsedStart, err := time.Parse("2006-01-02", profile.FirstDayOfWork); err == nil {
		startYear = parsedStart.Year()
		startMonth = int(parsedStart.Month())
	}

	now := time.Now()
	endYear := now.Year()
	endMonth := int(now.Month())

	// Include at least 6 months ahead for planning
	endMonth += 3
	for endMonth > 12 {
		endYear++
		endMonth -= 12
	}

	var optionsList []MonthOption
	idx := 1
	curYear := startYear
	curMonth := startMonth

	for {
		optionsList = append(optionsList, MonthOption{
			Year:       curYear,
			Month:      curMonth,
			MonthName:  fmt.Sprintf("%s %d", italianMonthName(curMonth), curYear),
			MonthIndex: idx,
		})

		if curYear == endYear && curMonth == endMonth {
			break
		}

		idx++
		curMonth++
		if curMonth > 12 {
			curMonth = 1
			curYear++
		}
	}

	// Sort most recent first or chronological
	sort.Slice(optionsList, func(i, j int) bool {
		if optionsList[i].Year == optionsList[j].Year {
			return optionsList[i].Month < optionsList[j].Month
		}
		return optionsList[i].Year < optionsList[j].Year
	})

	return optionsList, nil
}
