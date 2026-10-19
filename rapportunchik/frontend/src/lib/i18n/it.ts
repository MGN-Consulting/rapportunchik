import type { TranslationSchema } from './types';

export const it: TranslationSchema = {
  nav: {
    brandSubtitle: 'Collaborazione P.IVA',
    profileSelectorAria: 'Seleziona Profilo Committente',
    timesheet: 'Rapportino',
    stats: 'Statistiche',
    profile: 'Impostazioni',
    statusSaving: 'Salvataggio...',
    statusSaved: 'Salvato su MongoDB',
    statusError: 'Errore DB'
  },
  unsavedModal: {
    title: 'Modifiche non salvate',
    desc: 'Hai delle modifiche non salvate nelle impostazioni del profilo. Vuoi abbandonare la pagina e perdere le modifiche, oppure restare per salvarle?',
    stayBtn: 'Resta sulla pagina',
    discardBtn: 'Abbandona modifiche'
  },
  timeline: {
    weekdaysShort: ['DOM', 'LUN', 'MAR', 'MER', 'GIO', 'VEN', 'SAB'],
    months: [
      'Gennaio', 'Febbraio', 'Marzo', 'Aprile', 'Maggio', 'Giugno',
      'Luglio', 'Agosto', 'Settembre', 'Ottobre', 'Novembre', 'Dicembre'
    ],
    prevDay: 'Giorno precedente',
    nextDay: 'Giorno successivo',
    workedBadge: 'Giorno Lavorato'
  },
  schedule: {
    title: 'Orari & Luogo di Lavoro',
    editActive: 'Modifica attiva',
    editLocked: 'Precompilato (Clicca per modificare)',
    unlockTooltip: 'Sblocca per modificare i valori precompilati',
    lockTooltip: 'Blocca orari',
    dayType: 'Tipologia Giornata',
    workingDay: 'Giorno Lavorativo',
    nonWorkingDay: 'Giorno Non Lavorativo',
    workplace: 'Luogo di Lavoro',
    locations: {
      inOffice: 'In Sede (Ufficio)',
      smartworking: 'Smart-working (Remoto)',
      businessTrip: 'Trasferta'
    },
    morningShift: 'Mattina',
    afternoonShift: 'Pomeriggio',
    totalCalculatedHours: 'Totale Ore Lavorative Calcolate:',
    hoursUnit: 'ore',
    workedConfirmedTitle: 'Giorno Lavorato Confermato',
    workedConfirmedSub: '160 € accreditati al calcolo mensile',
    markAsWorkedTitle: 'Segna Giorno come Lavorato',
    markAsWorkedSub: 'Attiva il compenso di 160 € per oggi'
  },
  activity: {
    title: 'Attività & Obiettivo della Giornata',
    defaultEmptyBadge: 'Vuoto di default',
    objectiveLabel: 'Obiettivo / Task di Lavoro',
    objectivePlaceholder: 'Es. PROJ-1042, Formazione, Refactoring modulo auth...',
    quickTagsLabel: 'Suggeriti:',
    quickTags: ['JIRA-', 'Formazione', 'Refactoring', 'Bugfixing', 'Code Review', 'Allineamento'],
    descriptionLabel: 'Descrizione Dettagliata Attività',
    descriptionPlaceholder: 'Descrivi in dettaglio cosa hai svolto oggi per il committente...',
    autoSaveHint: 'I dati vengono salvati automaticamente su MongoDB locale ad ogni digitazione.'
  },
  stats: {
    activeMonthLabel: 'Mese di Lavoro Attivo',
    monthSelectorTitle: 'Clicca per selezionare un altro mese',
    monthIndexLabel: (n: number) => `${n}° Mese`,
    generatingPdf: 'Generazione PDF in corso...',
    printBtnLabel: (monthName: string) => `Stampa Rapportino (${monthName})`,
    monthlyProgressTitle: 'Progresso Incasso Mensile',
    workedDaysOfExpected: (w: number, e: number) => `${w} giorni lavorati su ${e} previsti`,
    totalHoursWorked: 'Ore Totali Lavorate (da inizio collaborazione)',
    totalDaysWorked: 'Giorni Totali Lavorati (da inizio collaborazione)',
    sinceContractStart: 'dall\'inizio del contratto (16/10/2026)',
    liquidTargetTitle: 'Target Finanziario Totale',
    liquidAccumulatedTitle: 'Compenso Totale Accumulato',
    liquidProgressBadge: 'Avanzamento',
    liquidCompletedBadge: 'Target Raggiunto'
  },
  profile: {
    headerTitle: 'Impostazioni & Profilo',
    headerSubtitle: 'Gestisci la lingua dell\'applicazione, i dati fiscali del Committente e del Collaboratore, e gli orari settimanali predefiniti.',
    saveBtn: 'Salva Modifiche',
    languageCardTitle: 'Lingua dell\'Applicazione',
    languageCardDesc: 'Seleziona la lingua per l\'interfaccia utente e per l\'esportazione del Rapportino in PDF.',
    languages: {
      it: 'Italiano',
      en: 'English',
      ru: 'Русский'
    },
    clientCardTitle: 'Dati Committente (Azienda)',
    contractorCardTitle: 'Dati Collaboratore (Partita IVA)',
    contractCardTitle: 'Parametri Contrattuali',
    presetsCardTitle: 'Orari Predefiniti Settimanali',
    labels: {
      clientCompanyName: 'Ragione Sociale Committente',
      clientPiva: 'Partita IVA / Codice Fiscale',
      clientAddress: 'Indirizzo Sede Legale',
      clientContact: 'Email Referente / Amministrazione',
      contractorName: 'Nome & Cognome / Ditta',
      contractorPiva: 'Partita IVA Collaboratore',
      contractorCf: 'Codice Fiscale Collaboratore',
      contractorAddress: 'Indirizzo di Fatturazione',
      contractorIban: 'Codice IBAN per Bonifici',
      dailyRate: 'Tariffa Giornaliera Concordata (€)',
      firstDayOfWork: 'Data Inizio Collaborazione',
      firstDayHint: 'Formato: YYYY-MM-DD (Default contrattuale: 2026-10-16)',
      revenueTarget: 'Obiettivo Compenso Totale (€)',
      colDay: 'Giorno',
      colWorking: 'Lavorativo',
      colLocation: 'Luogo',
      colMorning: 'Mattina',
      colAfternoon: 'Pomeriggio',
      colTotalHours: 'Ore Totali'
    },
    weekdays: ['Lunedì', 'Martedì', 'Mercoledì', 'Giovedì', 'Venerdì', 'Sabato', 'Domenica']
  },
  modal: {
    title: 'Seleziona Mese di Riferimento',
    closeBtn: 'Chiudi'
  },
  toasts: {
    mongoConnectError: 'Errore di connessione a MongoDB su localhost:27017',
    saveSuccess: 'Salvato con successo su MongoDB',
    saveError: 'Errore durante il salvataggio su MongoDB',
    dayWorkedActive: 'Giorno contrassegnato come Lavorato (💰 +160 €)',
    dayWorkedRemoved: 'Contrassegno giorno lavorato rimosso',
    profileSaved: 'Profilo e impostazioni salvate con successo!',
    profileError: 'Errore nel salvataggio del profilo',
    pdfSuccess: (path: string) => `Rapportino generato con successo! Salvato in: ${path}`,
    pdfError: 'Errore durante la generazione del PDF'
  }
};
