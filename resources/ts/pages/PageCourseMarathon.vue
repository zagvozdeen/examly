<template>
  <div
    class="flex flex-col gap-4"
  >
    <div class="flex flex-col gap-2 py-8">
      <!--      <span class="text-2xl text-center mb-6">Марафон по курсу</span>-->
      <span class="text-xl text-center">Все <n-skeleton
        v-if="total === -1"
        text
        width="40px"
        :sharp="false"
      /><span v-else>{{ total }}</span> вопросов из курса</span>
      <span class="text-sm text-center text-gray-400">Отличное испытание перед настоящим экзаменом!</span>
    </div>

    <n-button
      type="primary"
      class="sm:!w-52 sm:self-center"
      @click="onClick"
    >
      Начать марафон
    </n-button>

    <span class="text-xs text-center text-gray-400">Вы всегда можете остановиться и продолжить позже</span>
  </div>
</template>

<script lang="ts" setup>
import { PageExpose } from '@/types.ts'
import { useCourseStore } from '@/composables/useCourseStore.ts'
import { onMounted, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { NButton, NSkeleton, useMessage } from 'naive-ui'

const route = useRoute()
const router = useRouter()
const message = useMessage()
const courseStore = useCourseStore()

defineExpose<PageExpose>({
  title: 'Марафон по курсу',
  back: router.resolve({
    name: 'courses.show',
    params: { uuid: route.params.uuid },
  }),
})

const total = ref<number>(-1)

const onClick = () => {
  courseStore
    .createMarathon(route.params.uuid as string)
    .then(data => {
      message.success('Марафон начат!')
      
      router.push({
        name: 'tests.show',
        params: { uuid: data.data },
      })
    })
}

onMounted(() => {
  // courseStore
  //   .getCourseByUuid(route.params.uuid as string)
  //   .then(data => {
  //     course.value = data.data
  //     stats.value = data.stats
  //   })
})
</script>