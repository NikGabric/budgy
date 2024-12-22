<script setup lang="ts">
import ViewHeader from '@/components/ViewHeader.vue';
import { Gauge } from 'lucide-vue-next';
import TransactionsTable from '@/views/dashboard/components/TransactionsTable.vue';
import TransactionsSummary from './components/TransactionsSummary.vue';
import {
  type GetUserTransactionsParams,
  type Transaction,
  useGetUserTransactions,
} from '@/composables/useApi';
import { type Ref, ref } from 'vue';
import TransactionsFilters from './components/TransactionsFilters.vue';

const transactions: Ref<Transaction[]> = ref([]);

const handleUpdateTransactionTypeFilter = async (
  params?: GetUserTransactionsParams,
) => {
  transactions.value = await useGetUserTransactions(params);
};
</script>

<template>
  <div class="flex flex-col gap-4">
    <ViewHeader title="Dashboard" :icon="Gauge" />

    <TransactionsFilters
      @handle-filters-change="handleUpdateTransactionTypeFilter"
    />

    <TransactionsSummary :transactions="transactions" />

    <TransactionsTable :transactions="transactions" />
  </div>
</template>
