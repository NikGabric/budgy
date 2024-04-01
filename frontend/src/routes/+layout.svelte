<script lang="ts">
	import { AlignJustify } from 'lucide-svelte';
	import '../app.css';
	import { authToken } from '$lib/shared/stores/auth';

	let sidebarOpen: Boolean = true;
	const toggleSidebar = () => sidebarOpen != sidebarOpen;

	let token = '';
	authToken.subscribe((value: string | null) => {
		token = value ?? "";
	});
</script>

<div>
	<div>
		<!-- topbar -->
		<div class="flex px-2 items-center bg-gray-500 h-12">
			<button class="transition-all hover:bg-gray-200 rounded-full p-2" on:click={toggleSidebar}>
				<AlignJustify />
			</button>
			{token}
		</div>

		<!-- content -->
		<div>
			<slot />
		</div>
	</div>
</div>
