<script lang="ts">
  import type { LogEntry } from "../types/LogEntry";

  export let entry: LogEntry;

  function getLevelColor(level: string): string {
    switch (level.toLowerCase()) {
      case "error":
        return "text-red-500";
      case "warn":
        return "text-yellow-500";
      case "info":
        return "text-blue-500";
      case "debug":
        return "text-green-500";
      default:
        return "text-gray-500";
    }
  }
</script>

<div
  class="grid grid-cols-[auto_auto_auto_1fr] gap-x-2 py-1 font-mono text-sm hover:bg-gray-200"
>
  <span class="text-neutral-400 whitespace-nowrap">{entry.app_name}</span>
  <span class="text-neutral-500 whitespace-nowrap">{entry.timestamp}</span>
  <span class="whitespace-nowrap {getLevelColor(entry.level)}"
    >[{entry.level}]</span
  >
  <span class="text-black break-words">{entry.content}</span>
</div>
