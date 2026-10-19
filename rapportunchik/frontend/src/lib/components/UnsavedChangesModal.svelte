<script lang="ts">
  import { appState } from '../state.svelte';
  import { i18n } from '../i18n';
  import { AlertTriangle, X } from 'lucide-svelte';

  function handleBackdrop(e: MouseEvent) {
    if (e.target === e.currentTarget) {
      appState.cancelDiscard();
    }
  }
</script>

{#if appState.isUnsavedModalOpen}
  <div class="modal-backdrop" onclick={handleBackdrop} role="presentation">
    <div class="modal-card">
      <div class="modal-header">
        <div class="header-title">
          <div class="icon-circle">
            <AlertTriangle size={20} class="warning-icon" />
          </div>
          <h3>{i18n.t.unsavedModal.title}</h3>
        </div>
        <button
          class="close-btn"
          onclick={() => appState.cancelDiscard()}
          title="Close"
        >
          <X size={16} />
        </button>
      </div>

      <p class="modal-desc">
        {i18n.t.unsavedModal.desc}
      </p>

      <div class="modal-actions">
        <button
          type="button"
          class="btn btn-secondary"
          onclick={() => appState.cancelDiscard()}
        >
          {i18n.t.unsavedModal.stayBtn}
        </button>

        <button
          type="button"
          class="btn btn-danger"
          onclick={() => appState.confirmDiscardAndNavigate()}
        >
          {i18n.t.unsavedModal.discardBtn}
        </button>
      </div>
    </div>
  </div>
{/if}

<style>
  .modal-backdrop {
    position: fixed;
    inset: 0;
    background: rgba(0, 0, 0, 0.72);
    backdrop-filter: blur(8px);
    -webkit-backdrop-filter: blur(8px);
    display: flex;
    align-items: center;
    justify-content: center;
    z-index: 1100;
    animation: fadeIn 0.18s ease;
  }

  .modal-card {
    width: 90%;
    max-width: 440px;
    background: #0e1224;
    border: 1px solid rgba(245, 158, 11, 0.35);
    box-shadow: 0 16px 48px rgba(0, 0, 0, 0.7), 0 0 24px rgba(245, 158, 11, 0.15);
    border-radius: 14px;
    padding: 24px;
    animation: popIn 0.2s cubic-bezier(0.34, 1.56, 0.64, 1);
  }

  .modal-header {
    display: flex;
    align-items: center;
    justify-content: space-between;
    margin-bottom: 12px;
  }

  .header-title {
    display: flex;
    align-items: center;
    gap: 12px;
  }

  .icon-circle {
    width: 36px;
    height: 36px;
    border-radius: 50%;
    background: rgba(245, 158, 11, 0.12);
    display: flex;
    align-items: center;
    justify-content: center;
    flex-shrink: 0;
    border: 1px solid rgba(245, 158, 11, 0.25);
  }

  :global(.warning-icon) {
    color: #f59e0b;
  }

  .header-title h3 {
    font-size: 16px;
    font-weight: 700;
    color: var(--text-primary);
  }

  .close-btn {
    background: transparent;
    border: none;
    color: var(--text-muted);
    cursor: pointer;
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 4px;
    border-radius: 6px;
    transition: all 0.18s ease;
  }

  .close-btn:hover {
    color: var(--text-primary);
    background: rgba(147, 168, 255, 0.1);
  }

  .modal-desc {
    font-size: 13px;
    line-height: 1.55;
    color: var(--text-secondary);
    margin-bottom: 22px;
  }

  .modal-actions {
    display: flex;
    align-items: center;
    justify-content: flex-end;
    gap: 10px;
  }

  .btn {
    padding: 9px 16px;
    border-radius: 8px;
    font-size: 13px;
    font-weight: 600;
    cursor: pointer;
    transition: all 0.18s cubic-bezier(0.4, 0, 0.2, 1);
  }

  .btn-secondary {
    background: rgba(21, 26, 53, 0.8);
    border: 1px solid var(--border-subtle);
    color: var(--text-secondary);
  }

  .btn-secondary:hover {
    background: rgba(147, 168, 255, 0.1);
    color: var(--text-primary);
    border-color: rgba(147, 168, 255, 0.3);
  }

  .btn-danger {
    background: rgba(239, 68, 68, 0.18);
    border: 1px solid rgba(239, 68, 68, 0.4);
    color: #fca5a5;
  }

  .btn-danger:hover {
    background: #ef4444;
    border-color: #ef4444;
    color: #ffffff;
    box-shadow: 0 4px 16px rgba(239, 68, 68, 0.4);
    transform: translateY(-1px);
  }

  @keyframes fadeIn {
    from { opacity: 0; }
    to { opacity: 1; }
  }

  @keyframes popIn {
    from { opacity: 0; transform: scale(0.96) translateY(6px); }
    to { opacity: 1; transform: scale(1) translateY(0); }
  }
</style>
