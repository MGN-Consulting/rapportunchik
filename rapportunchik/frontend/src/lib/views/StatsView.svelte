<script lang="ts">
  import { appState } from '../state.svelte';
  import { i18n } from '../i18n';
  import { Printer, Calendar, Clock, Award, Loader2, ChevronDown } from 'lucide-svelte';
  import LiquidContainer from '../components/LiquidContainer.svelte';
  import MonthSelectorModal from '../components/MonthSelectorModal.svelte';

  // Calculate monthly cash progress percentage
  let monthlyExpected = $derived(appState.monthSummary?.expectedMonthlyRevenue || 0);
  let monthlyCurrent = $derived(appState.monthSummary?.currentMonthlyRevenue || 0);
  let monthlyPercent = $derived(
    monthlyExpected > 0 ? Math.min(100, Math.round((monthlyCurrent / monthlyExpected) * 100)) : 0
  );
</script>

<div class="stats-view">
  <!-- Left Half -->
  <div class="stats-left">
    <!-- Top Left: Current number of months worked -->
    <div class="month-indicator-zone">
      <div class="zone-label">{i18n.t.stats.activeMonthLabel}</div>
      <button
        class="month-selector-btn"
        onclick={() => (appState.isMonthModalOpen = true)}
        title={i18n.t.stats.monthSelectorTitle}
      >
        <div class="big-month-number">
          {i18n.t.stats.monthIndexLabel(appState.monthSummary?.monthIndex || 1)}
        </div>
        <div class="month-name-sub">
          <span>{i18n.getMonthName(appState.selectedMonth)} {appState.selectedYear}</span>
          <ChevronDown size={18} class="chevron-icon" />
        </div>
      </button>
    </div>

    <!-- Rounded Print Button with print icon -->
    <div class="print-zone">
      <button
        class="print-btn"
        onclick={() => appState.printMonthPDF()}
        disabled={appState.isPrinting}
      >
        {#if appState.isPrinting}
          <Loader2 size={18} class="spin" />
          <span>{i18n.t.stats.generatingPdf}</span>
        {:else}
          <Printer size={18} />
          <span>{i18n.t.stats.printBtnLabel(i18n.getMonthName(appState.selectedMonth))}</span>
        {/if}
      </button>
    </div>

    <!-- Monthly Cash Progress Bar -->
    <div class="progress-zone">
      <div class="progress-labels">
        <span class="p-title">{i18n.t.stats.monthlyProgressTitle}</span>
        <span class="p-val">
          € {monthlyCurrent.toLocaleString(i18n.currentLocale === 'ru' ? 'ru-RU' : i18n.currentLocale === 'en' ? 'en-US' : 'it-IT', { minimumFractionDigits: 2 })} / €{' '}
          {monthlyExpected.toLocaleString(i18n.currentLocale === 'ru' ? 'ru-RU' : i18n.currentLocale === 'en' ? 'en-US' : 'it-IT', { minimumFractionDigits: 2 })}
        </span>
      </div>

      <div class="progress-track">
        <div
          class="progress-fill"
          style="width: {monthlyPercent}%;"
        >
          <div class="progress-glow"></div>
        </div>
      </div>

      <div class="progress-caption">
        <span>{i18n.t.stats.workedDaysOfExpected(appState.monthSummary?.workedDays || 0, appState.monthSummary?.expectedWorkingDays || 0)}</span>
        <span class="p-pct">{monthlyPercent}%</span>
      </div>
    </div>

    <!-- 2 Big Numbers Stacked One Under The Other -->
    <div class="metrics-stacked">
      <!-- Metric 1: Absolute Total Hours Worked -->
      <div class="big-metric-card">
        <div class="metric-icon-box">
          <Clock size={24} />
        </div>
        <div class="metric-details">
          <div class="metric-num">
            {(appState.overallStats?.totalHoursWorked || 0).toLocaleString(i18n.currentLocale === 'ru' ? 'ru-RU' : i18n.currentLocale === 'en' ? 'en-US' : 'it-IT', {
              minimumFractionDigits: 1,
              maximumFractionDigits: 1
            })}
          </div>
          <div class="metric-label">{i18n.t.stats.totalHoursWorked}</div>
        </div>
      </div>

      <!-- Metric 2: Absolute Total Days Worked -->
      <div class="big-metric-card">
        <div class="metric-icon-box">
          <Award size={24} />
        </div>
        <div class="metric-details">
          <div class="metric-num">
            {appState.overallStats?.totalDaysWorked || 0}
          </div>
          <div class="metric-label">{i18n.t.stats.totalDaysWorked}</div>
        </div>
      </div>
    </div>
  </div>

  <!-- Vertical separator in 2 halves -->
  <div class="vertical-divider" role="separator" aria-orientation="vertical">
    <div class="divider-line"></div>
  </div>

  <!-- Right Half: Liquid Wave Container for 20K target -->
  <div class="stats-right">
    <LiquidContainer />
  </div>
</div>

<MonthSelectorModal />

<style>
  .stats-view {
    display: flex;
    height: 100%;
    overflow: hidden;
    position: relative;
  }

  .stats-left {
    flex: 1;
    min-width: 440px;
    padding: 32px 36px;
    display: flex;
    flex-direction: column;
    gap: 24px;
    overflow-y: auto;
  }

  .vertical-divider {
    width: 1px;
    background: var(--border-subtle);
    position: relative;
    display: flex;
    align-items: center;
    justify-content: center;
  }

  .divider-line {
    width: 1px;
    height: 100%;
    background: linear-gradient(
      180deg,
      rgba(147, 168, 255, 0.2) 0%,
      rgba(51, 82, 255, 0.45) 50%,
      rgba(147, 168, 255, 0.2) 100%
    );
    box-shadow: 0 0 8px rgba(51, 82, 255, 0.35);
  }

  .stats-right {
    flex: 1;
    display: flex;
    align-items: center;
    justify-content: center;
    background: rgba(14, 18, 36, 0.2);
    overflow: hidden;
  }

  /* Month Indicator */
  .month-indicator-zone {
    display: flex;
    flex-direction: column;
    gap: 4px;
  }

  .zone-label {
    font-size: 11px;
    font-weight: 700;
    text-transform: uppercase;
    letter-spacing: 0.6px;
    color: var(--text-muted);
  }

  .month-selector-btn {
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    gap: 4px;
    padding: 12px 18px;
    background: rgba(21, 26, 53, 0.6);
    border: 1px solid var(--border-subtle);
    border-radius: 14px;
    color: var(--text-primary);
    cursor: pointer;
    transition: all 0.25s ease;
  }

  .month-selector-btn:hover {
    background: var(--accent-surface);
    border-color: var(--accent-primary);
    transform: translateY(-2px);
    box-shadow: 0 8px 24px rgba(51, 82, 255, 0.25);
  }

  .big-month-number {
    font-size: 38px;
    font-weight: 800;
    letter-spacing: -1px;
    color: #ffffff;
    line-height: 1.1;
  }

  .month-name-sub {
    display: flex;
    align-items: center;
    gap: 8px;
    font-size: 14px;
    font-weight: 600;
    color: var(--accent-highlight);
  }

  .month-selector-btn:hover :global(.chevron-icon) {
    transform: translateY(2px);
  }

  /* Rounded Print Button */
  .print-zone {
    display: flex;
  }

  .print-btn {
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 10px;
    padding: 13px 24px;
    border-radius: 30px;
    background: linear-gradient(135deg, hsl(234, 100%, 60%) 0%, hsl(234, 100%, 48%) 100%);
    color: #ffffff;
    font-size: 14px;
    font-weight: 700;
    border: 1px solid rgba(255, 255, 255, 0.2);
    box-shadow: 0 6px 20px rgba(51, 82, 255, 0.4);
    transition: all 0.2s ease;
  }

  .print-btn:hover:not(:disabled) {
    transform: translateY(-2px);
    box-shadow: 0 8px 26px rgba(51, 82, 255, 0.55);
  }

  .print-btn:disabled {
    opacity: 0.6;
    cursor: not-allowed;
  }

  /* Progress Bar */
  .progress-zone {
    display: flex;
    flex-direction: column;
    gap: 8px;
    padding: 18px 20px;
    background: rgba(21, 26, 53, 0.5);
    border: 1px solid var(--border-subtle);
    border-radius: 12px;
  }

  .progress-labels {
    display: flex;
    justify-content: space-between;
    align-items: baseline;
  }

  .p-title {
    font-size: 13px;
    font-weight: 600;
    color: var(--text-primary);
  }

  .p-val {
    font-size: 13px;
    font-weight: 700;
    color: var(--accent-highlight);
  }

  .progress-track {
    width: 100%;
    height: 12px;
    background: rgba(8, 10, 20, 0.7);
    border-radius: 6px;
    overflow: hidden;
    position: relative;
    border: 1px solid rgba(147, 168, 255, 0.1);
  }

  .progress-fill {
    height: 100%;
    background: linear-gradient(90deg, hsl(234, 54%, 35%) 0%, hsl(234, 100%, 60%) 100%);
    border-radius: 6px;
    transition: width 0.6s cubic-bezier(0.4, 0, 0.2, 1);
    position: relative;
  }

  .progress-glow {
    position: absolute;
    right: 0;
    top: 0;
    bottom: 0;
    width: 14px;
    background: #ffffff;
    opacity: 0.6;
    filter: blur(4px);
  }

  .progress-caption {
    display: flex;
    justify-content: space-between;
    font-size: 12px;
    color: var(--text-secondary);
  }

  .p-pct {
    font-weight: 700;
    color: var(--accent-highlight);
  }

  /* Stacked 2 Big Numbers */
  .metrics-stacked {
    display: flex;
    flex-direction: column;
    gap: 14px;
    margin-top: auto;
  }

  .big-metric-card {
    display: flex;
    align-items: center;
    gap: 20px;
    padding: 18px 22px;
    background: rgba(21, 26, 53, 0.6);
    border: 1px solid var(--border-subtle);
    border-radius: 14px;
    transition: all 0.2s ease;
  }

  .big-metric-card:hover {
    background: rgba(30, 38, 77, 0.75);
    border-color: rgba(147, 168, 255, 0.25);
    transform: translateY(-2px);
  }

  .metric-icon-box {
    display: flex;
    align-items: center;
    justify-content: center;
    width: 52px;
    height: 52px;
    border-radius: 12px;
    background: var(--accent-surface);
    color: var(--accent-highlight);
    border: 1px solid rgba(147, 168, 255, 0.2);
    box-shadow: 0 4px 14px rgba(51, 82, 255, 0.2);
  }

  .metric-details {
    display: flex;
    flex-direction: column;
    gap: 2px;
  }

  .metric-num {
    font-size: 32px;
    font-weight: 800;
    letter-spacing: -0.5px;
    color: #ffffff;
    line-height: 1.1;
  }

  .metric-label {
    font-size: 12px;
    color: var(--text-secondary);
    font-weight: 500;
  }
</style>
