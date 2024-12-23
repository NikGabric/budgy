<script setup lang="ts">
import AddNewTransactionForm from './AddNewTransactionForm.vue';
</script>

<template>
  <div>
    <button class="btn btn-primary" onclick="add_new_modal.showModal()">
      Add new
    </button>
    <dialog id="add_new_modal" class="modal">
      <Suspense>
        <AddNewTransactionForm />

        <template #fallback>Loading</template>
      </Suspense>
    </dialog>
  </div>
</template>
