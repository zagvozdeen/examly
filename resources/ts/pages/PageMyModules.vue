<template>
  <div class="flex flex-col gap-4">
    <span class="text-gray-400 text-center text-xs">
      Модули тоже проходят проверку модерацией
      <br>
      перед тем как попасть в общий доступ.
    </span>

    <router-link
      :to="{ name: 'my.modules.create' }"
      class="sm:self-center"
    >
      <n-button
        type="info"
        class="sm:!px-10 !w-full"
      >
        <div class="flex items-center gap-2">
          <i class="bi bi-plus-square-fill" />
          <span>Создать модуль</span>
        </div>
      </n-button>
    </router-link>

    <n-table v-if="modules.length > 0">
      <thead>
        <tr>
          <th>Название</th>
          <th>Статус</th>
        </tr>
      </thead>
      <tbody>
        <tr
          v-for="module in modules"
          :key="module.id"
        >
          <td>{{ module.name }}</td>
          <td>{{ module.status }}</td>
        </tr>
      </tbody>
    </n-table>
    <span
      v-else
      class="text-center text-gray-400"
    >Пока ничего</span>
  </div>
</template>

<script lang="ts" setup>
import { NTable, NButton } from 'naive-ui'
import { onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'
import { Module, PageExpose } from '@/types.ts'
import { useModuleStore } from '@/composables/useModuleStore.ts'

const router = useRouter()
const moduleStore = useModuleStore()

const modules = ref<Module[]>([])

defineExpose<PageExpose>({
  title: 'Добавленные модули',
  back: router.resolve({ name: 'me' }),
})

onMounted(() => {
  moduleStore
    .getMyModules()
    .then(data => {
      modules.value = data.data
    })
})
</script>
