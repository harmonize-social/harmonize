<script lang="ts">
	import Panel from '../../components/Panel.svelte';
	import NavBar from '../../components/NavBar.svelte';
	import { post, throwError, get } from '../../fetch';
	import { errorMessage } from '../../store';
	import ErrorPopup from '../../components/ErrorPopup.svelte';
	import { onMount } from 'svelte';
	let caption = '';
	let error = '';
	let postData: {};
	let library;
	let id;
	let type;

	errorMessage.subscribe((value) => {
		error = value;
	});

	async function getContent(data: any){
		try{
			const response: any = await get(`me/library/${data.library}/${data.type}?id=${data.id}`);
			return response;
		}catch(e){
			throwError('Failed to get content');
		}
	}
	async function postPost(){
		try{
			const request = await post(`me/posts`, {postData, caption} );
			return request;
		}catch(e){
			throwError('Failed to post item');
		}
	}

	async function handleSubmit(event: Event){
		event.preventDefault();
		const response = await postPost();
		if(response){
			window.location.href = '/profile/library';
		}else{
			throwError('Failed to post item');
			window.location.href = '/profile/library';
		}
	}

onMount(async () => {
	const params = new URLSearchParams(window.location.search);
	 library = params.get('library');
	 id = params.get('id');
	 type = params.get('type');
	postData = {library, type, id};
	await getContent(postData);
});

</script>

<NavBar current_page="/newpost"></NavBar>
<Panel title="New Post">
	<div class="form">
		<form on:submit={handleSubmit}>
			<input type="text" placeholder="Insert caption here"  bind:value={caption} class="caption">
			<div class="submit">
				<input type="submit" value="Upload your post"/>
			</div>
		</form>
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
		height: 5rem;
		margin-bottom: 0.25rem;
	}

	.submit{
		margin-top: 3rem;
	}
</style>
