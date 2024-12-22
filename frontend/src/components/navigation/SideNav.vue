<script setup lang="ts">
import { endItems, startItems } from './side-nav';

defineProps({
  sidebarOpen: {
    type: Boolean,
    required: true,
  },
});
</script>

<template>
  <div
    class="flex flex-col items-center h-screen bg-base-200"
    :class="sidebarOpen ? 'w-72' : 'w-16'"
  >
    <div class="h-16 w-full text-3xl text-center">Budgy</div>

    <div class="flex flex-col w-full justify-between h-full">
      <div class="flex flex-col w-full p-1">
        <div
          v-for="item in startItems"
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
