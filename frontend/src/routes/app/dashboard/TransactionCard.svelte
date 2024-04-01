<script lang="ts">
	import type { Transaction } from "$lib/models/transaction";

  export let transaction: Transaction;
</script>

<div class="flex flex-col items-center justify-center border-2 hover:shadow-lg transition-shadow hover:cursor-pointer m-2 p-2 rounded-md">
  <div>{transaction.title}</div>
  <div>{transaction.description}</div>
  <div>{transaction.amount} eur</div>
</div>