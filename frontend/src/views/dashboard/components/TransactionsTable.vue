<script setup lang="ts">
import type { Transaction } from '@/composables/useApi';
import { formatDate } from '@/utils/formatters';

defineProps({
  transactions: {
    type: Array<Transaction>,
    required: true,
  },
});
</script>

<template>
  <table class="table-auto w-full border border-neutral-400">
    <thead>
      <tr class="">
        <th>Date</th>
        <th>Description</th>
        <th>Type</th>
        <th>Amount</th>
      </tr>
    </thead>

    <tbody>
      <tr
        v-for="t in transactions"
        :key="t.id"
        class="hover:bg-neutral-200 hover:cursor-pointer transition-all border-t border-neutral-200"
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
</template>
