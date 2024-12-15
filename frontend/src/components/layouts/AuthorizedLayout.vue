<script setup lang="ts">
import { RouterView } from 'vue-router';
import SideNav from '../navigation/SideNav.vue';
import TopNav from '../navigation/TopNav.vue';
</script>

<template>
  <div class="flex w-screen h-screen">
    <SideNav class="w-64" />
    <div class="flex-1">
      <header>
        <TopNav />
      </header>

      <RouterView />
    </div>
  </div>
</template>
