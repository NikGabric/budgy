<script setup lang="ts">
import ViewHeader from '@/components/ViewHeader.vue';
import { Gauge } from 'lucide-vue-next';
import TransactionsTable from '@/views/dashboard/components/TransactionsTable.vue';
import TransactionsSummary from './components/TransactionsSummary.vue';
import { type Transaction, useGetUserTransactions } from '@/composables/useApi';
import { type Ref, ref } from 'vue';
import TransactionsFilters from './components/TransactionsFilters.vue';

const transactions: Ref<Transaction[]> = ref(await useGetUserTransactions());

const handleUpdateTransactionTypeFilter = async (id?: number) => {
  transactions.value = await useGetUserTransactions({
    transaction_type_id: id,
  });
};
</script>

<template>
  <div>
    <ViewHeader title="Dashboard" :icon="Gauge" />

    <TransactionsFilters
      @handle-transaction-type-change="handleUpdateTransactionTypeFilter"
    />

    <TransactionsSummary :transactions="transactions" />

    <TransactionsTable :transactions="transactions" />
  </div>
</template>
