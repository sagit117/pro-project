<script setup lang="ts">
    import { fetchSideMenu } from "@/api/api.ts"

    console.log(await fetchSideMenu());
</script>

<template>
  <span>Меню приложения</span>
</template>