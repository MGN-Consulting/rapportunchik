# Rapportunchik Documentation Hub

Welcome to the **Rapportunchik** technical documentation center. Explore detailed architecture specifications, developer contribution guidelines, cross-platform build instructions, and database schemas.

---

## 📑 Documentation Index

| Guide | Description | Target Audience |
| :--- | :--- | :--- |
| **[Architecture & IPC Specifications](./architecture-and-ipc.md)** | Deep dive into Wails V3 IPC, Svelte 5 reactive Runes, Go backend services, and PDF compilation. | Architects, Go & Svelte Developers |
| **[Developer Guidelines & Code of Conduct](./contributing.md)** | Branching workflow, Conventional Commits, local setup, coding standards, and community pledge. | Open Source Contributors |
| **[Cross-Platform Build Recipes](./cross-platform-builds.md)** | Packaging instructions for macOS (DMG/.app), Windows (NSIS/MSIX), Android (APK/AAB), iOS, and Linux. | Release Engineers, Packagers |
| **[MongoDB Schema & Export Recipes](./database-schema-and-export.md)** | Data models, indexes, `mongodump`, `mongorestore`, and JSON/CSV accounting exports. | Database Administrators, Users |

---

## 📸 Visual UI Tour

Explore the core views of Rapportunchik:

### 1. Main View (Worked Day vs Unworked Day)
<div align="center">
  <img src="./img/worked-day.png" alt="Worked Day" width="48%" style="border-radius: 8px; margin-right: 2%;" />
  <img src="./img/unworked-day.png" alt="Unworked Day" width="48%" style="border-radius: 8px;" />
  <p><i>Left: Confirmed workday with cash bag (💰) on timeline. Right: Unconfirmed or weekend day.</i></p>
</div>

### 2. Stats & Real-Time Wave Physics
<div align="center">
  <img src="./img/stats-page.png" alt="Stats & Liquid Tank" width="90%" style="border-radius: 8px; margin-top: 10px;" />
  <p><i>Monthly earnings progress, all-time hour/day counters, and 60fps liquid tank wave simulation towards the 20K € target.</i></p>
</div>

### 3. Profile & Weekly Presets
<div align="center">
  <img src="./img/profile-page-top.png" alt="Profile Settings" width="48%" style="border-radius: 8px; margin-right: 2%;" />
  <img src="./img/profile-page-bottom.png" alt="Weekly Presets" width="48%" style="border-radius: 8px;" />
  <p><i>Left: Committente and Collaborator fiscal credentials. Right: Weekly schedule presets.</i></p>
</div>
