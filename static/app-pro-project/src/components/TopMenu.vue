<script setup lang="ts">
  import { goRoute } from '@/common';
</script>

<template>
  <div class="top_menu">
    <div class="link">
      <RouterLink :to="{ name: 'home' }">PRO-PROJECT</RouterLink>
    </div>
    <div class="link">
      <button class="btn btn_green" @click="goRoute({ name: 'auth' })">Войти в систему</button>
    </div>
  </div>
  
</template>

<style scoped>
  .top_menu {
    display: flex;
    justify-content: space-between;
    width: 100%;
  }
  .link {
    display: flex;
  }
  button {
    cursor: pointer;
  }
</style>