<template>
  <header>
    <div class="flex justify-between items-center">
      <router-link :to="{ name: 'main' }">
        <AppLogo />
      </router-link>

      <ul class="flex gap-2">
        <li>
          <router-link
            :to="{ name: 'main' }"
            class="uppercase text-xs font-medium"
            :class="{
              'text-orange-400': route.name === 'main'
            }"
          >
            Главная
          </router-link>
        </li>
        <li>
          <router-link
            :to="{ name: 'stats' }"
            class="uppercase text-xs font-medium"
            :class="{
              'text-orange-400': route.name === 'stats'
            }"
          >
            Статистика
          </router-link>
        </li>
        <li>
          <router-link
            :to="{ name: 'me' }"
            class="uppercase text-xs font-medium"
            :class="{
              'text-orange-400': route.name === 'me'
            }"
          >
            Личный кабинет
          </router-link>
        </li>
      </ul>
    </div>
  </header>
</template>

<script lang="ts" setup>
import AppLogo from '@/components/AppLogo.vue'
import { useRoute } from 'vue-router'

const route = useRoute()
</script>
