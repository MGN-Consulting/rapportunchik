<script lang="ts">
  import { appState } from '../state.svelte';
  import { Building2, Home, Clock, Lock, Unlock, Check, Sparkles } from 'lucide-svelte';
  import type { WorkDayLog } from '../../../bindings/changeme';

  let isUnlocked = $state(false);

  function calculateHours(mStart: string, mEnd: string, aStart: string, aEnd: string): number {
    const parseMins = (t: string) => {
      if (!t || !t.includes(':')) return 0;
      const [h, m] = t.split(':').map(Number);
      return (h || 0) * 60 + (m || 0);
    };
    const mDiff = Math.max(0, parseMins(mEnd) - parseMins(mStart));
    const aDiff = Math.max(0, parseMins(aEnd) - parseMins(aStart));
    return Math.round(((mDiff + aDiff) / 60) * 10) / 10;
  }

  function handleScheduleChange() {
    if (!appState.currentDayLog) return;
    const tot = calculateHours(
      appState.currentDayLog.morningStart,
      appState.currentDayLog.morningEnd,
      appState.currentDayLog.afternoonStart,
      appState.currentDayLog.afternoonEnd
    );
    appState.currentDayLog.totalHours = tot;
    appState.debounceSaveDay(appState.currentDayLog);
  }

  function toggleLocation(loc: string) {
    if (!appState.currentDayLog) return;
    appState.currentDayLog.workLocation = loc;
    appState.debounceSaveDay(appState.currentDayLog);
  }

  function toggleWorkingDay() {
    if (!appState.currentDayLog) return;
    appState.currentDayLog.isWorkingDay = !appState.currentDayLog.isWorkingDay;
    appState.debounceSaveDay(appState.currentDayLog);
  }
</script>

<div class="panel schedule-panel">
  <div class="panel-header">
    <div class="header-title">
      <Clock size={18} class="accent-icon" />
      <h3>Orari & Luogo di Lavoro</h3>
    </div>

    <button
      class="unlock-btn"
      class:unlocked={isUnlocked}
      onclick={() => (isUnlocked = !isUnlocked)}
      title={isUnlocked ? 'Blocca orari' : 'Sblocca per modificare i valori precompilati'}
    >
      {#if isUnlocked}
        <Unlock size={14} />
        <span>Modifica attiva</span>
      {:else}
        <Lock size={14} />
        <span>Precompilato (Clicca per modificare)</span>
      {/if}
    </button>
  </div>

  {#if appState.currentDayLog}
    <div class="panel-body" class:greyed-out={!isUnlocked}>
      <!-- Working day toggle -->
      <div class="form-row toggle-row">
        <label for="working-day-toggle" class="row-label">Tipologia Giornata</label>
        <button
          id="working-day-toggle"
          class="switch-btn"
          class:active={appState.currentDayLog.isWorkingDay}
          onclick={toggleWorkingDay}
        >
          <span class="switch-ball"></span>
          <span class="switch-text">
            {appState.currentDayLog.isWorkingDay ? 'Giorno Lavorativo' : 'Giorno Non Lavorativo'}
          </span>
        </button>
      </div>

      <!-- Workplace selector -->
      <div class="form-row">
        <span class="row-label">Luogo di Lavoro</span>
        <div class="location-group">
          <button
            class="loc-btn"
            class:active={appState.currentDayLog.workLocation === 'In Sede'}
            onclick={() => toggleLocation('In Sede')}
          >
            <Building2 size={16} />
            <span>In Sede (Ufficio)</span>
          </button>

          <button
            class="loc-btn"
            class:active={appState.currentDayLog.workLocation === 'Smart-working'}
            onclick={() => toggleLocation('Smart-working')}
          >
            <Home size={16} />
            <span>Smart-working (Remoto)</span>
          </button>
        </div>
      </div>

      <!-- Shifts time inputs -->
      <div class="shifts-grid">
        <!-- Morning shift -->
        <div class="shift-card">
          <div class="shift-label">Mattina</div>
          <div class="time-range">
            <input
              type="time"
              bind:value={appState.currentDayLog.morningStart}
              onchange={handleScheduleChange}
              disabled={!isUnlocked}
            />
            <span class="sep">–</span>
            <input
              type="time"
              bind:value={appState.currentDayLog.morningEnd}
              onchange={handleScheduleChange}
              disabled={!isUnlocked}
            />
          </div>
        </div>

        <!-- Afternoon shift -->
        <div class="shift-card">
          <div class="shift-label">Pomeriggio</div>
          <div class="time-range">
            <input
              type="time"
              bind:value={appState.currentDayLog.afternoonStart}
              onchange={handleScheduleChange}
              disabled={!isUnlocked}
            />
            <span class="sep">–</span>
            <input
              type="time"
              bind:value={appState.currentDayLog.afternoonEnd}
              onchange={handleScheduleChange}
              disabled={!isUnlocked}
            />
          </div>
        </div>
      </div>

      <!-- Total daily hours calculated -->
      <div class="hours-summary">
        <span class="hours-label">Totale Ore Lavorative Calcolate:</span>
        <span class="hours-value">{appState.currentDayLog.totalHours} ore</span>
      </div>
    </div>

    <!-- Mark as Worked Button -->
    <div class="work-action-zone">
      <button
        class="worked-toggle-btn"
        class:is-worked={appState.currentDayLog.isWorked}
        onclick={() => appState.toggleDayWorked()}
      >
        <div class="btn-content">
          {#if appState.currentDayLog.isWorked}
            <div class="check-icon">
              <Check size={18} strokeWidth={3} />
            </div>
            <div class="btn-text">
              <span class="main-title">Giorno Lavorato Confermato</span>
              <span class="sub-title">160 € accreditati al calcolo mensile</span>
            </div>
            <span class="cash-emoji">💰</span>
          {:else}
            <Sparkles size={18} />
            <div class="btn-text">
              <span class="main-title">Segna Giorno come Lavorato</span>
              <span class="sub-title">Attiva il compenso di 160 € per oggi</span>
            </div>
            <span class="cash-emoji">💰</span>
          {/if}
        </div>
      </button>
    </div>
  {/if}
</div>

<style>
  .schedule-panel {
    display: flex;
    flex-direction: column;
    height: 100%;
    padding: 24px;
    background: rgba(14, 18, 36, 0.4);
  }

  .panel-header {
    display: flex;
    align-items: center;
    justify-content: space-between;
    margin-bottom: 20px;
    padding-bottom: 12px;
    border-bottom: 1px solid var(--border-subtle);
  }

  .header-title {
    display: flex;
    align-items: center;
    gap: 10px;
  }

  .header-title h3 {
    font-size: 16px;
    font-weight: 600;
    color: var(--text-primary);
  }

  :global(.accent-icon) {
    color: var(--accent-highlight);
  }

  .unlock-btn {
    display: flex;
    align-items: center;
    gap: 6px;
    padding: 6px 12px;
    border-radius: 6px;
    font-size: 12px;
    background: rgba(147, 168, 255, 0.08);
    color: var(--text-secondary);
    border: 1px solid var(--border-subtle);
  }

  .unlock-btn:hover {
    background: rgba(147, 168, 255, 0.15);
    color: var(--text-primary);
  }

  .unlock-btn.unlocked {
    background: var(--accent-surface);
    color: var(--accent-highlight);
    border-color: var(--accent-primary);
  }

  .panel-body {
    display: flex;
    flex-direction: column;
    gap: 18px;
    transition: opacity 0.2s ease;
  }

  .panel-body.greyed-out {
    opacity: 0.68;
  }

  .form-row {
    display: flex;
    flex-direction: column;
    gap: 8px;
  }

  .row-label {
    font-size: 12px;
    font-weight: 600;
    color: var(--text-secondary);
    text-transform: uppercase;
    letter-spacing: 0.5px;
  }

  .toggle-row {
    flex-direction: row;
    align-items: center;
    justify-content: space-between;
  }

  .switch-btn {
    display: flex;
    align-items: center;
    gap: 10px;
    padding: 6px 14px;
    background: rgba(21, 26, 53, 0.8);
    border: 1px solid var(--border-subtle);
    border-radius: 20px;
    color: var(--text-secondary);
    font-size: 13px;
  }

  .switch-btn.active {
    background: rgba(16, 185, 129, 0.15);
    border-color: rgba(16, 185, 129, 0.4);
    color: #34d399;
  }

  .switch-ball {
    width: 12px;
    height: 12px;
    border-radius: 50%;
    background: var(--text-muted);
    transition: all 0.2s ease;
  }

  .switch-btn.active .switch-ball {
    background: #10b981;
    box-shadow: 0 0 8px #10b981;
  }

  .location-group {
    display: grid;
    grid-template-columns: 1fr 1fr;
    gap: 10px;
  }

  .loc-btn {
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 8px;
    padding: 10px 14px;
    border-radius: 8px;
    background: rgba(21, 26, 53, 0.6);
    border: 1px solid var(--border-subtle);
    color: var(--text-secondary);
    font-size: 13px;
    font-weight: 500;
  }

  .loc-btn:hover {
    background: rgba(30, 38, 77, 0.8);
    color: var(--text-primary);
  }

  .loc-btn.active {
    background: var(--accent-surface);
    color: var(--accent-highlight);
    border-color: var(--accent-primary);
    box-shadow: 0 0 12px rgba(51, 82, 255, 0.25);
  }

  .shifts-grid {
    display: grid;
    grid-template-columns: 1fr 1fr;
    gap: 12px;
  }

  .shift-card {
    background: rgba(21, 26, 53, 0.6);
    border: 1px solid var(--border-subtle);
    border-radius: 8px;
    padding: 10px 14px;
    display: flex;
    flex-direction: column;
    gap: 6px;
  }

  .shift-label {
    font-size: 11px;
    font-weight: 600;
    color: var(--text-muted);
    text-transform: uppercase;
  }

  .time-range {
    display: flex;
    align-items: center;
    gap: 6px;
  }

  .time-range input[type='time'] {
    flex: 1;
    background: rgba(8, 10, 20, 0.6);
    border: 1px solid var(--border-subtle);
    border-radius: 6px;
    color: var(--text-primary);
    padding: 6px 8px;
    font-family: inherit;
    font-size: 13px;
    font-weight: 600;
    outline: none;
  }

  .time-range input[type='time']:focus {
    border-color: var(--accent-primary);
  }

  .sep {
    color: var(--text-muted);
  }

  .hours-summary {
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 12px 16px;
    background: rgba(51, 82, 255, 0.08);
    border: 1px solid rgba(51, 82, 255, 0.2);
    border-radius: 8px;
  }

  .hours-label {
    font-size: 13px;
    color: var(--text-secondary);
  }

  .hours-value {
    font-size: 16px;
    font-weight: 700;
    color: var(--accent-highlight);
  }

  .work-action-zone {
    margin-top: auto;
    padding-top: 20px;
  }

  .worked-toggle-btn {
    width: 100%;
    padding: 14px 18px;
    border-radius: 12px;
    background: rgba(21, 26, 53, 0.9);
    border: 2px solid var(--border-subtle);
    color: var(--text-primary);
    transition: all 0.25s cubic-bezier(0.4, 0, 0.2, 1);
  }

  .worked-toggle-btn:hover {
    border-color: var(--accent-primary);
    transform: translateY(-2px);
    box-shadow: 0 6px 20px rgba(0, 0, 0, 0.3);
  }

  .worked-toggle-btn.is-worked {
    background: linear-gradient(135deg, hsl(234, 54%, 22%) 0%, hsl(234, 100%, 35%) 100%);
    border-color: var(--accent-primary);
    box-shadow: 0 6px 24px rgba(51, 82, 255, 0.35);
  }

  .btn-content {
    display: flex;
    align-items: center;
    gap: 14px;
    text-align: left;
  }

  .check-icon {
    display: flex;
    align-items: center;
    justify-content: center;
    width: 28px;
    height: 28px;
    border-radius: 50%;
    background: #10b981;
    color: #ffffff;
    box-shadow: 0 0 10px rgba(16, 185, 129, 0.5);
  }

  .btn-text {
    flex: 1;
    display: flex;
    flex-direction: column;
    gap: 2px;
  }

  .main-title {
    font-size: 14px;
    font-weight: 700;
    color: #ffffff;
  }

  .sub-title {
    font-size: 12px;
    color: var(--accent-highlight);
    opacity: 0.9;
  }

  .cash-emoji {
    font-size: 24px;
    filter: drop-shadow(0 0 8px rgba(255, 215, 0, 0.4));
  }
</style>
