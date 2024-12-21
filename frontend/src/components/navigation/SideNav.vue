<script setup lang="ts">
import { useUserStore } from '@/stores/user';
import {
  ChartColumnBig,
  LayoutDashboard,
  LogOut,
  Settings,
  type LucideIcon,
} from 'lucide-vue-next';
import { useRouter } from 'vue-router';

defineProps({
  sidebarOpen: {
    type: Boolean,
    required: true,
  },
});

interface SidebarItem {
  title: string;
  icon: LucideIcon;
  action: () => void;
  style?: string;
}

const router = useRouter();
const userStore = useUserStore();

const items: SidebarItem[] = [
  {
    title: 'Dashboard',
    icon: LayoutDashboard,
    action: () => router.push('/dashboard'),
  },
  {
    title: 'Budget',
    icon: ChartColumnBig,
    action: () => router.push('/budget'),
  },
];

const handleLogout = () => {
  userStore.logout();
  router.push('/login');
};

const endItems: SidebarItem[] = [
  {
    title: 'Settings',
    icon: Settings,
    action: () => router.push('/settings'),
  },
  {
    title: 'Logout',
    icon: LogOut,
    action: handleLogout,
    style: 'text-error',
  },
];
</script>

<template>
  <div class="flex flex-col h-screen bg-base-200">
    <div class="h-16 text-center text-3xl">Budgy</div>
    <div class="flex-1 min-h-0 flex flex-col justify-between">
      <div class="flex flex-col">
        <button
          v-for="item in items"
          :key="item.title"
          class="btn btn-ghost m-2"
          @click="item.action"
        >
          <component :is="item.icon" />
          <span v-if="sidebarOpen">{{ item.title }}</span>
        </button>
      </div>

      <div class="flex flex-col">
        <button
          v-for="item in endItems"
          :key="item.title"
          class="btn btn-ghost m-2"
          :class="item.style"
          @click="item.action"
        >
          <component :is="item.icon" />
          <span v-if="sidebarOpen">{{ item.title }}</span>
        </button>
      </div>
    </div>
  </div>
</template>
