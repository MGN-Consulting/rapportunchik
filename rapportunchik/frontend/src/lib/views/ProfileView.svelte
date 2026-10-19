<script lang="ts">
  import { appState } from '../state.svelte';
  import { i18n } from '../i18n';
  import type { Locale } from '../i18n/types';
  import {
    Building2,
    User,
    Calendar,
    Euro,
    Clock,
    Save,
    Sparkles,
    CheckCircle2,
    Shield,
    Globe
  } from 'lucide-svelte';
  import type { Profile, DayPreset } from '../../../bindings/changeme';

  // Local copy of profile for editing
  let form = $state<Profile>(
    appState.profile ? JSON.parse(JSON.stringify(appState.profile)) : ({} as Profile)
  );

  $effect(() => {
    if (appState.profile && !form.id) {
      form = JSON.parse(JSON.stringify(appState.profile));
    }
  });

  const dayKeyIndexMap: Record<string, number> = {
    '1': 0,
    '2': 1,
    '3': 2,
    '4': 3,
    '5': 4,
    '6': 5,
    '0': 6
  };

  const dayKeys = ['1', '2', '3', '4', '5', '6', '0'];

  interface ScheduleItem {
    key: string;
    name: string;
    preset: DayPreset;
  }

  let scheduleList = $derived.by<ScheduleItem[]>(() => {
    if (!form.weekdayPresets) return [];
    const list: ScheduleItem[] = [];
    for (const key of dayKeys) {
      const p = form.weekdayPresets[key];
      if (p) {
        const name = i18n.t.profile.weekdays[dayKeyIndexMap[key]] ?? key;
        list.push({ key, name, preset: p });
      }
    }
    return list;
  });

  let currentSelectedLang = $derived<Locale>(
    (form.language as Locale) || i18n.currentLocale || 'it'
  );

  function selectLanguage(loc: Locale) {
    form.language = loc;
    i18n.setLocale(loc);
  }

  function calcHours(mStart: string, mEnd: string, aStart: string, aEnd: string): number {
    const parseM = (t: string) => {
      if (!t || !t.includes(':')) return 0;
      const [h, m] = t.split(':').map(Number);
      return (h || 0) * 60 + (m || 0);
    };
    const mDiff = Math.max(0, parseM(mEnd) - parseM(mStart));
    const aDiff = Math.max(0, parseM(aEnd) - parseM(aStart));
    return Math.round(((mDiff + aDiff) / 60) * 10) / 10;
  }

  function handleScheduleRowChange(preset: DayPreset) {
    preset.totalHours = calcHours(preset.morningStart, preset.morningEnd, preset.afternoonStart, preset.afternoonEnd);
  }

  async function handleSave() {
    await appState.saveProfile(form);
  }
</script>

<div class="profile-view">
  <div class="profile-header">
    <div class="header-info">
      <h2>{i18n.t.profile.headerTitle}</h2>
      <p>{i18n.t.profile.headerSubtitle}</p>
    </div>

    <button class="save-main-btn" onclick={handleSave}>
      <Save size={16} />
      <span>{i18n.t.profile.saveBtn}</span>
    </button>
  </div>

  <div class="cards-grid">
    <!-- Language Selection Card -->
    <div class="settings-card full-width language-card">
      <div class="card-title">
        <Globe size={18} class="accent-icon" />
        <h3>{i18n.t.profile.languageCardTitle}</h3>
      </div>

      <p class="section-desc">
        {i18n.t.profile.languageCardDesc}
      </p>

      <div class="language-options">
        <button
          type="button"
          class="lang-choice-btn"
          class:selected={currentSelectedLang === 'it'}
          onclick={() => selectLanguage('it')}
        >
          <span class="flag">🇮🇹</span>
          <div class="lang-text">
            <span class="lang-title">{i18n.t.profile.languages.it}</span>
            <span class="lang-sub">Italiano</span>
          </div>
          {#if currentSelectedLang === 'it'}
            <div class="check-pill">
              <CheckCircle2 size={16} />
            </div>
          {/if}
        </button>

        <button
          type="button"
          class="lang-choice-btn"
          class:selected={currentSelectedLang === 'en'}
          onclick={() => selectLanguage('en')}
        >
          <span class="flag">🇬🇧</span>
          <div class="lang-text">
            <span class="lang-title">{i18n.t.profile.languages.en}</span>
            <span class="lang-sub">English</span>
          </div>
          {#if currentSelectedLang === 'en'}
            <div class="check-pill">
              <CheckCircle2 size={16} />
            </div>
          {/if}
        </button>

        <button
          type="button"
          class="lang-choice-btn"
          class:selected={currentSelectedLang === 'ru'}
          onclick={() => selectLanguage('ru')}
        >
          <span class="flag">🇷🇺</span>
          <div class="lang-text">
            <span class="lang-title">{i18n.t.profile.languages.ru}</span>
            <span class="lang-sub">Русский</span>
          </div>
          {#if currentSelectedLang === 'ru'}
            <div class="check-pill">
              <CheckCircle2 size={16} />
            </div>
          {/if}
        </button>
      </div>
    </div>

    <!-- Card 1: Dati Committente -->
    <div class="settings-card">
      <div class="card-title">
        <Building2 size={18} class="accent-icon" />
        <h3>{i18n.t.profile.clientCardTitle}</h3>
      </div>

      <div class="card-fields">
        <div class="field">
          <label for="client-name">{i18n.t.profile.labels.clientCompanyName}</label>
          <input
            id="client-name"
            type="text"
            placeholder="MGN Consulting S.r.l."
            bind:value={form.clientCompanyName}
          />
        </div>

        <div class="field">
          <label for="client-piva">{i18n.t.profile.labels.clientPiva}</label>
          <input
            id="client-piva"
            type="text"
            placeholder="IT12345678901"
            bind:value={form.clientPIVA}
          />
        </div>

        <div class="field">
          <label for="client-addr">{i18n.t.profile.labels.clientAddress}</label>
          <input
            id="client-addr"
            type="text"
            placeholder="Via Roma 1, Milano (MI)"
            bind:value={form.clientAddress}
          />
        </div>

        <div class="field">
          <label for="client-contact">{i18n.t.profile.labels.clientContact}</label>
          <input
            id="client-contact"
            type="text"
            placeholder="amministrazione@mgnconsulting.it"
            bind:value={form.clientContact}
          />
        </div>
      </div>
    </div>

    <!-- Card 2: Dati Collaboratore -->
    <div class="settings-card">
      <div class="card-title">
        <User size={18} class="accent-icon" />
        <h3>{i18n.t.profile.contractorCardTitle}</h3>
      </div>

      <div class="card-fields">
        <div class="field">
          <label for="contractor-name">{i18n.t.profile.labels.contractorName}</label>
          <input
            id="contractor-name"
            type="text"
            placeholder="Mario Rossi"
            bind:value={form.contractorName}
          />
        </div>

        <div class="fields-row">
          <div class="field">
            <label for="contractor-piva">{i18n.t.profile.labels.contractorPiva}</label>
            <input
              id="contractor-piva"
              type="text"
              placeholder="01234567890"
              bind:value={form.contractorPIVA}
            />
          </div>
          <div class="field">
            <label for="contractor-cf">{i18n.t.profile.labels.contractorCf}</label>
            <input
              id="contractor-cf"
              type="text"
              placeholder="RSSMRA80A01H501U"
              bind:value={form.contractorCF}
            />
          </div>
        </div>

        <div class="field">
          <label for="contractor-addr">{i18n.t.profile.labels.contractorAddress}</label>
          <input
            id="contractor-addr"
            type="text"
            placeholder="Via Garibaldi 10, Milano (MI)"
            bind:value={form.contractorAddress}
          />
        </div>

        <div class="field">
          <label for="contractor-iban">{i18n.t.profile.labels.contractorIban}</label>
          <input
            id="contractor-iban"
            type="text"
            placeholder="IT00X0000000000000000000000"
            bind:value={form.contractorIBAN}
          />
        </div>
      </div>
    </div>

    <!-- Card 3: Parametri Economici & Obiettivo -->
    <div class="settings-card full-width">
      <div class="card-title">
        <Euro size={18} class="accent-icon" />
        <h3>{i18n.t.profile.contractCardTitle}</h3>
      </div>

      <div class="contract-fields-grid">
        <div class="field">
          <label for="first-day">{i18n.t.profile.labels.firstDayOfWork}</label>
          <input
            id="first-day"
            type="date"
            bind:value={form.firstDayOfWork}
          />
          <span class="hint">{i18n.t.profile.labels.firstDayHint}</span>
        </div>

        <div class="field">
          <label for="daily-rate">{i18n.t.profile.labels.dailyRate}</label>
          <div class="input-with-affix">
            <input
              id="daily-rate"
              type="number"
              step="1"
              min="0"
              bind:value={form.dailyRate}
            />
            <span class="affix">€ / gg</span>
          </div>
          <span class="hint">
            {#if i18n.currentLocale === 'ru'}
              По умолчанию: 160 € за рабочий день
            {:else if i18n.currentLocale === 'en'}
              Default: 160 € per worked day
            {:else}
              Valore predefinito: 160 € per giorno lavorato
            {/if}
          </span>
        </div>

        <div class="field">
          <label for="target-revenue">{i18n.t.profile.labels.revenueTarget}</label>
          <div class="input-with-affix">
            <input
              id="target-revenue"
              type="number"
              step="500"
              min="1000"
              bind:value={form.revenueTarget}
            />
            <span class="affix">€ Totali</span>
          </div>
          <span class="hint">
            {#if i18n.currentLocale === 'ru'}
              Целевой объем для жидкого резервуара (по умолчанию: 20 000 €)
            {:else if i18n.currentLocale === 'en'}
              Target for the liquid wave reservoir (default: 20,000 €)
            {:else}
              Target per il serbatoio liquido (default: 20.000 €)
            {/if}
          </span>
        </div>
      </div>
    </div>

    <!-- Card 4: Orari Predefiniti per Giorno della Settimana -->
    <div class="settings-card full-width">
      <div class="card-title">
        <Clock size={18} class="accent-icon" />
        <h3>{i18n.t.profile.presetsCardTitle}</h3>
      </div>

      <p class="section-desc">
        {#if i18n.currentLocale === 'ru'}
          Эти значения используются для автозаполнения каждого дня в табеле:
        {:else if i18n.currentLocale === 'en'}
          These values are used to automatically pre-fill each day in the timesheet:
        {:else}
          Questi valori vengono utilizzati per precompilare automaticamente ogni giornata nel Rapportino:
        {/if}
      </p>

      <div class="schedule-table-wrapper">
        <table class="schedule-table">
          <thead>
            <tr>
              <th>{i18n.t.profile.labels.colDay}</th>
              <th>{i18n.t.profile.labels.colWorking}</th>
              <th>{i18n.t.profile.labels.colLocation}</th>
              <th>{i18n.t.profile.labels.colMorning}</th>
              <th>{i18n.t.profile.labels.colAfternoon}</th>
              <th>{i18n.t.profile.labels.colTotalHours}</th>
            </tr>
          </thead>
          <tbody>
            {#each scheduleList as item (item.key)}
              <tr class:is-off={!item.preset.isWorkingDay}>
                <td class="day-cell"><strong>{item.name}</strong></td>
                <td class="toggle-cell">
                  <input
                    type="checkbox"
                    bind:checked={item.preset.isWorkingDay}
                  />
                </td>
                <td>
                  <select bind:value={item.preset.workLocation}>
                    <option value="In Sede">{i18n.t.schedule.locations.inOffice}</option>
                    <option value="Smart-working">{i18n.t.schedule.locations.smartworking}</option>
                  </select>
                </td>
                <td>
                  <div class="time-pair">
                    <input
                      type="time"
                      bind:value={item.preset.morningStart}
                      onchange={() => handleScheduleRowChange(item.preset)}
                    />
                    <span>–</span>
                    <input
                      type="time"
                      bind:value={item.preset.morningEnd}
                      onchange={() => handleScheduleRowChange(item.preset)}
                    />
                  </div>
                </td>
                <td>
                  <div class="time-pair">
                    <input
                      type="time"
                      bind:value={item.preset.afternoonStart}
                      onchange={() => handleScheduleRowChange(item.preset)}
                    />
                    <span>–</span>
                    <input
                      type="time"
                      bind:value={item.preset.afternoonEnd}
                      onchange={() => handleScheduleRowChange(item.preset)}
                    />
                  </div>
                </td>
                <td class="hours-cell">
                  <span>{item.preset.totalHours} h</span>
                </td>
              </tr>
            {/each}
          </tbody>
        </table>
      </div>
    </div>
  </div>
</div>

<style>
  .profile-view {
    height: 100%;
    overflow-y: auto;
    padding: 30px 40px;
    display: flex;
    flex-direction: column;
    gap: 24px;
  }

  .profile-header {
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding-bottom: 16px;
    border-bottom: 1px solid var(--border-subtle);
  }

  .header-info h2 {
    font-size: 22px;
    font-weight: 700;
    color: var(--text-primary);
  }

  .header-info p {
    font-size: 13px;
    color: var(--text-secondary);
    margin-top: 4px;
  }

  .save-main-btn {
    display: flex;
    align-items: center;
    gap: 8px;
    padding: 10px 22px;
    background: var(--accent-surface);
    color: var(--accent-highlight);
    border: 1px solid var(--accent-primary);
    border-radius: 8px;
    font-size: 13px;
    font-weight: 700;
    box-shadow: 0 4px 14px rgba(51, 82, 255, 0.3);
  }

  .save-main-btn:hover {
    background: var(--accent-primary);
    color: #ffffff;
    transform: translateY(-2px);
    box-shadow: 0 6px 20px rgba(51, 82, 255, 0.45);
  }

  .cards-grid {
    display: grid;
    grid-template-columns: repeat(2, 1fr);
    gap: 20px;
  }

  .settings-card {
    background: rgba(21, 26, 53, 0.6);
    border: 1px solid var(--border-subtle);
    border-radius: 14px;
    padding: 22px;
    display: flex;
    flex-direction: column;
    gap: 16px;
  }

  .settings-card.full-width {
    grid-column: 1 / -1;
  }

  .language-card {
    border-color: rgba(147, 168, 255, 0.2);
    background: linear-gradient(135deg, rgba(21, 26, 53, 0.85) 0%, rgba(14, 18, 36, 0.85) 100%);
  }

  .language-options {
    display: grid;
    grid-template-columns: repeat(3, 1fr);
    gap: 16px;
  }

  .lang-choice-btn {
    display: flex;
    align-items: center;
    gap: 14px;
    padding: 14px 18px;
    background: rgba(8, 10, 20, 0.65);
    border: 1px solid var(--border-subtle);
    border-radius: 10px;
    cursor: pointer;
    transition: all 0.2s cubic-bezier(0.4, 0, 0.2, 1);
    text-align: left;
  }

  .lang-choice-btn:hover {
    border-color: rgba(147, 168, 255, 0.35);
    background: rgba(14, 18, 36, 0.85);
    transform: translateY(-2px);
  }

  .lang-choice-btn.selected {
    border-color: var(--accent-primary);
    background: rgba(51, 82, 255, 0.14);
    box-shadow: 0 0 16px rgba(51, 82, 255, 0.25);
  }

  .flag {
    font-size: 26px;
    line-height: 1;
    display: flex;
    align-items: center;
  }

  .lang-text {
    display: flex;
    flex-direction: column;
    gap: 2px;
    flex: 1;
  }

  .lang-title {
    font-size: 14px;
    font-weight: 600;
    color: var(--text-primary);
  }

  .lang-sub {
    font-size: 12px;
    color: var(--text-muted);
  }

  .check-pill {
    color: var(--accent-highlight);
    display: flex;
    align-items: center;
    justify-content: center;
  }

  .card-title {
    display: flex;
    align-items: center;
    gap: 10px;
    padding-bottom: 10px;
    border-bottom: 1px solid rgba(147, 168, 255, 0.1);
  }

  .card-title h3 {
    font-size: 15px;
    font-weight: 600;
    color: var(--text-primary);
  }

  .card-fields {
    display: flex;
    flex-direction: column;
    gap: 14px;
  }

  .fields-row {
    display: grid;
    grid-template-columns: 1fr 1fr;
    gap: 12px;
  }

  .field {
    display: flex;
    flex-direction: column;
    gap: 6px;
  }

  .field label {
    font-size: 11px;
    font-weight: 600;
    color: var(--text-secondary);
    text-transform: uppercase;
    letter-spacing: 0.4px;
  }

  .field input {
    background: rgba(8, 10, 20, 0.7);
    border: 1px solid var(--border-subtle);
    border-radius: 6px;
    padding: 9px 12px;
    color: var(--text-primary);
    font-family: inherit;
    font-size: 13px;
    outline: none;
    transition: all 0.2s ease;
  }

  .field input:focus {
    border-color: var(--accent-primary);
    background: rgba(14, 18, 36, 0.9);
    box-shadow: 0 0 10px rgba(51, 82, 255, 0.25);
  }

  .hint {
    font-size: 11px;
    color: var(--text-muted);
  }

  .contract-fields-grid {
    display: grid;
    grid-template-columns: repeat(3, 1fr);
    gap: 16px;
  }

  .input-with-affix {
    position: relative;
    display: flex;
    align-items: center;
  }

  .input-with-affix input {
    width: 100%;
    padding-right: 65px;
  }

  .affix {
    position: absolute;
    right: 12px;
    font-size: 12px;
    font-weight: 600;
    color: var(--accent-highlight);
    pointer-events: none;
  }

  .section-desc {
    font-size: 13px;
    color: var(--text-secondary);
    margin-top: -6px;
  }

  .schedule-table-wrapper {
    overflow-x: auto;
  }

  .schedule-table {
    width: 100%;
    border-collapse: collapse;
    font-size: 13px;
  }

  .schedule-table th {
    text-align: left;
    padding: 10px 12px;
    color: var(--text-muted);
    font-size: 11px;
    font-weight: 700;
    text-transform: uppercase;
    border-bottom: 1px solid var(--border-subtle);
  }

  .schedule-table td {
    padding: 10px 12px;
    border-bottom: 1px solid rgba(147, 168, 255, 0.06);
    vertical-align: middle;
  }

  .schedule-table tr.is-off {
    opacity: 0.55;
  }

  .day-cell {
    color: var(--text-primary);
  }

  .toggle-cell input[type='checkbox'] {
    width: 18px;
    height: 18px;
    accent-color: var(--accent-primary);
    cursor: pointer;
  }

  .schedule-table select {
    background: rgba(8, 10, 20, 0.8);
    border: 1px solid var(--border-subtle);
    border-radius: 6px;
    color: var(--text-primary);
    padding: 6px 10px;
    font-family: inherit;
    font-size: 12px;
    outline: none;
  }

  .time-pair {
    display: flex;
    align-items: center;
    gap: 6px;
  }

  .time-pair input[type='time'] {
    background: rgba(8, 10, 20, 0.8);
    border: 1px solid var(--border-subtle);
    border-radius: 6px;
    color: var(--text-primary);
    padding: 5px 8px;
    font-family: inherit;
    font-size: 12px;
    font-weight: 600;
    outline: none;
  }

  .hours-cell {
    font-weight: 700;
    color: var(--accent-highlight);
  }
</style>
