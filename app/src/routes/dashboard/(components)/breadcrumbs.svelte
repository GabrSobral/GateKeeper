<script lang="ts">
	  import * as Breadcrumb from "$lib/components/ui/breadcrumb/index.js";
      import * as Sidebar from "$lib/components/ui/sidebar/";

      let { items }: { items: { name: string; path: string }[] } = $props();
</script>

<div class="flex gap-4 items-center">
    <Sidebar.Trigger />

    <Breadcrumb.Root>
        <Breadcrumb.List>

            {#each items as item, i (item.name) }
                <Breadcrumb.Item>
                    <Breadcrumb.Link href={item.path}>{item.name}</Breadcrumb.Link>
                </Breadcrumb.Item>
        
                {#if i !== items.length -1}
                    <Breadcrumb.Separator />
                {/if}
            {/each}
        </Breadcrumb.List>
    </Breadcrumb.Root>
</div>
