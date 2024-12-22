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
  <div
    class="flex flex-col items-center h-screen bg-base-200"
    :class="sidebarOpen ? 'w-72' : 'w-16'"
  >
    <div class="h-16 bg-primary bg-opacity-15 w-full text-3xl text-center">
      Budgy
    </div>

    <div class="flex flex-col w-full justify-between h-full">
      <div class="flex flex-col w-full p-1">
        <div
          v-for="item in items"
          :key="item.title"
          :class="{ 'tooltip tooltip-right tooltip-accent': !sidebarOpen }"
          :data-tip="item.title"
        >
          <button
            class="btn btn-ghost flex w-full"
            :class="sidebarOpen ? 'justify-start' : 'justify-center btn-square'"
            @click="item.action"
          >
            <component :is="item.icon" />
            <div v-if="sidebarOpen">
              {{ item.title }}
            </div>
          </button>
        </div>
      </div>

      <div class="flex flex-col w-full p-1">
        <div
          v-for="item in endItems"
          :key="item.title"
          :class="{ 'tooltip tooltip-right tooltip-accent': !sidebarOpen }"
          :data-tip="item.title"
        >
          <button
            class="btn btn-ghost flex w-full"
            :class="
              sidebarOpen
                ? `justify-start ${item.style}`
                : `justify-center btn-square ${item.style}`
            "
            @click="item.action"
          >
            <component :is="item.icon" />
            <div v-if="sidebarOpen">
              {{ item.title }}
            </div>
          </button>
        </div>
      </div>
    </div>
  </div>
</template>
