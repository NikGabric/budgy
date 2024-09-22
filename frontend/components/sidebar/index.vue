<script setup lang="ts">
import { LogOut, Settings } from "lucide-vue-next";
import type { SidebarItem } from ".";

defineProps({
  sidebarOpen: {
    type: Boolean,
    required: true,
  },
});

const defaultItems: SidebarItem[] = [
  {
    label: "Settings",
    to: "/",
    icon: Settings,
  },
  {
    label: "Logout",
    to: "/",
    icon: LogOut,
    styles: "text-error",
  },
];
</script>

<template>
  <div class="bg-primary p-x-2 text-base-100">
    <transition
      name="fade"
      mode="out-in"
      class="flex items-center justify-center"
    >
      <NuxtLink
        class="w-full text-center text-3xl h-16"
        v-if="sidebarOpen"
        to="/"
      >
        Budgy
      </NuxtLink>
      <NuxtLink class="w-full text-center text-3xl h-16" v-else to="/"
        >BG</NuxtLink
      >
    </transition>

    <div class="w-full h-full flex flex-col justify-between">
      <div>
        <sidebar-item
          v-for="item in navigationItems"
          :item="item"
          :sidebar-open="sidebarOpen"
        />
      </div>

      <div>
        <sidebar-item
          v-for="item in defaultItems"
          :item="item"
          :sidebar-open="sidebarOpen"
        />
      </div>
    </div>
  </div>
</template>
