<script setup lang="ts">
import { ChevronRight } from "lucide-vue-next";
import { type BreadcrumbItem } from "./breadcrumbs";
defineProps({
  items: {
    type: Object as () => BreadcrumbItem[],
    required: true,
  },
});
</script>

<template>
  <div v-for="(item, i) in items" class="flex items-center gap-2 mr-2">
    <NuxtLink :to="item.to" class="group">
      {{ item.label }}
      <span
        class="block max-w-0 group-hover:max-w-full transition-all duration-200 h-0.5 bg-neutral-300"
      />
    </NuxtLink>
    <ChevronRight v-if="i < items.length - 1" :size="16" />
  </div>
</template>
