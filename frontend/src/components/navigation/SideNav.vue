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
  <div class="h-full bg-primary-200 border-r">
    <RouterLink
      v-for="item in items"
      :key="item.title"
      :to="item.to"
      class="flex gap-2 px-4 py-3 m-1 h-fit rounded-lg hover:bg-neutral-400 bg-opacity-50 transition-all"
    >
      <!-- TODO: fix icon behavior when expanding -->
      <component :is="item.icon" />
      <span v-if="sidebarOpen">
        {{ item.title }}
      </span>
    </RouterLink>
  </div>
</template>
