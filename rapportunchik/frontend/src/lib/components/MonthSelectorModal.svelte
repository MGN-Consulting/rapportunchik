<script lang="ts">
  import { appState } from '../state.svelte';
  import { i18n } from '../i18n';
  import { Calendar, X, Check } from 'lucide-svelte';

  function handleBackdrop(e: MouseEvent) {
    if (e.target === e.currentTarget) {
      appState.isMonthModalOpen = false;
    }
  }
</script>

{#if appState.isMonthModalOpen}
  <div class="modal-backdrop" onclick={handleBackdrop} role="presentation">
    <div class="modal-card">
      <div class="modal-header">
        <div class="header-title">
          <Calendar size={18} class="accent-icon" />
          <h3>{i18n.t.stats.monthSelectorTitle}</h3>
        </div>
        <button
          class="close-btn"
          onclick={() => (appState.isMonthModalOpen = false)}
          title="Close"
        >
          <X size={16} />
        </button>
      </div>

      <p class="modal-desc">
        {#if i18n.currentLocale === 'en'}
          Choose a month to review history, verify attendance, or generate the PDF timesheet:
        {:else if i18n.currentLocale === 'ru'}
          Выберите месяц для просмотра истории, проверки табеля или экспорта PDF-отчета:
        {:else}
          Scegli il mese per consultare lo storico, verificare le presenze o generare il relativo Rapportino PDF:
        {/if}
      </p>

      <div class="month-grid">
        {#each appState.availableMonths as opt}
          {@const isCurrent =
            opt.year === appState.selectedYear && opt.month === appState.selectedMonth}
          <button
            class="month-card"
            class:active={isCurrent}
            onclick={() => appState.selectMonth(opt.year, opt.month)}
          >
            <div class="month-badge">{i18n.t.stats.monthIndexLabel(opt.monthIndex)}</div>
            <div class="month-name">{i18n.getMonthName(opt.month)} {opt.year}</div>
            {#if isCurrent}
              <div class="active-check">
                <Check size={14} strokeWidth={3} />
              </div>
            {/if}
          </button>
        {/each}
      </div>
    </div>
  </div>
{/if}

<style>
  .modal-backdrop {
    position: fixed;
    inset: 0;
    background: rgba(0, 0, 0, 0.7);
    backdrop-filter: blur(8px);
    -webkit-backdrop-filter: blur(8px);
    display: flex;
    align-items: center;
    justify-content: center;
    z-index: 1000;
    animation: fadeIn 0.18s ease;
  }

  .modal-card {
    width: 90%;
    max-width: 480px;
    background: #0e1224;
    border: 1px solid var(--border-accent);
    box-shadow: 0 16px 48px rgba(0, 0, 0, 0.6), 0 0 24px rgba(51, 82, 255, 0.2);
    border-radius: 14px;
    padding: 24px;
    animation: popIn 0.2s cubic-bezier(0.34, 1.56, 0.64, 1);
  }

  .modal-header {
    display: flex;
    align-items: center;
    justify-content: space-between;
    margin-bottom: 8px;
  }

  .header-title {
    display: flex;
    align-items: center;
    gap: 8px;
  }

  .header-title h3 {
    font-size: 16px;
    font-weight: 600;
    color: var(--text-primary);
  }

  .close-btn {
    display: flex;
    align-items: center;
    justify-content: center;
    width: 28px;
    height: 28px;
    border-radius: 6px;
    background: rgba(147, 168, 255, 0.08);
    color: var(--text-secondary);
  }

  .close-btn:hover {
    background: rgba(239, 68, 68, 0.2);
    color: #f87171;
  }

  .modal-desc {
    font-size: 13px;
    color: var(--text-secondary);
    margin-bottom: 18px;
    line-height: 1.4;
  }

  .month-grid {
    display: grid;
    grid-template-columns: repeat(2, 1fr);
    gap: 10px;
    max-height: 320px;
    overflow-y: auto;
    padding-right: 4px;
  }

  .month-card {
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    gap: 4px;
    padding: 12px 14px;
    background: rgba(21, 26, 53, 0.7);
    border: 1px solid var(--border-subtle);
    border-radius: 10px;
    color: var(--text-primary);
    position: relative;
    text-align: left;
    transition: all 0.2s ease;
  }

  .month-card:hover {
    background: rgba(30, 38, 77, 0.9);
    border-color: var(--accent-primary);
    transform: translateY(-2px);
  }

  .month-card.active {
    background: var(--accent-surface);
    border-color: var(--accent-primary);
    box-shadow: 0 0 16px rgba(51, 82, 255, 0.3);
  }

  .month-badge {
    font-size: 11px;
    font-weight: 700;
    color: var(--accent-highlight);
    text-transform: uppercase;
  }

  .month-name {
    font-size: 14px;
    font-weight: 600;
  }

  .active-check {
    position: absolute;
    top: 10px;
    right: 10px;
    width: 20px;
    height: 20px;
    border-radius: 50%;
    background: var(--accent-primary);
    color: #ffffff;
    display: flex;
    align-items: center;
    justify-content: center;
  }

  @keyframes fadeIn {
    from { opacity: 0; }
    to { opacity: 1; }
  }

  @keyframes popIn {
    from { opacity: 0; transform: scale(0.95); }
    to { opacity: 1; transform: scale(1); }
  }
</style>
