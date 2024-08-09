<script lang="ts">
    import { Input } from "$lib/components/ui/input";
    import { Button } from "$lib/components/ui/button";
    import { Label } from "$lib/components/ui/label";

    let {
        label,
        input = $bindable(),
        removeAction,
        addAction,
        itemList,
    }: {
        label: string;
        input: string;
        itemList: string[];
        addAction: () => void;
        removeAction: (name: string) => void;
    } = $props();
</script>

<div>
    <div class="flex w-full max-w-sm flex-col gap-1.5">
        <Label for="email">{label}</Label>
        <div class="flex w-full max-w-sm items-center space-x-2">
            <Input
                type="text"
                placeholder="Input value"
                bind:value={input}
                onkeyup={(event) => event.key === "Enter" && addAction()}
            />
            <Button
                onclick={addAction}
                class="px-4 py-2 bg-blue-600 text-white rounded-r-md hover:bg-blue-700"
            >
                Add
            </Button>
        </div>
    </div>
    <div class="mt-2 flex flex-wrap gap-2">
        {#each itemList as item}
            <span
                class="bg-blue-500 text-white px-2 py-1 rounded-full text-sm flex items-center"
            >
                {item}
                <button
                    onclick={() => removeAction(item)}
                    class="ml-2 text-xs bg-blue-600 rounded-full w-4 h-4 flex items-center justify-center hover:bg-blue-700"
                >
                    ×
                </button>
            </span>
        {/each}
    </div>
</div>
