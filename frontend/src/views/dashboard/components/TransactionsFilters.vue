<script setup lang="ts">
import {
  useGetUserTransactionTypes,
  type GetUserTransactionsParams,
  type TransactionType,
} from '@/composables/useApi';
import { onMounted, ref, watch, type Ref } from 'vue';

const emit = defineEmits<{
  (e: 'handleFiltersChange', params?: GetUserTransactionsParams): void;
}>();

const params: Ref<GetUserTransactionsParams> = ref({});
const emitChange = () => emit('handleFiltersChange', params.value);
watch(params.value, () => {
  emitChange();
});

const transactionTypes = await useGetUserTransactionTypes();
const transactionType: Ref<TransactionType | null> = ref(null);

const handleTransactionTypeChange = (e: Event) => {
  const el = e.target as HTMLSelectElement;
  const selected = el.options[el.selectedIndex];
  if (selected.value === 'All') {
    params.value.transactionTypeId = undefined;
    return;
  }

  params.value.transactionTypeId = transactionTypes.find(
    (el) => el.id.toString() === selected.value,
  )?.id;
};

const d = new Date();
d.setDate(d.getDate() - 7);
params.value.fromDate = d.toISOString().slice(0, 10);
params.value.toDate = new Date().toISOString().slice(0, 10);

const clearFilters = () => {
  transactionType.value = null;
  params.value.fromDate = undefined;
  params.value.toDate = new Date().toISOString().slice(0, 10);
};

const handleThisMonth = () => {
  const d = new Date();
  params.value.fromDate = new Date(d.getFullYear(), d.getMonth(), 1, 1)
    .toISOString()
    .slice(0, 10);
  params.value.toDate = new Date(d.getFullYear(), d.getMonth() + 1, 1)
    .toISOString()
    .slice(0, 10);
};

onMounted(() => emitChange());
</script>

<template>
  <div class="flex justify-between gap-2 items-end">
    <label>
      <div class="label"><span class="label-text">Transaction type</span></div>
      <select
        class="select select-primary select-sm w-32"
        @change="handleTransactionTypeChange"
      >
        <option value="all" :selected="!params.transactionTypeId">All</option>
        <option
          v-for="item in transactionTypes"
          :key="item.id"
          :value="item.id"
          :selected="params.transactionTypeId === item.id"
        >
          {{ item.name }}
        </option>
      </select>
    </label>

    <input
      type="date"
      v-model.lazy="params.fromDate"
      @change="emitChange"
      class="h-8 w-40 px-2 rounded-lg bg-base-100 border border-primary"
    />
    <input
      type="date"
      v-model="params.toDate"
      @change="emitChange"
      class="h-8 w-40 px-2 rounded-lg bg-base-100 border border-primary"
    />

    <div class="flex gap-2">
      <div class="dropdown dropdown-end">
        <div
          tabindex="0"
          role="button"
          class="btn btn-outline btn-primary btn-sm"
        >
          Quick filters
        </div>
        <ul
          tabindex="0"
          class="dropdown-content menu bg-base-300 rounded-box z-[1] w-52 p-2 shadow my-2"
        >
          <li><button @click="handleThisMonth">This month</button></li>
          <li><a>Item 2</a></li>
        </ul>
      </div>

      <button
        class="btn btn-secondary btn-sm justify-self-end"
        @click="clearFilters"
      >
        Clear filters
      </button>
    </div>
  </div>
</template>
