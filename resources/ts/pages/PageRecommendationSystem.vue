<template>
  <div class="flex flex-col gap-4">
    <div
      class="flex flex-col items-center gap-1.5 rounded-md bg-obscure-500 bg-opacity-50 p-3"
    >
      <div class="bg-blue-400 rounded w-9 py-0.5 text-center">
        <i class="bi bi-info-lg text-2xl" />
      </div>
      <span class="text-lg">Рекомендованные кандидаты</span>
      <span class="text-xs text-gray-100 text-center">
        Если вы компания, то вам доступен список рекомендованных кандидатов.
        Если пользователь указал свой контакт для связи, то вы можете связаться с ним.
        Если пользователь не указал своего контакта, то вы можете попробовать найти его самостоятельно.
        Составляйте тесты, и пользователи, которые прошли тест лучше всего, будут рекомендованы вам.
      </span>
    </div>

    <n-table
      v-if="users.length > 0"
      size="small"
    >
      <thead>
        <tr>
          <th>Имя и фамилия</th>
          <th>Контакт</th>
          <th class="!text-right">
            Действия
          </th>
        </tr>
      </thead>
      <tbody>
        <tr
          v-for="module in users"
          :key="module.id"
        >
          <td>{{ module.full_name }}</td>
          <td>{{ module.contact || 'Не указан' }}</td>
          <td class="text-right">
            —
            <!--          <router-link :to="{ name: 'users.edit', params: {uuid: module.uuid} }">-->
            <!--            <n-button-->
            <!--              type="warning"-->
            <!--              size="tiny"-->
            <!--            >-->
            <!--              Редактировать-->
            <!--            </n-button>-->
            <!--          </router-link>-->
          </td>
        </tr>
      </tbody>
    </n-table>
    <span
      v-else
      class="text-center text-gray-400"
    >Пользователи не найдены</span>
  </div>
</template>

<script lang="ts" setup>
import { NTable } from 'naive-ui'
import { onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'
import { PageExpose, User } from '@/types.ts'
import { useUserStore } from '@/composables/useUserStore.ts'

const router = useRouter()
const userStore = useUserStore()
const users = ref<User[]>([])

defineExpose<PageExpose>({
  title: 'Система рекомендаций кандидатов',
  back: router.resolve({ name: 'me' }),
})

onMounted(() => {
  userStore
    .getUsers()
    .then(data => {
      users.value = data.data || []
    })
})
</script>
