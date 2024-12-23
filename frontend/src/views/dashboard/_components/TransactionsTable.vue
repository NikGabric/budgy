<script setup lang="ts">
import type { PaginationParams, Transaction } from '@/composables/useApi';
import { formatDate } from '@/utils/formatters';
import { type PaginationResponse } from '../../../api/types';
import { type PropType } from 'vue';
import PaginationFilters from '@/components/common/PaginationFilters.vue';
import { useRouter } from 'vue-router';

defineProps({
  transactions: {
    type: Object as PropType<PaginationResponse<Transaction[]>>,
    required: true,
  },
});

const emit = defineEmits<{
  (e: 'paginationChange', params?: PaginationParams): void;
}>();
const emitChange = (params: PaginationParams) =>
  emit('paginationChange', params);

const router = useRouter();
const navigateTo = (id: number) => router.push(`/transaction/${id}`);
</script>

<template>
  <div class="w-full border rounded-lg p-2">
    <table class="table table-lg">
      <thead class="">
        <tr>
          <th>Date</th>
          <th>Description</th>
          <th>Type</th>
          <th>Amount</th>
        </tr>
      </thead>

      <tbody>
        <tr
          v-for="t in transactions?.data"
          :key="t.id"
          class="hover:bg-base-200 hover:cursor-pointer"
          @click="() => navigateTo(t.id)"
        >
          <td class="p-2 text-center">
            {{ formatDate(t.createdAt) }}
          </td>
          <td class="p-2 text-center">{{ t.description }}</td>
          <td class="p-2 text-center">{{ t.transactionTypeName }}</td>
          <td class="p-2 text-center">{{ t.amount }} eur</td>
        </tr>
      </tbody>
    </table>

    <PaginationFilters
      :total-count="transactions.totalCount"
      @pagination-change="emitChange"
    />
  </div>
</template>
