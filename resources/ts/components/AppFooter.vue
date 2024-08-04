<template>
  <footer class="bg-obscure-700 border-t border-obscure-500">
    <ul class="grid grid-cols-3 gap-2">
      <li>
        <router-link
          :to="{ name: 'main' }"
          class="text-center text-xs px-1 py-2 font-medium flex flex-col"
          :class="{
            'text-orange-400': route.name === 'main'
          }"
        >
          <i class="bi bi-house-door-fill text-sm" />
          <span>Главная</span>
        </router-link>
      </li>
      <li>
        <router-link
          :to="{ name: 'register' }"
          class="text-center text-xs px-1 py-2 font-medium flex flex-col"
          :class="{
            'text-orange-400': route.name === 'register'
          }"
        >
          <i class="bi bi-bar-chart-line-fill text-sm" />
          <span>Статистика</span>
        </router-link>
      </li>
      <li>
        <router-link
          :to="{ name: 'login' }"
          class="text-center text-xs px-1 py-2 font-medium flex flex-col"
          :class="{
            'text-orange-400': route.name === 'login'
          }"
        >
          <i class="bi bi-person-circle text-sm" />
          <span>Личный кабинет</span>
        </router-link>
      </li>
    </ul>
  </footer>
</template>

<script lang="ts" setup>
import { useRoute } from 'vue-router'

const route = useRoute()
</script>