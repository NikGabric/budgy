<script setup lang="ts">
import { LogOut, Settings } from "lucide-vue-next";
import { useUserStore } from "../../stores/user";

defineProps({
  sidebarOpen: {
    type: Boolean,
    required: true,
  },
});

const { getNavigationItems } = useUserStore();

const defaultItems: NavigationItem[] = [
  {
    label: "Settings",
    to: "/",
    iconName: "Settings",
  },
  {
    label: "Logout",
    to: "/",
    iconName: "LogOut",
    styles: "text-error",
  },
];
</script>

<template>
  <div class="bg-primary rounded-r-xl p-x-2">
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
          v-for="item in getNavigationItems"
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
