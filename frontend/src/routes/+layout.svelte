<script lang="ts">
	import { AlignJustify, ChevronLeft } from 'lucide-svelte';
	import '../app.css';
	import Sidebar from './Sidebar.svelte';
	import { page } from '$app/stores';

	let sidebarOpen: boolean = true;
	const toggleSidebar = () => (sidebarOpen = !sidebarOpen);

	$: isAuthPage = $page.url.pathname.startsWith('/auth');
</script>

<div class="flex w-screen">
	{#if !isAuthPage}
		<Sidebar {sidebarOpen} />
	{/if}

	<div class="w-full">
		{#if !isAuthPage}
			<div class="sticky top-0 flex px-2 items-center bg-gray-500 h-12">
				<button class="transition-all hover:bg-gray-200 rounded-full p-2" on:click={toggleSidebar}>
					{#if sidebarOpen}
						<ChevronLeft />
					{:else}
						<AlignJustify />
					{/if}
				</button>
			</div>
		{/if}

		<div class="">
			<slot />
		</div>
	</div>
</div>
