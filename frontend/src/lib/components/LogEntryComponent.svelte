<script lang="ts">
  import type { LogEntry } from "../types/LogEntry";

  export let entry: LogEntry;
</script>

<div class="flex flex-row py-1 font-mono text-sm">
  <div class="flex">
    <span class="text-gray-600">{entry.app_name}</span>
    <span class="ml-2 text-gray-500">{entry.timestamp}</span>
    <span
      class="ml-2 {entry.level === 'ERROR'
        ? 'text-red-500'
        : entry.level === 'WARN'
          ? 'text-yellow-500'
          : 'text-green-500'}">[{entry.level}]</span
    >
  </div>
  <div class="ml-2 text-white">
    {entry.content}
  </div>
</div>
