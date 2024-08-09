<script lang="ts">
  import { filters } from "../stores/Filters.svelte";

  let appNameInput = "";
  let levelInput = "";

  function addAppName() {
    if (appNameInput && !filters.appNames.includes(appNameInput)) {
      filters.appNames = [...filters.appNames, appNameInput];
      appNameInput = "";
    }
  }

  function removeAppName(name: string) {
    filters.appNames = filters.appNames.filter((app) => app !== name);
  }

  function addLevel() {
    if (levelInput && !filters.levels.includes(levelInput)) {
      filters.levels = [...filters.levels, levelInput];
      levelInput = "";
    }
  }

  function removeLevel(level: string) {
    filters.levels = filters.levels.filter((l) => l !== level);
  }

  function resetFilters() {
    filters.reset();
  }
</script>

<div class="bg-gray-800 p-4 rounded-lg">
  <h2 class="text-xl font-bold mb-4 text-white">Filters</h2>

  <button
    onclick={resetFilters}
    class="mt-4 w-full px-4 py-2 bg-red-600 text-white rounded-md hover:bg-red-700"
  >
    Reset Filters
  </button>

  <div class="mb-4">
    <label class="block text-sm font-medium text-gray-300">App Names</label>
    <div class="flex mt-1">
      <input
        type="text"
        bind:value={appNameInput}
        onkeyup={(event) => event.key === "Enter" && addAppName()}
        class="flex-grow px-3 py-2 bg-gray-700 text-white rounded-l-md"
        placeholder="Enter app name"
      />
      <button
        onclick={addAppName}
        class="px-4 py-2 bg-blue-600 text-white rounded-r-md hover:bg-blue-700"
      >
        Add
      </button>
    </div>
    <div class="mt-2 flex flex-wrap gap-2">
      {#each filters.appNames as appName}
        <span
          class="bg-blue-500 text-white px-2 py-1 rounded-full text-sm flex items-center"
        >
          {appName}
          <button
            onclick={() => removeAppName(appName)}
            class="ml-2 text-xs bg-blue-600 rounded-full w-4 h-4 flex items-center justify-center hover:bg-blue-700"
          >
            ×
          </button>
        </span>
      {/each}
    </div>
  </div>

  <div class="mb-4">
    <label class="block text-sm font-medium text-gray-300">Levels</label>
    <div class="flex mt-1">
      <input
        type="text"
        bind:value={levelInput}
        onkeyup={(event) => event.key === "Enter" && addLevel()}
        class="flex-grow px-3 py-2 bg-gray-700 text-white rounded-l-md"
        placeholder="Enter log level"
      />
      <button
        onclick={addLevel}
        class="px-4 py-2 bg-blue-600 text-white rounded-r-md hover:bg-blue-700"
      >
        Add
      </button>
    </div>
    <div class="mt-2 flex flex-wrap gap-2">
      {#each filters.levels as level}
        <span
          class="bg-green-500 text-white px-2 py-1 rounded-full text-sm flex items-center"
        >
          {level}
          <button
            onclick={() => removeLevel(level)}
            class="ml-2 text-xs bg-green-600 rounded-full w-4 h-4 flex items-center justify-center hover:bg-green-700"
          >
            ×
          </button>
        </span>
      {/each}
    </div>
  </div>

  <div class="mb-4">
    <label class="block text-sm font-medium text-gray-300">Start Date</label>
    <input
      type="datetime-local"
      step="1"
      bind:value={filters.startDate}
      onkeyup={(event) => console.log("value", event?.target?.value)}
      class="mt-1 block w-full px-3 py-2 bg-gray-700 text-white rounded-md"
    />
  </div>

  <div class="mb-4">
    <label class="block text-sm font-medium text-gray-300">End Date</label>
    <input
      type="datetime-local"
      bind:value={filters.endDate}
      class="mt-1 block w-full px-3 py-2 bg-gray-700 text-white rounded-md"
    />
  </div>

  <div class="mt-4 text-sm text-gray-400">
    <strong>Query Params:</strong>
    {filters.queryParams}
  </div>
</div>
