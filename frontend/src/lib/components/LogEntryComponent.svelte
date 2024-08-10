<script lang="ts">
  import type { LogEntry } from "../types/LogEntry";
  import * as DropdownMenu from "$lib/components/ui/dropdown-menu";
  import { filters } from "$lib/stores/Filters.svelte";

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

  function setAsStartDate() {
    console.log("set as start date", entry.timestamp);
    // filters.startDate = entry.timestamp;
  }
</script>

<div
  class="grid grid-cols-[auto_auto_auto_1fr] gap-x-2 py-1 font-mono text-sm hover:bg-gray-200"
>
  <span class="text-neutral-400 whitespace-nowrap">{entry.app_name}</span>
  <DropdownMenu.Root>
    <DropdownMenu.Trigger>
      <span
        class="text-neutral-500 whitespace-nowrap hover:text-neutral-600 hover:font-bold hover:cursor-pointer"
        >{entry.timestamp}
      </span>
    </DropdownMenu.Trigger>
    <DropdownMenu.Content>
      <DropdownMenu.Group>
        <DropdownMenu.Item onclick={setAsStartDate}
          >Set as start date</DropdownMenu.Item
        >
        <DropdownMenu.Item>Set as end date</DropdownMenu.Item>
      </DropdownMenu.Group>
    </DropdownMenu.Content>
  </DropdownMenu.Root>
  <span class="whitespace-nowrap {getLevelColor(entry.level)}"
    >[{entry.level}]</span
  >
  <span class="text-black break-words">{entry.content}</span>
</div>
