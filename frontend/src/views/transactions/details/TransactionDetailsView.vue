<script setup lang="ts">
import ViewHeader from '@/components/ViewHeader.vue';
import TransactionDetails from './TransactionDetails.vue';
import { BadgeEuro } from 'lucide-vue-next';
</script>

<template>
  <div>
    <ViewHeader :icon="BadgeEuro" title="Transaction details" />

    <Suspense>
      <TransactionDetails />

      <template #fallback>Loading</template>
    </Suspense>
  </div>
</template>
