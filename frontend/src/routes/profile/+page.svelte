<script lang="ts">
	import Panel from '../../components/Panel.svelte';
	import NavBar from '../../components/NavBar.svelte';
	import Button from '../../components/Button.svelte';
	import { get, throwError } from '../../fetch';
	import type PostModel from '../../models/post';
	import Post from '../../components/Post.svelte';
	import { onMount } from 'svelte';
	import Followers from '../../components/Followers.svelte';
	import Following from '../../components/Following.svelte';

	let posts: PostModel[] = [];
	let followers: string[] = [];
	let following: string[] = [];

	async function getData() {
		try {
			const response: string = await get('/me');
			posts = JSON.parse(response);
		} catch (e) {
			throwError('Internal server error');
		}
	}

	async function getFollowers() {
		try {
			followers = await get('/me/followers');
			console.log(followers);
		} catch (e) {
			throwError('Internal server error');
		}
	}

	async function getFollowing() {
		try {
			following = await get('/me/following');
		} catch (e) {
			throwError('Internal server error');
		}
	}
</script>

<div class="nav">
	<NavBar current_page="/profile"></NavBar>
</div>

<div class="profile-container">
	<div class="user-container">
		<div class="user">
			<h2 class="username">Username</h2>
			<div class="following" on:click={getFollowing}>
				<Button buttonText="Following"></Button>
				{#each following as item}
					<p>
						<a href="/user/{item}">{item}</a>
					</p>
				{/each}
			</div>
			<div class="followers" on:click={getFollowers}>
				<Button buttonText="Followers"></Button>
			</div>
			{#each followers as item}
					<p>
						<a href="/user/{item}">{item}</a>
					</p>
				{/each}

			<div class="library">
				<Button buttonText="Library" link="/profile/library"></Button>
			</div>
			<div class="liked">
				<Button buttonText="Liked" link="/profile/liked"></Button>
			</div>
			<div class="bio">
				<p>
					Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt
					ut labore et dolore magna aliqua.
				</p>
			</div>
		</div>
	</div>
</div>

<!-- personal feed -->
<div class="feed-container">
	<Panel title="Your feed">
		<div class="feed">
			{#each posts as post, i}
				<div class="post" id={'post' + (i + 1)}>
					<Post
						content={post.content}
						caption={post.caption}
						likes={post.likeCount}
						id={post.id}
						type={post.type}
					/>
				</div>
			{/each}
		</div>
	</Panel>
</div>

<!-- TODO: Check dropdowns -->

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
	.following {
		grid-area: following;
		margin-top: 0.5rem;
	}
	.followers {
		grid-area: followers;
		margin-top: 0.5rem;
	}
	.library {
		grid-area: library;
		margin-top: 1rem;
	}
	.liked {
		grid-area: liked;
		margin-top: 1rem;
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
