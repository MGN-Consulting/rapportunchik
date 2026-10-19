package main

type DayPreset struct {
	IsWorkingDay   bool    `bson:"isWorkingDay" json:"isWorkingDay"`
	WorkLocation   string  `bson:"workLocation" json:"workLocation"`
	MorningStart   string  `bson:"morningStart" json:"morningStart"`
	MorningEnd     string  `bson:"morningEnd" json:"morningEnd"`
	AfternoonStart string  `bson:"afternoonStart" json:"afternoonStart"`
	AfternoonEnd   string  `bson:"afternoonEnd" json:"afternoonEnd"`
	TotalHours     float64 `bson:"totalHours" json:"totalHours"`
}

type Profile struct {
	ID                string               `bson:"_id" json:"id"`
	ClientCompanyName string               `bson:"clientCompanyName" json:"clientCompanyName"`
	ClientAddress     string               `bson:"clientAddress" json:"clientAddress"`
	ClientPIVA        string               `bson:"clientPIVA" json:"clientPIVA"`
	ClientContact     string               `bson:"clientContact" json:"clientContact"`
	ContractorName    string               `bson:"contractorName" json:"contractorName"`
	ContractorPIVA    string               `bson:"contractorPIVA" json:"contractorPIVA"`
	ContractorCF      string               `bson:"contractorCF" json:"contractorCF"`
	ContractorAddress string               `bson:"contractorAddress" json:"contractorAddress"`
	ContractorIBAN    string               `bson:"contractorIBAN" json:"contractorIBAN"`
	FirstDayOfWork    string               `bson:"firstDayOfWork" json:"firstDayOfWork"` // YYYY-MM-DD (2026-10-16)
	DailyRate         float64              `bson:"dailyRate" json:"dailyRate"`           // 160.0
	RevenueTarget     float64              `bson:"revenueTarget" json:"revenueTarget"`   // 20000.0
	Language          string               `bson:"language,omitempty" json:"language"`   // "it" | "en" | "ru"
	WeekdayPresets    map[string]DayPreset `bson:"weekdayPresets" json:"weekdayPresets"` // "0".."6", 0=Sunday
}

type WorkDayLog struct {
	Date           string  `bson:"_id" json:"date"` // YYYY-MM-DD
	Year           int     `bson:"year" json:"year"`
	Month          int     `bson:"month" json:"month"`
	Day            int     `bson:"day" json:"day"`
	IsWorkingDay   bool    `bson:"isWorkingDay" json:"isWorkingDay"`
	WorkLocation   string  `bson:"workLocation" json:"workLocation"`
	MorningStart   string  `bson:"morningStart" json:"morningStart"`
	MorningEnd     string  `bson:"morningEnd" json:"morningEnd"`
	AfternoonStart string  `bson:"afternoonStart" json:"afternoonStart"`
	AfternoonEnd   string  `bson:"afternoonEnd" json:"afternoonEnd"`
	TotalHours     float64 `bson:"totalHours" json:"totalHours"`
	Objective      string  `bson:"objective" json:"objective"`
	Description    string  `bson:"description" json:"description"`
	IsWorked       bool    `bson:"isWorked" json:"isWorked"`
	UpdatedAt      string  `bson:"updatedAt" json:"updatedAt"`
}

type MonthSummary struct {
	Year                   int          `json:"year"`
	Month                  int          `json:"month"`
	MonthName              string       `json:"monthName"`
	MonthIndex             int          `json:"monthIndex"`
	TotalCalendarDays      int          `json:"totalCalendarDays"`
	ExpectedWorkingDays    int          `json:"expectedWorkingDays"`
	WorkedDays             int          `json:"workedDays"`
	ExpectedMonthlyRevenue float64      `json:"expectedMonthlyRevenue"`
	CurrentMonthlyRevenue  float64      `json:"currentMonthlyRevenue"`
	TotalHours             float64      `json:"totalHours"`
	Days                   []WorkDayLog `json:"days"`
}

type OverallStats struct {
	MonthsWorkedCount int     `json:"monthsWorkedCount"`
	TotalHoursWorked  float64 `json:"totalHoursWorked"`
	TotalDaysWorked   int     `json:"totalDaysWorked"`
	TotalMoneyEarned  float64 `json:"totalMoneyEarned"`
	RevenueTarget     float64 `json:"revenueTarget"`
	LiquidFillPercent float64 `json:"liquidFillPercent"`
}

type MonthOption struct {
	Year       int    `json:"year"`
	Month      int    `json:"month"`
	MonthName  string `json:"monthName"`
	MonthIndex int    `json:"monthIndex"`
}
