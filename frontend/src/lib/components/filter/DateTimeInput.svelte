<script lang="ts">
    let {
        label,
        date = $bindable(),
    }: {
        label: string;
        date: Date | null;
    } = $props();
</script>

<div class="mb-4">
    <label class="block text-sm font-medium">{label}</label>
    <input
        type="datetime-local"
        step="1"
        bind:value={date}
        class="mt-1 block w-full px-3 py-2 rounded-md"
    />
</div>
