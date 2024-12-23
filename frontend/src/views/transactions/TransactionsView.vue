<script setup lang="ts">
import ViewHeader from '@/components/ViewHeader.vue';
import { BadgeEuro } from 'lucide-vue-next';
import Transactions from './Transactions.vue';
</script>

<template>
  <div class="flex flex-col gap-4">
    <ViewHeader :icon="BadgeEuro" title="Transactions" />

    <Suspense>
      <Transactions />

      <template #fallback>Loading</template>
    </Suspense>
  </div>
</template>
