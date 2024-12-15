<script setup lang="ts">
import type { Component, PropType } from 'vue';

defineProps({
  title: {
    type: String,
    required: true,
  },
  icon: {
    type: [Object, Function] as PropType<Component>,
    required: false,
  },
});
</script>

<template>
  <div class="flex items-center gap-4 mb-8">
    <component :is="icon" />
    <h3 class="text-3xl">{{ title }}</h3>
  </div>
</template>
