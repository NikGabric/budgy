<script setup lang="ts">
import { anonGet } from '@/api/api';
import { type ApiResponse } from '@/api/types';
import { ref } from 'vue';

const data = ref<ApiResponse | null>(null);
const getData = async () => {
  data.value = await anonGet('/user/5');
};
</script>

<template>
  <div>
    {{ data }}
    <button class="btn" @click="getData">
      Get data
    </button>
  </div>
</template>
