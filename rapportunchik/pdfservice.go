package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/jung-kurt/gofpdf"
)

type PDFService struct{}

func NewPDFService() *PDFService {
	return &PDFService{}
}

func (s *PDFService) GenerateRapportinoPDF(profile *Profile, summary *MonthSummary) (string, error) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	tr := pdf.UnicodeTranslatorFromDescriptor("")

	pdf.SetMargins(14, 14, 14)
	pdf.SetAutoPageBreak(true, 18)
	pdf.AddPage()

	// Colors
	// Primary: #2638c4 (approx hsl(234, 100%, 46%))
	// Slate surface: #f0f2fa
	// Dark text: #1a1e36
	// Gray text: #555d78

	// Header Box
	pdf.SetFillColor(245, 247, 252)
	pdf.SetDrawColor(210, 218, 238)
	pdf.RoundedRect(14, 14, 182, 38, 3, "1234", "FD")

	// Title
	pdf.SetXY(20, 18)
	pdf.SetFont("Helvetica", "B", 16)
	pdf.SetTextColor(30, 42, 120)
	pdf.CellFormat(100, 7, tr("RAPPORTINO MENSILE ATTIVITÀ"), "", 0, "L", false, 0, "")

	// Month badge
	pdf.SetXY(135, 18)
	pdf.SetFont("Helvetica", "B", 12)
	pdf.SetTextColor(50, 60, 90)
	pdf.CellFormat(55, 7, tr(summary.MonthName), "", 0, "R", false, 0, "")

	// Collaborator & Client info columns
	pdf.SetFont("Helvetica", "", 8.5)
	pdf.SetTextColor(40, 45, 60)

	// Left: Collaboratore
	pdf.SetXY(20, 27)
	contractorName := profile.ContractorName
	if contractorName == "" {
		contractorName = "Collaboratore a P.IVA"
	}
	piva := profile.ContractorPIVA
	if piva != "" {
		piva = "P.IVA: " + piva
	}
	cf := profile.ContractorCF
	if cf != "" {
		cf = "CF: " + cf
	}
	pdf.SetFont("Helvetica", "B", 9)
	pdf.CellFormat(80, 4.5, tr(contractorName), "", 1, "L", false, 0, "")
	pdf.SetFont("Helvetica", "", 8)
	pdf.SetX(20)
	pdf.CellFormat(80, 4, tr(strings.TrimSpace(piva+" "+cf)), "", 1, "L", false, 0, "")
	if profile.ContractorIBAN != "" {
		pdf.SetX(20)
		pdf.CellFormat(80, 4, tr("IBAN: "+profile.ContractorIBAN), "", 1, "L", false, 0, "")
	}

	// Right: Committente
	pdf.SetXY(110, 27)
	clientName := profile.ClientCompanyName
	if clientName == "" {
		clientName = "Committente"
	}
	pdf.SetFont("Helvetica", "B", 9)
	pdf.CellFormat(80, 4.5, tr("Committente: "+clientName), "", 1, "L", false, 0, "")
	pdf.SetFont("Helvetica", "", 8)
	if profile.ClientPIVA != "" {
		pdf.SetX(110)
		pdf.CellFormat(80, 4, tr("P.IVA / CF: "+profile.ClientPIVA), "", 1, "L", false, 0, "")
	}
	if profile.ClientAddress != "" {
		pdf.SetX(110)
		pdf.CellFormat(80, 4, tr("Sede: "+profile.ClientAddress), "", 1, "L", false, 0, "")
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
		{"GIORNI LAVORATI", fmt.Sprintf("%d gg", summary.WorkedDays)},
		{"ORE TOTALI", fmt.Sprintf("%.1f h", summary.TotalHours)},
		{"TARIFFA GIORNO", fmt.Sprintf("%.2f €", profile.DailyRate)},
		{"COMPENSO TOTALE", fmt.Sprintf("%.2f €", summary.CurrentMonthlyRevenue)},
	}

	for i, c := range cards {
		x := startX + float64(i)*(cardWidth+gap)
		pdf.SetFillColor(250, 251, 255)
		pdf.SetDrawColor(215, 222, 240)
		pdf.RoundedRect(x, 56, cardWidth, 16, 2, "1234", "FD")

		pdf.SetXY(x+2, 58)
		pdf.SetFont("Helvetica", "B", 7)
		pdf.SetTextColor(110, 118, 140)
		pdf.CellFormat(cardWidth-4, 4, tr(c.Label), "", 1, "C", false, 0, "")

		pdf.SetXY(x+2, 63)
		pdf.SetFont("Helvetica", "B", 11)
		pdf.SetTextColor(25, 35, 95)
		pdf.CellFormat(cardWidth-4, 7, tr(c.Value), "", 1, "C", false, 0, "")
	}

	// Section Divider / Subtitle
	pdf.SetY(76)
	pdf.SetFont("Helvetica", "B", 10)
	pdf.SetTextColor(35, 45, 85)
	pdf.CellFormat(182, 6, tr("DETTAGLIO ATTIVITÀ GIORNALIERE"), "", 1, "L", false, 0, "")

	// Table Header
	pdf.SetY(83)
	pdf.SetFillColor(38, 56, 180)
	pdf.SetTextColor(255, 255, 255)
	pdf.SetFont("Helvetica", "B", 8)
	pdf.SetDrawColor(38, 56, 180)

	colW := []float64{24, 25, 34, 14, 30, 55}
	headers := []string{"DATA", "LUOGO", "ORARI", "ORE", "OBIETTIVO", "DESCRIZIONE ATTIVITÀ"}

	for i, h := range headers {
		align := "L"
		if i == 3 {
			align = "C"
		}
		pdf.CellFormat(colW[i], 7, tr(h), "1", 0, align, true, 0, "")
	}
	pdf.Ln(-1)

	// Table Rows: filter only worked days
	var workedList []WorkDayLog
	for _, day := range summary.Days {
		if day.IsWorked {
			workedList = append(workedList, day)
		}
	}

	pdf.SetFont("Helvetica", "", 7.5)
	pdf.SetDrawColor(225, 230, 242)

	if len(workedList) == 0 {
		pdf.SetFillColor(255, 255, 255)
		pdf.SetTextColor(120, 125, 140)
		pdf.CellFormat(182, 12, tr("Nessuna attività registrata per il mese selezionato."), "1", 1, "C", false, 0, "")
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
			hoursTotal := fmt.Sprintf("%.1fh", row.TotalHours)

			// Truncate or wrap objective and description
			objText := row.Objective
			if objText == "" {
				objText = "-"
			}
			descText := row.Description
			if descText == "" {
				descText = "Attività di sviluppo e supporto progetti."
			}

			// Clean line breaks in description for single/double line fit
			descText = strings.ReplaceAll(descText, "\n", " ")

			// Check auto page break height
			if pdf.GetY() > 255 {
				pdf.AddPage()
				// Re-print header on new page
				pdf.SetFillColor(38, 56, 180)
				pdf.SetTextColor(255, 255, 255)
				pdf.SetFont("Helvetica", "B", 8)
				for i, h := range headers {
					align := "L"
					if i == 3 {
						align = "C"
					}
					pdf.CellFormat(colW[i], 7, tr(h), "1", 0, align, true, 0, "")
				}
				pdf.Ln(-1)
				pdf.SetFont("Helvetica", "", 7.5)
				pdf.SetDrawColor(225, 230, 242)
			}

			rowH := 7.5
			pdf.CellFormat(colW[0], rowH, tr(dateLabel), "1", 0, "L", true, 0, "")
			pdf.CellFormat(colW[1], rowH, tr(row.WorkLocation), "1", 0, "L", true, 0, "")
			pdf.CellFormat(colW[2], rowH, tr(hoursSchedule), "1", 0, "L", true, 0, "")
			pdf.CellFormat(colW[3], rowH, tr(hoursTotal), "1", 0, "C", true, 0, "")
			pdf.CellFormat(colW[4], rowH, tr(truncate(objText, 20)), "1", 0, "L", true, 0, "")
			pdf.CellFormat(colW[5], rowH, tr(truncate(descText, 38)), "1", 1, "L", true, 0, "")
		}
	}

	// Signatures Section at bottom
	pdf.SetY(248)
	if pdf.GetY() > 250 {
		pdf.AddPage()
		pdf.SetY(240)
	}

	pdf.SetFont("Helvetica", "", 8.5)
	pdf.SetTextColor(60, 65, 80)
	pdf.SetXY(14, 250)
	nowStr := time.Now().Format("02/01/2006")
	pdf.CellFormat(90, 5, tr(fmt.Sprintf("Data di compilazione: %s", nowStr)), "", 0, "L", false, 0, "")

	pdf.SetY(260)
	pdf.SetFont("Helvetica", "B", 8)
	pdf.CellFormat(85, 4, tr("Firma del Collaboratore"), "T", 0, "C", false, 0, "")
	pdf.SetX(111)
	pdf.CellFormat(85, 4, tr("Firma del Committente per Approvazione"), "T", 1, "C", false, 0, "")

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
	if len(str) <= maxLen {
		return str
	}
	return str[:maxLen-3] + "..."
}
