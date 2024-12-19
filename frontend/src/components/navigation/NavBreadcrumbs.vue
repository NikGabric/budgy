<script setup lang="ts">
import { ChevronRight } from 'lucide-vue-next';
import { computed } from 'vue';
import { useRoute } from 'vue-router';

const route = useRoute();

const items = computed(() => route.fullPath.split('/').filter(Boolean));
</script>

<template>
  <div class="flex items-center">
    <div v-for="(item, i) in items" :key="item">
      <RouterLink to="/dashboard" class="capitalize hover:underline">
        {{ item }}
      </RouterLink>
      <ChevronRight v-if="i < items.length - 1" />
    </div>
  </div>
</template>
