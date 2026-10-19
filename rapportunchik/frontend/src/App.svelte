<script lang="ts">
  import { onMount } from 'svelte';
  import { appState } from './lib/state.svelte';
  import Navbar from './lib/components/Navbar.svelte';
  import MainView from './lib/views/MainView.svelte';
  import StatsView from './lib/views/StatsView.svelte';
  import ProfileView from './lib/views/ProfileView.svelte';
  import Toast from './lib/components/Toast.svelte';
  import UnsavedChangesModal from './lib/components/UnsavedChangesModal.svelte';
  import ProfileSelectorModal from './lib/components/ProfileSelectorModal.svelte';

  onMount(() => {
    appState.init();
  });
</script>

<div class="app-layout">
  <Navbar />

  <main class="app-content">
    {#if appState.activeTab === 'main'}
      <MainView />
    {:else if appState.activeTab === 'stats'}
      <StatsView />
    {:else if appState.activeTab === 'profile'}
      <ProfileView />
    {/if}
  </main>

  <Toast />
  <UnsavedChangesModal />
  <ProfileSelectorModal />
</div>

<style>
  .app-layout {
    width: 100vw;
    height: 100vh;
    display: flex;
    flex-direction: column;
    overflow: hidden;
    background: var(--bg-app);
  }

  .app-content {
    flex: 1;
    overflow: hidden;
    position: relative;
  }
</style>
