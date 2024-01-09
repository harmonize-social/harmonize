<script lang="ts">
    import type CommentModel from '../models/comment';
    import type PostModel from '../models/post';
    import Comment from './Comment.svelte';
    import Song from './Song.svelte';
    import Album from './Album.svelte';
    import Playlist from './Playlist.svelte';
    import Artist from './Artist.svelte';
    import ActionButton from './ActionButton.svelte';
    import { delete_, post, throwError } from '../fetch';
    import ErrorPopup from './ErrorPopup.svelte';
    import { errorMessage } from '../store';

    export let content: PostModel;
    let comments: CommentModel[] = [];
    let error = '';
    errorMessage.subscribe((value) => {
        error = value;
    });

    async function postLike(): Promise<string> {
        try {
            const response: string = await post<string, string>(`/me/liked?id=${content.id}`, content.id);
            return response;
        } catch (e) {
            throwError('Error posting like');
            return e as string;
        }
    }
    async function postSave(): Promise<string> {
        try {
            const response: string = await post<string, string>(`/me/saved?id=${content.id}`, content.id);
            return response;
        } catch (e) {
            throwError('Error posting save');
            return e as string;
        }
    }

    async function deleteLike(): Promise<string> {
        try {
            const response: string = await delete_<string>(`/me/liked?id=${content.id}`);
            return response;
        } catch (e) {
            throwError('Error deleting like');
            return e as string;
        }
    }

    async function deleteSave(): Promise<string> {
        try {
            const response: string = await delete_<string>(`/me/saved?id=${content.id}`);
            return response;
        } catch (e) {
            throwError('Error deleting save');
            return e as string;
        }
    }

    function toggleLikeButton() {
        if (content.hasLiked) {
            deleteLike();
            content.likeCount--;
        } else {
            postLike();
            content.likeCount++;
        }
        content.hasLiked = !content.hasLiked;
    }

    function toggleSaveButton() {
        if (content.hasSaved) {
            deleteSave();
        } else {
            postSave();
        }
        content.hasSaved = !content.hasSaved;
    }
</script>

<div class="post">
    <p>{content.username}</p>
    {#if content.type == 'song'}
        <Song content={content.content} />
    {:else if content.type == 'album'}
        <Album content={content.content} />
    {:else if content.type == 'playlist'}
        <Playlist content={content.content} />
    {:else if content.type == 'artist'}
        <Artist content={content.content} />
    {:else}
        <p>Invalid content type</p>
    {/if}
    <!--<h4>Comments:</h4>
    {#if comments.length == 0}
        <p>No comments yet</p>
    {/if}
    {#each comments as comment}
        <Comment content={comment} />
    {/each}-->
    <div class="action-buttons">
        <h3 class="caption">{content.caption}</h3>
        <div class="interactions">
            <div class="likes">
                <p>{content.likeCount}</p>
                <ActionButton state={content.hasLiked} type="like" action={toggleLikeButton} />
            </div>

            <div class="saves">
                <ActionButton state={content.hasSaved} type="save" action={toggleSaveButton} />
            </div>
        </div>
    </div>

    {#if error}
        <ErrorPopup message={error}></ErrorPopup>
    {/if}
</div>

<style>
    .post {
        border: 1px solid black;
        padding: 0px;
        margin: 50px;
        border-radius: 10px;
        color: black;
        text-align: center;
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
    }

    .likes {
        display: flex;
        flex-direction: row;
        align-items: center;
        justify-content: center;
    }

    .interactions {
        display: flex;
        flex-direction: row;
        align-items: center;
        justify-content: flex-end;
        width: 100%;
    }

    h3 {
        max-width: 75%;
        justify-items: flex-start;
    }

    .caption {
        margin-left: 20px;
    }

    .saves {
        margin-right: 10px;
    }

    .action-buttons {
        display: flex;
        flex-direction: row;
        justify-content: space-between;
        width: 100%;
        margin: 20px;
    }
</style>
