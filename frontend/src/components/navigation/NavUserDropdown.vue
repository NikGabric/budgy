<script setup lang="ts">
import { useUserStore } from '@/stores/user';
import { LogOut } from 'lucide-vue-next';
import { ref } from 'vue';
import { useRouter } from 'vue-router';

const router = useRouter();
const userStore = useUserStore();
const firstLetter = userStore.user?.username.charAt(0).toUpperCase();

const isOpen = ref(false);

const toggleOpen = () => (isOpen.value = !isOpen.value);
const handleLogout = async () => {
  userStore.logout();
  router.push('/login');
};
</script>

<template>
  <div class="relative">
    <button
      class="hover:bg-secondary-100 transition-all w-12 h-12 rounded-full"
      @click="toggleOpen"
    >
      {{ firstLetter }}
    </button>
    <div
      v-if="isOpen"
      class="absolute right-0 min-w-48 mt-4 p-2 bg-neutral-400 border"
    >
      <button
        class="hover:bg-neutral-200 flex gap-2 items-center w-full p-2 transition-all"
        @click="handleLogout"
      >
        <LogOut />
        Logout
      </button>
    </div>
  </div>
</template>
