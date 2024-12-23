<script setup lang="ts">
import ViewHeader from '@/components/ViewHeader.vue';
import { BadgeEuro } from 'lucide-vue-next';
</script>

<template>
  <div class="flex flex-col gap-4">
    <ViewHeader :icon="BadgeEuro" title="Transactions" />
  </div>
</template>
