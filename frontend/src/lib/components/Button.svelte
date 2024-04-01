<script lang="ts">
	export let onClick: () => void = () => {};
	export let type: 'button' | 'submit' | 'reset' | null | undefined;
	export let variant: 'primary' | 'secondary' = 'primary';

	let btnClass =
		variant === 'primary'
			? 'w-full bg-gray-700 hover:bg-gray-500 transition-all hover:scale-105 p-2 rounded-lg shadow-md text-white'
			: 'w-full bg-white hover:bg-gray-200 transition-all hover:scale-105 p-2 rounded-lg shadow-md';
</script>

<button on:click={onClick} {type} class={btnClass}> 
  <slot />
</button>
