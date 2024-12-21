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
  <div class="w-full border rounded-lg p-2">
    <table class="table">
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
          v-for="t in transactions"
          :key="t.id"
          class="hover:bg-base-200 hover:cursor-pointer"
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
  </div>
</template>
