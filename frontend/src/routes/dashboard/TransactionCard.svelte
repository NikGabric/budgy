<script lang="ts">
	import type { Transaction } from "$lib/models/transaction";

  export let transaction: Transaction;
</script>

<div>
  <div>{transaction.title}</div>
</div>