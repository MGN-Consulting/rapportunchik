import { it } from './it';
import { en } from './en';
import { ru } from './ru';
import type { Locale, TranslationSchema } from './types';

const dictionaries: Record<Locale, TranslationSchema> = { it, en, ru };

class I18nService {
  currentLocale = $state<Locale>('it');

  get t(): TranslationSchema {
    return dictionaries[this.currentLocale] || dictionaries.it;
  }

  setLocale(locale: Locale) {
    if (dictionaries[locale]) {
      this.currentLocale = locale;
      try {
        localStorage.setItem('rapportunchik_locale', locale);
      } catch (e) {
        console.warn('localStorage not available:', e);
      }
    }
  }

  init(initialLocale?: string) {
    if (initialLocale && initialLocale in dictionaries) {
      this.currentLocale = initialLocale as Locale;
      return;
    }
    try {
      const saved = localStorage.getItem('rapportunchik_locale') as Locale | null;
      if (saved && saved in dictionaries) {
        this.currentLocale = saved;
      }
    } catch (e) {
      console.warn('Could not read locale from localStorage:', e);
    }
  }

  getLocationLabel(canonicalLocation: string): string {
    switch (canonicalLocation) {
      case 'In Sede':
        return this.t.schedule.locations.inOffice;
      case 'Smart-working':
        return this.t.schedule.locations.smartworking;
      case 'Trasferta':
        return this.t.schedule.locations.businessTrip;
      default:
        return canonicalLocation;
    }
  }

  getMonthName(monthNum: number): string {
    if (monthNum >= 1 && monthNum <= 12) {
      return this.t.timeline.months[monthNum - 1];
    }
    return '';
  }

  formatMonthYear(year: number, month: number): string {
    const name = this.getMonthName(month);
    return `${name} ${year}`;
  }
}

export const i18n = new I18nService();
