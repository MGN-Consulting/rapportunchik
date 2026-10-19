<script lang="ts">
  import Timeline from '../components/Timeline.svelte';
  import SchedulePanel from '../components/SchedulePanel.svelte';
  import ActivityPanel from '../components/ActivityPanel.svelte';
  import { appState } from '../state.svelte';
</script>

<div class="main-view">
  <!-- Top Full-Width Timeline -->
  <Timeline />

  <!-- Dual Vertically Separated Sections -->
  <div class="split-sections">
    <div class="section-left">
      <SchedulePanel />
    </div>

    <!-- Vertical separator line with subtle accent glow -->
    <div class="vertical-divider" role="separator" aria-orientation="vertical">
      <div class="divider-line"></div>
    </div>

    <div class="section-right">
      <ActivityPanel />
    </div>
  </div>
</div>

<style>
  .main-view {
    display: flex;
    flex-direction: column;
    height: 100%;
    overflow: hidden;
  }

  .split-sections {
    display: flex;
    flex: 1;
    overflow: hidden;
    position: relative;
  }

  .section-left {
    flex: 1;
    min-width: 380px;
    max-width: 50%;
    height: 100%;
    overflow-y: auto;
  }

  .vertical-divider {
    width: 1px;
    position: relative;
    background: var(--border-subtle);
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
      rgba(51, 82, 255, 0.4) 50%,
      rgba(147, 168, 255, 0.2) 100%
    );
    box-shadow: 0 0 6px rgba(51, 82, 255, 0.3);
  }

  .section-right {
    flex: 1;
    height: 100%;
    overflow-y: auto;
  }
</style>
