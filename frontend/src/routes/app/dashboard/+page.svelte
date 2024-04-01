<script lang="ts">
	import { getAuth } from '$lib/api/api';
	import type { Transaction } from '$lib/models/transaction';
	import TransactionCard from './TransactionCard.svelte';
	import { authTokenString, authToken } from '$lib/shared/stores/auth';

	const getTransactions = async () => {
		const res: Transaction[] = await getAuth($authTokenString ?? '', '/transactions');
		return res;
	};
</script>

<div>
	{#await getTransactions()}
		Loading...
	{:then transactions}
		<div class="grid grid-cols-4">
			{#each transactions as transaction}
				<TransactionCard {transaction} />
			{/each}
		</div>
	{:catch error}
		{error}
	{/await}
</div>
