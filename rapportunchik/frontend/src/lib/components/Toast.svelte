<script lang="ts">
  import { appState } from '../state.svelte';
  import { CheckCircle2, AlertCircle, Info, X } from 'lucide-svelte';
</script>

{#if appState.toast}
  <div class="toast-wrapper">
    <div class="toast-card" class:success={appState.toast.type === 'success'} class:error={appState.toast.type === 'error'}>
      <div class="toast-icon">
        {#if appState.toast.type === 'success'}
          <CheckCircle2 size={18} />
        {:else if appState.toast.type === 'error'}
          <AlertCircle size={18} />
        {:else}
          <Info size={18} />
        {/if}
      </div>

      <div class="toast-message">
        {appState.toast.message}
      </div>

      <button class="toast-close" onclick={() => (appState.toast = null)}>
        <X size={14} />
      </button>
    </div>
  </div>
{/if}

<style>
  .toast-wrapper {
    position: fixed;
    bottom: 24px;
    right: 24px;
    z-index: 2000;
    pointer-events: none;
    animation: toastSlide 0.25s cubic-bezier(0.16, 1, 0.3, 1);
  }

  .toast-card {
    pointer-events: auto;
    display: flex;
    align-items: center;
    gap: 12px;
    padding: 12px 18px;
    border-radius: 10px;
    background: #0e1224;
    border: 1px solid var(--border-subtle);
    box-shadow: 0 10px 30px rgba(0, 0, 0, 0.5), 0 0 15px rgba(51, 82, 255, 0.2);
    color: var(--text-primary);
    max-width: 420px;
  }

  .toast-card.success {
    border-color: rgba(16, 185, 129, 0.4);
    box-shadow: 0 10px 30px rgba(0, 0, 0, 0.5), 0 0 15px rgba(16, 185, 129, 0.25);
  }

  .toast-card.success .toast-icon {
    color: #10b981;
  }

  .toast-card.error {
    border-color: rgba(239, 68, 68, 0.4);
    box-shadow: 0 10px 30px rgba(0, 0, 0, 0.5), 0 0 15px rgba(239, 68, 68, 0.25);
  }

  .toast-card.error .toast-icon {
    color: #ef4444;
  }

  .toast-icon {
    color: var(--accent-highlight);
    flex-shrink: 0;
  }

  .toast-message {
    font-size: 13px;
    font-weight: 500;
    line-height: 1.4;
  }

  .toast-close {
    background: transparent;
    border: none;
    color: var(--text-muted);
    cursor: pointer;
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 2px;
    margin-left: 6px;
  }

  .toast-close:hover {
    color: var(--text-primary);
  }

  @keyframes toastSlide {
    from {
      opacity: 0;
      transform: translateY(12px) scale(0.96);
    }
    to {
      opacity: 1;
      transform: translateY(0) scale(1);
    }
  }
</style>
