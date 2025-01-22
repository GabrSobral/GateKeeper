<script lang="ts">
	import { on } from "svelte/events";
	import Badge from "../badge/badge.svelte";
    import * as Tooltip from "$lib/components/ui/tooltip";

    type Props = {
        items: string[];
        placeholder?: string;
    }
    let { items, placeholder }: Props = $props();

    let value = $state("")
    let errorMessage = $state("");

    let inputRef: HTMLInputElement | null = null;

    $effect(() => {
        if(inputRef) {
            on(inputRef, "keydown", (e) => {
                if (e.key === "Enter") {
                    e.preventDefault();
                    
                    if(value === "") {
                        errorMessage = "Item is empty"

                        setTimeout(() => (errorMessage = ""), 2000)
                        return;
                    }

                    if (!items.includes(value)) {
                        items = [...items, value]
                        value = ""
                    } else {
                        errorMessage = "Item already exists"
                        setTimeout(() => (errorMessage = ""), 2000)
                    }
                }
            })
        }
    }) 
</script>

<div class="flex flex-col gap-2">
    <!-- svelte-ignore a11y_unknown_aria_attribute -->
    <!-- svelte-ignore a11y_click_events_have_key_events -->
    <!-- svelte-ignore a11y_no_static_element_interactions -->
    <div 
        onclick={() => inputRef?.focus()}
        class="hover:cursor-text flex-wrap gap-1 border-input bg-background ring-offset-background placeholder:text-muted-foreground focus-visible:ring-ring flex min-h-10 w-full rounded-md border px-3 py-2 text-base file:border-0 file:bg-transparent file:text-sm file:font-medium focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50 md:text-sm"
    >
        {#each items as item (item)}
        <Tooltip.Root>
            <Tooltip.Trigger>
                <Badge variant="default" class="whitespace-normal" onclick={(e) => {
                    e.stopPropagation();
                    items = items.filter(i => i !== item)
                }}>
                    {item}
                </Badge>  
            </Tooltip.Trigger>
    
            <Tooltip.Content>
                <p>Clito to delete</p>
            </Tooltip.Content>
        </Tooltip.Root>
              
        {/each}
        
        <input 
            bind:this={inputRef}
            placeholder={placeholder || "Type and press enter to add"}
            type='text'
            bind:value={value}
            class="min-w-[175px] outline-none"
        />
    
    </div>
    
    {#if errorMessage !== ""}
        <p class="text-red-500 text-sm">{errorMessage}</p>
    {/if}
</div>