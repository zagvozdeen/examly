<template>
  <label
    :for="`apple-checkbox-${$.uid}`"
    class="flex has-[:checked]:pl-[18px] has-[:checked]:bg-green-500 p-0.5 transition-all rounded-full w-10 h-6 bg-obscure-500 cursor-pointer"
  >
    <input
      :id="`apple-checkbox-${$.uid}`"
      v-model="modelValue"
      class="hidden"
      type="checkbox"
    >
    <span class="block w-5 h-5 rounded-full bg-white" />
  </label>
</template>

<script setup lang="ts">
const modelValue = defineModel('modelValue')
</script>