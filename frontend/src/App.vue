<script setup lang="ts">
import AuthorizedLayout from './components/layouts/AuthorizedLayout.vue';
import UnauthorizedLayout from './components/layouts/UnauthorizedLayout.vue';
import { useUserStore } from './stores/user';

const userStore = useUserStore();
</script>

<template>
  <div>
    <UnauthorizedLayout v-if="!userStore.isAuthenticated" />
    <AuthorizedLayout v-else />
  </div>
</template>

<style scoped></style>
