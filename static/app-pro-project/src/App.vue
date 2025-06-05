<script setup lang="ts">
  import { RouterView } from 'vue-router'
  import TopMenu from "@/components/TopMenu.vue";
  import SideMenu from "@/components/SideMenu.vue"
</script>

<template>
  <header>
    <div class="top_menu_container shadow">
      <TopMenu />
    </div>
  </header>
  <div class="workspace">
    <aside>
      <div class="side_bar_container">
        <SideMenu />
      </div>
    </aside>

    <div class="context">
      <RouterView />
    </div>
  </div>
  
</template>

<style scoped>
  .workspace {
    display: flex;
  }
  .side_bar_container {
    padding: 1rem;
    width: 20rem;
    display: flex;
    flex-direction: column;
    min-height: 100vh;
    box-shadow: 0px 14px 6px 0px var(--vt-c-text-dark-2);
  }
  .context {
    min-width: 100vw;
    min-height: 100vh;
    display: flex;
    flex-direction: column;
    padding: 1rem;
  }
  .top_menu_container {
    min-width: 100vw;
    display: flex;
    padding: 1rem .5rem;
    box-shadow: 0px 0px 10px var(--vt-c-text-dark-2);
  }
  nav a.router-link-exact-active {
    color: var(--color-text);
  }
  nav a.router-link-exact-active:hover {
    background-color: transparent;
  }
  nav a {
    display: inline-block;
    padding: 0 1rem;
    border-left: 1px solid var(--color-border);
  }
  nav a:first-of-type {
    border: 0;
  }
</style>
