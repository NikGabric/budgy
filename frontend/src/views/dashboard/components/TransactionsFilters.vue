<script setup lang="ts">
import {
  useGetUserTransactionTypes,
  type TransactionType,
} from '@/composables/useApi';
import { ref, type Ref } from 'vue';

const emit = defineEmits<{
  (e: 'handleTransactionTypeChange', id?: number): void;
  (e: 'handleFromDateChange', date?: Date): void;
  (e: 'handleToDateChange', date?: Date): void;
}>();

const transactionTypes = await useGetUserTransactionTypes();
const transactionType: Ref<TransactionType | null> = ref(null);

const handleChangeTransactionType = (e: Event) => {
  const el = e.target as HTMLSelectElement;
  const selectedOption = el.options[el.selectedIndex];

  if (selectedOption.value === 'All') {
    emit('handleTransactionTypeChange');
    return;
  }

  transactionType.value =
    transactionTypes.find((el) => el.id.toString() === selectedOption.value) ??
    null;
  emit('handleTransactionTypeChange', transactionType.value?.id);
};

//TODO: handle change
const d = new Date();
d.setDate(d.getDate() - 7);
const fromDate: Ref<string | null> = ref(d.toISOString().slice(0, 10));
const toDate: Ref<string | null> = ref(new Date().toISOString().slice(0, 10));

const clearFilters = () => {
  transactionType.value = null;
  fromDate.value = null;
  toDate.value = null;
  // TODO: handle change
};
</script>

<template>
  <div class="flex gap-2 items-end">
    <label>
      <div class="label"><span class="label-text">Transaction type</span></div>
      <select
        class="select select-primary select-sm w-32"
        @change="handleChangeTransactionType"
      >
        <option value="all" :selected="transactionType === null">All</option>
        <option
          v-for="item in transactionTypes"
          :key="item.id"
          :value="item.id"
          :selected="transactionType?.id === item.id"
        >
          {{ item.name }}
        </option>
      </select>
    </label>

    <input
      type="date"
      v-model.lazy="fromDate"
      class="h-8 px-2 rounded-lg bg-base-100 border border-primary"
    />
    <input
      type="date"
      v-model="toDate"
      class="h-8 px-2 rounded-lg bg-base-100 border border-primary"
    />

    <button
      class="btn btn-secondary btn-sm justify-self-end"
      @click="clearFilters"
    >
      Clear filters
    </button>
  </div>
</template>
