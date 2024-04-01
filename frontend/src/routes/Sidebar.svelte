<script lang="ts">
	import { goto } from "$app/navigation";
	import { Gauge, Home } from "lucide-svelte";

  export let sidebarOpen: boolean;
  $: sidebarClass = sidebarOpen ? 'w-1/6 min-w-48' : 'w-16 min-w-16';

  const logout = () => {
    goto("/auth/login");
  };
</script>

<div class="transition-all bg-indigo-200 h-screen sticky top-0 left-0 p-1 {sidebarClass}">
  <a 
    href="/" 
    class="w-full block text-center hover:bg-red-200 transition-all p-2 rounded-lg"
  >
    {#if sidebarOpen}
      Home
    {:else}
      <Home />
    {/if}
  </a>
  <a
    href="/app/dashboard"
    class="w-full block text-center hover:bg-red-200 transition-all p-2 rounded-lg"
  >
  {#if sidebarOpen}
    <span>Dashboard</span>
  {:else}
    <Gauge />
  {/if}
  </a>

  <button
    on:click={logout}
    class="self-end w-full block text-center hover:bg-red-200 transition-all p-2 rounded-lg"
  >
  {#if sidebarOpen}
    <span>Logout</span>
  {:else}
    <Gauge />
  {/if}
</button>
</div>