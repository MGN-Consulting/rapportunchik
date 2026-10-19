<script lang="ts">
  import { appState } from '../state.svelte';
  import { i18n } from '../i18n';
  import { Building2, X, Check, Plus, Briefcase } from 'lucide-svelte';

  function handleBackdrop(e: MouseEvent) {
    if (e.target === e.currentTarget) {
      appState.isProfileModalOpen = false;
    }
  }

  async function handleSelectProfile(id: string) {
    appState.isProfileModalOpen = false;
    await appState.switchProfile(id);
  }

  function handleAddNewProfile() {
    appState.isProfileModalOpen = false;
    appState.startNewProfile();
  }
</script>

{#if appState.isProfileModalOpen}
  <div class="modal-backdrop" onclick={handleBackdrop} role="presentation">
    <div class="modal-card">
      <div class="modal-header">
        <div class="header-title">
          <Building2 size={19} class="accent-icon" />
          <h3>{i18n.t.nav.profileModalTitle}</h3>
        </div>
        <button
          class="close-btn"
          onclick={() => (appState.isProfileModalOpen = false)}
          title="Close"
        >
          <X size={16} />
        </button>
      </div>

      <p class="modal-desc">
        {i18n.t.nav.profileModalDesc}
      </p>

      <div class="profile-list">
        {#each appState.allProfiles as p}
          {@const isCurrent = p.id === appState.profile?.id}
          <button
            class="profile-card"
            class:active={isCurrent}
            onclick={() => handleSelectProfile(p.id)}
          >
            <div class="profile-card-left">
              <div class="profile-avatar">
                <Briefcase size={17} />
              </div>
              <div class="profile-info">
                <div class="company-name">{p.clientCompanyName || 'Committente'}</div>
                <div class="company-sub">
                  {#if p.clientPIVA}
                    <span class="sub-item">P.IVA: {p.clientPIVA}</span>
                  {:else if p.clientAddress}
                    <span class="sub-item">{p.clientAddress}</span>
                  {/if}
                  <span class="rate-badge">{i18n.t.nav.dailyRateBadge(p.dailyRate)}</span>
                </div>
              </div>
            </div>

            {#if isCurrent}
              <div class="active-badge-wrap">
                <div class="active-pill">{i18n.t.nav.activeProfileBadge}</div>
                <div class="active-check">
                  <Check size={14} strokeWidth={3} />
                </div>
              </div>
            {/if}
          </button>
        {/each}
      </div>

      <div class="modal-footer">
        <button class="add-profile-btn" onclick={handleAddNewProfile}>
          <Plus size={16} />
          <span>{i18n.t.nav.addNewProfileBtn}</span>
        </button>
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
    width: 92%;
    max-width: 520px;
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
    gap: 9px;
  }

  .header-title h3 {
    font-size: 17px;
    font-weight: 600;
    color: var(--text-primary);
  }

  :global(.accent-icon) {
    color: var(--accent-highlight);
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
    border: none;
    cursor: pointer;
    transition: all 0.15s ease;
  }

  .close-btn:hover {
    background: rgba(239, 68, 68, 0.2);
    color: #f87171;
  }

  .modal-desc {
    font-size: 13px;
    color: var(--text-secondary);
    margin-bottom: 18px;
    line-height: 1.45;
  }

  .profile-list {
    display: flex;
    flex-direction: column;
    gap: 10px;
    max-height: 320px;
    overflow-y: auto;
    padding-right: 4px;
    margin-bottom: 18px;
  }

  .profile-card {
    position: relative;
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 12px 14px;
    border-radius: 10px;
    background: rgba(21, 26, 53, 0.6);
    border: 1px solid var(--border-subtle);
    cursor: pointer;
    text-align: left;
    transition: all 0.2s cubic-bezier(0.4, 0, 0.2, 1);
  }

  .profile-card:hover {
    background: rgba(26, 33, 67, 0.85);
    border-color: rgba(147, 168, 255, 0.3);
    transform: translateY(-1px);
    box-shadow: 0 4px 16px rgba(0, 0, 0, 0.3);
  }

  .profile-card.active {
    background: rgba(51, 82, 255, 0.12);
    border-color: var(--accent-primary);
    box-shadow: 0 0 16px rgba(51, 82, 255, 0.25);
  }

  .profile-card-left {
    display: flex;
    align-items: center;
    gap: 12px;
    min-width: 0;
  }

  .profile-avatar {
    display: flex;
    align-items: center;
    justify-content: center;
    width: 36px;
    height: 36px;
    border-radius: 8px;
    background: rgba(51, 82, 255, 0.15);
    color: var(--accent-highlight);
    border: 1px solid rgba(147, 168, 255, 0.2);
    flex-shrink: 0;
  }

  .profile-card.active .profile-avatar {
    background: var(--accent-primary);
    color: #ffffff;
    box-shadow: 0 0 12px rgba(51, 82, 255, 0.5);
  }

  .profile-info {
    min-width: 0;
  }

  .company-name {
    font-size: 14px;
    font-weight: 600;
    color: var(--text-primary);
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
    margin-bottom: 3px;
  }

  .company-sub {
    display: flex;
    align-items: center;
    gap: 8px;
    font-size: 11px;
    color: var(--text-secondary);
  }

  .sub-item {
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
  }

  .rate-badge {
    background: rgba(147, 168, 255, 0.1);
    color: var(--text-primary);
    padding: 1px 6px;
    border-radius: 4px;
    font-weight: 500;
    font-size: 10px;
  }

  .active-badge-wrap {
    display: flex;
    align-items: center;
    gap: 8px;
    flex-shrink: 0;
  }

  .active-pill {
    font-size: 11px;
    font-weight: 600;
    color: var(--accent-highlight);
    background: rgba(51, 82, 255, 0.2);
    padding: 3px 8px;
    border-radius: 12px;
  }

  .active-check {
    display: flex;
    align-items: center;
    justify-content: center;
    width: 22px;
    height: 22px;
    border-radius: 50%;
    background: var(--accent-primary);
    color: #ffffff;
    box-shadow: 0 0 10px rgba(51, 82, 255, 0.6);
  }

  .modal-footer {
    display: flex;
    justify-content: center;
  }

  .add-profile-btn {
    width: 100%;
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 8px;
    padding: 11px 16px;
    border-radius: 10px;
    border: 1px dashed rgba(147, 168, 255, 0.35);
    background: rgba(21, 26, 53, 0.5);
    color: var(--text-primary);
    font-size: 13px;
    font-weight: 600;
    cursor: pointer;
    transition: all 0.2s cubic-bezier(0.4, 0, 0.2, 1);
  }

  .add-profile-btn:hover {
    background: rgba(51, 82, 255, 0.15);
    border-color: var(--accent-highlight);
    color: #ffffff;
    box-shadow: 0 0 16px rgba(51, 82, 255, 0.2);
  }

  @keyframes fadeIn {
    from {
      opacity: 0;
    }
    to {
      opacity: 1;
    }
  }

  @keyframes popIn {
    from {
      opacity: 0;
      transform: scale(0.95);
    }
    to {
      opacity: 1;
      transform: scale(1);
    }
  }
</style>
