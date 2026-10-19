# Cross-Platform Build Recipes

Rapportunchik is powered by the **Wails V3** desktop & mobile application framework. Wails compiles your Go backend into native machine code while embedding a lightweight webview driven by modern HTML5/Svelte.

This guide provides exhaustive build recipes, prerequisites, and packaging instructions for **macOS**, **Windows**, **Android**, **iOS**, and **Linux**.

---

## 1. Multi-Platform Architecture Matrix

| Platform | Web Engine | Go Compiler Architecture | Package Output | Status |
| :--- | :--- | :--- | :--- | :--- |
| **macOS** | WebKit (Cocoa) | `arm64`, `amd64`, Universal | `.app`, `.dmg` | ✅ Tested & Production Ready |
| **Windows** | Microsoft Edge WebView2 | `amd64`, `arm64` | `.exe`, NSIS Installer, `.msix` | 🛠️ Fully Configured |
| **Android** | Android System WebView | `arm64-v8a`, `armeabi-v7a` | `.apk`, `.aab` | 🛠️ Fully Configured |
| **iOS** | Apple WKWebView | `arm64` | `.app`, `.ipa` | 🛠️ Fully Configured |
| **Linux** | WebKit2GTK | `amd64`, `arm64` | Binary, `.deb`, `.rpm`, AppImage | 🛠️ Fully Configured |

---

## 2. macOS Build Recipes (Verified Platform)

### Prerequisites
- macOS 12+ (Monterey, Ventura, Sonoma, Sequoia)
- Xcode Command Line Tools (`xcode-select --install`)
- Go 1.24+ and Bun 1.2+

### 2.1 Native Development Build
```bash
cd rapportunchik
wails3 dev
```

### 2.2 Production App Bundle (`.app`)
Compiles stripped production Go binary and bundles macOS application structure:
```bash
wails3 task darwin:package
```
*Output*: `rapportunchik/bin/rapportunchik.app`

### 2.3 Universal Binary (Apple Silicon + Intel)
Creates a universal Mach-O binary combining `arm64` and `amd64` using macOS `lipo`:
```bash
wails3 task darwin:package:universal
```

### 2.4 Styled macOS DMG Installer
Creates a customized `.dmg` installer with dark branding background, drag-to-Applications layout, and volume icons:
```bash
wails3 task darwin:package:dmg
```
*Output*: `rapportunchik/bin/rapportunchik.dmg` (~13 MB)

### 2.5 Regenerating Brand Assets
To rebuild macOS `.icns`, `.iconset`, and DMG graphics from vector code:
```bash
swift build/generate_assets.swift .
```

---

## 3. Windows 10 & 11 Build Recipes

### Prerequisites
- **On Windows**:
  - Go 1.24+ (Windows 64-bit)
  - WebView2 Runtime (installed by default on Windows 10/11)
  - [NSIS 3+](https://nsis.sourceforge.io/) (for creating `.exe` setup installers)
  - Git for Windows / MinGW-w64 (if CGO required)
- **Cross-compiling from macOS / Linux**:
  - Docker Desktop (for containerized builds)

### 3.1 Native Windows Build
```bash
cd rapportunchik
wails3 task windows:build
```
*Output*: `bin/rapportunchik.exe`

### 3.2 NSIS Installer Packaging
Builds a standalone installer with uninstallation support, desktop shortcut, and registry associations:
```bash
wails3 task windows:package:nsis
```
*Output*: `bin/rapportunchik-setup.exe`

### 3.3 Windows MSIX Modern Store Package
Prepares signed `.msix` container package ready for Microsoft Store or enterprise sideloading:
```bash
# First-time tools setup
wails3 task windows:install:msix:tools

# Package MSIX
wails3 task windows:package:msix
```

### 3.4 Docker Cross-Compilation (from Mac/Linux to Windows)
```bash
wails3 task setup:docker
wails3 task windows:build:docker
```

---

## 4. Android Build Recipes

### Prerequisites
- **Android Studio** (Hedgehog, Iguana, or later)
- **Android SDK** (API 34+ / Upside Down Cake)
- **Android NDK** (version 26.1+ recommended)
- **Java Development Kit (JDK 17+)**
- Environment variables:
  ```bash
  export ANDROID_HOME=$HOME/Library/Android/sdk
  export ANDROID_NDK_HOME=$ANDROID_HOME/ndk/26.1.10909125
  ```

### 4.1 Debug APK Build
```bash
cd rapportunchik
wails3 task android:build:debug
```
*Output*: `build/android/app/build/outputs/apk/debug/app-debug.apk`

### 4.2 Release Android App Bundle (AAB for Google Play)
```bash
wails3 task android:build:release
```
*Output*: `build/android/app/build/outputs/bundle/release/app-release.aab`

### 4.3 Running on Connected Phone / Emulator
```bash
# Ensure device is connected via ADB
adb devices

# Build, install, and stream logcat
wails3 task android:run
```

---

## 5. iOS Build Recipes

### Prerequisites
- macOS running Xcode 15 or 16
- Apple Developer Account (free or paid)
- Connected iPhone/iPad or iOS Simulator

### 5.1 Listing Connected Devices
```bash
wails3 task ios:device:list
```

### 5.2 Build & Run on Physical iPhone
```bash
wails3 task ios:run:device \
  PROJECT=build/ios/xcode/App.xcodeproj \
  SCHEME=ios.dev \
  UDID=<YOUR_DEVICE_UDID> \
  BUNDLE_ID=com.mdmgnp.rapportunchik
```

### 5.3 Xcode Direct Development
You can also open the generated Xcode project directly:
```bash
open build/ios/xcode/App.xcodeproj
```

---

## 6. Linux Build Recipes

### Prerequisites (Ubuntu / Debian)
```bash
sudo apt update
sudo apt install -y build-essential libgtk-3-dev libwebkit2gtk-4.1-dev libayatana-appindicator3-dev
```

### Prerequisites (Fedora / RHEL)
```bash
sudo dnf install -y gcc gcc-c++ gtk3-devel webkit2gtk4.1-devel libayatana-appindicator-gtk3-devel
```

### 6.1 Native Binary Compilation
```bash
cd rapportunchik
wails3 task linux:build
```
*Output*: `bin/rapportunchik`

### 6.2 Portable AppImage Packaging
Generates an all-in-one portable executable compatible across modern Linux distributions:
```bash
wails3 task linux:package:appimage
```
*Output*: `bin/Rapportunchik-x86_64.AppImage`

### 6.3 Native `.deb` and `.rpm` Packages (via nfpm)
```bash
wails3 task linux:package:deb
wails3 task linux:package:rpm
```

---

## 7. Headless / Server Mode

Rapportunchik can also run in **Server Mode** (without a local native GUI window), exposing the Svelte UI over HTTP for browser access on your local network:

```bash
# Build headless server binary
wails3 task build:server

# Run development server on port 8080
wails3 task run:server

# Build production containerized server
wails3 task build:docker
```
