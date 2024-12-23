<script setup lang="ts">
import { useGetUserTransactions } from '@/composables/useApi';
import AddNewTransactionModalButton from './_components/AddNewTransactionModalButton.vue';

const transactions = await useGetUserTransactions();
</script>

<template>
  <div>
    <AddNewTransactionModalButton />

    {{ transactions }}
  </div>
</template>
