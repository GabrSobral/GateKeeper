<script lang="ts">
	import * as Tooltip from '$lib/components/ui/tooltip';
	import Label from '$lib/components/ui/label/label.svelte';
	import Checkbox from '$lib/components/ui/checkbox/checkbox.svelte';

	type Props = {  };

	let {  }: Props = $props();
</script>

<div class="flex flex-col gap-1">
	<span class="text-sm font-medium"> Multi Factor Authentication </span>

	<span class="text-muted-foreground my-2 text-sm">
		Choose the methods that will be used for multi factor authentication.
	</span>

	<div class="flex items-center space-x-2">
		<Checkbox id="e-mail-mfa-1" />

		<Tooltip.Provider>
			<Tooltip.Root>
				<Tooltip.Trigger>
					<Label
						for="e-mail-mfa-1"
						class="text-sm leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
					>
						E-mail
					</Label>
				</Tooltip.Trigger>

				<Tooltip.Content>
					Send a verification code to the user's email address.
				</Tooltip.Content>
			</Tooltip.Root>
		</Tooltip.Provider>
	</div>

	<div class="flex items-center space-x-2">
		<Checkbox id="auth-app-mfa-2" />

		<Tooltip.Provider>
			<Tooltip.Root>
				<Tooltip.Trigger>
					<Label
						for="auth-app-mfa-2"
						class="text-sm leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
					>
						Authenticator App (Microsoft, Google, etc)
					</Label>
				</Tooltip.Trigger>

				<Tooltip.Content>
					Use an authenticator app to generate a verification code.
				</Tooltip.Content>
			</Tooltip.Root>
		</Tooltip.Provider>
	</div>
</div>
