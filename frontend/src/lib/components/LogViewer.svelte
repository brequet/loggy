<script lang="ts">
    import { mount, onMount } from "svelte";
    import type { LogEntry } from "../types/LogEntry";
    import { fetchLogs } from "../services/logService";
    import LogEntryComponent from "./LogEntryComponent.svelte";
    import { filters } from "../stores/Filters.svelte";
    import FilterPanel from "./FilterPanel.svelte";

    const PAGE_SIZE = 50;
    let page = $state(1);

    let logs: LogEntry[] = $state([]);

    let loading = $state(false);
    let error = $state<string | null>("null");

    let filtersChangedDate = $derived.by(() => {
        const { appNames, levels, startDate, endDate } = filters;

        return Date.now();
    });

    let lastFiltersChanged: number | null = null;

    $effect(() => {
        console.log("loading logs", page);
        if (lastFiltersChanged && filtersChangedDate === lastFiltersChanged) {
            return;
        }

        window.scrollTo({
            top: 0,
            behavior: "smooth",
        });

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
            const newLogs = await fetchLogs(fullQueryParams.toString());
            logs = append ? [...logs, ...newLogs] : newLogs;
        } catch (err) {
            error = "Failed to load logs";
        } finally {
            loading = false;
        }
    }

    function loadMoreLogs() {
        page++;
        loadLogs(true);
    }

    function handleScroll(event: Event) {
        const target = event.target as HTMLDivElement;
        if (
            target.scrollHeight - target.scrollTop <=
            target.clientHeight + 100
        ) {
            loadMoreLogs();
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
