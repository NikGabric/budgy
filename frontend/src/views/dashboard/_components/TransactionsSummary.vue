<script setup lang="ts">
import { type Transaction } from '../../../composables/useApi';

defineProps({
  transactions: {
    type: Array<Transaction>,
    required: false,
  },
});
</script>

<template>
  <div>
    <!-- TODO: summary -->
    Summary!
  </div>
</template>
