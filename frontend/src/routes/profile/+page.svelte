<script lang="ts">
    import Panel from '../../components/Panel.svelte';
    import NavBar from '../../components/NavBar.svelte';
    import Button from '../../components/Button.svelte';
    import { get, throwError } from '../../fetch';
    import type PostModel from '../../models/post';
    import Post from '../../components/Post.svelte';
    import { errorMessage } from '../../store';
    import { onMount } from 'svelte';
    import ErrorPopup from '../../components/ErrorPopup.svelte';

    let posts: PostModel[] = [];
    let followers: string[] = [];
    let following: string[] = [];
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

    async function getFollowersOrFollowing(item: string): Promise<void> {
        try{
            if(item == 'following'){
                following = await get(`me/${item}`);
                console.log(following)
            }
            else if(item == 'followers'){
                followers = await get(`me/${item}`);
                console.log(followers)
            }
        }catch(e){
            throwError('Could not get followers/following');
        }
    }

    let isFollowersOpen = false;
    let isFollowingOpen = false;
    const handleFollowers = async (): Promise<void> => {
        if(!isFollowersOpen){
            await getFollowersOrFollowing('followers');
        }
        isFollowersOpen != isFollowersOpen;
        isFollowingOpen = false;
    }

    const handleFollowing = async (): Promise<void> => {
        if(!isFollowingOpen){
            await getFollowersOrFollowing('following');
        }
        isFollowingOpen != isFollowingOpen;
        isFollowersOpen = false;
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

    onMount(() => {
        fetchPosts().then((fetchedPosts) => {
            posts = fetchedPosts;
        });
    });
</script>

<div class="nav">
    <NavBar current_page="/profile"></NavBar>
</div>

<div class="profile-container">
    <div class="user-container">
        <div class="user">
            <h2 class="username">Username</h2>
                <button class="button" on:click={async () => await handleFollowing()}>
                    FOLLOWING
                </button>
                {#if isFollowingOpen}
                <div class="following">
                {#each following as item}
                    <p>
                        <a href="/user/{item}">{item}</a>
                    </p>
                {/each}
                {#if following.length == 0}
                    <p>Not following anyone</p>
                {/if}
                {#if error}
                    <ErrorPopup message = {error}/>
                {/if}
            </div>
            {/if}
            <div class="followers">
                <button class="button" on:click={async () => await handleFollowers()}></button>
                {#each followers as item}
                <p>
                    <a href="/user/{item}">{item}</a>
                </p>
                {/each}
                {#if followers.length == 0}
                    <p>No followers</p>
                {/if}
                {#if error}
                    <ErrorPopup message = {error}/>
                {/if}
            </div>

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
                <ErrorPopup message = {error}/>
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
        height: calc(100vh - var(--navbar-height));
        overflow-y: auto;
        padding: 1rem;
    }
    .button{
        width: 7rem;
    }
</style>
