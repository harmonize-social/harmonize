<script lang="ts">
    import ErrorPopup from '../../../components/ErrorPopup.svelte';
    import Button from '../../../components/Button.svelte';
    import Panel from '../../../components/Panel.svelte';
    import NavBar from '../../../components/NavBar.svelte';
    import { delete_, get, post, throwError } from '../../../fetch';
    import type PostModel from '../../../models/post';

    import Post from '../../../components/Post.svelte';
    import { errorMessage } from '../../../store';
    import type { PageData } from '../user/$types';
    import { onMount } from 'svelte';

    let followed: string[] = [];
    let following: string[] = [];
    let posts: PostModel[] = [];
    let follows: boolean;
    let loading = false;
    let error = '';
    $: text = follows ? 'Unfollow' : 'Follow';
    export let data: PageData;
    let username: any = data.user;

    errorMessage.subscribe((value) => {
        error = value;
    });

    async function getPosts() {
        try {
            posts = await get<PostModel[]>(`/posts?username=${username}`);
        } catch (error) {
            throwError('Error fetching user posts');
        }
    }

    async function getFollowing() {
        try {
            followed = await get<string[]>(`/me/following`);
            following = await get<string[]>(`/me/followers`);
            follows = followed.includes(username);
        } catch (e) {
            throwError('Error fetching user posts');
        }
    }

    async function deleteFollow(username: string) {
        try {
            await delete_<string>(`/me/follow?username=${username}`);
            following = following.filter((user) => user !== username);
            follows = false;
        } catch (e) {
            throwError('Error deleting follow');
        }
    }

    async function postFollow(username: string) {
        try {
            await post<any, string>(`/me/follow?username=${username}`, {});
            following = following.filter((user) => user !== username);
            follows = true;
        } catch (e) {
            throwError('Error posting follow');
        }
    }

    // function onScroll(event: Event) {
    //     const target = event.target as HTMLElement;
    //     if (target.scrollHeight - target.scrollTop === target.clientHeight) {
    //         loadMorePosts(name);
    //     }
    // }
    // async function loadMorePosts(name: string) {
    //     if (loading) return;
    //     const morePosts = await _fetchUserData(name);
    //     posts = [...posts, ...morePosts];
    // }

    let isClicked = false;
    const handleButtonClick = async () => {
        if (follows) {
            await deleteFollow(username);
        } else {
            await postFollow(username);
        }
        isClicked = !isClicked;
    };
    onMount(async () => {
        await getFollowing();
        if (follows) {
            await getPosts();
        }
    });
</script>

<!-- navbar -->
<div class="nav">
    <NavBar current_page={`/profile/${username}`}></NavBar>
</div>
<!-- profile -->
<div class="profile-container">
    <h2 class="username">{username}</h2>

    <!-- personal feed -->
    <div class="feed-container">
        <div class="follow_button">
            <Button buttonText={text} action={handleButtonClick}></Button>
        </div>
        <Panel title={`${username}'s posts`}>
            <div class="feed">
                {#if posts.length === 0}
                    <p>{username} did not post yet!</p>
                {/if}
                {#each posts as post}
                    <div class="post">
                        <Post
                            content={post}
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

    .username {
        grid-area: username;
    }

    .feed {
        display: flex;
    }
</style>
