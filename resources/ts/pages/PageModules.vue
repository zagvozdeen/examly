<template>
  <div class="flex flex-col gap-4">
    <span class="text-gray-400 text-center text-xs">
      Модули тоже проходят проверку модерацией
      <br>
      перед тем как попасть в общий доступ.
    </span>

    <router-link
      :to="{ name: 'modules.create' }"
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

    <n-table
      v-if="modules.length > 0"
      size="small"
    >
      <thead>
        <tr>
          <th>Название</th>
          <th>Курс</th>
          <th>Статус</th>
          <th class="!text-right">
            Действия
          </th>
        </tr>
      </thead>
      <tbody>
        <tr
          v-for="module in modules"
          :key="module.id"
        >
          <td>{{ module.name }}</td>
          <td>{{ module.course?.name }}</td>
          <td>
            <span
              class="rounded-full text-xs px-2 py-1 font-medium"
              :class="{
                [StatusBackgroundColors[module.status]]: true,
                [StatusTextColors[module.status]]: true,
              }"
            >{{ StatusTranslates[module.status] }}</span>
          </td>
          <td class="text-right">
            <router-link :to="{ name: 'modules.edit', params: {uuid: module.uuid} }">
              <n-button
                type="warning"
                size="tiny"
              >
                Редактировать
              </n-button>
            </router-link>
          </td>
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
import { StatusTranslates, Module, PageExpose, StatusBackgroundColors, StatusTextColors } from '@/types.ts'
import { useModuleStore } from '@/composables/useModuleStore.ts'
import { isAdminMode, me } from '@/composables/useAuthStore.ts'

const router = useRouter()
const moduleStore = useModuleStore()

const modules = ref<Module[]>([])

defineExpose<PageExpose>({
  title: 'Добавленные модули',
  back: router.resolve({ name: 'me' }),
})

onMounted(() => {
  moduleStore
    .getModules({
      created_by: me.value?.id,
      all: isAdminMode.value,
    })
    .then(data => {
      if (data.data) {
        modules.value = data.data
      }
    })
})
</script>
