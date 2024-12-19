<script setup lang="ts">
import { LayoutDashboard, type LucideIcon } from 'lucide-vue-next';

defineProps({
  sidebarOpen: {
    type: Boolean,
    required: true,
  },
});

interface SidebarItem {
  title: string;
  icon: LucideIcon;
  to: string;
}

const items: SidebarItem[] = [
  {
    title: 'Dashboard',
    icon: LayoutDashboard,
    to: '/dashboard',
  },
];
</script>

<template>
  <div class="h-full border-r border-primary">
    <RouterLink
      v-for="item in items"
      :key="item.title"
      :to="item.to"
      class="flex justify-start btn btn-ghost m-2"
    >
      <!-- TODO: fix icon behavior when expanding -->
      <component :is="item.icon" />
      <span v-if="sidebarOpen">
        {{ item.title }}
      </span>
    </RouterLink>
  </div>
</template>
