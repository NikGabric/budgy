<script setup lang="ts">
const sidebarOpen: Ref<boolean> = ref(true);
const toggleSidebarOpen = () => (sidebarOpen.value = !sidebarOpen.value);
</script>

<template>
  <div class="flex">
    <Sidebar
      :class="sidebarOpen ? 'w-64' : 'w-16'"
      class="transition-all sticky top-0 self-start h-screen flex flex-col"
      :sidebar-open="sidebarOpen"
    />
    <div class="w-full">
      <Navbar
        class="sticky top-0"
        :sidebar-open="sidebarOpen"
        @toggle-sidebar="toggleSidebarOpen"
      />
      <div class="px-4 py-2">
        <slot />
      </div>
    </div>
  </div>
</template>
