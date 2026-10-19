package main

import (
	_ "embed"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/jung-kurt/gofpdf"
)

//go:embed frontend/public/Inter-Medium.ttf
var interFontBytes []byte

type PDFService struct{}

func NewPDFService() *PDFService {
	return &PDFService{}
}

type pdfLocaleStrings struct {
	Title             string
	DefaultContractor string
	ClientPrefix      string
	PivaPrefix        string
	VatPrefix         string
	CfPrefix          string
	AddressPrefix     string
	WorkedDays        string
	TotalHours        string
	DailyRate         string
	GrossTotal        string
	DaysUnit          string
	HoursUnit         string
	DetailsSection    string
	ColDate           string
	ColLocation       string
	ColSchedule       string
	ColHours          string
	ColObjective      string
	ColDescription    string
	EmptyMsg          string
	CompiledDate      string
	ContractorSig     string
	ClientSig         string
	DefaultDesc       string
	MonthNames        [12]string
	LocationMap       map[string]string
}

var pdfLocales = map[string]pdfLocaleStrings{
	"it": {
		Title:             "RAPPORTINO MENSILE ATTIVITÀ",
		DefaultContractor: "Collaboratore a P.IVA",
		ClientPrefix:      "Committente: ",
		PivaPrefix:        "P.IVA: ",
		VatPrefix:         "P.IVA / CF: ",
		CfPrefix:          "CF: ",
		AddressPrefix:     "Sede: ",
		WorkedDays:        "GIORNI LAVORATI",
		TotalHours:        "ORE TOTALI",
		DailyRate:         "TARIFFA GIORNO",
		GrossTotal:        "COMPENSO TOTALE",
		DaysUnit:          "gg",
		HoursUnit:         "h",
		DetailsSection:    "DETTAGLIO ATTIVITÀ GIORNALIERE",
		ColDate:           "DATA",
		ColLocation:       "LUOGO",
		ColSchedule:       "ORARI",
		ColHours:          "ORE",
		ColObjective:      "OBIETTIVO",
		ColDescription:    "DESCRIZIONE ATTIVITÀ",
		EmptyMsg:          "Nessuna attività registrata per il mese selezionato.",
		CompiledDate:      "Data di compilazione: ",
		ContractorSig:     "Firma del Collaboratore",
		ClientSig:         "Firma del Committente per Approvazione",
		DefaultDesc:       "Attività di sviluppo e supporto progetti.",
		MonthNames: [12]string{
			"GENNAIO", "FEBBRAIO", "MARZO", "APRILE", "MAGGIO", "GIUGNO",
			"LUGLIO", "AGOSTO", "SETTEMBRE", "OTTOBRE", "NOVEMBRE", "DICEMBRE",
		},
		LocationMap: map[string]string{
			"Smart-working": "Smart-working",
			"In Sede":       "In Sede",
			"Trasferta":     "Trasferta",
		},
	},
	"en": {
		Title:             "MONTHLY TIMESHEET REPORT",
		DefaultContractor: "Contractor (Freelance)",
		ClientPrefix:      "Client: ",
		PivaPrefix:        "VAT: ",
		VatPrefix:         "VAT / Tax ID: ",
		CfPrefix:          "Tax ID: ",
		AddressPrefix:     "Office: ",
		WorkedDays:        "WORKED DAYS",
		TotalHours:        "TOTAL HOURS",
		DailyRate:         "DAILY RATE",
		GrossTotal:        "TOTAL COMPENSATION",
		DaysUnit:          "d",
		HoursUnit:         "h",
		DetailsSection:    "DAILY WORK ACTIVITIES BREAKDOWN",
		ColDate:           "DATE",
		ColLocation:       "LOCATION",
		ColSchedule:       "HOURS",
		ColHours:          "TOT",
		ColObjective:      "OBJECTIVE",
		ColDescription:    "TASK DESCRIPTION",
		EmptyMsg:          "No work activity logged for the selected month.",
		CompiledDate:      "Date compiled: ",
		ContractorSig:     "Contractor Signature",
		ClientSig:         "Client Approval Signature",
		DefaultDesc:       "Software development and technical project support.",
		MonthNames: [12]string{
			"JANUARY", "FEBRUARY", "MARCH", "APRIL", "MAY", "JUNE",
			"JULY", "AUGUST", "SEPTEMBER", "OCTOBER", "NOVEMBER", "DECEMBER",
		},
		LocationMap: map[string]string{
			"Smart-working": "Remote",
			"In Sede":       "Office",
			"Trasferta":     "On-site / Trip",
		},
	},
	"ru": {
		Title:             "МЕСЯЧНЫЙ ТАБЕЛЬ УЧЕТА РАБОЧЕГО ВРЕМЕНИ",
		DefaultContractor: "Исполнитель",
		ClientPrefix:      "Заказчик: ",
		PivaPrefix:        "ИНН: ",
		VatPrefix:         "ИНН / КПП: ",
		CfPrefix:          "Код: ",
		AddressPrefix:     "Адрес: ",
		WorkedDays:        "ОТРАБОТАНО ДНЕЙ",
		TotalHours:        "ВСЕГО ЧАСОВ",
		DailyRate:         "ДНЕВНАЯ СТАВКА",
		GrossTotal:        "ИТОГО К ВЫПЛАТЕ",
		DaysUnit:          "дн",
		HoursUnit:         "ч",
		DetailsSection:    "ДЕТАЛИЗАЦИЯ ВЫПОЛНЕННЫХ РАБОТ",
		ColDate:           "ДАТА",
		ColLocation:       "МЕСТО",
		ColSchedule:       "ЧАСЫ",
		ColHours:          "ИТОГО",
		ColObjective:      "ЦЕЛЬ / ЗАДАЧА",
		ColDescription:    "ОПИСАНИЕ ДЕЯТЕЛЬНОСТИ",
		EmptyMsg:          "Нет зарегистрированных записей за выбранный месяц.",
		CompiledDate:      "Дата составления: ",
		ContractorSig:     "Подпись исполнителя",
		ClientSig:         "Подпись заказчика для утверждения",
		DefaultDesc:       "Разработка программного обеспечения и поддержка проектов.",
		MonthNames: [12]string{
			"ЯНВАРЬ", "ФЕВРАЛЬ", "МАРТ", "АПРЕЛЬ", "МАЙ", "ИЮНЬ",
			"ИЮЛЬ", "АВГУСТ", "СЕНТЯБРЬ", "ОКТЯБРЬ", "НОЯБРЬ", "ДЕКАБРЬ",
		},
		LocationMap: map[string]string{
			"Smart-working": "Удаленно",
			"In Sede":       "В офисе",
			"Trasferta":     "Командировка",
		},
	},
}

func (s *PDFService) GenerateRapportinoPDF(profile *Profile, summary *MonthSummary) (string, error) {
	pdf := gofpdf.New("P", "mm", "A4", "")

	// Load Inter font for universal UTF-8 & Cyrillic support
	pdf.AddUTF8FontFromBytes("Inter", "", interFontBytes)
	pdf.AddUTF8FontFromBytes("Inter", "B", interFontBytes)

	pdf.SetMargins(14, 14, 14)
	pdf.SetAutoPageBreak(true, 18)
	pdf.AddPage()

	// Resolve language
	lang := profile.Language
	if lang == "" {
		lang = "it"
	}
	loc, exists := pdfLocales[lang]
	if !exists {
		loc = pdfLocales["it"]
	}

	// Month badge text
	monthTitle := summary.MonthName
	if summary.Month >= 1 && summary.Month <= 12 {
		monthTitle = fmt.Sprintf("%s %04d", loc.MonthNames[summary.Month-1], summary.Year)
	}

	// Header Box
	pdf.SetFillColor(245, 247, 252)
	pdf.SetDrawColor(210, 218, 238)
	pdf.RoundedRect(14, 14, 182, 38, 3, "1234", "FD")

	// Title
	pdf.SetXY(20, 18)
	pdf.SetFont("Inter", "B", 13)
	pdf.SetTextColor(30, 42, 120)
	pdf.CellFormat(110, 7, loc.Title, "", 0, "L", false, 0, "")

	// Month badge
	pdf.SetXY(130, 18)
	pdf.SetFont("Inter", "B", 11)
	pdf.SetTextColor(50, 60, 90)
	pdf.CellFormat(60, 7, monthTitle, "", 0, "R", false, 0, "")

	// Collaborator & Client info columns
	pdf.SetFont("Inter", "", 8)
	pdf.SetTextColor(40, 45, 60)

	// Left: Collaboratore / Contractor
	pdf.SetXY(20, 27)
	contractorName := profile.ContractorName
	if contractorName == "" {
		contractorName = loc.DefaultContractor
	}
	piva := profile.ContractorPIVA
	if piva != "" {
		piva = loc.PivaPrefix + piva
	}
	cf := profile.ContractorCF
	if cf != "" {
		cf = loc.CfPrefix + cf
	}
	pdf.SetFont("Inter", "B", 8.5)
	pdf.CellFormat(85, 4.5, contractorName, "", 1, "L", false, 0, "")
	pdf.SetFont("Inter", "", 7.5)
	pdf.SetX(20)
	pdf.CellFormat(85, 4, strings.TrimSpace(piva+" "+cf), "", 1, "L", false, 0, "")
	if profile.ContractorIBAN != "" {
		pdf.SetX(20)
		pdf.CellFormat(85, 4, "IBAN: "+profile.ContractorIBAN, "", 1, "L", false, 0, "")
	}

	// Right: Committente / Client
	pdf.SetXY(110, 27)
	clientName := profile.ClientCompanyName
	if clientName == "" {
		clientName = "-"
	}
	pdf.SetFont("Inter", "B", 8.5)
	pdf.CellFormat(80, 4.5, loc.ClientPrefix+clientName, "", 1, "L", false, 0, "")
	pdf.SetFont("Inter", "", 7.5)
	if profile.ClientPIVA != "" {
		pdf.SetX(110)
		pdf.CellFormat(80, 4, loc.VatPrefix+profile.ClientPIVA, "", 1, "L", false, 0, "")
	}
	if profile.ClientAddress != "" {
		pdf.SetX(110)
		pdf.CellFormat(80, 4, loc.AddressPrefix+profile.ClientAddress, "", 1, "L", false, 0, "")
	}

	// KPI Summary Cards
	pdf.SetY(56)
	cardWidth := 43.5
	gap := 2.6
	startX := 14.0

	cards := []struct {
		Label string
		Value string
	}{
		{loc.WorkedDays, fmt.Sprintf("%d %s", summary.WorkedDays, loc.DaysUnit)},
		{loc.TotalHours, fmt.Sprintf("%.1f %s", summary.TotalHours, loc.HoursUnit)},
		{loc.DailyRate, fmt.Sprintf("%.2f €", profile.DailyRate)},
		{loc.GrossTotal, fmt.Sprintf("%.2f €", summary.CurrentMonthlyRevenue)},
	}

	for i, c := range cards {
		x := startX + float64(i)*(cardWidth+gap)
		pdf.SetFillColor(250, 251, 255)
		pdf.SetDrawColor(215, 222, 240)
		pdf.RoundedRect(x, 56, cardWidth, 16, 2, "1234", "FD")

		pdf.SetXY(x+2, 58)
		pdf.SetFont("Inter", "B", 6.5)
		pdf.SetTextColor(110, 118, 140)
		pdf.CellFormat(cardWidth-4, 4, c.Label, "", 1, "C", false, 0, "")

		pdf.SetXY(x+2, 63)
		pdf.SetFont("Inter", "B", 10.5)
		pdf.SetTextColor(25, 35, 95)
		pdf.CellFormat(cardWidth-4, 7, c.Value, "", 1, "C", false, 0, "")
	}

	// Section Divider / Subtitle
	pdf.SetY(76)
	pdf.SetFont("Inter", "B", 9.5)
	pdf.SetTextColor(35, 45, 85)
	pdf.CellFormat(182, 6, loc.DetailsSection, "", 1, "L", false, 0, "")

	// Table Header
	pdf.SetY(83)
	pdf.SetFillColor(38, 56, 180)
	pdf.SetTextColor(255, 255, 255)
	pdf.SetFont("Inter", "B", 7.5)
	pdf.SetDrawColor(38, 56, 180)

	colW := []float64{22, 25, 34, 14, 32, 55}
	headers := []string{loc.ColDate, loc.ColLocation, loc.ColSchedule, loc.ColHours, loc.ColObjective, loc.ColDescription}

	for i, h := range headers {
		align := "L"
		if i == 3 {
			align = "C"
		}
		pdf.CellFormat(colW[i], 7, h, "1", 0, align, true, 0, "")
	}
	pdf.Ln(-1)

	// Table Rows: filter only worked days
	var workedList []WorkDayLog
	for _, day := range summary.Days {
		if day.IsWorked {
			workedList = append(workedList, day)
		}
	}

	pdf.SetFont("Inter", "", 7.5)
	pdf.SetDrawColor(225, 230, 242)

	if len(workedList) == 0 {
		pdf.SetFillColor(255, 255, 255)
		pdf.SetTextColor(120, 125, 140)
		pdf.CellFormat(182, 12, loc.EmptyMsg, "1", 1, "C", false, 0, "")
	} else {
		for idx, row := range workedList {
			if idx%2 == 0 {
				pdf.SetFillColor(255, 255, 255)
			} else {
				pdf.SetFillColor(247, 249, 254)
			}
			pdf.SetTextColor(35, 40, 55)

			parsedDate, _ := time.Parse("2006-01-02", row.Date)
			dateLabel := row.Date
			if !parsedDate.IsZero() {
				dateLabel = fmt.Sprintf("%02d/%02d/%04d", parsedDate.Day(), int(parsedDate.Month()), parsedDate.Year())
			}

			hoursSchedule := fmt.Sprintf("%s-%s / %s-%s", row.MorningStart, row.MorningEnd, row.AfternoonStart, row.AfternoonEnd)
			hoursTotal := fmt.Sprintf("%.1f%s", row.TotalHours, loc.HoursUnit)

			locLabel := row.WorkLocation
			if translatedLoc, ok := loc.LocationMap[row.WorkLocation]; ok {
				locLabel = translatedLoc
			}

			objText := row.Objective
			if objText == "" {
				objText = "-"
			}
			descText := row.Description
			if descText == "" {
				descText = loc.DefaultDesc
			}
			descText = strings.ReplaceAll(descText, "\n", " ")

			// Check auto page break height
			if pdf.GetY() > 255 {
				pdf.AddPage()
				pdf.SetFillColor(38, 56, 180)
				pdf.SetTextColor(255, 255, 255)
				pdf.SetFont("Inter", "B", 7.5)
				for i, h := range headers {
					align := "L"
					if i == 3 {
						align = "C"
					}
					pdf.CellFormat(colW[i], 7, h, "1", 0, align, true, 0, "")
				}
				pdf.Ln(-1)
				pdf.SetFont("Inter", "", 7.5)
				pdf.SetDrawColor(225, 230, 242)
			}

			rowH := 7.5
			pdf.CellFormat(colW[0], rowH, dateLabel, "1", 0, "L", true, 0, "")
			pdf.CellFormat(colW[1], rowH, locLabel, "1", 0, "L", true, 0, "")
			pdf.CellFormat(colW[2], rowH, hoursSchedule, "1", 0, "L", true, 0, "")
			pdf.CellFormat(colW[3], rowH, hoursTotal, "1", 0, "C", true, 0, "")
			pdf.CellFormat(colW[4], rowH, truncate(objText, 20), "1", 0, "L", true, 0, "")
			pdf.CellFormat(colW[5], rowH, truncate(descText, 38), "1", 1, "L", true, 0, "")
		}
	}

	// Signatures Section at bottom
	pdf.SetY(248)
	if pdf.GetY() > 250 {
		pdf.AddPage()
		pdf.SetY(240)
	}

	pdf.SetFont("Inter", "", 8)
	pdf.SetTextColor(60, 65, 80)
	pdf.SetXY(14, 250)
	nowStr := time.Now().Format("02/01/2006")
	pdf.CellFormat(90, 5, loc.CompiledDate+nowStr, "", 0, "L", false, 0, "")

	pdf.SetY(260)
	pdf.SetFont("Inter", "B", 8)
	pdf.CellFormat(85, 4, loc.ContractorSig, "T", 0, "C", false, 0, "")
	pdf.SetX(111)
	pdf.CellFormat(85, 4, loc.ClientSig, "T", 1, "C", false, 0, "")

	// Save to disk: ~/Documents/Rapportini/Rapportino_YYYY_MM.pdf
	homeDir, err := os.UserHomeDir()
	if err != nil {
		homeDir = "."
	}
	outDir := filepath.Join(homeDir, "Documents", "Rapportini")
	_ = os.MkdirAll(outDir, 0755)

	outPath := filepath.Join(outDir, fmt.Sprintf("Rapportino_%04d_%02d.pdf", summary.Year, summary.Month))
	err = pdf.OutputFileAndClose(outPath)
	if err != nil {
		return "", fmt.Errorf("failed to write PDF file: %w", err)
	}

	return outPath, nil
}

func (s *PDFService) OpenFile(path string) error {
	cmd := exec.Command("open", path)
	return cmd.Start()
}

func truncate(str string, maxLen int) string {
	runes := []rune(str)
	if len(runes) <= maxLen {
		return str
	}
	return string(runes[:maxLen-3]) + "..."
}
