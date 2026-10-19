<script lang="ts">
  import { appState } from '../state.svelte';
  import { FileText, Target, Tag, Sparkles } from 'lucide-svelte';
  import type { WorkDayLog } from '../../../bindings/changeme';

  const quickObjectives = ['JIRA-', 'Formazione', 'Refactoring', 'Bugfixing', 'Code Review', 'Allineamento'];

  function handleInput() {
    if (!appState.currentDayLog) return;
    appState.debounceSaveDay(appState.currentDayLog);
  }

  function applyQuickObjective(tag: string) {
    if (!appState.currentDayLog) return;
    if (tag === 'JIRA-') {
      appState.currentDayLog.objective = tag;
    } else {
      appState.currentDayLog.objective = tag;
    }
    handleInput();
  }
</script>

<div class="panel activity-panel">
  <div class="panel-header">
    <div class="header-title">
      <FileText size={18} class="accent-icon" />
      <h3>Attività & Obiettivo della Giornata</h3>
    </div>
    <span class="badge-empty">Vuoto di default</span>
  </div>

  {#if appState.currentDayLog}
    <div class="panel-body">
      <!-- Objective input -->
      <div class="form-group">
        <label for="objective-input" class="group-label">
          <Target size={14} />
          <span>Obiettivo / Task di Lavoro</span>
        </label>
        <div class="input-wrapper">
          <input
            id="objective-input"
            type="text"
            placeholder="Es. PROJ-1042, Formazione, Refactoring modulo auth..."
            bind:value={appState.currentDayLog.objective}
            oninput={handleInput}
          />
        </div>

        <!-- Quick Tags -->
        <div class="quick-tags">
          <span class="tags-label"><Tag size={12} /> Suggeriti:</span>
          {#each quickObjectives as tag}
            <button
              type="button"
              class="tag-btn"
              onclick={() => applyQuickObjective(tag)}
            >
              {tag}
            </button>
          {/each}
        </div>
      </div>

      <!-- Description input -->
      <div class="form-group flex-1">
        <label for="desc-input" class="group-label">
          <FileText size={14} />
          <span>Descrizione Dettagliata Attività</span>
        </label>
        <div class="textarea-wrapper">
          <textarea
            id="desc-input"
            placeholder="Descrivi dettagliatamente cosa hai realizzato oggi per il Rapportino mensile..."
            bind:value={appState.currentDayLog.description}
            oninput={handleInput}
          ></textarea>
        </div>
      </div>

      <!-- Tips footer -->
      <div class="activity-tips">
        <Sparkles size={14} class="tips-icon" />
        <span>
          Le modifiche vengono salvate automaticamente in background nel database MongoDB locale.
        </span>
      </div>
    </div>
  {/if}
</div>

<style>
  .activity-panel {
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

  .badge-empty {
    font-size: 11px;
    color: var(--text-muted);
    background: rgba(147, 168, 255, 0.06);
    padding: 3px 8px;
    border-radius: 4px;
  }

  .panel-body {
    display: flex;
    flex-direction: column;
    gap: 20px;
    flex: 1;
  }

  .form-group {
    display: flex;
    flex-direction: column;
    gap: 8px;
  }

  .form-group.flex-1 {
    flex: 1;
    display: flex;
    flex-direction: column;
  }

  .group-label {
    display: flex;
    align-items: center;
    gap: 7px;
    font-size: 12px;
    font-weight: 600;
    color: var(--text-secondary);
    text-transform: uppercase;
    letter-spacing: 0.5px;
  }

  .input-wrapper input {
    width: 100%;
    background: rgba(8, 10, 20, 0.7);
    border: 1px solid var(--border-subtle);
    border-radius: 8px;
    padding: 11px 14px;
    color: var(--text-primary);
    font-family: inherit;
    font-size: 14px;
    outline: none;
    transition: all 0.2s ease;
  }

  .input-wrapper input:focus {
    border-color: var(--accent-primary);
    background: rgba(14, 18, 36, 0.9);
    box-shadow: 0 0 12px rgba(51, 82, 255, 0.25);
  }

  .quick-tags {
    display: flex;
    align-items: center;
    gap: 6px;
    flex-wrap: wrap;
    margin-top: 4px;
  }

  .tags-label {
    display: flex;
    align-items: center;
    gap: 4px;
    font-size: 11px;
    color: var(--text-muted);
  }

  .tag-btn {
    padding: 3px 9px;
    border-radius: 6px;
    background: rgba(147, 168, 255, 0.08);
    border: 1px solid rgba(147, 168, 255, 0.12);
    color: var(--text-secondary);
    font-size: 11px;
    font-weight: 500;
  }

  .tag-btn:hover {
    background: var(--accent-surface);
    color: var(--accent-highlight);
    border-color: var(--accent-primary);
  }

  .textarea-wrapper {
    flex: 1;
    display: flex;
  }

  .textarea-wrapper textarea {
    width: 100%;
    min-height: 180px;
    height: 100%;
    resize: none;
    background: rgba(8, 10, 20, 0.7);
    border: 1px solid var(--border-subtle);
    border-radius: 8px;
    padding: 12px 14px;
    color: var(--text-primary);
    font-family: inherit;
    font-size: 14px;
    line-height: 1.5;
    outline: none;
    transition: all 0.2s ease;
  }

  .textarea-wrapper textarea:focus {
    border-color: var(--accent-primary);
    background: rgba(14, 18, 36, 0.9);
    box-shadow: 0 0 12px rgba(51, 82, 255, 0.25);
  }

  .activity-tips {
    display: flex;
    align-items: center;
    gap: 8px;
    padding: 10px 14px;
    background: rgba(21, 26, 53, 0.4);
    border-radius: 6px;
    font-size: 12px;
    color: var(--text-muted);
  }

  :global(.tips-icon) {
    color: var(--accent-highlight);
    flex-shrink: 0;
  }
</style>
