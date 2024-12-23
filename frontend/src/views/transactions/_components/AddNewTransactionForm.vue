<script setup lang="ts">
import { type TransactionDto, postTransaction } from '@/api/api';
import { useGetUserTransactionTypes } from '@/composables/useApi';
import { type Ref, ref } from 'vue';
import { useRouter } from 'vue-router';
import { X } from 'lucide-vue-next';

const transactionTypes = await useGetUserTransactionTypes();

const transaction: Ref<TransactionDto> = ref({
  amount: 10,
  description: '',
  transactionDate: new Date().toISOString().slice(0, 10),
  transactionTypeId: transactionTypes[0].id,
});

const handleTransactionTypeChange = (e: Event) => {
  const el = e.target as HTMLSelectElement;
  const selected = el.options[el.selectedIndex];

  transaction.value.transactionTypeId = parseInt(selected.value);
};

const router = useRouter();
const handleSubmit = async () => {
  try {
    const res = await postTransaction(transaction.value);
    router.push({ name: 'transaction-details', params: { id: res.id } });
  } catch (e) {
    console.error(e);
  }
};
</script>

<template>
  <div class="modal-box w-3/4 max-w-[1560px]">
    <form method="dialog">
      <button class="btn btn-sm btn-circle btn-ghost absolute right-2 top-2">
        <X />
      </button>
    </form>
    <h3 class="text-lg font-bold">Add new transaction</h3>

    <form @submit.prevent="handleSubmit" class="flex flex-col gap-2 mt-2">
      <input
        type="text"
        placeholder="Description"
        class="input input-bordered w-96"
        v-model="transaction.description"
      />
      <div class="join">
        <input
          type="number"
          class="input input-bordered w-96 mr-2"
          v-model="transaction.amount"
        />
        <span>EUR</span>
      </div>
      <input
        type="date"
        class="input input-bordered w-96"
        v-model="transaction.transactionDate"
      />

      <select
        class="select select-primary w-96"
        @change="handleTransactionTypeChange"
      >
        <option
          v-for="item in transactionTypes"
          :key="item.id"
          :value="item.id"
          :selected="transaction.transactionTypeId === item.id"
        >
          {{ item.name }}
        </option>
      </select>

      <button type="submit">Submit</button>
    </form>
  </div>

  <form method="dialog" class="modal-backdrop">
    <button>close</button>
  </form>
</template>
