<script lang="ts">
	import type { PageData } from './$types';
	import Panel from '../../components/Panel.svelte';
	import NavBar from '../../components/NavBar.svelte';
	import Button from '../../components/Button.svelte';
	import TextInput from '../../components/TextInput.svelte';
	import { page } from '$app/stores';
	import { post } from '../../fetch';
	let data: PageData;
	const url = $page.url.pathname;
	let caption = 'Caption';
	data = { caption };

	async function handleInput() {
		try {
			const response: PageData = await post('/newpost', data);
			data = response;
		} catch (e) {
			throw new Error('Internal server error');
		}
	}
</script>

<NavBar current_page="/newpost"></NavBar>
<Panel title="New Post">
	<div class="form">
		<div class="caption">
			<TextInput placeholder="Insert a caption" on:input={handleInput}></TextInput>
		</div>
		<Button buttonText="Get the music on your platform!" link="/api/platform"></Button>
	</div>
</Panel>

<style>
	.form {
		display: flex;
		flex-direction: column;
	}

	.caption {
		width: 25rem;
		height: 2rem;
		margin-bottom: 0.25rem;
	}
</style>
