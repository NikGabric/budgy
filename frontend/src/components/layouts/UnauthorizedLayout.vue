<script setup lang="ts">
import { RouterView } from 'vue-router';
</script>

<template>
  <div class="flex w-screen h-screen">
    <div class="flex-1">
      <div class="flex w-screen justify-center p-4">
        <RouterView />
      </div>
    </div>
  </div>
</template>
