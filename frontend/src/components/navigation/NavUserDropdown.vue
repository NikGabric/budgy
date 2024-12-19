<script setup lang="ts">
import { useUserStore } from '@/stores/user';
import { useRouter } from 'vue-router';

const router = useRouter();
const userStore = useUserStore();
const firstLetter = userStore.user?.username.charAt(0).toUpperCase();

const handleLogout = async () => {
  userStore.logout();
  router.push('/login');
};
</script>

<template>
  <div class="relative">
    <div class="dropdown dropdown-end">
      <div tabindex="0" role="button" class="btn btn-circle btn-ghost">
        {{ firstLetter }}
      </div>
      <ul
        tabindex="0"
        class="dropdown-content menu bg-base-200 rounded-box z-[1] w-52 p-2 shadow"
      >
        <li>
          <button @click="handleLogout" class="text-error">Logout</button>
        </li>
      </ul>
    </div>
  </div>
</template>
