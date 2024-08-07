<script lang="ts">
    import { onMount } from "svelte";
    import type { LogEntry } from "../types/LogEntry";
    import { fetchLogs } from "../services/logService";
    import LogEntryComponent from "./LogEntryComponent.svelte";

    const PAGE_SIZE = 50;

    let logs: LogEntry[] = $state([]);
    let page = $state(1);
    let loading = $state(false);
    let error = $state<string | null>("null");
    let containerRef: HTMLDivElement;

    async function loadMoreLogs() {
        if (loading) return;
        loading = true;
        error = null;
        try {
            const newLogs = await fetchLogs(page, PAGE_SIZE);
            console.log("new logs:", newLogs);
            logs = [...logs, ...newLogs];
            console.log("logs:", logs);
            page++;
        } catch (err) {
            error = "Failed to load logs";
        } finally {
            loading = false;
        }
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

<div
    bind:this={containerRef}
    onscroll={handleScroll}
    class="h-screen overflow-y-auto overflow-x-auto bg-gray-900 p-4"
>
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
