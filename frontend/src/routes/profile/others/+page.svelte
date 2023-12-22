<script lang="ts">
	import Button from './../../../components/Button.svelte';
	import Panel from '../../../components/Panel.svelte';
	import NavBar from './../../../components/NavBar.svelte';
	import { get, throwError } from '../../../fetch';
	import type PostModel from '../../../models/post';
	import Post from '../../../components/Post.svelte';
	import { onMount } from 'svelte';
	import Followers from '../../../components/Followers.svelte';
	import Following from '../../../components/Following.svelte';
	let posts: PostModel[] = [];
	let followers: string[] = [];
	let following: string []= [];
    let follows: boolean = true;
	async function getData() {
		try {
			const response: string = await get('/api/v1/{user_id}');
			posts = JSON.parse(response);
		} catch (e) {
			throwError('Internal server error');
		}
	}
	async function getFollowers() {
		try {
			const response: string = await get('/api/v1/{user_id}/followers');
			followers = JSON.parse(response);
		} catch (e) {
			throwError('Internal server error');
		}
	}
	async function getFollowing() {
		try {
			const response: string = await get('/api/v1/{user_id}/following');
			following = JSON.parse(response);
		} catch (e) {
			throwError('Internal server error');
		}
	}
	onMount(getData);
	onMount(getFollowers);
	onMount(getFollowing);

	//https://svelte.dev/repl/4c5dfd34cc634774bd242725f0fc2dab?version=3.46.4 (dropdown handling)
	let isDropdownOpen = false;
	const handleDropdownClick = () => {
		isDropdownOpen = !isDropdownOpen;
	};

	const handleDropdownFocusLoss = (event: FocusEvent) => {
		const { currentTarget, relatedTarget } = event; // relatedTarget: HTMLElement;
		// use "focusout" event to ensure that we can close the dropdown when clicking outside or when we leave the dropdown with the button
		if (relatedTarget instanceof HTMLElement && (currentTarget as Node).contains(relatedTarget))
			return; // check if the new focus target doesn't present in the dropdown tree
		isDropdownOpen = false;
	};

    let isClicked = false;
    const handleButtonClick = () => {
        isClicked = !isClicked;
        follows = !follows;

    }
</script>

<!-- navbar -->
<div class="nav">
	<NavBar current_page={'/{user_id}'}></NavBar>
</div>
<!-- profile -->
<div class="profile-container">
	<div class="user-container">
		<!-- username + followers/following + link to saved and library -->
		<div class="user">
			<h2 class="username">Username</h2>
			<div class="following" on:focusout={handleDropdownFocusLoss}>
				<Button buttonText="Following" on:click={handleDropdownClick}
				></Button><!-- generate a dropdown with all the following-->
				<div class="followingDropdown" style:visibility={isDropdownOpen ? 'visible' : 'hidden'}>
					<Following {following} />
				</div>
			</div>
			<div class="followers" on:focusout={handleDropdownFocusLoss}>
				<Button buttonText="Followers" on:click={handleDropdownClick}></Button>
				<!-- generate a dropdown with all the followers-->
				<div class="followersDropdown" style:visibility={isDropdownOpen ? 'visible' : 'hidden'}>
					<Followers {followers} />
				</div>
			</div>
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
	<div class="feed-container">
		<Panel title="Your feed">
			<div class="feed">
				{#each posts as post, i}
					<div class="post" id={'post' + (i + 1)}>
						<Post caption={post.caption} likes={post.likes}></Post>
					</div>
				{/each}
			</div>
		</Panel>
	</div>
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
	.bio {
		margin-top: 2rem;
		grid-area: bio;
		width: 100rem;
	}
	.feed {
		display: flex;
	}
</style>
