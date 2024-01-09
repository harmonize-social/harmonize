<script lang="ts">
    import Panel from '../../components/Panel.svelte';
    import NavBar from '../../components/NavBar.svelte';
    import Button from './../../components/Button.svelte';
    import { get, throwError } from '../../fetch';
    import type PostModel from '../../models/post';
    import Post from '../../components/Post.svelte';
    import { errorMessage } from '../../store';
    import { onMount } from 'svelte';
    import ErrorPopup from '../../components/ErrorPopup.svelte';

    let posts: PostModel[] = [];
    let followers: string[] = [];
    let following: string[] = [];
    let selectedList: string = '';
    let loading = false;
    let error = '';
    errorMessage.subscribe((value) => {
        error = value;
    });

    async function fetchPosts(): Promise<PostModel[]> {
        try {
            const response: PostModel[] = await get<PostModel[]>('/me/posts');
            return response;
        } catch (e) {
            throwError('Error fetching posts');
            return [];
        } finally {
            loading = false;
        }
    }

    async function getFollowers(): Promise<string[]> {
        try {
            const response: string[] = await get<string[]>('/me/followers');
            return response;
        } catch (e) {
            throwError('Error fetching followers');
            return [];
        }
    }

    async function getFollowing(): Promise<any> {
        try {
            const response: string[] = await get<any>('/me/following');
            return response;
        } catch (e) {
            throwError('Error fetching following');
            return [];
        }
    }

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

    onMount(async () => {
        try {
            fetchPosts().then((fetchedPosts) => {
                posts = fetchedPosts;
            });
            followers = await getFollowers();
            following = await getFollowing();
        } catch (e) {
            throwError('Could not get profile information');
        }
    });
</script>

<div class="nav">
    <NavBar current_page="/profile"></NavBar>

    <div class="buttons">
        <div class="library">
            <Button buttonText="Library" link="/profile/library"></Button>
        </div>
        <div class="liked">
            <Button buttonText="Liked" link="/profile/liked"></Button>
        </div>
    </div>
</div>

<div class="profile-container">
    <div class="user-container">
        <h2 class="username">My Profile</h2>
        <div class="followers-list-container">
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
        <div class="following-list-container">
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
    </div>
</div>

<!-- personal feed -->
<div class="feed-container">
    <Panel title="My feed">
        <div class="feed" on:scroll={onScroll}>
            {#each posts as post}
                <div class="post">
                    <Post
                        content={post.content}
                        caption={post.caption}
                        likes={post.likeCount}
                        id={post.id}
                        typez={post.type}
                        isLiked={post.hasLiked}
                        isSaved={post.hasSaved}
                    />
                </div>
            {/each}
            {#if loading}
                <p>Loading more posts...</p>
            {/if}
            {#if error}
                <ErrorPopup message={error} />
            {/if}
        </div>
    </Panel>
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
        justify-content: flex-start;
        height: 10rem;
    }

    .username {
        padding: 0;
        margin-left: 0rem;
        margin-right: 10rem;
    }

    .followers-list-container,
    .following-list-container {
        margin-left: 5rem;
        margin-top: 1.8rem;
    }

    .nav {
        display: flex;
        flex-direction: row;
        align-items: flex-start;
    }

    .buttons {
        margin-left: 25rem;
        display: flex;
        flex-direction: row;
        align-items: center;
    }

    .library {
        margin-left: 1rem;
        margin-top: 1rem;
    }

    .liked {
        margin-left: 8rem;
        margin-top: 1rem;
    }

    .bio {
        margin-top: 2rem;
        width: 100rem;
    }

    .feed {
        height: calc(100vh - var(--navbar-height));
        overflow-y: auto;
        padding: 1rem;
    }

    .following-list-content,
    .followers-list-content {
        position: absolute;
        z-index: 1000;
        background-color: white;
        border: black solid 1px;
        border-radius: 1rem;
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
        padding: 0.5rem;
        margin-top: 0.5rem;
        margin-left: 0rem;
    }
    .followers-list-content a,
    .following-list-content a {
        display: block;
        padding: 0.2rem;
        margin: 0.5rem;
        background-color: white;
        border-radius: 5px;
        text-decoration: none;
    }

    .followers-list-content a:hover,
    .following-list-content a:hover {
        text-transform: uppercase;
    }
</style>
