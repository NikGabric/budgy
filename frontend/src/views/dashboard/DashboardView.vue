<script setup lang="ts">
import ViewHeader from '@/components/ViewHeader.vue';
import { Gauge } from 'lucide-vue-next';
import TransactionsTable from '@/views/dashboard/_components/TransactionsTable.vue';
import TransactionsSummary from './_components/TransactionsSummary.vue';
import {
  type GetUserTransactionsParams,
  type PaginationParams,
  type Transaction,
  useGetUserTransactions,
} from '@/composables/useApi';
import { type Ref, ref } from 'vue';
import TransactionsFilters from './_components/TransactionsFilters.vue';
import { type PaginationResponse } from '../../api/types';

const transactionParams: Ref<GetUserTransactionsParams & PaginationParams> =
  ref({});
const transactions: Ref<PaginationResponse<Transaction[]>> = ref({
  data: [],
  totalCount: 0,
});

// TODO: fix 2 api calls on mount
const handleUpdateTransactions = async (
  params?: GetUserTransactionsParams & PaginationParams,
) => {
  transactionParams.value = { ...transactionParams.value, ...params };
  transactions.value = await useGetUserTransactions(transactionParams.value);
};
</script>

<template>
  <div class="flex flex-col gap-4">
    <ViewHeader title="Dashboard" :icon="Gauge" />

    <TransactionsFilters @handle-filters-change="handleUpdateTransactions" />

    <TransactionsSummary :transactions="transactions?.data" />

    <TransactionsTable
      :transactions="transactions"
      @pagination-change="handleUpdateTransactions"
    />
  </div>
</template>
