<script setup lang="ts">
import {
  AlignJustify,
  ChevronLeft,
  ChevronLeftCircle,
  MoveLeft,
  ToggleLeft,
  X,
} from "lucide-vue-next";
import Breadcrumbs from "../common/breadcrumbs.vue";
import type { BreadcrumbItem } from "../common/breadcrumbs";

defineProps({
  sidebarOpen: {
    type: Boolean,
    required: true,
  },
});

defineEmits(["toggleSidebar"]);

const router = useRouter();
const breadcrumbItems: ComputedRef<BreadcrumbItem[]> = computed(() => {
  const routeParts: string[] = router.currentRoute.value.fullPath
    .split("/")
    .filter(Boolean);
  const breadcrumbs: BreadcrumbItem[] = [];

  routeParts.forEach((part, index) => {
    const fullPath = "/" + routeParts.slice(0, index + 1).join("/");

    breadcrumbs.push({
      label: part.charAt(0).toUpperCase() + part.slice(1),
      to: fullPath,
    });
  });

  breadcrumbs.unshift({
    label: "Home",
    to: "/",
  });

  return breadcrumbs;
});
</script>

<template>
  <div class="flex items-center justify-between h-16 px-2 bg-base-100">
    <div class="flex items-center">
      <button class="btn btn-circle mr-4" @click="$emit('toggleSidebar')">
        <transition name="fade" mode="out-in">
          <ChevronLeft v-if="sidebarOpen" />
          <AlignJustify v-else />
        </transition>
      </button>

      <Breadcrumbs :items="breadcrumbItems" />
    </div>

    <div>
      <NavbarAvatarDropdown />
    </div>
  </div>
</template>
