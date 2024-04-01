<script lang="ts">
  import { getAuth } from "$lib/api/api";
	import type { Transaction } from "$lib/models/transaction";
	import { authToken } from '$lib/shared/stores/auth';
	import TransactionCard from "./TransactionCard.svelte";

  let token = "";
  authToken.subscribe((val: string | null) => token = val ?? "");

  const getTransactions = async () => {
    const res: Transaction[] = await getAuth(token, "/transactions");
    return res;
  }
</script>

<div>
  {#await getTransactions()}
    Loading...
  {:then transactions} 
    {#each transactions as transaction}
      <TransactionCard {transaction} />
    {/each}
  {:catch error}
    {error}
  {/await}
</div>