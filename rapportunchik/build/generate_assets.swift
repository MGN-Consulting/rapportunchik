import AppKit
import CoreGraphics
import Foundation

// MARK: - Color Palette
// Primary Accents: hsl(234, 100%, 60%), hsl(234, 54%, 20%), hsl(234, 100%, 87%)
let darkBgTop = NSColor(calibratedRed: 0.05, green: 0.07, blue: 0.16, alpha: 1.0)
let darkBgBottom = NSColor(calibratedRed: 0.02, green: 0.03, blue: 0.08, alpha: 1.0)
let slateNavy = NSColor(calibratedHue: 234.0 / 360.0, saturation: 0.54, brightness: 0.22, alpha: 1.0)
let electricBlue = NSColor(calibratedHue: 234.0 / 360.0, saturation: 1.0, brightness: 0.85, alpha: 1.0)
let glowBlue = NSColor(calibratedHue: 234.0 / 360.0, saturation: 0.85, brightness: 0.98, alpha: 1.0)
let iceBlue = NSColor(calibratedHue: 234.0 / 360.0, saturation: 0.35, brightness: 0.95, alpha: 1.0)

let goldBright = NSColor(calibratedRed: 0.99, green: 0.82, blue: 0.25, alpha: 1.0)
let goldMid = NSColor(calibratedRed: 0.93, green: 0.65, blue: 0.10, alpha: 1.0)
let goldDark = NSColor(calibratedRed: 0.72, green: 0.46, blue: 0.05, alpha: 1.0)
let goldDeep = NSColor(calibratedRed: 0.48, green: 0.28, blue: 0.02, alpha: 1.0)

// Helper to save NSImage to PNG
func savePNG(_ image: NSImage, to path: String) {
    guard let tiff = image.tiffRepresentation,
          let bitmap = NSBitmapImageRep(data: tiff),
          let png = bitmap.representation(using: .png, properties: [:]) else {
        fatalError("Failed to convert image to PNG for \(path)")
    }
    try! png.write(to: URL(fileURLWithPath: path))
    print("Saved: \(path)")
}

// MARK: - 1. Render App Icon (1024 x 1024)
func renderAppIcon() -> NSImage {
    let size = NSSize(width: 1024, height: 1024)
    let image = NSImage(size: size)
    image.lockFocus()
    let ctx = NSGraphicsContext.current!.cgContext

    // 1. Clear transparent background
    ctx.clear(CGRect(origin: .zero, size: size))

    // 2. macOS Standard Squircle Path
    let inset: CGFloat = 80.0
    let squircleRect = CGRect(x: inset, y: inset, width: size.width - inset * 2, height: size.height - inset * 2)
    let cornerRadius: CGFloat = 195.0
    let squirclePath = NSBezierPath(roundedRect: squircleRect, xRadius: cornerRadius, yRadius: cornerRadius)

    // Squircle Drop Shadow
    ctx.saveGState()
    let shadow = NSShadow()
    shadow.shadowColor = NSColor(calibratedWhite: 0.0, alpha: 0.55)
    shadow.shadowOffset = NSSize(width: 0, height: -24)
    shadow.shadowBlurRadius = 38
    shadow.set()
    darkBgBottom.setFill()
    squirclePath.fill()
    ctx.restoreGState()

    // Squircle Clip for Internal Contents
    ctx.saveGState()
    squirclePath.addClip()

    // Background Gradient (Deep Navy / Slate)
    let bgGradient = NSGradient(colors: [
        NSColor(calibratedHue: 234.0/360.0, saturation: 0.60, brightness: 0.24, alpha: 1.0),
        NSColor(calibratedHue: 234.0/360.0, saturation: 0.55, brightness: 0.12, alpha: 1.0),
        NSColor(calibratedHue: 234.0/360.0, saturation: 0.65, brightness: 0.05, alpha: 1.0)
    ], atLocations: [0.0, 0.45, 1.0], colorSpace: .deviceRGB)!
    bgGradient.draw(in: squircleRect, angle: -45)

    // Subtle Radial Glow in Center
    let radialGlow = NSGradient(starting: NSColor(calibratedHue: 234.0/360.0, saturation: 0.90, brightness: 0.70, alpha: 0.28),
                                ending: NSColor(calibratedHue: 234.0/360.0, saturation: 0.70, brightness: 0.10, alpha: 0.0))!
    radialGlow.draw(fromCenter: CGPoint(x: 512, y: 512), radius: 0,
                    toCenter: CGPoint(x: 512, y: 512), radius: 460, options: [])

    // Subtle grid/concentric tech ring in background
    ctx.setStrokeColor(NSColor(calibratedHue: 234.0/360.0, saturation: 0.5, brightness: 0.7, alpha: 0.12).cgColor)
    ctx.setLineWidth(2)
    ctx.strokeEllipse(in: CGRect(x: 180, y: 180, width: 664, height: 664))
    ctx.strokeEllipse(in: CGRect(x: 270, y: 270, width: 484, height: 484))

    // 3. Draw Stylized Modern Letters "R" and "K"
    // "R" on left side
    let font = NSFont.systemFont(ofSize: 125, weight: .heavy)
    let letterAttrs: [NSAttributedString.Key: Any] = [
        .font: font,
        .foregroundColor: NSColor(calibratedHue: 234.0/360.0, saturation: 0.4, brightness: 0.95, alpha: 0.40)
    ]
    let strR = NSAttributedString(string: "R", attributes: letterAttrs)
    strR.draw(at: NSPoint(x: 175, y: 645))

    // "K" on right side
    let strK = NSAttributedString(string: "K", attributes: letterAttrs)
    strK.draw(at: NSPoint(x: 755, y: 645))

    // 4. Draw Center Money Bag (💰)
    let cx: CGFloat = 512
    let cy: CGFloat = 475

    // Money Bag Body Shadow
    ctx.saveGState()
    let bagShadow = NSShadow()
    bagShadow.shadowColor = NSColor(calibratedHue: 234.0/360.0, saturation: 0.8, brightness: 0.1, alpha: 0.7)
    bagShadow.shadowOffset = NSSize(width: 0, height: -18)
    bagShadow.shadowBlurRadius = 32
    bagShadow.set()

    // Bag main round belly path
    let bagBelly = NSBezierPath()
    let bw: CGFloat = 260 // half width

    // Organic teardrop sack belly
    bagBelly.move(to: CGPoint(x: cx, y: cy + 130)) // top neck
    bagBelly.curve(to: CGPoint(x: cx + bw, y: cy - 50),
                   controlPoint1: CGPoint(x: cx + 110, y: cy + 120),
                   controlPoint2: CGPoint(x: cx + bw, y: cy + 40))
    bagBelly.curve(to: CGPoint(x: cx + 160, y: cy - 250),
                   controlPoint1: CGPoint(x: cx + bw, y: cy - 170),
                   controlPoint2: CGPoint(x: cx + 220, y: cy - 250))
    bagBelly.curve(to: CGPoint(x: cx - 160, y: cy - 250),
                   controlPoint1: CGPoint(x: cx + 50, y: cy - 265),
                   controlPoint2: CGPoint(x: cx - 50, y: cy - 265))
    bagBelly.curve(to: CGPoint(x: cx - bw, y: cy - 50),
                   controlPoint1: CGPoint(x: cx - 220, y: cy - 250),
                   controlPoint2: CGPoint(x: cx - bw, y: cy - 170))
    bagBelly.curve(to: CGPoint(x: cx, y: cy + 130),
                   controlPoint1: CGPoint(x: cx - bw, y: cy + 40),
                   controlPoint2: CGPoint(x: cx - 110, y: cy + 120))
    bagBelly.close()

    let bellyGradient = NSGradient(colors: [
        goldBright,
        goldMid,
        goldDark,
        goldDeep
    ], atLocations: [0.0, 0.35, 0.75, 1.0], colorSpace: .deviceRGB)!

    bellyGradient.draw(in: bagBelly, angle: 105)
    ctx.restoreGState()

    // Subtle 3D Highlight on top-left of bag belly
    ctx.saveGState()
    bagBelly.addClip()
    let highlightGrad = NSGradient(starting: NSColor(calibratedWhite: 1.0, alpha: 0.45),
                                   ending: NSColor(calibratedWhite: 1.0, alpha: 0.0))!
    highlightGrad.draw(fromCenter: CGPoint(x: cx - 90, y: cy + 20), radius: 0,
                       toCenter: CGPoint(x: cx - 90, y: cy + 20), radius: 240, options: [])
    ctx.restoreGState()

    // Bag Neck & Flared Top Ruff
    let topRuff = NSBezierPath()
    let ruffY = cy + 130
    topRuff.move(to: CGPoint(x: cx - 55, y: ruffY))
    topRuff.curve(to: CGPoint(x: cx - 130, y: ruffY + 120),
                  controlPoint1: CGPoint(x: cx - 90, y: ruffY + 40),
                  controlPoint2: CGPoint(x: cx - 150, y: ruffY + 80))
    topRuff.curve(to: CGPoint(x: cx - 40, y: ruffY + 140),
                  controlPoint1: CGPoint(x: cx - 110, y: ruffY + 145),
                  controlPoint2: CGPoint(x: cx - 60, y: ruffY + 130))
    topRuff.curve(to: CGPoint(x: cx + 40, y: ruffY + 145),
                  controlPoint1: CGPoint(x: cx - 10, y: ruffY + 155),
                  controlPoint2: CGPoint(x: cx + 20, y: ruffY + 140))
    topRuff.curve(to: CGPoint(x: cx + 130, y: ruffY + 115),
                  controlPoint1: CGPoint(x: cx + 70, y: ruffY + 150),
                  controlPoint2: CGPoint(x: cx + 115, y: ruffY + 135))
    topRuff.curve(to: CGPoint(x: cx + 55, y: ruffY),
                  controlPoint1: CGPoint(x: cx + 145, y: ruffY + 75),
                  controlPoint2: CGPoint(x: cx + 90, y: ruffY + 35))
    topRuff.close()

    let ruffGrad = NSGradient(colors: [goldBright, goldMid, goldDark], atLocations: [0.0, 0.4, 1.0], colorSpace: .deviceRGB)!
    ruffGrad.draw(in: topRuff, angle: 90)

    // Bag Tie Ribbon / Cord (Electric Blue theme accent!)
    let tieRect = CGRect(x: cx - 65, y: ruffY - 14, width: 130, height: 26)
    let tiePath = NSBezierPath(roundedRect: tieRect, xRadius: 13, yRadius: 13)
    let tieGrad = NSGradient(colors: [electricBlue, slateNavy], atLocations: [0.0, 1.0], colorSpace: .deviceRGB)!
    tieGrad.draw(in: tiePath, angle: 0)

    // Tie Hanging Tassels
    let tasselL = NSBezierPath()
    tasselL.move(to: CGPoint(x: cx - 25, y: ruffY - 5))
    tasselL.curve(to: CGPoint(x: cx - 55, y: ruffY - 70), controlPoint1: CGPoint(x: cx - 45, y: ruffY - 25), controlPoint2: CGPoint(x: cx - 65, y: ruffY - 45))
    tasselL.lineWidth = 10
    tasselL.lineCapStyle = .round
    electricBlue.setStroke()
    tasselL.stroke()

    let tasselR = NSBezierPath()
    tasselR.move(to: CGPoint(x: cx + 10, y: ruffY - 5))
    tasselR.curve(to: CGPoint(x: cx + 35, y: ruffY - 80), controlPoint1: CGPoint(x: cx + 25, y: ruffY - 30), controlPoint2: CGPoint(x: cx + 15, y: ruffY - 55))
    tasselR.lineWidth = 10
    tasselR.lineCapStyle = .round
    glowBlue.setStroke()
    tasselR.stroke()

    // Tassel golden beads
    let beadL = NSBezierPath(ovalIn: CGRect(x: cx - 65, y: ruffY - 85, width: 20, height: 20))
    goldBright.setFill()
    beadL.fill()
    let beadR = NSBezierPath(ovalIn: CGRect(x: cx + 25, y: ruffY - 95, width: 20, height: 20))
    goldBright.setFill()
    beadR.fill()

    // 5. Embossed Euro (€) Symbol on Bag Belly
    let euroFont = NSFont.systemFont(ofSize: 180, weight: .black)
    let euroShadow = NSShadow()
    euroShadow.shadowColor = NSColor(calibratedRed: 0.35, green: 0.18, blue: 0.02, alpha: 0.8)
    euroShadow.shadowOffset = NSSize(width: 0, height: -6)
    euroShadow.shadowBlurRadius = 8

    let euroAttrs: [NSAttributedString.Key: Any] = [
        .font: euroFont,
        .foregroundColor: NSColor(calibratedRed: 0.40, green: 0.22, blue: 0.02, alpha: 0.85),
        .shadow: euroShadow
    ]
    let euroStr = NSAttributedString(string: "€", attributes: euroAttrs)
    let euroSize = euroStr.size()
    euroStr.draw(at: NSPoint(x: cx - euroSize.width / 2, y: cy - 135))

    // Golden gleam sparkles
    func drawSparkle(at pt: CGPoint, size: CGFloat) {
        let sp = NSBezierPath()
        sp.move(to: CGPoint(x: pt.x, y: pt.y - size))
        sp.curve(to: CGPoint(x: pt.x + size, y: pt.y), controlPoint1: CGPoint(x: pt.x, y: pt.y), controlPoint2: CGPoint(x: pt.x, y: pt.y))
        sp.curve(to: CGPoint(x: pt.x, y: pt.y + size), controlPoint1: CGPoint(x: pt.x, y: pt.y), controlPoint2: CGPoint(x: pt.x, y: pt.y))
        sp.curve(to: CGPoint(x: pt.x - size, y: pt.y), controlPoint1: CGPoint(x: pt.x, y: pt.y), controlPoint2: CGPoint(x: pt.x, y: pt.y))
        sp.curve(to: CGPoint(x: pt.x, y: pt.y - size), controlPoint1: CGPoint(x: pt.x, y: pt.y), controlPoint2: CGPoint(x: pt.x, y: pt.y))
        sp.close()
        NSColor(calibratedWhite: 1.0, alpha: 0.85).setFill()
        sp.fill()
    }
    drawSparkle(at: CGPoint(x: cx - 130, y: cy + 50), size: 18)
    drawSparkle(at: CGPoint(x: cx + 140, y: cy - 20), size: 14)

    ctx.restoreGState() // Unclip squircle

    // 6. Glowing Squircle Border
    ctx.saveGState()
    squirclePath.lineWidth = 5.0
    let borderGrad = NSGradient(colors: [
        NSColor(calibratedHue: 234.0/360.0, saturation: 0.4, brightness: 1.0, alpha: 0.55),
        NSColor(calibratedHue: 234.0/360.0, saturation: 0.9, brightness: 0.7, alpha: 0.35),
        NSColor(calibratedHue: 234.0/360.0, saturation: 0.3, brightness: 0.9, alpha: 0.15)
    ], atLocations: [0.0, 0.5, 1.0], colorSpace: .deviceRGB)!
    borderGrad.draw(in: squirclePath, angle: -45)
    ctx.restoreGState()

    image.unlockFocus()
    return image
}

// MARK: - 2. Render DMG Background (540 x 380)
func renderDMGBackground() -> NSImage {
    let size = NSSize(width: 540, height: 380)
    let image = NSImage(size: size)
    image.lockFocus()
    let ctx = NSGraphicsContext.current!.cgContext

    // Deep Dark Theme Background
    let bgGrad = NSGradient(colors: [
        NSColor(calibratedRed: 0.05, green: 0.07, blue: 0.15, alpha: 1.0),
        NSColor(calibratedRed: 0.03, green: 0.04, blue: 0.09, alpha: 1.0)
    ], atLocations: [0.0, 1.0], colorSpace: .deviceRGB)!
    bgGrad.draw(in: NSRect(origin: .zero, size: size), angle: -90)

    // Center Radial Spotlight
    let centerSpot = NSGradient(starting: NSColor(calibratedHue: 234.0/360.0, saturation: 0.95, brightness: 0.75, alpha: 0.16),
                                ending: NSColor(calibratedHue: 234.0/360.0, saturation: 0.8, brightness: 0.1, alpha: 0.0))!
    centerSpot.draw(fromCenter: CGPoint(x: 270, y: 190), radius: 0,
                    toCenter: CGPoint(x: 270, y: 190), radius: 280, options: [])

    // Top Brand Header
    let titleAttrs: [NSAttributedString.Key: Any] = [
        .font: NSFont.systemFont(ofSize: 18, weight: .bold),
        .foregroundColor: NSColor(calibratedWhite: 0.96, alpha: 1.0)
    ]
    let titleStr = NSAttributedString(string: "Rapportunchik", attributes: titleAttrs)
    titleStr.draw(at: NSPoint(x: 32, y: 330))

    let subAttrs: [NSAttributedString.Key: Any] = [
        .font: NSFont.systemFont(ofSize: 11, weight: .medium),
        .foregroundColor: NSColor(calibratedHue: 234.0/360.0, saturation: 0.35, brightness: 0.85, alpha: 0.8)
    ]
    let subStr = NSAttributedString(string: "Collaborazione Partita IVA • Timesheet & Compensi", attributes: subAttrs)
    subStr.draw(at: NSPoint(x: 32, y: 312))

    // Top right decorative money bag icon mini
    let miniLogoAttrs: [NSAttributedString.Key: Any] = [
        .font: NSFont.systemFont(ofSize: 24),
    ]
    let miniLogoStr = NSAttributedString(string: "💰", attributes: miniLogoAttrs)
    miniLogoStr.draw(at: NSPoint(x: 480, y: 322))

    // Left Icon (App) & Right Icon (Applications) dropzone bases
    let appX: CGFloat = 135
    let appsX: CGFloat = 405
    let iconY: CGFloat = 175

    // Left base ellipse
    let baseL = NSBezierPath(ovalIn: CGRect(x: appX - 55, y: iconY - 65, width: 110, height: 26))
    NSColor(calibratedWhite: 0.0, alpha: 0.35).setFill()
    baseL.fill()

    // Right base ellipse with glowing accent
    let baseR = NSBezierPath(ovalIn: CGRect(x: appsX - 55, y: iconY - 65, width: 110, height: 26))
    NSColor(calibratedHue: 234.0/360.0, saturation: 0.7, brightness: 0.5, alpha: 0.25).setFill()
    baseR.fill()
    ctx.setStrokeColor(NSColor(calibratedHue: 234.0/360.0, saturation: 0.7, brightness: 0.8, alpha: 0.4).cgColor)
    ctx.setLineWidth(1.5)
    ctx.strokeEllipse(in: CGRect(x: appsX - 55, y: iconY - 65, width: 110, height: 26))

    // Middle Flow Pathway (Connecting App to Applications)
    ctx.saveGState()
    let dash: [CGFloat] = [6.0, 4.0]
    ctx.setLineDash(phase: 0, lengths: dash)
    ctx.setStrokeColor(NSColor(calibratedHue: 234.0/360.0, saturation: 0.6, brightness: 0.7, alpha: 0.45).cgColor)
    ctx.setLineWidth(2.0)
    ctx.move(to: CGPoint(x: appX + 60, y: iconY - 5))
    ctx.addLine(to: CGPoint(x: appsX - 70, y: iconY - 5))
    ctx.strokePath()
    ctx.restoreGState()

    // Center Glowing Action Pill
    let pillRect = CGRect(x: 215, y: iconY - 22, width: 110, height: 34)
    let pillPath = NSBezierPath(roundedRect: pillRect, xRadius: 17, yRadius: 17)
    let pillGrad = NSGradient(colors: [
        NSColor(calibratedHue: 234.0/360.0, saturation: 0.54, brightness: 0.26, alpha: 0.95),
        NSColor(calibratedHue: 234.0/360.0, saturation: 1.0, brightness: 0.45, alpha: 0.95)
    ], atLocations: [0.0, 1.0], colorSpace: .deviceRGB)!
    pillGrad.draw(in: pillPath, angle: 0)

    pillPath.lineWidth = 1.5
    NSColor(calibratedHue: 234.0/360.0, saturation: 0.4, brightness: 1.0, alpha: 0.6).setStroke()
    pillPath.stroke()

    // Arrow icon in pill
    let arrowTextAttrs: [NSAttributedString.Key: Any] = [
        .font: NSFont.systemFont(ofSize: 15, weight: .bold),
        .foregroundColor: NSColor(calibratedWhite: 1.0, alpha: 1.0)
    ]
    let arrowStr = NSAttributedString(string: "Trascina ➔", attributes: arrowTextAttrs)
    let arrowSize = arrowStr.size()
    arrowStr.draw(at: NSPoint(x: 270 - arrowSize.width / 2, y: iconY - 14))

    // Instruction caption below
    let guideAttrs: [NSAttributedString.Key: Any] = [
        .font: NSFont.systemFont(ofSize: 11, weight: .medium),
        .foregroundColor: NSColor(calibratedHue: 234.0/360.0, saturation: 0.3, brightness: 0.85, alpha: 0.8)
    ]
    let guideStr = NSAttributedString(string: "Trascina Rapportunchik nella cartella Applicazioni", attributes: guideAttrs)
    let guideSize = guideStr.size()
    guideStr.draw(at: NSPoint(x: 270 - guideSize.width / 2, y: 55))

    // Footer info
    let footerAttrs: [NSAttributedString.Key: Any] = [
        .font: NSFont.systemFont(ofSize: 10, weight: .regular),
        .foregroundColor: NSColor(calibratedWhite: 0.45, alpha: 0.6)
    ]
    let footerStr = NSAttributedString(string: "Wails V3 • macOS 12+ • Architettura Apple Silicon & Intel", attributes: footerAttrs)
    footerStr.draw(at: NSPoint(x: 32, y: 24))

    image.unlockFocus()
    return image
}

// MARK: - 3. Render DMG File Icon (1024 x 1024)
// Sleek Apple external volume / installer disk with theme accents & money bag emblem
func renderDMGFileIcon(appIcon: NSImage) -> NSImage {
    let size = NSSize(width: 1024, height: 1024)
    let image = NSImage(size: size)
    image.lockFocus()
    let ctx = NSGraphicsContext.current!.cgContext

    ctx.clear(CGRect(origin: .zero, size: size))

    // Drive enclosure outer rect
    let driveRect = CGRect(x: 140, y: 110, width: 744, height: 804)
    let drivePath = NSBezierPath(roundedRect: driveRect, xRadius: 85, yRadius: 85)

    // Drive shadow
    ctx.saveGState()
    let shadow = NSShadow()
    shadow.shadowColor = NSColor(calibratedWhite: 0.0, alpha: 0.5)
    shadow.shadowOffset = NSSize(width: 0, height: -22)
    shadow.shadowBlurRadius = 36
    shadow.set()
    NSColor.black.setFill()
    drivePath.fill()
    ctx.restoreGState()

    // Outer Bezel Gradient (Dark Titanium)
    let bezelGrad = NSGradient(colors: [
        NSColor(calibratedHue: 234.0/360.0, saturation: 0.25, brightness: 0.45, alpha: 1.0),
        NSColor(calibratedHue: 234.0/360.0, saturation: 0.35, brightness: 0.20, alpha: 1.0),
        NSColor(calibratedHue: 234.0/360.0, saturation: 0.30, brightness: 0.12, alpha: 1.0)
    ], atLocations: [0.0, 0.5, 1.0], colorSpace: .deviceRGB)!
    bezelGrad.draw(in: drivePath, angle: -45)

    // Inner Face Inset
    let innerRect = driveRect.insetBy(dx: 18, dy: 18)
    let innerPath = NSBezierPath(roundedRect: innerRect, xRadius: 72, yRadius: 72)
    let innerGrad = NSGradient(colors: [
        NSColor(calibratedHue: 234.0/360.0, saturation: 0.55, brightness: 0.20, alpha: 1.0),
        NSColor(calibratedHue: 234.0/360.0, saturation: 0.60, brightness: 0.08, alpha: 1.0)
    ], atLocations: [0.0, 1.0], colorSpace: .deviceRGB)!
    innerGrad.draw(in: innerPath, angle: -90)

    // Center Logo Badge: Draw appIcon scaled down into center
    let iconTargetRect = CGRect(x: 262, y: 262, width: 500, height: 500)
    appIcon.draw(in: iconTargetRect, from: .zero, operation: .sourceOver, fraction: 1.0)

    // Front Drive vent grill and blue power LED at bottom
    let ventY = driveRect.origin.y + 45
    let ventRect = CGRect(x: driveRect.origin.x + 80, y: ventY, width: 220, height: 35)
    let ventPath = NSBezierPath(roundedRect: ventRect, xRadius: 8, yRadius: 8)
    NSColor(calibratedWhite: 0.06, alpha: 1.0).setFill()
    ventPath.fill()

    // Blue Power LED
    let ledCircle = NSBezierPath(ovalIn: CGRect(x: driveRect.origin.x + driveRect.width - 120, y: ventY + 7, width: 20, height: 20))
    electricBlue.setFill()
    ledCircle.fill()
    let ledGlow = NSBezierPath(ovalIn: CGRect(x: driveRect.origin.x + driveRect.width - 124, y: ventY + 3, width: 28, height: 28))
    ctx.setStrokeColor(glowBlue.withAlphaComponent(0.6).cgColor)
    ctx.setLineWidth(3)
    ctx.addPath(ledGlow.cgPath)
    ctx.strokePath()

    image.unlockFocus()
    return image
}

// MARK: - Helper to Generate ICNS using iconutil
func generateICNS(from sourceImage: NSImage, targetPath: String) {
    let fm = FileManager.default
    let iconsetName = targetPath.replacingOccurrences(of: ".icns", with: ".iconset")
    try? fm.removeItem(atPath: iconsetName)
    try! fm.createDirectory(atPath: iconsetName, withIntermediateDirectories: true, attributes: nil)

    let sizes: [(String, Int, Int)] = [
        ("icon_16x16.png", 16, 1),
        ("icon_16x16@2x.png", 16, 2),
        ("icon_32x32.png", 32, 1),
        ("icon_32x32@2x.png", 32, 2),
        ("icon_128x128.png", 128, 1),
        ("icon_128x128@2x.png", 128, 2),
        ("icon_256x256.png", 256, 1),
        ("icon_256x256@2x.png", 256, 2),
        ("icon_512x512.png", 512, 1),
        ("icon_512x512@2x.png", 512, 2)
    ]

    for (name, pt, scale) in sizes {
        let px = pt * scale
        let resized = NSImage(size: NSSize(width: px, height: px))
        resized.lockFocus()
        NSGraphicsContext.current?.imageInterpolation = .high
        sourceImage.draw(in: NSRect(x: 0, y: 0, width: px, height: px),
                         from: NSRect(origin: .zero, size: sourceImage.size),
                         operation: .copy, fraction: 1.0)
        resized.unlockFocus()
        savePNG(resized, to: "\(iconsetName)/\(name)")
    }

    // Run iconutil -c icns <iconset>
    let p = Process()
    p.executableURL = URL(fileURLWithPath: "/usr/bin/iconutil")
    p.arguments = ["-c", "icns", iconsetName, "-o", targetPath]
    try! p.run()
    p.waitUntilExit()

    if p.terminationStatus == 0 {
        print("Successfully generated ICNS: \(targetPath)")
        try? fm.removeItem(atPath: iconsetName)
    } else {
        print("Error generating ICNS: \(p.terminationStatus)")
    }
}

// MARK: - Main Execution
let basePath = CommandLine.arguments.count > 1 ? CommandLine.arguments[1] : "."
let buildDir = "\(basePath)/build"
let darwinDir = "\(buildDir)/darwin"
let frontendPublicDir = "\(basePath)/frontend/public"

print("Rendering Rapportunchik Icon Assets...")

// 1. App Icon
let appIcon = renderAppIcon()
savePNG(appIcon, to: "\(buildDir)/appicon.png")
savePNG(appIcon, to: "\(frontendPublicDir)/wails.png") // also update frontend webview logo

// 2. Generate darwin/icons.icns from appIcon
generateICNS(from: appIcon, targetPath: "\(darwinDir)/icons.icns")

// 3. DMG Background
let dmgBg = renderDMGBackground()
savePNG(dmgBg, to: "\(darwinDir)/dmg-background.png")

// 4. DMG File Icon
let dmgFileIcon = renderDMGFileIcon(appIcon: appIcon)
savePNG(dmgFileIcon, to: "\(darwinDir)/dmg-file-icon.png")
generateICNS(from: dmgFileIcon, targetPath: "\(darwinDir)/dmg-file-icon.icns")

print("All icons and DMG assets successfully updated!")
