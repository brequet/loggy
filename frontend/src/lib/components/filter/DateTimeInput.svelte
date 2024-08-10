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
    if (!isFocused) return; // Only handle paste if this input is focused

    console.log("handle paste");
    event.preventDefault();
    const pastedText = event.clipboardData?.getData("text");
    console.log("pasted text", pastedText);

    if (pastedText) {
      const isValidFormat = /^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}$/.test(
        pastedText
      );

      if (isValidFormat) {
        console.log(
          `valid format: ${pastedText}, formatDate: ${formatDate(new Date(pastedText))}`
        );
        date = pastedText;
      } else {
        // If the pasted text is not in the correct format, you can show an error or handle it as needed
        console.error("Invalid datetime format");
      }
    }
  }

  function formatDate(d: Date | null): string {
    if (!d) return "";
    return d.toISOString().slice(0, 19);
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
    on:focus={handleFocus}
    on:blur={handleBlur}
    class="mt-1 block w-full px-3 py-2 rounded-md"
  />
</div>
