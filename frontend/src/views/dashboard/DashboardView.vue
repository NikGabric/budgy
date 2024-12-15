<script setup lang="ts">
import ViewHeader from '@/components/ViewHeader.vue';
import { Gauge } from 'lucide-vue-next';
import TransactionsTable from '@/views/dashboard/components/TransactionsTable.vue';
import TransactionsSummary from './components/TransactionsSummary.vue';
</script>

<template>
  <div>
    <ViewHeader title="Dashboard" :icon="Gauge" />

    <Suspense>
      <TransactionsSummary />
      <template #fallback>Loading...</template>
    </Suspense>

    <Suspense>
      <TransactionsTable />
      <template #fallback>Loading...</template>
    </Suspense>
  </div>
</template>
