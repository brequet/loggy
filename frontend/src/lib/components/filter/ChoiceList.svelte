<script lang="ts">
    import { Checkbox } from "$lib/components/ui/checkbox";
    import { Label } from "$lib/components/ui/label";

    let {
        label,
        availableChoices,
        selectedChoices = $bindable([]),
    }: {
        label: string;
        availableChoices: string[];
        selectedChoices: string[];
    } = $props();

    function updateSelection(choice: string, isChecked: boolean) {
        if (isChecked) {
            selectedChoices = [...selectedChoices, choice];
        } else {
            selectedChoices = selectedChoices.filter((c) => c !== choice);
        }
        console.log(selectedChoices);
    }
</script>

<div>
    <div class="flex w-full max-w-sm flex-col gap-1.5">
        <Label>{label}</Label>
        {#each availableChoices as choice}
            <div class="flex items-center space-x-2">
                <Checkbox
                    id={choice}
                    checked={selectedChoices.includes(choice)}
                    onCheckedChange={(v) => {
                        updateSelection(choice, v);
                    }}
                />
                <Label for={choice}>{choice}</Label>
            </div>
        {/each}
    </div>
</div>
