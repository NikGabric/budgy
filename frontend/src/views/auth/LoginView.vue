<script setup lang="ts">
import { useUserStore } from '@/stores/user';
import { ref, type Ref } from 'vue';
import { useRouter } from 'vue-router';

const username: Ref<string> = ref('');
const password: Ref<string> = ref('');

const userStore = useUserStore();
const router = useRouter();

const handleLogin = async () => {
  try {
    await userStore.login({
      username: username.value,
      password: password.value,
    });
    router.push('/dashboard');
  } catch (err) {
    alert(err);
  }
};
</script>

<template>
  <div class="mt-16 w-96 flex flex-col gap-4 text-center">
    <h1 class="text-3xl">Welcome to Budgy!</h1>

    <form @submit.prevent="handleLogin" class="flex flex-col gap-4 text-center">
      <input
        type="text"
        v-model="username"
        class="border rounded-md border-neutral-400 p-2"
      />
      <input
        type="password"
        v-model="password"
        class="border rounded-md border-neutral-400 p-2"
      />
      <button type="submit">Login</button>
    </form>
  </div>
</template>
