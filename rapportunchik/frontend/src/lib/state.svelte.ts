import { RapportService } from '../../bindings/changeme';
import type {
  Profile,
  WorkDayLog,
  MonthSummary,
  OverallStats,
  MonthOption
} from '../../bindings/changeme';
import { i18n, type Locale } from './i18n';

export class AppState {
  activeTab = $state<'main' | 'stats' | 'profile'>('main');
  selectedDate = $state<string>('2026-10-19');
  selectedYear = $state<number>(2026);
  selectedMonth = $state<number>(10);

  profile = $state<Profile | null>(null);
  allProfiles = $state<Profile[]>([]);
  monthSummary = $state<MonthSummary | null>(null);
  overallStats = $state<OverallStats | null>(null);
  availableMonths = $state<MonthOption[]>([]);
  currentDayLog = $state<WorkDayLog | null>(null);

  isMonthModalOpen = $state<boolean>(false);
  isProfileModalOpen = $state<boolean>(false);
  isUnsavedModalOpen = $state<boolean>(false);
  hasUnsavedProfileChanges = $state<boolean>(false);
  pendingTab = $state<'main' | 'stats' | 'profile' | null>(null);
  pendingProfileId = $state<string | null>(null);
  pendingNewProfile = $state<boolean>(false);
  previousSavedProfile: Profile | null = null;
  onDiscardProfileCallback: (() => void) | null = null;
  isPrinting = $state<boolean>(false);
  saveStatus = $state<'idle' | 'saving' | 'saved' | 'error'>('idle');
  toast = $state<{ message: string; type: 'info' | 'success' | 'error' } | null>(null);

  private toastTimeout: ReturnType<typeof setTimeout> | null = null;
  private autoSaveDebounce: ReturnType<typeof setTimeout> | null = null;

  async init() {
    i18n.init();
    try {
      await this.loadProfile();
      await this.loadAllProfiles();
      await this.loadAvailableMonths();
      await this.loadMonthSummary(this.selectedYear, this.selectedMonth);
      await this.loadDay(this.selectedDate);
      await this.loadOverallStats();
    } catch (err) {
      console.error('Initialization error:', err);
      this.showToast(i18n.t.toasts.mongoConnectError, 'error');
    }
  }

  async loadProfile() {
    try {
      const p = await RapportService.GetProfile();
      if (p) {
        this.profile = JSON.parse(JSON.stringify(p));
        this.previousSavedProfile = JSON.parse(JSON.stringify(p));
        if (p.language) {
          i18n.init(p.language);
        }
      }
    } catch (err) {
      console.error('Failed to load profile:', err);
    }
  }

  async loadAllProfiles() {
    try {
      const list = await RapportService.GetAllProfiles();
      if (list && list.length > 0) {
        this.allProfiles = list;
      }
    } catch (err) {
      console.error('Failed to load all profiles:', err);
    }
  }

  async loadAvailableMonths() {
    try {
      const list = await RapportService.GetAvailableMonths();
      if (list) {
        this.availableMonths = list;
      }
    } catch (err) {
      console.error('Failed to load available months:', err);
    }
  }

  async loadMonthSummary(year: number, month: number) {
    try {
      const summary = await RapportService.GetMonthSummary(year, month);
      if (summary) {
        this.monthSummary = summary;
      }
    } catch (err) {
      console.error('Failed to load month summary:', err);
    }
  }

  async loadOverallStats() {
    try {
      const stats = await RapportService.GetOverallStats();
      if (stats) {
        this.overallStats = stats;
      }
    } catch (err) {
      console.error('Failed to load overall stats:', err);
    }
  }

  async loadDay(dateStr: string) {
    try {
      const log = await RapportService.GetDayLog(dateStr);
      if (log) {
        this.currentDayLog = log;
      }
    } catch (err) {
      console.error('Failed to load day log:', err);
    }
  }

  async selectDate(dateStr: string) {
    if (this.selectedDate === dateStr && this.currentDayLog) return;
    this.selectedDate = dateStr;
    const parts = dateStr.split('-');
    if (parts.length === 3) {
      const y = parseInt(parts[0], 10);
      const m = parseInt(parts[1], 10);
      if (y !== this.selectedYear || m !== this.selectedMonth) {
        this.selectedYear = y;
        this.selectedMonth = m;
        await this.loadMonthSummary(y, m);
      }
    }
    await this.loadDay(dateStr);
  }

  async selectMonth(year: number, month: number) {
    this.selectedYear = year;
    this.selectedMonth = month;
    await this.loadMonthSummary(year, month);
    // Align selected date to the same month
    const parts = this.selectedDate.split('-');
    let targetDay = '19';
    if (parts.length === 3) {
      targetDay = parts[2];
    }
    const maxDay = new Date(year, month, 0).getDate();
    let numDay = parseInt(targetDay, 10);
    if (isNaN(numDay) || numDay < 1) numDay = 1;
    if (numDay > maxDay) numDay = maxDay;
    const formattedDate = `${year}-${String(month).padStart(2, '0')}-${String(numDay).padStart(2, '0')}`;
    this.selectedDate = formattedDate;
    await this.loadDay(formattedDate);
    this.isMonthModalOpen = false;
  }

  debounceSaveDay(log: WorkDayLog) {
    this.saveStatus = 'saving';
    if (this.autoSaveDebounce) {
      clearTimeout(this.autoSaveDebounce);
    }
    this.autoSaveDebounce = setTimeout(async () => {
      await this.saveDay(log);
    }, 450);
  }

  async saveDay(log: WorkDayLog) {
    try {
      this.saveStatus = 'saving';
      await RapportService.SaveDayLog(log);
      this.currentDayLog = log;
      this.saveStatus = 'saved';
      // Refresh summaries in background
      await this.loadMonthSummary(this.selectedYear, this.selectedMonth);
      await this.loadOverallStats();
      setTimeout(() => {
        if (this.saveStatus === 'saved') {
          this.saveStatus = 'idle';
        }
      }, 2000);
    } catch (err) {
      console.error('Failed to save day log:', err);
      this.saveStatus = 'error';
      this.showToast(i18n.t.toasts.saveError, 'error');
    }
  }

  async toggleDayWorked() {
    if (!this.currentDayLog) return;
    const updated = {
      ...this.currentDayLog,
      isWorked: !this.currentDayLog.isWorked
    };
    await this.saveDay(updated);
    this.showToast(
      updated.isWorked ? i18n.t.toasts.dayWorkedActive : i18n.t.toasts.dayWorkedRemoved,
      'success'
    );
  }

  async saveProfile(p: Profile) {
    try {
      if (p.language) {
        i18n.setLocale(p.language as Locale);
      }
      if (!p.id) {
        p.id = 'profile_' + Date.now();
      }
      await RapportService.SaveProfile(p);
      this.profile = JSON.parse(JSON.stringify(p));
      this.previousSavedProfile = JSON.parse(JSON.stringify(p));
      this.hasUnsavedProfileChanges = false;
      await this.loadAllProfiles();
      await this.loadMonthSummary(this.selectedYear, this.selectedMonth);
      await this.loadOverallStats();
      await this.loadDay(this.selectedDate);
      this.showToast(i18n.t.toasts.profileSaved, 'success');
    } catch (err) {
      console.error('Failed to save profile:', err);
      this.showToast(i18n.t.toasts.profileError, 'error');
    }
  }

  requestTabChange(targetTab: 'main' | 'stats' | 'profile') {
    if (this.activeTab === targetTab) return;
    if (this.activeTab === 'profile' && this.hasUnsavedProfileChanges) {
      this.pendingTab = targetTab;
      this.isUnsavedModalOpen = true;
      return;
    }
    // If navigating away from an unsaved new profile draft, revert back to previous saved profile
    if (this.profile && !this.allProfiles.some((p) => p.id === this.profile?.id)) {
      if (this.previousSavedProfile) {
        this.profile = JSON.parse(JSON.stringify(this.previousSavedProfile));
      }
    }
    this.activeTab = targetTab;
  }

  async confirmDiscardAndNavigate() {
    if (this.onDiscardProfileCallback) {
      this.onDiscardProfileCallback();
    }
    this.hasUnsavedProfileChanges = false;
    this.isUnsavedModalOpen = false;

    // Check if we were editing an unsaved new profile draft
    if (this.profile && !this.allProfiles.some((p) => p.id === this.profile?.id)) {
      if (this.previousSavedProfile) {
        this.profile = JSON.parse(JSON.stringify(this.previousSavedProfile));
      }
    }

    if (this.pendingProfileId) {
      const pid = this.pendingProfileId;
      this.pendingProfileId = null;
      await this.applySwitchProfile(pid);
      return;
    }

    if (this.pendingNewProfile) {
      this.pendingNewProfile = false;
      this.prepareNewProfile();
      return;
    }

    if (this.pendingTab) {
      this.activeTab = this.pendingTab;
      this.pendingTab = null;
    }
  }

  cancelDiscard() {
    this.pendingTab = null;
    this.pendingProfileId = null;
    this.pendingNewProfile = false;
    this.isUnsavedModalOpen = false;
  }

  async switchProfile(id: string) {
    if (this.profile?.id === id) return;
    if (this.activeTab === 'profile' && this.hasUnsavedProfileChanges) {
      this.pendingProfileId = id;
      this.isUnsavedModalOpen = true;
      return;
    }
    await this.applySwitchProfile(id);
  }

  async applySwitchProfile(id: string) {
    try {
      await RapportService.SetActiveProfile(id);
      await this.loadProfile();
      await this.loadAllProfiles();
      await this.loadMonthSummary(this.selectedYear, this.selectedMonth);
      await this.loadOverallStats();
      await this.loadDay(this.selectedDate);
    } catch (err) {
      console.error('Failed to switch profile:', err);
    }
  }

  startNewProfile() {
    if (this.activeTab === 'profile' && this.hasUnsavedProfileChanges) {
      this.pendingNewProfile = true;
      this.isUnsavedModalOpen = true;
      return;
    }
    this.prepareNewProfile();
  }

  prepareNewProfile() {
    this.previousSavedProfile = this.profile ? JSON.parse(JSON.stringify(this.profile)) : null;
    const newId = 'profile_' + Date.now();
    const defaultPresets = this.profile?.weekdayPresets
      ? JSON.parse(JSON.stringify(this.profile.weekdayPresets))
      : {
          '0': { isWorkingDay: false, workLocation: 'Smart-working', morningStart: '09:00', morningEnd: '14:00', afternoonStart: '14:30', afternoonEnd: '17:00', totalHours: 0.0 },
          '1': { isWorkingDay: true, workLocation: 'In Sede', morningStart: '09:10', morningEnd: '14:00', afternoonStart: '14:30', afternoonEnd: '17:10', totalHours: 7.5 },
          '2': { isWorkingDay: true, workLocation: 'Smart-working', morningStart: '09:00', morningEnd: '14:00', afternoonStart: '14:30', afternoonEnd: '17:00', totalHours: 7.5 },
          '3': { isWorkingDay: true, workLocation: 'Smart-working', morningStart: '09:00', morningEnd: '14:00', afternoonStart: '14:30', afternoonEnd: '17:30', totalHours: 8.0 },
          '4': { isWorkingDay: true, workLocation: 'In Sede', morningStart: '09:10', morningEnd: '14:00', afternoonStart: '14:30', afternoonEnd: '17:10', totalHours: 7.5 },
          '5': { isWorkingDay: true, workLocation: 'Smart-working', morningStart: '09:00', morningEnd: '14:00', afternoonStart: '14:30', afternoonEnd: '17:00', totalHours: 7.5 },
          '6': { isWorkingDay: false, workLocation: 'Smart-working', morningStart: '09:00', morningEnd: '14:00', afternoonStart: '14:30', afternoonEnd: '17:00', totalHours: 0.0 }
        };

    const newProfile: Profile = {
      id: newId,
      clientCompanyName: '',
      clientAddress: '',
      clientPIVA: '',
      clientContact: '',
      contractorName: this.profile?.contractorName || '',
      contractorPIVA: this.profile?.contractorPIVA || '',
      contractorCF: this.profile?.contractorCF || '',
      contractorAddress: this.profile?.contractorAddress || '',
      contractorIBAN: this.profile?.contractorIBAN || '',
      firstDayOfWork: new Date().toISOString().split('T')[0],
      dailyRate: this.profile?.dailyRate || 160.0,
      revenueTarget: this.profile?.revenueTarget || 20000.0,
      language: (this.profile?.language as Locale) || i18n.currentLocale || 'it',
      weekdayPresets: defaultPresets
    };

    this.profile = newProfile;
    this.hasUnsavedProfileChanges = false;
    this.activeTab = 'profile';
  }

  async printMonthPDF() {
    try {
      this.isPrinting = true;
      const path = await RapportService.PrintRapportino(this.selectedYear, this.selectedMonth);
      this.showToast(i18n.t.toasts.pdfSuccess(path), 'success');
    } catch (err) {
      console.error('Failed to generate PDF:', err);
      this.showToast(i18n.t.toasts.pdfError, 'error');
    } finally {
      this.isPrinting = false;
    }
  }

  showToast(message: string, type: 'info' | 'success' | 'error' = 'info') {
    this.toast = { message, type };
    if (this.toastTimeout) {
      clearTimeout(this.toastTimeout);
    }
    this.toastTimeout = setTimeout(() => {
      this.toast = null;
    }, 4000);
  }
}

export const appState = new AppState();
