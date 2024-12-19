<script setup lang="ts">
import { AlignJustify, X } from 'lucide-vue-next';
import NavBreadcrumbs from './NavBreadcrumbs.vue';
import { useUserStore } from '@/stores/user';
import { RouterLink } from 'vue-router';
import NavUserDropdown from './NavUserDropdown.vue';

defineProps({
  sidebarOpen: {
    type: Boolean,
    required: true,
  },
});
defineEmits(['toggleSidebar']);

const userStore = useUserStore();
const { isAuthenticated } = userStore;
</script>

<template>
  <nav
    class="flex items-center justify-between h-16 px-4 w-full border-b bg-primary-200 border-b-primary"
  >
    <div class="flex items-center gap-2" v-if="isAuthenticated">
      <button
        class="rounded-full p-2 transition-all hover:bg-neutral-400"
        @click="$emit('toggleSidebar')"
      >
        <AlignJustify v-if="!sidebarOpen" />
        <X v-else />
      </button>

      <NavBreadcrumbs />
    </div>

    <div v-if="isAuthenticated">
      <NavUserDropdown />
    </div>
    <div v-else>
      <RouterLink
        class="rounded-lg transition-all bg-secondary-400 hover:bg-secondary-200 p-2"
        to="/login"
      >
        Login
      </RouterLink>
    </div>
  </nav>
</template>
