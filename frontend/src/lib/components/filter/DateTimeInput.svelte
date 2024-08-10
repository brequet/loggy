<script lang="ts">
  import { onMount } from "svelte";

  let {
    label,
    date = $bindable(),
  }: {
    label: string;
    date: string | null;
  } = $props();

  let inputElement: HTMLInputElement;
  let isFocused = false;

  onMount(() => {
    if (inputElement) {
      document.addEventListener("paste", handlePaste);
    }

    return () => {
      document.removeEventListener("paste", handlePaste);
    };
  });

  function handlePaste(event: ClipboardEvent) {
    if (!isFocused) return;

    event.preventDefault();
    const pastedText = event.clipboardData?.getData("text");

    if (pastedText) {
      const isValidFormat = /^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}$/.test(
        pastedText,
      );

      if (isValidFormat) {
        date = pastedText;
      } else {
        // TODO: show error on input
        console.error("Invalid datetime format");
      }
    }
  }

  function handleFocus() {
    isFocused = true;
  }

  function handleBlur() {
    isFocused = false;
  }
</script>

<div class="mb-4">
  <label class="block text-sm font-medium">{label}</label>
  <input
    type="datetime-local"
    step="1"
    bind:this={inputElement}
    bind:value={date}
    onfocus={handleFocus}
    onblur={handleBlur}
    class="mt-1 block w-full px-3 py-2 rounded-md"
  />
</div>
