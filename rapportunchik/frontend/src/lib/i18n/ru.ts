import type { TranslationSchema } from './types';

export const ru: TranslationSchema = {
  nav: {
    brandSubtitle: 'Табель учета времени',
    profileSelectorAria: 'Выберите профиль заказчика',
    profileModalTitle: 'Профили заказчиков',
    profileModalDesc: 'Выберите активного заказчика для просмотра табеля или добавьте нового:',
    addNewProfileBtn: '+ Добавить нового заказчика',
    activeProfileBadge: 'Активен',
    dailyRateBadge: (rate: number) => `€ ${rate}/день`,
    timesheet: 'Табель',
    stats: 'Статистика',
    profile: 'Настройки',
    statusSaving: 'Сохранение...',
    statusSaved: 'Сохранено в MongoDB',
    statusError: 'Ошибка БД'
  },
  unsavedModal: {
    title: 'Несохраненные изменения',
    desc: 'В настройках профиля есть несохраненные изменения. Вы хотите покинуть страницу и сбросить их, или остаться, чтобы сохранить?',
    stayBtn: 'Остаться на странице',
    discardBtn: 'Сбросить изменения'
  },
  timeline: {
    weekdaysShort: ['ВОС', 'ПОН', 'ВТО', 'СРЕ', 'ЧЕТ', 'ПЯТ', 'СУБ'],
    months: [
      'Январь', 'Февраль', 'Март', 'Апрель', 'Май', 'Июнь',
      'Июль', 'Август', 'Сентябрь', 'Октябрь', 'Ноябрь', 'Декабрь'
    ],
    prevDay: 'Предыдущий день',
    nextDay: 'Следующий день',
    workedBadge: 'Отработанный день'
  },
  schedule: {
    title: 'Рабочие часы и место работы',
    editActive: 'Редактирование активно',
    editLocked: 'По умолчанию (Нажмите для изменения)',
    unlockTooltip: 'Разблокировать для настройки предустановленных часов',
    lockTooltip: 'Заблокировать часы графика',
    dayType: 'Тип дня',
    workingDay: 'Рабочий день',
    nonWorkingDay: 'Выходной день',
    workplace: 'Место работы',
    locations: {
      inOffice: 'В офисе (In Sede)',
      smartworking: 'Удаленно (Smart-working)',
      businessTrip: 'Командировка (Trasferta)'
    },
    morningShift: 'Утро',
    afternoonShift: 'День/Вечер',
    totalCalculatedHours: 'Всего рабочих часов за день:',
    hoursUnit: 'ч.',
    workedConfirmedTitle: 'День отмечен как отработанный',
    workedConfirmedSub: '160 € начислено в расчет за месяц',
    markAsWorkedTitle: 'Отметить день как отработанный',
    markAsWorkedSub: 'Активировать дневную ставку 160 € на сегодня'
  },
  activity: {
    title: 'Задачи и цели рабочего дня',
    defaultEmptyBadge: 'По умолчанию пусто',
    objectiveLabel: 'Цель / Номер задачи',
    objectivePlaceholder: 'Напр. PROJ-1042, Обучение Wails, Рефакторинг модуля авторизации...',
    quickTagsLabel: 'Быстрые теги:',
    quickTags: ['JIRA-', 'Обучение', 'Рефакторинг', 'Багфикс', 'Код-ревью', 'Созвон'],
    descriptionLabel: 'Подробное описание деятельности',
    descriptionPlaceholder: 'Опишите подробно, какие задачи вы выполнили сегодня для заказчика...',
    autoSaveHint: 'Данные автоматически сохраняются в локальную MongoDB в реальном времени.'
  },
  stats: {
    activeMonthLabel: 'Текущий расчетный месяц',
    monthSelectorTitle: 'Нажмите для выбора другого месяца',
    monthIndexLabel: (n: number) => `${n}-й Месяц`,
    generatingPdf: 'Формирование отчета в PDF...',
    printBtnLabel: (monthName: string) => `Печать табеля (${monthName})`,
    monthlyProgressTitle: 'Прогресс дохода за месяц',
    workedDaysOfExpected: (w: number, e: number) => `${w} отработано из ${e} запланированных дней`,
    totalHoursWorked: 'Всего отработано часов (за все время)',
    totalDaysWorked: 'Всего отработано дней (за все время)',
    sinceContractStart: 'с начала сотрудничества (16.10.2026)',
    liquidTargetTitle: 'Общая целевая сумма',
    liquidAccumulatedTitle: 'Итого накопленный доход',
    liquidProgressBadge: 'Прогресс',
    liquidCompletedBadge: 'Цель достигнута'
  },
  profile: {
    headerTitle: 'Настройки и профиль',
    headerSubtitle: 'Управление языком интерфейса, реквизитами заказчика и исполнителя, а также недельным графиком по умолчанию.',
    saveBtn: 'Сохранить изменения',
    languageCardTitle: 'Язык приложения',
    languageCardDesc: 'Выберите предпочитаемый язык интерфейса и экспортируемых PDF-отчетов.',
    languages: {
      it: 'Italiano',
      en: 'English',
      ru: 'Русский'
    },
    clientCardTitle: 'Данные заказчика (Компания)',
    contractorCardTitle: 'Данные исполнителя (Partita IVA / ИП)',
    contractCardTitle: 'Условия контракта',
    presetsCardTitle: 'Еженедельный график по умолчанию',
    labels: {
      clientCompanyName: 'Название компании заказчика',
      clientPiva: 'Номер плательщика НДС / ИНН (P.IVA)',
      clientAddress: 'Юридический адрес заказчика',
      clientContact: 'Email контактного лица / бухгалтерии',
      contractorName: 'ФИО / Наименование исполнителя',
      contractorPiva: 'Partita IVA (Номер НДС)',
      contractorCf: 'Codice Fiscale (Налоговый номер)',
      contractorAddress: 'Адрес выставления счетов',
      contractorIban: 'Банковский счет (IBAN)',
      dailyRate: 'Согласованная дневная ставка (€)',
      firstDayOfWork: 'Дата начала сотрудничества',
      firstDayHint: 'Формат: YYYY-MM-DD (По договору: 2026-10-16)',
      revenueTarget: 'Целевой совокупный доход (€)',
      colDay: 'День',
      colWorking: 'Рабочий',
      colLocation: 'Место',
      colMorning: 'Утро',
      colAfternoon: 'День/Вечер',
      colTotalHours: 'Всего часов'
    },
    weekdays: ['Понедельник', 'Вторник', 'Среда', 'Четверг', 'Пятница', 'Суббота', 'Воскресенье']
  },
  modal: {
    title: 'Выбор расчетного месяца',
    closeBtn: 'Закрыть'
  },
  toasts: {
    mongoConnectError: 'Ошибка подключения к MongoDB на localhost:27017',
    saveSuccess: 'Успешно сохранено в MongoDB',
    saveError: 'Ошибка при сохранении в MongoDB',
    dayWorkedActive: 'День отмечен как отработанный (💰 +160 €)',
    dayWorkedRemoved: 'Отметка об отработанном дне снята',
    profileSaved: 'Профиль и настройки успешно сохранены!',
    profileError: 'Ошибка при сохранении профиля',
    pdfSuccess: (path: string) => `Табель PDF сформирован! Сохранен в: ${path}`,
    pdfError: 'Ошибка при генерации PDF табеля'
  }
};
