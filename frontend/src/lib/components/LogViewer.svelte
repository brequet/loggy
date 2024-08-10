<script lang="ts">
  import { fetchLogEntriesResult } from "../services/logService";
  import { filters } from "../stores/Filters.svelte";
  import type { LogEntry } from "../types/LogEntry";
  import LogEntryComponent from "./LogEntryComponent.svelte";

  const PAGE_SIZE = 50;
  let page = $state(1);

  let logs: LogEntry[] = $state([]);
  let totalAvailableEntries = $state(0);

  let loading = $state(false);
  let error = $state<string | null>("null");

  let filtersChangedDate = $derived.by(() => {
    // trick to force recomputation of the derived store, else variable are ignored => optimized by compiler
    const { appNames, levels, startDate, endDate } = filters;
    JSON.stringify({ appNames, levels, startDate, endDate });
    return Date.now();
  });

  let lastFiltersChanged: number | null = null;

  $effect(() => {
    if (lastFiltersChanged && filtersChangedDate === lastFiltersChanged) {
      return;
    }

    logs = [];
    page = 1;
    loadLogs();

    lastFiltersChanged = filtersChangedDate;
  });

  async function loadLogs(append = false) {
    if (loading) return;
    loading = true;
    error = null;
    try {
      const fullQueryParams = new URLSearchParams({
        appNames: filters.appNames.join(","),
        levels: filters.levels.join(","),
        startDate: filters.startDate?.toString() || "",
        endDate: filters.endDate?.toString() || "",
        page: page.toString(),
        pageSize: PAGE_SIZE.toString(),
      });
      const logEntriesResult = await fetchLogEntriesResult(
        fullQueryParams.toString(),
      );
      if (append) {
        if (logEntriesResult.entries.length > 0) {
          logs = [...logs, ...logEntriesResult.entries];
          page++;
        }
      } else {
        logs = logEntriesResult.entries;
      }
      totalAvailableEntries = logEntriesResult.total_count;
    } catch (err) {
      error = "Failed to load logs";
    } finally {
      loading = false;
    }
  }

  function handleScroll(event: Event) {
    const target = event.target as HTMLDivElement;
    if (target.scrollHeight - target.scrollTop <= target.clientHeight + 100) {
      loadLogs(true);
    }
  }
</script>

<div onscroll={handleScroll} class=" bg-gray-100 p-4 flex-1 overflow-y-auto">
  {#each logs as log}
    <LogEntryComponent entry={log} />
  {/each}
  {#if loading}
    <div class="text-center text-gray-500">Loading...</div>
  {/if}
  {#if error}
    <div class="text-center text-red-500">{error}</div>
  {/if}
</div>

<p class="text-sm text-neutral-400">
  {totalAvailableEntries} entries available
</p>
