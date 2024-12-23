<script setup lang="ts">
import { getTransaction } from '@/api/api';
import { useRoute } from 'vue-router';

const route = useRoute();

const transaction = await getTransaction(parseInt(route.params.id as string));
</script>

<template>
  <div>
    {{ transaction }}
  </div>
</template>
