<script lang="ts">
	import ErrorPopup from './../../../components/ErrorPopup.svelte';
	import Button from './../../../components/Button.svelte';
	import Panel from '../../../components/Panel.svelte';
	import NavBar from './../../../components/NavBar.svelte';
	import { get, throwError } from '../../../fetch';
	import type PostModel from '../../../models/post';
	import Post from '../../../components/Post.svelte';
	import { onMount } from 'svelte';
	import { errorMessage } from '../../../store';

	let posts: PostModel[] = [];
	let username: string;
	let follows: boolean = true;
	let loading = false;
	let error = '';

	errorMessage.subscribe((value) => {
		error = value;
	});

	async function fetchPosts(): Promise<PostModel[]> {
		try {
			const response: PostModel[] = await get<PostModel[]>(`/posts?username=${username}`);
			posts = response;
			return posts;
		} catch (e) {
			throwError('Error fetching posts');
			return [];
		} finally {
			loading = false;
		}
	}
	onMount(() => {
		fetchPosts().then((fetchedPosts) => {
			posts = fetchedPosts;
		});
	});

	function onScroll(event: Event) {
		const target = event.target as HTMLElement;
		if (target.scrollHeight - target.scrollTop === target.clientHeight) {
			loadMorePosts();
		}
	}
	async function loadMorePosts() {
		if (loading) return;
		const morePosts = await fetchPosts();
		posts = [...posts, ...morePosts];
	}

	let isClicked = false;
	const handleButtonClick = () => {
		isClicked = !isClicked;
		follows = !follows;
	};
</script>

<!-- navbar -->
<div class="nav">
	<NavBar current_page={`/user/${username}`}></NavBar>
</div>
<!-- profile -->
<div class="profile-container">
	<div class="user-container">
		<div class="user">
			<h2 class="username">{username}</h2>
			<div class="bio">
				<p>
					Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt
					ut labore et dolore magna aliqua.
				</p>
			</div>
			<div class="follow_button">
				{#if isClicked && follows}
					<Button buttonText="Unfollow" on:click={handleButtonClick}></Button>
				{:else}
					<Button buttonText="Follow" on:click={handleButtonClick}></Button>
				{/if}
			</div>
		</div>
	</div>
	<!-- personal feed -->
	<div class="feed-container" on:scroll={onScroll}>
		<Panel title={`${username}'s posts`}>
			<div class="feed">
				{#each posts as post}
					<div class="post">
						<Post
							content={post.content}
							caption={post.caption}
							likes={post.likeCount}
							id={post.id}
							typez={post.type}
						/>
					</div>
				{/each}
				{#if loading}
					<p>Loading more posts...</p>
				{/if}
				{#if error}
					<ErrorPopup message={error}></ErrorPopup>
				{/if}
			</div>
		</Panel>
	</div>
</div>

<style>
	.profile-container {
		display: flex;
		flex-direction: column;
		justify-content: flex-start;
	}
	.user-container {
		display: flex;
		flex-wrap: wrap-reverse;
		flex-direction: row-reverse;
		justify-content: center;
		height: 10rem;
	}
	.user {
		display: grid;
		grid-template-rows: repeat(3, 2rem);
		grid-template-columns: repeat(5, 20rem);
		grid-template-areas:
			'username followers following library liked'
			'bio bio bio bio bio';
		width: 100rem;
		gap: 2rem 1rem;
		padding-top: 0;
	}
	.username {
		grid-area: username;
	}
	.bio {
		margin-top: 2rem;
		grid-area: bio;
		width: 100rem;
	}
	.feed {
		display: flex;
	}
</style>
