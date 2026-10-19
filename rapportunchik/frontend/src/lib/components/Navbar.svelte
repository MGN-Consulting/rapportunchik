<script lang="ts">
  import { appState } from '../state.svelte';
  import { i18n } from '../i18n';
  import {
    Calendar,
    BarChart3,
    Settings,
    CheckCircle2,
    Loader2,
    AlertCircle,
    Building2,
    ChevronDown
  } from 'lucide-svelte';
</script>

<header class="navbar window-drag-region">
  <div class="profile-switch-zone no-drag">
    <div class="profile-select-wrapper">
      <Building2 size={15} class="client-icon" />
      <select
        class="profile-select"
        value={appState.profile?.id || 'default'}
        onchange={(e) => appState.switchProfile(e.currentTarget.value)}
        aria-label={i18n.t.nav.profileSelectorAria}
        title={appState.profile?.clientCompanyName || 'Committente'}
      >
        {#if appState.allProfiles && appState.allProfiles.length > 0}
          {#each appState.allProfiles as p}
            <option value={p.id}>
              {p.clientCompanyName || 'Committente'}
            </option>
          {/each}
        {:else}
          <option value="default">
            {appState.profile?.clientCompanyName || 'MGN Consulting'}
          </option>
        {/if}
      </select>
      <ChevronDown size={13} class="select-chevron" />
    </div>
  </div>

  <nav class="nav-tabs no-drag">
    <button
      class="tab-btn"
      class:active={appState.activeTab === 'main'}
      onclick={() => appState.requestTabChange('main')}
    >
      <Calendar size={15} />
      <span>{i18n.t.nav.timesheet}</span>
    </button>

    <button
      class="tab-btn"
      class:active={appState.activeTab === 'stats'}
      onclick={() => appState.requestTabChange('stats')}
    >
      <BarChart3 size={15} />
      <span>{i18n.t.nav.stats}</span>
    </button>

    <button
      class="tab-btn"
      class:active={appState.activeTab === 'profile'}
      onclick={() => appState.requestTabChange('profile')}
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
    position: relative;
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

  .profile-switch-zone {
    display: flex;
    align-items: center;
    z-index: 2;
  }

  .profile-select-wrapper {
    position: relative;
    display: flex;
    align-items: center;
    background: rgba(21, 26, 53, 0.7);
    border: 1px solid rgba(147, 168, 255, 0.15);
    border-radius: 8px;
    padding: 4px 10px 4px 8px;
    transition: all 0.2s cubic-bezier(0.4, 0, 0.2, 1);
    max-width: 260px;
  }

  .profile-select-wrapper:hover {
    background: rgba(21, 26, 53, 0.95);
    border-color: rgba(147, 168, 255, 0.35);
    box-shadow: 0 0 12px rgba(51, 82, 255, 0.2);
  }

  :global(.client-icon) {
    color: var(--accent-highlight);
    margin-right: 7px;
    flex-shrink: 0;
  }

  .profile-select {
    appearance: none;
    -webkit-appearance: none;
    background: transparent;
    border: none;
    color: var(--text-primary);
    font-family: inherit;
    font-size: 13px;
    font-weight: 600;
    padding-right: 18px;
    outline: none;
    cursor: pointer;
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
    max-width: 190px;
  }

  .profile-select option {
    background: #0e1224;
    color: var(--text-primary);
    padding: 8px;
  }

  :global(.select-chevron) {
    position: absolute;
    right: 8px;
    pointer-events: none;
    color: var(--text-muted);
  }

  .nav-tabs {
    position: absolute;
    left: 50%;
    transform: translateX(-50%);
    display: flex;
    gap: 4px;
    background: rgba(8, 10, 20, 0.6);
    padding: 3px;
    border-radius: 9px;
    border: 1px solid var(--border-subtle);
    z-index: 2;
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
