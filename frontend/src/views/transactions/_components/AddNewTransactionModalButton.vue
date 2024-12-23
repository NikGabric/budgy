<script setup lang="ts">
import { postTransaction, type TransactionDto } from '@/api/api';
import { X } from 'lucide-vue-next';
import { ref, type Ref } from 'vue';
import { useRouter } from 'vue-router';

const transaction: Ref<TransactionDto> = ref({
  amount: 10,
  description: '',
  transactionDate: new Date().toISOString().slice(0, 10),
  transactionTypeId: 1,
});

const router = useRouter();
const handleSubmit = async () => {
  try {
    const res = await postTransaction(transaction.value);
    router.push({ name: 'transaction-details', params: { id: res.id } });
  } catch (e) {
    console.error(e);
  }
};
</script>

<template>
  <div>
    <button class="btn btn-primary" onclick="add_new_modal.showModal()">
      Add new
    </button>
    <dialog id="add_new_modal" class="modal">
      <div class="modal-box w-3/4 max-w-[1560px]">
        <form method="dialog">
          <button
            class="btn btn-sm btn-circle btn-ghost absolute right-2 top-2"
          >
            <X />
          </button>
        </form>
        <h3 class="text-lg font-bold">Add new transaction</h3>

        <form @submit.prevent="handleSubmit" class="flex flex-col gap-2 mt-2">
          <input
            type="text"
            placeholder="Description"
            class="input input-bordered w-96"
            v-model="transaction.description"
          />
          <div class="join">
            <input
              type="number"
              class="input input-bordered w-96 mr-2"
              v-model="transaction.amount"
            />
            <span>EUR</span>
          </div>
          <input
            type="date"
            class="input input-bordered w-96"
            v-model="transaction.transactionDate"
          />
          <select>
            <option></option>
          </select>

          <button type="submit">Submit</button>
        </form>
      </div>

      <form method="dialog" class="modal-backdrop">
        <button>close</button>
      </form>
    </dialog>
  </div>
</template>
