import type { TranslationSchema } from './types';

export const en: TranslationSchema = {
  nav: {
    brandSubtitle: 'Contractor Timesheet',
    profileSelectorAria: 'Select Client Profile',
    profileModalTitle: 'Client Profiles',
    profileModalDesc: 'Select an active client to review timesheets or add a new one:',
    addNewProfileBtn: 'Add new client',
    activeProfileBadge: 'Active',
    dailyRateBadge: (rate: number) => `€ ${rate}/day`,
    timesheet: 'Timesheet',
    stats: 'Statistics',
    profile: 'Settings',
    statusSaving: 'Saving...',
    statusSaved: 'Saved to MongoDB',
    statusError: 'DB Error'
  },
  unsavedModal: {
    title: 'Unsaved Changes',
    desc: 'You have unsaved changes in your profile settings. Do you want to leave this page and discard your changes, or stay to save them?',
    stayBtn: 'Stay on Page',
    discardBtn: 'Discard Changes'
  },
  timeline: {
    weekdaysShort: ['SUN', 'MON', 'TUE', 'WED', 'THU', 'FRI', 'SAT'],
    months: [
      'January', 'February', 'March', 'April', 'May', 'June',
      'July', 'August', 'September', 'October', 'November', 'December'
    ],
    prevDay: 'Previous day',
    nextDay: 'Next day',
    workedBadge: 'Worked Day'
  },
  schedule: {
    title: 'Hours & Work Location',
    editActive: 'Editing active',
    editLocked: 'Preset (Click to edit)',
    unlockTooltip: 'Unlock to customize preset hours',
    lockTooltip: 'Lock schedule hours',
    dayType: 'Day Classification',
    workingDay: 'Working Day',
    nonWorkingDay: 'Non-Working Day',
    workplace: 'Work Location',
    locations: {
      inOffice: 'In Office',
      smartworking: 'Smart-working (Remote)',
      businessTrip: 'Business Trip'
    },
    morningShift: 'Morning',
    afternoonShift: 'Afternoon',
    totalCalculatedHours: 'Calculated Daily Hours:',
    hoursUnit: 'hours',
    workedConfirmedTitle: 'Confirmed Worked Day',
    workedConfirmedSub: '160 € credited to monthly total',
    markAsWorkedTitle: 'Mark Day as Worked',
    markAsWorkedSub: 'Activate the 160 € daily rate for today'
  },
  activity: {
    title: 'Daily Tasks & Objectives',
    defaultEmptyBadge: 'Empty by default',
    objectiveLabel: 'Work Objective / Ticket',
    objectivePlaceholder: 'e.g. PROJ-1042, Framework Training, Auth Refactor...',
    quickTagsLabel: 'Quick tags:',
    quickTags: ['JIRA-', 'Training', 'Refactoring', 'Bugfixing', 'Code Review', 'Meeting'],
    descriptionLabel: 'Detailed Activity Description',
    descriptionPlaceholder: 'Describe in detail the tasks you completed today for the client...',
    autoSaveHint: 'Entries are automatically persisted to local MongoDB in real-time.'
  },
  stats: {
    activeMonthLabel: 'Active Work Month',
    monthSelectorTitle: 'Click to select another month',
    monthIndexLabel: (n: number) => `Month ${n}`,
    generatingPdf: 'Generating PDF report...',
    printBtnLabel: (monthName: string) => `Export Timesheet (${monthName})`,
    monthlyProgressTitle: 'Monthly Earnings Progress',
    workedDaysOfExpected: (w: number, e: number) => `${w} days worked out of ${e} expected`,
    totalHoursWorked: 'Total Hours Worked (since contract start)',
    totalDaysWorked: 'Total Days Worked (since contract start)',
    sinceContractStart: 'since start of contract (16/10/2026)',
    liquidTargetTitle: 'Total Milestone Target',
    liquidAccumulatedTitle: 'Total Accumulated Earnings',
    liquidProgressBadge: 'Progress',
    liquidCompletedBadge: 'Target Achieved'
  },
  profile: {
    headerTitle: 'Settings & Profile',
    headerSubtitle: 'Manage application language, client company details, contractor tax information, and weekly schedule defaults.',
    saveBtn: 'Save Changes',
    languageCardTitle: 'Application Language',
    languageCardDesc: 'Choose your preferred language for the interface and generated timesheet PDF reports.',
    languages: {
      it: 'Italiano',
      en: 'English',
      ru: 'Русский'
    },
    clientCardTitle: 'Client Company Details',
    contractorCardTitle: 'Contractor Tax & Banking Data',
    contractCardTitle: 'Contract Parameters',
    presetsCardTitle: 'Default Weekly Schedule Presets',
    labels: {
      clientCompanyName: 'Client Company Name',
      clientPiva: 'VAT Number / Tax ID',
      clientAddress: 'Registered Office Address',
      clientContact: 'Contact Person / Admin Email',
      contractorName: 'Full Name / Company',
      contractorPiva: 'Contractor VAT Number (P.IVA)',
      contractorCf: 'Fiscal Code / Tax ID',
      contractorAddress: 'Billing Address',
      contractorIban: 'Bank IBAN Code',
      dailyRate: 'Agreed Daily Rate (€)',
      firstDayOfWork: 'Contract Start Date',
      firstDayHint: 'Format: YYYY-MM-DD (Contractual default: 2026-10-16)',
      revenueTarget: 'Total Revenue Target (€)',
      colDay: 'Day',
      colWorking: 'Working',
      colLocation: 'Location',
      colMorning: 'Morning',
      colAfternoon: 'Afternoon',
      colTotalHours: 'Total Hours'
    },
    weekdays: ['Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday', 'Sunday']
  },
  modal: {
    title: 'Select Target Billing Month',
    closeBtn: 'Close'
  },
  toasts: {
    mongoConnectError: 'Failed to connect to local MongoDB on localhost:27017',
    saveSuccess: 'Successfully saved to MongoDB',
    saveError: 'Error saving data to MongoDB',
    dayWorkedActive: 'Day marked as Worked (💰 +160 €)',
    dayWorkedRemoved: 'Worked day status removed',
    profileSaved: 'Profile and preferences saved successfully!',
    profileError: 'Error saving profile configuration',
    pdfSuccess: (path: string) => `Timesheet PDF successfully generated! Saved to: ${path}`,
    pdfError: 'Error generating PDF timesheet report'
  }
};
