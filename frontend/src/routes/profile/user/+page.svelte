<script lang="ts">
    import ErrorPopup from '../../../components/ErrorPopup.svelte';
    import Button from '../../../components/Button.svelte';
    import Panel from '../../../components/Panel.svelte';
    import NavBar from '../../../components/NavBar.svelte';
    import { delete_, get, post, throwError } from '../../../fetch';
    import type PostModel from '../../../models/post';
    import Post from '../../../components/Post.svelte';
    import { onMount } from 'svelte';
    import { errorMessage } from '../../../store';

    let posts: PostModel[] = [];
    let name: string = '';
    let follows: boolean = true;
    let loading = false;
    let error = '';

    errorMessage.subscribe((value) => {
        error = value;
    });

    async function fetchPosts(username: string | null): Promise<PostModel[]> {
        try {
            const response: PostModel[] = await get<PostModel[]>(`/posts?username=${username}`);
            return response;
        } catch (e) {
            throwError('Error fetching posts');
            return [];
        } finally {
            loading = false;
        }
    }

    async function deleteFollow(name: string){
        try{
            const response: string = await delete_<string>(`/follow?username=${name}`);
            return response;
        }catch(e){
            throwError('Error deleting follow');
            return 0;
        }
    }
    async function postFollow(name: string){
        try{
            const response: string = await post<string, string>(`/follow?username=${name}`, name);
            return response;
        }catch(e){
            throwError('Error posting follow');
            return 0;
        }
    }

    function onScroll(event: Event) {
        const target = event.target as HTMLElement;
        if (target.scrollHeight - target.scrollTop === target.clientHeight) {
            loadMorePosts(name);
        }
    }
    async function loadMorePosts(name: string) {
        if (loading) return;
        const morePosts = await fetchPosts(name);
        posts = [...posts, ...morePosts];
    }

    let isClicked = false;
    const handleButtonClick = async () => {
        if(follows){
            await deleteFollow(name);
            isClicked = !isClicked;
        }else{
            await postFollow(name);
            isClicked = !isClicked;
        }
        follows = !follows;
    };
    onMount(() => {
        const params = new URLSearchParams(window.location.search);
        let username = params.get('username');
        fetchPosts(username).then((fetchedPosts) => {
            posts = fetchedPosts;
        });
    });
</script>

<!-- navbar -->
<div class="nav">
    <NavBar current_page={`/user/${name}`}></NavBar>
</div>
<!-- profile -->
<div class="profile-container">
    <div class="user-container">
        <div class="user">
            <h2 class="username">{name}</h2>
            <div class="bio">
                <p>
                    Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt
                    ut labore et dolore magna aliqua.
                </p>
            </div>
            <div class="follow_button">
                {#if follows}
                    <Button buttonText="Unfollow" action={handleButtonClick}></Button>
                {:else}
                    <Button buttonText="Follow" action={handleButtonClick}></Button>
                {/if}
            </div>
        </div>
    </div>
    <!-- personal feed -->
    <div class="feed-container" on:scroll={onScroll}>
        <Panel title={`${name}'s posts`}>
            <div class="feed">
                {#if posts.length === 0}
                    <p>{name} did not post yet!</p>
                {/if}
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
