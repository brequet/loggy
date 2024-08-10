<script lang="ts">
  import * as DropdownMenu from "$lib/components/ui/dropdown-menu";
  import { filters } from "$lib/stores/Filters.svelte";
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

  function setAsStartDate() {
    const formattedDate = entry.timestamp.split(".")[0];
    filters.startDate = formattedDate;
  }

  function setAsEndDate() {
    const formattedDate = entry.timestamp.split(".")[0];
    filters.endDate = formattedDate;
  }
</script>

<tr class="hover:bg-gray-200">
  <td class="py-1 pr-2 align-top text-neutral-400 whitespace-nowrap"
    >{entry.app_name}</td
  >

  <DropdownMenu.Root>
    <DropdownMenu.Trigger>
      <td
        class="my-auto py-1 pr-2 text-neutral-500 whitespace-nowrap hover:text-neutral-600 hover:font-bold hover:cursor-pointer"
      >
        {entry.timestamp}
      </td>
    </DropdownMenu.Trigger>
    <DropdownMenu.Content>
      <DropdownMenu.Group>
        <DropdownMenu.Label>Set filters</DropdownMenu.Label>
        <DropdownMenu.Separator />
        <DropdownMenu.Item onclick={setAsStartDate}
          >Set as start date</DropdownMenu.Item
        >
        <DropdownMenu.Item onclick={setAsEndDate}
          >Set as end date</DropdownMenu.Item
        >
      </DropdownMenu.Group>
    </DropdownMenu.Content>
  </DropdownMenu.Root>
  <td
    class="py-1 pr-2 align-top whitespace-nowrap {getLevelColor(entry.level)}"
  >
    [{entry.level}]
  </td>
  <td class="py-1 align-top text-black break-words">{entry.content}</td>
</tr>
