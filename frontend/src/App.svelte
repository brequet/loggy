<script lang="ts">
  import FilterPanel from "$lib/components/FilterPanel.svelte";
  import LogViewer from "./lib/components/LogViewer.svelte";
</script>

<main class="h-screen max-h-screen w-screen bg-gray-100 text-black">
  <div
    class="container h-full mx-auto px-4 py-8 overflow-y-hidden flex flex-col"
  >
    <header class="mb-4 flex justify-between">
      <h1 class="text-3xl font-bold">Loggy</h1>
      <FilterPanel />
    </header>
    <LogViewer />
  </div>
</main>
