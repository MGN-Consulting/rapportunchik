<script lang="ts">
  import { appState } from '../state.svelte';
  import { ChevronLeft, ChevronRight } from 'lucide-svelte';

  let containerEl: HTMLElement | null = null;

  const italianDaysShort = ['DOM', 'LUN', 'MAR', 'MER', 'GIO', 'VEN', 'SAB'];

  function getWeekdayShort(dateStr: string): string {
    const d = new Date(dateStr);
    return isNaN(d.getTime()) ? '' : italianDaysShort[d.getDay()];
  }

  function scrollToActive() {
    if (!containerEl) return;
    const activeEl = containerEl.querySelector('.tick-item.active') as HTMLElement;
    if (activeEl) {
      activeEl.scrollIntoView({ behavior: 'smooth', inline: 'center', block: 'nearest' });
    }
  }

  $effect(() => {
    // Whenever selectedDate changes, center the timeline
    if (appState.selectedDate) {
      setTimeout(scrollToActive, 60);
    }
  });

  function prevDay() {
    if (!appState.monthSummary?.days) return;
    const currentIdx = appState.monthSummary.days.findIndex(
      (d) => d.date === appState.selectedDate
    );
    if (currentIdx > 0) {
      appState.selectDate(appState.monthSummary.days[currentIdx - 1].date);
    }
  }

  function nextDay() {
    if (!appState.monthSummary?.days) return;
    const currentIdx = appState.monthSummary.days.findIndex(
      (d) => d.date === appState.selectedDate
    );
    if (currentIdx < appState.monthSummary.days.length - 1) {
      appState.selectDate(appState.monthSummary.days[currentIdx + 1].date);
    }
  }
</script>

<div class="timeline-wrapper">
  <button class="nav-arrow left" onclick={prevDay} title="Giorno precedente">
    <ChevronLeft size={18} />
  </button>

  <div class="timeline-container" bind:this={containerEl}>
    {#if appState.monthSummary?.days}
      {#each appState.monthSummary.days as day (day.date)}
        {@const isSelected = day.date === appState.selectedDate}
        <button
          class="tick-item"
          class:active={isSelected}
          class:worked={day.isWorked}
          class:non-working={!day.isWorkingDay}
          onclick={() => appState.selectDate(day.date)}
        >
          <!-- Cash Bag badge if worked -->
          {#if day.isWorked}
            <div class="cash-badge" title="Giorno Lavorato (+160 €)">
              <span class="cash-icon">💰</span>
            </div>
          {:else}
            <div class="tick-spacer"></div>
          {/if}

          <!-- Main tick body -->
          <div class="tick-body">
            <span class="day-name">{getWeekdayShort(day.date)}</span>
            <span class="day-number">{day.day}</span>
            <div class="tick-indicator"></div>
          </div>

          <!-- Workplace tag or non-working status -->
          <div class="tick-status">
            {#if !day.isWorkingDay}
              <span class="dot off" title="Giorno non lavorativo"></span>
            {:else if day.workLocation === 'In Sede'}
              <span class="dot office" title="In Sede"></span>
            {:else}
              <span class="dot remote" title="Smart-working"></span>
            {/if}
          </div>
        </button>
      {/each}
    {/if}
  </div>

  <button class="nav-arrow right" onclick={nextDay} title="Giorno successivo">
    <ChevronRight size={18} />
  </button>
</div>

<style>
  .timeline-wrapper {
    position: relative;
    display: flex;
    align-items: center;
    background: linear-gradient(180deg, rgba(14, 18, 36, 0.95) 0%, rgba(8, 10, 20, 0.95) 100%);
    border-bottom: 1px solid var(--border-subtle);
    padding: 10px 14px;
    user-select: none;
  }

  .nav-arrow {
    display: flex;
    align-items: center;
    justify-content: center;
    width: 32px;
    height: 64px;
    background: rgba(21, 26, 53, 0.8);
    border: 1px solid var(--border-subtle);
    border-radius: 8px;
    color: var(--text-secondary);
    z-index: 10;
    flex-shrink: 0;
  }

  .nav-arrow:hover {
    background: var(--accent-surface);
    color: var(--accent-highlight);
    border-color: rgba(147, 168, 255, 0.3);
  }

  .timeline-container {
    display: flex;
    align-items: flex-end;
    gap: 8px;
    overflow-x: auto;
    padding: 8px 12px;
    scroll-behavior: smooth;
    scrollbar-width: none;
    width: 100%;
  }

  .timeline-container::-webkit-scrollbar {
    display: none;
  }

  .tick-item {
    display: flex;
    flex-direction: column;
    align-items: center;
    background: transparent;
    border: none;
    padding: 4px 6px;
    border-radius: 10px;
    color: var(--text-muted);
    cursor: pointer;
    min-width: 52px;
    transition: all 0.22s cubic-bezier(0.4, 0, 0.2, 1);
    position: relative;
    flex-shrink: 0;
  }

  .tick-item:hover {
    color: var(--text-primary);
    transform: translateY(-2px);
  }

  .tick-item.non-working {
    opacity: 0.55;
  }

  .cash-badge {
    height: 22px;
    display: flex;
    align-items: center;
    justify-content: center;
    animation: bounce 0.3s ease;
  }

  .cash-icon {
    font-size: 17px;
    filter: drop-shadow(0 0 6px rgba(255, 215, 0, 0.5));
  }

  .tick-spacer {
    height: 22px;
  }

  .tick-body {
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 2px;
    padding: 7px 10px;
    border-radius: 8px;
    background: rgba(21, 26, 53, 0.5);
    border: 1px solid rgba(147, 168, 255, 0.08);
    transition: all 0.2s ease;
    width: 100%;
  }

  .tick-item:hover .tick-body {
    background: rgba(30, 38, 77, 0.7);
    border-color: rgba(147, 168, 255, 0.2);
  }

  .day-name {
    font-size: 10px;
    font-weight: 600;
    letter-spacing: 0.5px;
    text-transform: uppercase;
  }

  .day-number {
    font-size: 15px;
    font-weight: 700;
    color: var(--text-secondary);
  }

  .tick-indicator {
    width: 14px;
    height: 3px;
    border-radius: 2px;
    background: rgba(147, 168, 255, 0.2);
    margin-top: 2px;
    transition: all 0.2s ease;
  }

  /* Active Main Tick Styling */
  .tick-item.active {
    transform: translateY(-4px) scale(1.08);
  }

  .tick-item.active .tick-body {
    background: var(--accent-surface);
    border-color: var(--accent-primary);
    box-shadow: 0 4px 16px rgba(51, 82, 255, 0.35);
  }

  .tick-item.active .day-name {
    color: var(--accent-highlight);
  }

  .tick-item.active .day-number {
    font-size: 18px;
    color: #ffffff;
    text-shadow: 0 0 10px var(--accent-primary);
  }

  .tick-item.active .tick-indicator {
    width: 24px;
    height: 3.5px;
    background: var(--accent-highlight);
    box-shadow: 0 0 8px var(--accent-highlight);
  }

  .tick-status {
    margin-top: 5px;
    display: flex;
    justify-content: center;
  }

  .dot {
    width: 6px;
    height: 6px;
    border-radius: 50%;
  }

  .dot.office {
    background: #38bdf8;
    box-shadow: 0 0 6px #38bdf8;
  }

  .dot.remote {
    background: #a855f7;
    box-shadow: 0 0 6px #a855f7;
  }

  .dot.off {
    background: #64748b;
  }

  @keyframes bounce {
    0% { transform: scale(0.6); }
    50% { transform: scale(1.2); }
    100% { transform: scale(1); }
  }
</style>
