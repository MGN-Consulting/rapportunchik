export type Locale = 'it' | 'en' | 'ru';

export interface TranslationSchema {
  nav: {
    brandSubtitle: string;
    profileSelectorAria: string;
    timesheet: string;
    stats: string;
    profile: string;
    statusSaving: string;
    statusSaved: string;
    statusError: string;
  };
  unsavedModal: {
    title: string;
    desc: string;
    stayBtn: string;
    discardBtn: string;
  };
  timeline: {
    weekdaysShort: string[]; // 7 elements starting Sunday (0) to Saturday (6)
    months: string[];        // 12 elements
    prevDay: string;
    nextDay: string;
    workedBadge: string;
  };
  schedule: {
    title: string;
    editActive: string;
    editLocked: string;
    unlockTooltip: string;
    lockTooltip: string;
    dayType: string;
    workingDay: string;
    nonWorkingDay: string;
    workplace: string;
    locations: {
      inOffice: string;
      smartworking: string;
      businessTrip: string;
    };
    morningShift: string;
    afternoonShift: string;
    totalCalculatedHours: string;
    hoursUnit: string;
    workedConfirmedTitle: string;
    workedConfirmedSub: string;
    markAsWorkedTitle: string;
    markAsWorkedSub: string;
  };
  activity: {
    title: string;
    defaultEmptyBadge: string;
    objectiveLabel: string;
    objectivePlaceholder: string;
    quickTagsLabel: string;
    quickTags: string[];
    descriptionLabel: string;
    descriptionPlaceholder: string;
    autoSaveHint: string;
  };
  stats: {
    activeMonthLabel: string;
    monthSelectorTitle: string;
    monthIndexLabel: (index: number) => string;
    generatingPdf: string;
    printBtnLabel: (monthName: string) => string;
    monthlyProgressTitle: string;
    workedDaysOfExpected: (worked: number, expected: number) => string;
    totalHoursWorked: string;
    totalDaysWorked: string;
    sinceContractStart: string;
    liquidTargetTitle: string;
    liquidAccumulatedTitle: string;
    liquidProgressBadge: string;
    liquidCompletedBadge: string;
  };
  profile: {
    headerTitle: string;
    headerSubtitle: string;
    saveBtn: string;
    languageCardTitle: string;
    languageCardDesc: string;
    languages: {
      it: string;
      en: string;
      ru: string;
    };
    clientCardTitle: string;
    contractorCardTitle: string;
    contractCardTitle: string;
    presetsCardTitle: string;
    labels: {
      clientCompanyName: string;
      clientPiva: string;
      clientAddress: string;
      clientContact: string;
      contractorName: string;
      contractorPiva: string;
      contractorCf: string;
      contractorAddress: string;
      contractorIban: string;
      dailyRate: string;
      firstDayOfWork: string;
      firstDayHint: string;
      revenueTarget: string;
      colDay: string;
      colWorking: string;
      colLocation: string;
      colMorning: string;
      colAfternoon: string;
      colTotalHours: string;
    };
    weekdays: string[]; // 7 elements: Mon..Sun or indexed
  };
  modal: {
    title: string;
    closeBtn: string;
  };
  toasts: {
    mongoConnectError: string;
    saveSuccess: string;
    saveError: string;
    dayWorkedActive: string;
    dayWorkedRemoved: string;
    profileSaved: string;
    profileError: string;
    pdfSuccess: (path: string) => string;
    pdfError: string;
  };
}
