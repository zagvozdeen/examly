<template>
  <div
    v-if="me"
    class="flex flex-col gap-4"
  >
    <div class="flex flex-col items-center gap-1.5 rounded-md bg-obscure-500 bg-opacity-50 p-3">
      <div class="bg-blue-400 rounded w-9 py-0.5 text-center">
        <i class="bi bi-info-lg text-2xl" />
      </div>
      <span class="text-lg">Список рефералов и работодателей</span>
      <span class="text-xs text-gray-100 text-center">
        Здесь будет список рефералов и работодателей.
        Чтобы открыть раздел, вам необходимо потратить 50 баллов.
        200 баллов можно получить за прохождение опроса о пользе приложения.
      </span>
    </div>

    <n-button
      v-if="!me.can_view_referrals"
      type="primary"
      @click="unlock"
    >
      Разблокировать раздел за 50 баллов
    </n-button>

    <template v-else>
      <n-table
        v-if="referrals.length > 0"
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
            v-for="module in referrals"
            :key="module.id"
          >
            <td>{{ module.full_name }}</td>
            <td>{{ module.contact || 'Не указан' }}</td>
            <td class="text-right">
              —
            </td>
          </tr>
        </tbody>
      </n-table>
      <span
        v-else
        class="text-center text-gray-400"
      >Пользователи не найдены</span>
    </template>
  </div>
</template>

<script lang="ts" setup>
import { NButton,useMessage, NTable } from 'naive-ui'
import { onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'
import { PageExpose, User } from '@/types.ts'
import { useUserStore } from '@/composables/useUserStore.ts'
import { me, useAuthStore } from '@/composables/useAuthStore.ts'

const router = useRouter()
const message = useMessage()
const userStore = useUserStore()
const authStore = useAuthStore()
const referrals = ref<User[]>([])

defineExpose<PageExpose>({
  title: 'Реферальная система и работодатели',
  back: router.resolve({ name: 'me' }),
})

const getReferrals = async () => {
  return userStore
    .getReferrals()
    .then(data => {
      referrals.value = data.data || []
    })
}

const unlock = () => {
  if (!me.value) return

  if (me.value.account < 50) {
    message.error(`У вас недостаточно баллов для разблокировки раздела: у вас ${me.value.account}, необходимо 50`)
    return
  }

  userStore
    .unlockReferrals()
    .then(async () => {
      await authStore.getMe()
      await getReferrals()
      
      message.success('Раздел разблокирован')
    })
}

onMounted(() => {
  if (me.value?.can_view_referrals) {
    getReferrals()
  }
})
</script>
