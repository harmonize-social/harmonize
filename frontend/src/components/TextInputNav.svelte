<script lang="ts">
	import { onDestroy } from 'svelte';
	import { get, throwError } from '../fetch';

	export let placeholder = 'Search';
	let input: string = '';
	let list: string[] = [];

	function handleTextInput(event: InputEvent) {
		if (event.inputType === 'insertText' && event.data === input)  {
			handleInput();
		}
	}

	async function handleInput() {
		try {
			const response = (await get('/search?username=' + input)) as any;
			if (response.error) throwError(response.error);
			list = response;
		} catch (e) {
			console.log(e);
			throwError('Internal server error');
		}
		console.log(list);
	}

	let isDropdownOpen = true;
	const handleDropdownFocusLoss = (event: FocusEvent) => {
		const { currentTarget, relatedTarget } = event;
		if (relatedTarget instanceof HTMLElement && (currentTarget as Node).contains(relatedTarget))
			return;
		isDropdownOpen = false;
	};

	const timeout = setTimeout(() => {
		isDropdownOpen = false;
	}, 3000);

	onDestroy(() => {
		clearTimeout(timeout);
	});
  //TODO: fix dropdown list and add search functionality
</script>

<div class="search-container">
	<input
		type="text"
		class="textInputNav"
		{placeholder}
on:input={handleTextInput}
		on:focusout={handleDropdownFocusLoss}
		bind:value={input}
	/>

	{#each list as item}
		<p class="list">
			<a href="/profile/{item}">{item}</a>
		</p>
	{/each}
</div>

<style>
	.search-container {
		position: relative;
		display: inline-block;
	}

	.textInputNav {
		grid-area: searchbox;
		margin-top: 6px;
		margin-right: 8px;
		width: 200px;
		height: 55px;
		border-radius: 50px;
		border: 1px solid black;
		color: rebeccapurple;
	}

	.list {
		position: absolute;
		top: 100%;
		left: 0;
		width: 200px;
		border: 1px solid #ddd;
		border-radius: 4px;
		box-shadow: 0 2px 4px rgba(0, 0, 0, 0.2);
		background-color: white;
		z-index: 1000;
	}

	.list a {
		margin: 10px;
		cursor: pointer;
		text-decoration: none;
	}

</style>
