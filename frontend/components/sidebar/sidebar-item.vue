<script setup lang="ts">
defineProps({
  item: {
    type: Object as () => NavigationItem,
    required: true,
  },
  sidebarOpen: {
    type: Boolean,
    required: true,
  },
});
</script>

<template>
  <div
    class="w-full"
    :class="{ 'tooltip tooltip-right tooltip-secondary': !sidebarOpen }"
    :data-tip="item.label"
  >
    <div v-if="!item.to" class="text-sm text-neutral-300">
      {{ item.label }}
    </div>

    <NuxtLink
      class="btn btn-ghost overflow-hidden w-full justify-start"
      :to="item.to"
      v-else
    >
      <div class="flex items-center space-x-2" :class="item.styles">
        <component :is="iconMap[item.iconName]" />
        <span v-if="sidebarOpen" class="whitespace-nowrap">{{
          item.label
        }}</span>
      </div>
    </NuxtLink>
  </div>
</template>
