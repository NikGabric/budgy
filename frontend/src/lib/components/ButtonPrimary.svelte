<script lang="ts">
  export let onClick: () => void = () => {};
</script>

<button 
  on:click={onClick}
  class="w-full bg-gray-700 hover:bg-gray-500 transition-all hover:scale-105 p-2 rounded-lg shadow-md text-white"
>
  Login
</button>