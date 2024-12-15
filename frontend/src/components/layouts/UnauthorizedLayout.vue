<script setup lang="ts">
import { RouterView } from 'vue-router';
import TopNav from '../navigation/TopNav.vue';
</script>

<template>
  <div class="flex w-screen h-screen">
    <div class="flex-1">
      <header>
        <TopNav />
      </header>

      <div class="flex w-screen justify-center p-4">
        <RouterView />
      </div>
    </div>
  </div>
</template>
