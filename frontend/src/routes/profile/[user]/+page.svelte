<script lang="ts">
	import ErrorPopup from '../../../components/ErrorPopup.svelte';
	import Button from '../../../components/Button.svelte';
	import NavBar from '../../../components/NavBar.svelte';
	import { delete_, get, post, throwError } from '../../../fetch';
	import type PostModel from '../../../models/post';
	import { errorMessage } from '../../../store';

	import Post from '../../../components/Post.svelte';
	import type { PageData } from '../user/$types';
	import { goto } from '$app/navigation';
	import { onMount } from 'svelte';

	let followers: string[] = [];
	let following: string[] = [];
	let posts: PostModel[] = [];
	let follows = false;
	let loading = false;
	let error = '';
	$: text = following.includes(username) ? 'Unfollow' : 'Follow';
	export let data: PageData;
	$: username = data.user;
	let old = username;
	let me = '';
	errorMessage.subscribe((value) => {
		error = value;
	});

	onMount(async () => {
		errorMessage.set('');
	});
	$: if (username != old) {
		refresh();
	}

	async function refresh() {
		posts = [];
		await getInfo();
		await getFollowing();
		if (follows) await getPosts();
		if (username === me) await getPosts(true);
		old = username;
	}

	async function getPosts(me = false) {
		try {
			if (me) posts = await get<PostModel[]>(`/me/posts`);
			else posts = await get<PostModel[]>(`/posts?username=${username}`);
		} catch (e) {
			throwError('Error fetching user posts');
		}
	}

	async function getInfo() {
		try {
			const info = await get<any>(`/me/info`);
			me = info.username;
		} catch (e) {
			throwError('Error fetching user info');
		}
	}

	async function getFollowing() {
		try {
			following = await get<string[]>(`/me/following`);
			followers = await get<string[]>(`/me/followers`);
			follows = followers.includes(username);
		} catch (e) {
			throwError('Error fetching user lists');
		}
	}

	async function deleteFollow(username: string) {
		try {
			await delete_<string>(`/me/follow?username=${username}`);
			following = following.filter((user) => user !== username);
		} catch (e) {
			throwError('Error deleting user follow');
		}
	}

	async function postFollow(username: string) {
		try {
			await post<any, string>(`/me/follow?username=${username}`, {});
			following = following.concat(username);
		} catch (e) {
			throwError('Error posting user follow');
		}
	}

	// function onScroll(event: Event) {
	//	   const target = event.target as HTMLElement;
	//	   if (target.scrollHeight - target.scrollTop === target.clientHeight) {
	//		   loadMorePosts(name);
	//	   }
	// }
	// async function loadMorePosts(name: string) {
	//	   if (loading) return;
	//	   const morePosts = await _fetchUserData(name);
	//	   posts = [...posts, ...morePosts];
	// }

	let isClicked = false;
	const handleButtonClick = async () => {
		if (following.includes(username)) {
			await deleteFollow(username);
		} else {
			await postFollow(username);
		}
		isClicked = !isClicked;
	};

	let selectedList: string = '';
</script>

<NavBar />

{#if username === me}
	<div class="profile-container">
		<div class="user-container">
			<div class="followers">
				<Button
					action={() => (selectedList = selectedList === 'Followers' ? '' : 'Followers')}
					buttonText="Followers: {followers.length}"
				/>
				{#if selectedList == 'Followers'}
					<div class="followers-list-content">
						<h4>Users who follow you:</h4>
						{#if followers.length == 0}
							<p>No followers</p>
						{/if}
						{#if followers.length > 0}
							{#each followers as follower}
								<a href="/profile/{follower}">{follower}</a>
							{/each}
						{/if}
					</div>
				{/if}
			</div>
			<div class="following">
				<Button
					action={() => (selectedList = selectedList === 'Following' ? '' : 'Following')}
					buttonText="Following: {following.length}"
				/>

				{#if selectedList == 'Following'}
					<div class="following-list-content">
						<h4>Users you follow:</h4>
						{#if following.length == 0}
							<p class="following-list">Not following anyone</p>
						{/if}
						{#if following.length > 0}
							{#each following as follow}
								<a href="/profile/{follow}">{follow}</a>
							{/each}
						{/if}
					</div>
				{/if}
			</div>
			<Button buttonText="Library" action={() => goto('/profile/library')} />
			<Button buttonText="Liked" action={() => goto('/profile/liked')}></Button>
		</div>
	</div>
{/if}

<!-- profile -->
<div class="profile-container">
	<!-- personal feed -->
	<div class="feed-container">
		{#if me !== username}
			<div class="follow_button">
				<Button buttonText={text} action={handleButtonClick}></Button>
			</div>
		{/if}
		<div class="feed">
			{#if username !== me}
				<h2 class="username">{username}</h2>
			{/if}
			{#if !follows && username !== me}
				<p>Get followed by {username} to see their posts!</p>
			{:else if posts.length === 0}
				<p>{username} did not post yet!</p>
			{/if}
			{#each posts as post}
				<div class="post">
					<Post content={post} />
				</div>
			{/each}
			{#if loading}
				<p>Loading more posts...</p>
			{/if}
			{#if error}
				<ErrorPopup message={error}></ErrorPopup>
			{/if}
		</div>
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
		flex-wrap: wrap;
		flex-direction: row;
		justify-content: space-evenly;
		margin: 20px;
	}

	.username {
		padding: 0;
		margin-left: 0rem;
		margin-right: 10rem;
	}

	.profile-container {
		display: flex;
		flex-direction: column;
		justify-content: flex-start;
	}

	.username {
		grid-area: username;
	}

	.feed-container {
		display: flex;
		flex-direction: column;
		justify-content: center;
	}

	.feed {
		margin: 0 auto;
	}

	.following-list-content,
	.followers-list-content {
		position: fixed;
		z-index: 1000;
		border: black solid 1px;
		border-radius: 1rem;
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
		padding: 0.5rem;
		margin-top: 0.2rem;
		margin-left: 0rem;
		background: rgb(234, 185, 255);
		background: linear-gradient(
			0deg,
			rgba(234, 185, 255, 1) 24%,
			rgba(241, 207, 243, 1) 41%,
			rgba(248, 231, 231, 1) 100%
		);
	}

	.followers-list-content a,
	.following-list-content a {
		display: block;
		padding: 0.2rem;
		margin: 0.5rem;
		border-radius: 5px;
		text-decoration: none;
	}

	.followers-list-content a:hover,
	.following-list-content a:hover {
		text-transform: uppercase;
	}

	.follow_button{
		margin: 0 auto;
	}
</style>
