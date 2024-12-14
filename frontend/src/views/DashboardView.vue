<script setup lang="ts">
import { useUserStore } from '@/stores/user';
import { storeToRefs } from 'pinia';
import { ref, type Ref } from 'vue';

const userStore = useUserStore();
const { user } = storeToRefs(userStore);
const { login, logout } = userStore;

const username: Ref<string> = ref('');
const password: Ref<string> = ref('');

const handleLogin = () => {
  login({
    username: username.value,
    password: password.value,
  });
};
</script>

<template>
  <div>
    <div v-if="user">
      Logged in as: {{ user.username }}
      <button @click="logout">Logout</button>
    </div>
    <div v-else>
      <input v-model="username" />
      <input v-model="password" />
      <button class="btn" @click="handleLogin">Login</button>
    </div>
  </div>
</template>
