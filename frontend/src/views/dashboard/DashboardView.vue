<script setup lang="ts">
import ViewHeader from '@/components/ViewHeader.vue';
import Dashboard from './Dashboard.vue';
import { Gauge } from 'lucide-vue-next';
</script>

<template>
  <div class="flex flex-col gap-4">
    <ViewHeader title="Dashboard" :icon="Gauge" />

    <Suspense>
      <Dashboard />

      <template #fallback>Loading</template>
    </Suspense>
  </div>
</template>
