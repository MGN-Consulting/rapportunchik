<script lang="ts">
  import { appState } from '../state.svelte';
  import { i18n } from '../i18n';
  import { Calendar, BarChart3, Settings, CheckCircle2, Loader2, AlertCircle } from 'lucide-svelte';
</script>

<header class="navbar window-drag-region">
  <div class="brand">
    <div class="logo-mark">
      <div class="logo-circle"></div>
      <span class="logo-text">Rapportunchik</span>
    </div>
    <span class="subtitle">{i18n.t.nav.brandSubtitle}</span>
  </div>

  <nav class="nav-tabs no-drag">
    <button
      class="tab-btn"
      class:active={appState.activeTab === 'main'}
      onclick={() => (appState.activeTab = 'main')}
    >
      <Calendar size={15} />
      <span>{i18n.t.nav.timesheet}</span>
    </button>

    <button
      class="tab-btn"
      class:active={appState.activeTab === 'stats'}
      onclick={() => (appState.activeTab = 'stats')}
    >
      <BarChart3 size={15} />
      <span>{i18n.t.nav.stats}</span>
    </button>

    <button
      class="tab-btn"
      class:active={appState.activeTab === 'profile'}
      onclick={() => (appState.activeTab = 'profile')}
    >
      <Settings size={15} />
      <span>{i18n.t.nav.profile}</span>
    </button>
  </nav>

  <div class="status-zone no-drag">
    {#if appState.saveStatus === 'saving'}
      <div class="status-badge saving">
        <Loader2 size={13} class="spin" />
        <span>{i18n.t.nav.statusSaving}</span>
      </div>
    {:else if appState.saveStatus === 'saved'}
      <div class="status-badge saved">
        <CheckCircle2 size={13} />
        <span>{i18n.t.nav.statusSaved}</span>
      </div>
    {:else if appState.saveStatus === 'error'}
      <div class="status-badge error">
        <AlertCircle size={13} />
        <span>{i18n.t.nav.statusError}</span>
      </div>
    {/if}
  </div>
</header>

<style>
  .navbar {
    display: flex;
    align-items: center;
    justify-content: space-between;
    height: 52px;
    padding: 0 18px 0 85px; /* offset for macOS window traffic lights */
    background: rgba(14, 18, 36, 0.75);
    backdrop-filter: blur(20px);
    -webkit-backdrop-filter: blur(20px);
    border-bottom: 1px solid var(--border-subtle);
    z-index: 100;
  }

  .brand {
    display: flex;
    align-items: center;
    gap: 10px;
  }

  .logo-mark {
    display: flex;
    align-items: center;
    gap: 8px;
  }

  .logo-circle {
    width: 10px;
    height: 10px;
    border-radius: 50%;
    background: var(--accent-primary);
    box-shadow: 0 0 10px var(--accent-primary);
  }

  .logo-text {
    font-weight: 700;
    font-size: 15px;
    letter-spacing: -0.3px;
    color: var(--text-primary);
  }

  .subtitle {
    font-size: 11px;
    color: var(--text-muted);
    padding: 2px 7px;
    background: rgba(147, 168, 255, 0.08);
    border-radius: 6px;
    font-weight: 500;
  }

  .nav-tabs {
    display: flex;
    gap: 4px;
    background: rgba(8, 10, 20, 0.6);
    padding: 3px;
    border-radius: 9px;
    border: 1px solid var(--border-subtle);
  }

  .tab-btn {
    display: flex;
    align-items: center;
    gap: 7px;
    padding: 6px 14px;
    border-radius: 7px;
    background: transparent;
    color: var(--text-secondary);
    font-size: 13px;
    font-weight: 500;
    border: 1px solid transparent;
  }

  .tab-btn:hover {
    color: var(--text-primary);
    background: rgba(147, 168, 255, 0.06);
  }

  .tab-btn.active {
    background: var(--accent-surface);
    color: var(--accent-highlight);
    border-color: rgba(147, 168, 255, 0.25);
    box-shadow: 0 2px 8px rgba(0, 0, 0, 0.25);
  }

  .status-zone {
    min-width: 130px;
    display: flex;
    justify-content: flex-end;
  }

  .status-badge {
    display: flex;
    align-items: center;
    gap: 6px;
    font-size: 11px;
    font-weight: 500;
    padding: 4px 10px;
    border-radius: 6px;
    animation: fadeIn 0.2s ease;
  }

  .status-badge.saving {
    background: rgba(245, 158, 11, 0.15);
    color: #fbbf24;
    border: 1px solid rgba(245, 158, 11, 0.3);
  }

  .status-badge.saved {
    background: rgba(16, 185, 129, 0.15);
    color: #34d399;
    border: 1px solid rgba(16, 185, 129, 0.3);
  }

  .status-badge.error {
    background: rgba(239, 68, 68, 0.15);
    color: #f87171;
    border: 1px solid rgba(239, 68, 68, 0.3);
  }

  :global(.spin) {
    animation: spin 1s linear infinite;
  }

  @keyframes spin {
    from { transform: rotate(0deg); }
    to { transform: rotate(360deg); }
  }

  @keyframes fadeIn {
    from { opacity: 0; transform: translateY(-2px); }
    to { opacity: 1; transform: translateY(0); }
  }
</style>
