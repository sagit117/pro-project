<script setup lang="ts">
  import { fetchAppMenu, ErrorFetch } from "@/api/api.ts"
  import type { ISideMenuItem } from "@/api/api_interfaces"
  import { ref } from "vue";

  const menu = ref<ISideMenuItem[]>([])
  
  fetchAppMenu()
    .then(response => menu.value = response)
    .catch(error => {
        if (error instanceof ErrorFetch) {
          console.error("Данные по составу меню приложения не получены, ошибка: " + error.message)
        } else {
          console.error("Не предвиденная ошибка: " + error)
        }
      }
    )
</script>

<template>
  <span>Меню приложения</span>
  <ul>
    <li v-for="items in menu">
      {{ items.name }}
    </li>
  </ul>
</template>