<script setup lang="ts">
import { DEFAULT_PAGINATION_LIMIT } from '@/app-constants';
import type { PaginationParams } from '@/composables/useApi';
import { ref, type Ref, computed, type ComputedRef } from 'vue';

const props = defineProps({
  totalCount: {
    type: Number,
    required: true,
  },
});

const emit = defineEmits<{
  (e: 'paginationChange', pagination: PaginationParams): void;
}>();
const options = [DEFAULT_PAGINATION_LIMIT, 10, 20];
const params: Ref<PaginationParams> = ref({
  page: 1,
  limit: DEFAULT_PAGINATION_LIMIT,
});

const pages: ComputedRef<number[]> = computed(() => {
  const pagesCount = Math.ceil(props.totalCount / params.value.limit!);
  return Array.from({ length: pagesCount }, (_, i) => i + 1);
});
const handlePageClick = (page: number) => {
  params.value.page = page;
  emitChange();
};

const emitChange = () => emit('paginationChange', params.value);
</script>

<template>
  <div class="flex gap-2 justify-between m-2">
    <select
      class="select select-bordered select-sm"
      v-model="params.limit"
      @change="emitChange"
    >
      <option v-for="option in options" :key="option">
        {{ option }}
      </option>
    </select>

    <div class="join">
      <button
        v-for="page in pages"
        :key="page"
        @click="() => handlePageClick(page)"
        class="join-item btn btn-sm btn-outline"
        :class="{ 'btn-primary': page === params.page }"
        :disabled="page === params.page"
      >
        {{ page }}
      </button>
    </div>
  </div>
</template>
