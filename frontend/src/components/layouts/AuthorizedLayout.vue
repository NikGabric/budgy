<script setup lang="ts">
import { RouterView } from 'vue-router';
import SideNav from '../navigation/SideNav.vue';
import TopNav from '../navigation/TopNav.vue';
import { ref, type Ref } from 'vue';

const localSidebarOpen = localStorage.getItem('budgy_sidebar_open');
const sidebarOpen: Ref<boolean> = ref(
  localSidebarOpen ? JSON.parse(localSidebarOpen) : true,
);

const toggleSidebarOpen = () => (sidebarOpen.value = !sidebarOpen.value);
</script>

<template>
  <div class="flex w-screen h-screen">
    <SideNav
      class="transition-all"
      :class="sidebarOpen ? 'w-72' : 'w-16'"
      :sidebar-open="sidebarOpen"
    />
    <div class="flex-1">
      <header>
        <TopNav
          :sidebar-open="sidebarOpen"
          @toggle-sidebar="toggleSidebarOpen"
        />
      </header>

      <RouterView />
    </div>
  </div>
</template>
