<template>
  <n-config-provider :theme="darkTheme">
    <n-message-provider>
      <n-loading-bar-provider>
        <router-view />
      </n-loading-bar-provider>
    </n-message-provider>
  </n-config-provider>
</template>

<script lang="ts" setup>
import { NConfigProvider, NMessageProvider, NLoadingBarProvider, darkTheme } from 'naive-ui'
</script>