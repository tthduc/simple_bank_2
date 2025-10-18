// This is included 3 sessions in the file
// 1. Script
// 2. Template
// 3. Style
// The scoped keyword is used to apply the style only to the current component

<script setup lang="ts">
import {RouterView} from 'vue-router'
</script>

<template>
  <RouterView/>
</template>