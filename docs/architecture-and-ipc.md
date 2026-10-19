# Full Application Architecture & IPC Specifications

This document outlines the system architecture, inter-process communication (IPC) protocol, backend services, and frontend reactive state flow for **Rapportunchik**.

---

## 1. High-Level Architecture

Rapportunchik is built on an asynchronous, multi-threaded desktop architecture combining **Wails V3 (Go 1.24+)** for native system orchestration with **Svelte 5 (Bun + Vite)** for high-framerate, reactive user interface rendering, backed by **local-first MongoDB** persistence.

```mermaid
flowchart TD
    subgraph Frontend ["Frontend Layer (Svelte 5 + Bun)"]
        UI[Views & Components]
        State["Reactive State Store (Runes: $state, $derived)"]
        Canvas["Liquid Wave Simulation (HTML5 Canvas @ 60fps)"]
        UI <--> State
        State --> Canvas
    end

    subgraph IPC ["Wails V3 IPC Bridge"]
        Bindings["TypeScript Typed Bindings (window.go)"]
        Dispatcher["Native Event & Call Dispatcher"]
        State <--> Bindings
        Bindings <--> Dispatcher
    end

    subgraph Backend ["Go Backend Runtime (Go 1.24+)"]
        RS["RapportService (Go Struct with Exported RPC Methods)"]
        PDF["PDF Generator Engine (gofpdf / A4 Italian Layout)"]
        DB["MongoDB Repository Client (go.mongodb.org/mongo-driver)"]
        Dispatcher <--> RS
        RS --> DB
        RS --> PDF
    end

    subgraph OS ["Operating System & Filesystem"]
        FS["~/Documents/Rapportini/*.pdf"]
        Preview["macOS Preview / Native PDF Viewer"]
        MongoProc[("mongod Community Instance on localhost:27017")]
        PDF --> FS
        FS --> Preview
        DB <--> MongoProc
    end
```

---

## 2. Visual Workflow & State Mapping

The UI reflects real-time synchronization between the Svelte 5 frontend and the Go backend services.

### 2.1 Main Timesheet & Shift Flow
When a user navigates dates via the animated timeline ribbon, the system queries the MongoDB repository via IPC (`GetWorkDayLog`). If no explicit record exists for that date, default presets (`DayPreset`) are retrieved dynamically based on the weekday.

<div align="center">
  <img src="./img/worked-day.png" alt="Worked Day View" width="850" style="border-radius: 12px; box-shadow: 0 8px 24px rgba(0,0,0,0.4); margin: 16px 0;" />
  <p><i>Figure 1: Main View showing a confirmed worked day with cash bag badge (💰) on the timeline and pre-compiled shift inputs.</i></p>
</div>

When a day is unconfirmed or falls on a weekend, the "Giorno Lavorato" toggle remains off, keeping the day excluded from monthly billing calculations:

<div align="center">
  <img src="./img/unworked-day.png" alt="Unworked Day View" width="850" style="border-radius: 12px; box-shadow: 0 8px 24px rgba(0,0,0,0.4); margin: 16px 0;" />
  <p><i>Figure 2: Main View displaying an unworked weekend day with unlocked editing capabilities.</i></p>
</div>

---

## 3. Wails V3 IPC Bridge Specifications

Wails V3 provides compile-time TypeScript binding generation. Every public method on `RapportService` is automatically exported to `frontend/bindings/changeme/rapportservice.ts`.

All IPC calls return standard JavaScript `Promise<T>` and execute asynchronously without blocking the UI thread.

### 3.1 IPC API Method Catalog

#### `GetProfile(): Promise<models.Profile>`
Retrieves collaborator fiscal data, client metadata, contract parameters, and weekly day presets.
- **Backend Handler**: `RapportService.GetProfile()`
- **Return Type**: `models.Profile` (JSON-serialized struct)
- **Error Handling**: Initializes and seeds default MongoDB document if database is unseeded.

#### `SaveProfile(profile: models.Profile): Promise<void>`
Persists modified fiscal, contract, or schedule preset settings.
- **Backend Handler**: `RapportService.SaveProfile(p models.Profile)`
- **Behavior**: Upserts the singleton configuration record in the `profiles` collection.

#### `GetWorkDayLog(dateStr: string): Promise<models.WorkDayLog>`
Fetches the log for a specific date (`YYYY-MM-DD`).
- **Parameters**: `dateStr` (e.g., `"2026-10-19"`)
- **Behavior**: Checks `work_logs` collection. If absent, derives default hours and location from the weekday preset stored in `Profile`.

#### `SetWorkDayLog(log: models.WorkDayLog): Promise<void>`
Saves or updates daily activity logs with automatic total hour calculation.
- **Parameters**: `models.WorkDayLog` object
- **Behavior**: MongoDB upsert based on the unique `date` index. Debounced by 400ms on the frontend during typing.

#### `GetMonthSummary(year: number, month: number): Promise<models.MonthSummary>`
Calculates all worked days, total hours, and gross pay for the selected month.
- **Parameters**: `year` (int), `month` (1–12)
- **Aggregation**: Queries all logs between `YYYY-MM-01` and `YYYY-MM-LastDay` where `is_worked == true`.

#### `GetOverallStats(): Promise<models.OverallStats>`
Aggregates all-time performance metrics since collaboration inception (`2026-10-16`).
- **Return Fields**:
  - `TotalHours`: Absolute hours worked across all months.
  - `TotalDays`: Absolute number of confirmed worked days.
  - `TotalMoney`: Accumulated gross earnings (`TotalDays × DailyRate`).
  - `TargetRevenue`: Financial milestone target (default: `20,000.00 €`).
  - `TargetPercent`: Progress percentage towards the target milestone.
  - `MonthsWorked`: Count of active collaboration months.

<div align="center">
  <img src="./img/stats-page.png" alt="Stats View & Liquid Canvas" width="850" style="border-radius: 12px; box-shadow: 0 8px 24px rgba(0,0,0,0.4); margin: 16px 0;" />
  <p><i>Figure 3: Stats View showing monthly expected progress, total hours/days counters, and real-time 60fps liquid tank wave simulation.</i></p>
</div>

#### `GeneratePDF(year: number, month: number): Promise<string>`
Triggers pure-Go PDF compilation of the monthly Italian *Rapportino*.
- **Parameters**: `year` (int), `month` (int)
- **Output**: Returns absolute path to the generated PDF (`~/Documents/Rapportini/Rapportino_YYYY_MM.pdf`).
- **System Action**: Invokes native `open` (macOS Preview) on completion.

---

## 4. Frontend State Flow (Svelte 5 Runes)

State management in `src/lib/state.svelte.ts` leverages modern **Svelte 5 Runes**:

```typescript
// Core reactive state container
export const appState = $state({
  activeView: 'main' as 'main' | 'stats' | 'profile',
  selectedDate: new Date().toISOString().split('T')[0],
  currentLog: null as WorkDayLog | null,
  profile: null as Profile | null,
  monthSummary: null as MonthSummary | null,
  overallStats: null as OverallStats | null,
  isSaving: false,
  saveStatus: 'Pronto'
});
```

### 4.1 Unidirectional Event Loop
1. **User Interaction**: Day selected on Timeline ribbon or shift hours modified.
2. **Local Mutation**: `$state` values update immediately for zero input lag.
3. **Debounced Persistence**: Background timer triggers `RapportService.SetWorkDayLog(log)`.
4. **Broadcast & Metrics Sync**: On save completion, `overallStats` and `monthSummary` refresh reactively, updating the liquid tank level and progress bars.

---

## 5. Security & Isolation

- **Local-Only Communication**: Wails V3 IPC operates strictly over local inter-thread memory buffers; no exposed HTTP ports or external network attack surfaces.
- **Privacy Guarantee**: All timesheet entries, client tax IDs, and financial metrics stay confined to the local MongoDB database on `localhost:27017`.
