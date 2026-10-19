# Rapportunchik — Walkthrough & Verification Summary

**Rapportunchik** has been successfully designed, implemented, and packaged as a native macOS cross-platform application using **Wails V3**, **Go backend**, **Svelte 5 frontend (Bun + Vite)**, and **MongoDB**.

---

## 1. Accomplished Features & Architecture

### 1.1 Dark Theme & Custom Accents
Configured with the exact requested color tokens:
- **Primary Electric Indigo**: `hsl(234, 100%, 60%)`
- **Deep Slate Accent Surface**: `hsl(234, 54%, 20%)`
- **Glow Highlight Ice-Blue**: `hsl(234, 100%, 87%)`
- macOS hidden-inset translucent title bar with customized drag region and navigation tabs.

---

### 1.2 Page 1: Main Page (`MainView.svelte`)
- **Full-Width Animated Timeline (`Timeline.svelte`)**:
  - Horizontal scrolling calendar ribbon centered on the active day.
  - **Main Tick**: Current/selected day displayed prominently with active glow, date number, and day name (`LUN 19`).
  - **Smaller Ticks**: Adjacent days dynamically rendered with day numbers and weekday abbreviations.
  - **Cash Bag Badge (💰)**: Days explicitly marked as worked display a cash bag badge directly on the tick!
  - Chevron navigation controls (prev/next day) and click-to-select.
- **Dual Vertically Separated Sections (50/50 Split)**:
  - **Left Section (`SchedulePanel.svelte`)**:
    - Pre-compiled based on the weekday schedule:
      - **Monday & Thursday**: In Sede (Office), 09:10 – 14:00 and 14:30 – 17:10 (7.5h).
      - **Tuesday & Friday**: Smart-working, 09:00 – 14:00 and 14:30 – 17:00 (7.5h).
      - **Wednesday**: Smart-working, 09:00 – 14:00 and 14:30 – 17:30 (8.0h).
      - **Saturday & Sunday**: Off (non-working day).
    - Greyed-out pre-compiled appearance by default with a dedicated **"Modifica attiva" (Unlock)** toggle for custom day adjustments.
    - Automatic daily total hours calculation.
    - **"Giorno Lavorato" (Mark as Worked)** explicit toggle button with 💰 icon: only when toggled on does the day become an official worked day, granting the 💰 timeline badge and contributing 160 € to metrics.
  - **Right Section (`ActivityPanel.svelte`)**:
    - **Empty by default**.
    - Input field for **Obiettivo** (e.g. `JIRA-1042`, `Formazione`, `Refactoring`) with quick tag helpers.
    - Multi-line auto-expanding textarea for **Descrizione Attività**.
    - Debounced auto-save directly to MongoDB with visual indicator.

---

### 1.3 Page 2: Stats Page (`StatsView.svelte`)
Divided into two vertical halves:
- **Left Half**:
  - **Months Worked Indicator (Top Left)**: Big bold typography (`1° Mese`), hoverable and clickable. On click, opens `MonthSelectorModal.svelte` to switch to any month since collaboration start (`16/10/2026`).
  - **Rounded Print Button**: Prominent rounded action button with printer icon that compiles the official Italian Rapportino PDF and opens it in macOS Preview.
  - **Monthly Cash Progress Bar**: Displays cash progress towards the monthly return (`Worked Days × 160 € / Expected Working Days × 160 €`) with live euro readout and percentage.
  - **Stacked All-Time Metrics**:
    1. **Absolute Total Hours Worked** since start date (16/10/2026).
    2. **Total Days Worked** since start date.
- **Right Half: Animated Liquid Wave Container (`LiquidContainer.svelte`)**:
  - Organic volumetric tank with dynamic canvas-based double sine waves and micro-bubble physics running smoothly at 60fps.
  - Slowly fills up based on **all money accumulated across all months worked** until reaching the **20,000 € target**.
  - Depth gradients using `hsl(234, ...)` accents, glowing crest, milestone markers, and live percentage readout.

---

### 1.4 Page 3: Profile Page (`ProfileView.svelte`)
- **Dati Committente**: Company Name (`MGN Consulting`), P.IVA, legal address, reference contact email.
- **Dati Collaboratore**: Name, P.IVA, Codice Fiscale, billing address, IBAN.
- **Parametri Contrattuali**: First day of work (`2026-10-16`), Daily rate (`160 €`), Revenue target (`20.000 €`).
- **Orari Predefiniti per Giorno della Settimana**: Editable schedule table for all 7 days of the week (working day toggle, workplace selector, morning shift, afternoon shift, total hours).
- Instant MongoDB persistence with feedback toasts.

---

### 1.5 Official Italian Rapportino PDF Engine (`pdfservice.go`)
- Built with pure Go (`github.com/jung-kurt/gofpdf`) with cp1252 accent support.
- Automatically compiles:
  - Header: Collaborator and Committente tax and contact details.
  - KPI Summary: Giorni Lavorati, Ore Totali, Tariffa Giornaliera (160 €), Compenso Totale (€).
  - Clean table of all worked days with Date, Luogo, Orari, Ore, Obiettivo, Descrizione Attività.
  - Signature lines: *Firma del Collaboratore* and *Firma del Committente per Approvazione*.
  - Saves file to `~/Documents/Rapportini/Rapportino_YYYY_MM.pdf` and opens in macOS Preview.

---

## 2. Verification & Validation Results

### 2.1 Backend Automated Tests
- Command: `go test -v .`
- Results:
  ```
  === RUN   TestBackendFlow
      backend_test.go:107: PDF successfully generated at /Users/chillin.killer/Documents/Rapportini/Rapportino_2026_10.pdf (2761 bytes)
  --- PASS: TestBackendFlow (0.04s)
  PASS
  ok   changeme   0.734s
  ```
- Verified:
  - MongoDB connection and Profile initialization.
  - Wednesday preset correctly set to `09:00 - 14:00` and `14:30 - 17:30` (8.0h).
  - Mon/Thu presets set to `09:10 - 14:00` and `14:30 - 17:10` (7.5h).
  - Tue/Fri presets set to `09:00 - 14:00` and `14:30 - 17:00` (7.5h).
  - Explicit mark as worked logic and daily rate calculations.
  - PDF generation produces a valid non-empty PDF file.

### 2.2 Frontend Diagnostics & Build
- `bun run check`: **0 errors, 0 warnings**.
- `bun run build:dev`: Build succeeded in 1.50s.
- `bun run build`: Production build succeeded in 1.52s.

### 2.3 Wails V3 Application Build & Packaging
- Command: `wails3 task darwin:package`
- Result:
  - Native binary compiled with strip flags `-w -s`.
  - macOS App bundle created at `rapportunchik/bin/rapportunchik.app`.
  - App bundle code-signed with ad-hoc signing (`codesign --verify --deep --strict "bin/rapportunchik.app"`: PASS).

---

## 3. How to Run & Develop

### Run Development Mode (Live Reload):
```bash
cd rapportunchik
export PATH=$PATH:/usr/local/go/bin:$HOME/go/bin:$HOME/.bun/bin
wails3 dev
```

### Launch the Packaged Application:
```bash
open rapportunchik/bin/rapportunchik.app
```
