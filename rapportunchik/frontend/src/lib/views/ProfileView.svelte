<script lang="ts">
  import { appState } from '../state.svelte';
  import {
    Building2,
    User,
    Calendar,
    Euro,
    Clock,
    Save,
    Sparkles,
    CheckCircle2,
    Shield
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

  const dayNames = [
    { key: '1', name: 'Lunedì' },
    { key: '2', name: 'Martedì' },
    { key: '3', name: 'Mercoledì' },
    { key: '4', name: 'Giovedì' },
    { key: '5', name: 'Venerdì' },
    { key: '6', name: 'Sabato' },
    { key: '0', name: 'Domenica' }
  ];

  interface ScheduleItem {
    key: string;
    name: string;
    preset: DayPreset;
  }

  let scheduleList = $derived.by<ScheduleItem[]>(() => {
    if (!form.weekdayPresets) return [];
    const list: ScheduleItem[] = [];
    for (const d of dayNames) {
      const p = form.weekdayPresets[d.key];
      if (p) {
        list.push({ key: d.key, name: d.name, preset: p });
      }
    }
    return list;
  });

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
      <h2>Impostazioni & Profilo</h2>
      <p>Gestisci i dati fiscali del Committente, del Collaboratore e le regole di orario predefinite.</p>
    </div>

    <button class="save-main-btn" onclick={handleSave}>
      <Save size={16} />
      <span>Salva Modifiche</span>
    </button>
  </div>

  <div class="cards-grid">
    <!-- Card 1: Dati Committente -->
    <div class="settings-card">
      <div class="card-title">
        <Building2 size={18} class="accent-icon" />
        <h3>Dati Committente (Azienda)</h3>
      </div>

      <div class="card-fields">
        <div class="field">
          <label for="client-name">Ragione Sociale Committente</label>
          <input
            id="client-name"
            type="text"
            placeholder="MGN Consulting S.r.l."
            bind:value={form.clientCompanyName}
          />
        </div>

        <div class="field">
          <label for="client-piva">Partita IVA / Codice Fiscale</label>
          <input
            id="client-piva"
            type="text"
            placeholder="IT12345678901"
            bind:value={form.clientPIVA}
          />
        </div>

        <div class="field">
          <label for="client-addr">Indirizzo Sede Legale</label>
          <input
            id="client-addr"
            type="text"
            placeholder="Via Roma 1, Milano (MI)"
            bind:value={form.clientAddress}
          />
        </div>

        <div class="field">
          <label for="client-contact">Referente / Email Amministrazione</label>
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
        <h3>Dati Collaboratore (Partita IVA)</h3>
      </div>

      <div class="card-fields">
        <div class="field">
          <label for="contractor-name">Nome e Cognome</label>
          <input
            id="contractor-name"
            type="text"
            placeholder="Mario Rossi"
            bind:value={form.contractorName}
          />
        </div>

        <div class="fields-row">
          <div class="field">
            <label for="contractor-piva">Partita IVA</label>
            <input
              id="contractor-piva"
              type="text"
              placeholder="01234567890"
              bind:value={form.contractorPIVA}
            />
          </div>
          <div class="field">
            <label for="contractor-cf">Codice Fiscale</label>
            <input
              id="contractor-cf"
              type="text"
              placeholder="RSSMRA80A01H501U"
              bind:value={form.contractorCF}
            />
          </div>
        </div>

        <div class="field">
          <label for="contractor-addr">Indirizzo di Fatturazione</label>
          <input
            id="contractor-addr"
            type="text"
            placeholder="Via Garibaldi 10, Milano (MI)"
            bind:value={form.contractorAddress}
          />
        </div>

        <div class="field">
          <label for="contractor-iban">Coordinate di Pagamento (IBAN)</label>
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
        <h3>Parametri di Contratto & Target Finanziario</h3>
      </div>

      <div class="contract-fields-grid">
        <div class="field">
          <label for="first-day">Data Inizio Lavoro</label>
          <input
            id="first-day"
            type="date"
            bind:value={form.firstDayOfWork}
          />
          <span class="hint">Inizio collaborazione stabilito al 16/10/2026</span>
        </div>

        <div class="field">
          <label for="daily-rate">Tariffa Giornaliera (€ / giorno)</label>
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
          <span class="hint">Valore predefinito: 160 € per giorno lavorato</span>
        </div>

        <div class="field">
          <label for="target-revenue">Obiettivo Fatturato Complessivo (€)</label>
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
          <span class="hint">Target per il serbatoio liquido (default: 20.000 €)</span>
        </div>
      </div>
    </div>

    <!-- Card 4: Orari Predefiniti per Giorno della Settimana -->
    <div class="settings-card full-width">
      <div class="card-title">
        <Clock size={18} class="accent-icon" />
        <h3>Regole Orarie Predefinite per Giorno della Settimana</h3>
      </div>

      <p class="section-desc">
        Questi valori vengono utilizzati per precompilare automaticamente ogni giornata nel Rapportino:
      </p>

      <div class="schedule-table-wrapper">
        <table class="schedule-table">
          <thead>
            <tr>
              <th>Giorno</th>
              <th>Lavorativo</th>
              <th>Luogo Predefinito</th>
              <th>Fascia Mattina</th>
              <th>Fascia Pomeriggio</th>
              <th>Ore Totali</th>
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
                    <option value="In Sede">In Sede (Ufficio)</option>
                    <option value="Smart-working">Smart-working (Remoto)</option>
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
