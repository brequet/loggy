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

    let fullQueryParams = $derived(
        new URLSearchParams({
            ...Object.fromEntries(filters.queryParams),
            page: page.toString(),
            pageSize: PAGE_SIZE.toString(),
        }),
    );

    let previousQueryParams = "";

    $effect(() => {
        if (
            previousQueryParams !== "" &&
            fullQueryParams.toString() !== previousQueryParams
        ) {
            const fullQueryPage = fullQueryParams.get("page");
            const previousQueryPage = previousQueryParams
                .split("page=")[1]
                .split("&")[0];
            if (fullQueryPage === previousQueryPage) {
                loadLogs();
            } else {
                loadMoreLogs();
            }
            previousQueryParams = fullQueryParams.toString();
        }

        if (previousQueryParams === "") {
            previousQueryParams = fullQueryParams.toString();
        }
    });

    async function loadLogs(append = false) {
        if (loading) return;
        loading = true;
        error = null;
        try {
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

    onMount(() => {
        loadMoreLogs();
    });

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
