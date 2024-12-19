<script setup lang="ts">
import {
  useGetUserTransactionTypes,
  type TransactionType,
} from '@/composables/useApi';
import { ref, type Ref } from 'vue';

const emit = defineEmits<{
  (e: 'handleTransactionTypeChange', id?: number): void;
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
</script>

<template>
  <div>
    <select
      class="p-3 outline outline-1 bg-neutral-200 hover:bg-neutral-300 hover:shadow-sm transition-colors rounded-md"
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
  </div>
</template>
