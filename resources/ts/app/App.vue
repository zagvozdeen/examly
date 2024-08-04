<template>
  <n-config-provider :theme="darkTheme">
    <router-view />
  </n-config-provider>
</template>

<script lang="ts" setup>
import { NConfigProvider, darkTheme } from 'naive-ui'
</script>