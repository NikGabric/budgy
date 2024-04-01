<script lang="ts">
	import { post } from '$lib/api/api';
	import ButtonPrimary from '$lib/components/ButtonPrimary.svelte';
	import { authToken } from '$lib/shared/stores/auth';

	let username: String = '';
	let password: String = '';

	const loginUser = async () => {
    const token = await post("/users/login", { username, password });
    authToken.set(token);
	};
</script>

<div class="flex flex-col w-full gap-4 my-4 justify-center items-center">
	<div>Login to Budgy</div>

	<form class="flex flex-col gap-4 w-1/3" on:submit={loginUser}>
		<label>
			Username:
			<input class="w-full border-2 rounded-md p-2" type="text" bind:value={username} />
		</label>
		<label>
			Password:
			<input class="w-full border-2 rounded-md p-2" type="password" bind:value={password} />
		</label>
		<ButtonPrimary />
	</form>

</div>
