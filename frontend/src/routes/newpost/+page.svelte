<script lang="ts">
	import Panel from '../../components/Panel.svelte';
	import NavBar from '../../components/NavBar.svelte';
	import Button from '../../components/Button.svelte';
	import TextInput from '../../components/TextInput.svelte';
	import { post, throwError } from '../../fetch';
	import { errorMessage } from '../../store';
	import ErrorPopup from '../../components/ErrorPopup.svelte';
	import { onMount } from 'svelte';
	let caption = '';
	let error = '';
	let postData = {};

	errorMessage.subscribe((value) => {
		error = value;
	});

	async function postPost(){
		try{
			const request = await post(`me/posts`, postData);
			return request;
		}catch(e){
			throwError('Failed to post item');
		}
	}

onMount(async () => {
	const params = new URLSearchParams(window.location.search);
	const library = params.get('library');
	const id = params.get('id');
	const type = params.get('type');
	postData = {library, id, type};
});

</script>

<NavBar current_page="/newpost"></NavBar>
<Panel title="New Post">
	<div class="form">
		<div class="caption">
			<TextInput placeholder="Insert a caption" bind:value={caption}></TextInput>
		</div>
		<Button buttonText="Upload post" on:click={async () => await postPost()}/>
		{#if error}
			<ErrorPopup message={error}></ErrorPopup>
		{/if}
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
