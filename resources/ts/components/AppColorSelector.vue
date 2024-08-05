<template>
  <ul class="flex gap-3 px-1 flex-wrap">
    <li
      v-for="color in colors"
      :key="color"
    >
      <input
        :id="`app-color-${color}-selector`"
        v-model="value"
        :value="color"
        type="radio"
        name="app-color-selector"
        class="hidden peer"
      >
      <label
        :for="`app-color-${color}-selector`"
        class="block peer-checked:outline rounded w-8 h-8 cursor-pointer hover:outline outline-offset-2 outline-2 outline-orange-400"
        :class="{
          [color]: true,
        }"
      />
    </li>
  </ul>
</template>

<script lang="ts" setup>
const value = defineModel('value')

const colors = [
  'bg-slate-400',
  'bg-gray-400',
  'bg-zinc-400',
  'bg-neutral-400',
  'bg-stone-400',
  'bg-red-400',
  'bg-orange-400',
  'bg-amber-400',
  'bg-yellow-400',
  'bg-lime-400',
  'bg-green-400',
  'bg-emerald-400',
  'bg-teal-400',
  'bg-cyan-400',
  'bg-sky-400',
  'bg-blue-400',
  'bg-indigo-400',
  'bg-violet-400',
  'bg-purple-400',
  'bg-fuchsia-400',
  'bg-pink-400',
  'bg-rose-400',
]
</script>
