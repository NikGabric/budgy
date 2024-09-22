<script setup lang="ts">
const toggleValue: Ref<boolean> = ref(false);
watch(toggleValue, async (newValue, oldValue) => {
  document
    .getElementById("theme-wrapper")
    ?.setAttribute("data-theme", newValue ? "budgy-dark" : "budgy");
});
</script>

<template>
  <div>
    <input type="checkbox" class="toggle" v-model="toggleValue" />
  </div>
</template>
