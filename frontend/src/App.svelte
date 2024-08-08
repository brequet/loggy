<script lang="ts">
  import LogViewer from "./lib/components/LogViewer.svelte";
</script>

<main class="h-screen max-h-screen w-screen bg-gray-900 text-white">
  <div
    class="container h-full mx-auto px-4 py-8 overflow-y-hidden flex flex-col"
  >
    <h1 class="mb-6 text-3xl font-bold">Loggy</h1>
    <LogViewer />
  </div>
</main>
