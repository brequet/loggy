<script lang="ts">
  import { Badge } from "$lib/components/ui/badge";
  import { Button } from "$lib/components/ui/button";
  import * as Sheet from "$lib/components/ui/sheet";
  import { fetchAppNames } from "$lib/services/logService";
  import { allLogLevels } from "$lib/types/LogEntry";
  import { Filter } from "lucide-svelte";
  import { onMount } from "svelte";
  import { filters } from "../stores/Filters.svelte";
  import ChoiceList from "./filter/ChoiceList.svelte";
  import DateTimeInput from "./filter/DateTimeInput.svelte";

  let allAppsName: string[] = [];

  onMount(async () => {
    allAppsName = await fetchAppNames();
  });

  function resetFilters() {
    filters.reset();
  }
</script>

<Sheet.Root>
  <Sheet.Trigger>
    <Button variant="outline" size="icon" class="relative">
      <Filter class="w-5 h-5" />
      {#if filters.activeFilterCount > 0}
        <Badge
          class="absolute -top-2 -right-2 px-2 min-w-[1.2rem] h-[1.2rem] text-[0.7rem] flex items-center justify-center"
        >
          {filters.activeFilterCount}
        </Badge>
      {/if}
    </Button>
  </Sheet.Trigger>
  <Sheet.Content shouldBlur={false}>
    <Sheet.Header>
      <Sheet.Title>Logs Filtering</Sheet.Title>
      <Sheet.Description>
        Apply filter to your logs to get a better overview of what's going on
      </Sheet.Description>
    </Sheet.Header>

    <div class="w-full pt-4 flex flex-col gap-4">
      <ChoiceList
        label="By application name"
        availableChoices={allAppsName}
        bind:selectedChoices={filters.appNames}
      ></ChoiceList>

      <ChoiceList
        label="By log level"
        availableChoices={allLogLevels}
        bind:selectedChoices={filters.levels}
      ></ChoiceList>

      <DateTimeInput label="Start Date" bind:date={filters.startDate} />

      <DateTimeInput label="End Date" bind:date={filters.endDate} />

      <Button class="mt-4 w-full px-4 py-2" onclick={resetFilters}>
        Reset Filters
      </Button>
    </div>
  </Sheet.Content>
</Sheet.Root>

<!-- 
<div class="bg-gray-800 p-4 rounded-lg">
  <h2 class="text-xl font-bold mb-4 text-white">Filters</h2>

  <Button
    onclick={resetFilters}
    variant="destructive"
    class="mt-4 w-full px-4 py-2 "
  >
    Reset Filters
  </Button>

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
      step="1"
      bind:value={filters.endDate}
      class="mt-1 block w-full px-3 py-2 bg-gray-700 text-white rounded-md"
    />
  </div>

  <div class="mt-4 text-sm text-gray-400">
    <strong>Query Params:</strong>
    {filters.queryParams}
  </div>
</div> -->
