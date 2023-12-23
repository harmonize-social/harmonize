<script lang="ts">
	import type CommentModel from '../models/comment';
	import type PostModel from '../models/post';
	import Comment from './Comment.svelte';
	import Song from './Song.svelte';
	import Album from './Album.svelte';
	import Playlist from './Playlist.svelte';
	import Artist from './Artist.svelte';
	import ActionButton from './ActionButton.svelte';
	import { get, post, throwError } from '../fetch.js';
	import { onMount } from 'svelte';
	import type ArtistModel from '../models/artist';
	export let content: any;
	let comments: CommentModel[] = [];
	let artists: ArtistModel[] = [];
	export let caption: PostModel['caption'];
	export let likes: PostModel['likeCount'];
	export let id: PostModel['id'];
	export let type: PostModel['type'];

	async function fetchData(){
		try{
			artists = await getArtists();
			comments = await getComments();
		}
		catch(error){
			throwError('Error fetching data');
		}
	}
	async function getArtists(): Promise<ArtistModel[]> {
		try {
			const response: ArtistModel[] = await get<ArtistModel[]>(`/artists?id=${id}`);
			artists = response;
			return artists;
		} catch (error) {
			throwError('Error fetching artists');
			return [];
		}
	}
	async function getComments(): Promise<CommentModel[]> {
		try {
			const response: CommentModel[] = await get<CommentModel[]>(`/comments?id=${id}`);
			comments = response;
			return comments;
		} catch (error) {
			throwError('Error fetching comments');
			return [];
		}
	}
	async function postLike(): Promise<number> {
		try {
			const response: number = await post<number, number>(`/likes?id=${id}`, 0);
			likes++;
			return response;
		} catch (error) {
			throwError('Error posting like');
			return 0;
		}
	}
	async function postSave(): Promise<number> {
		try {
			const response: number = await post<number, number>(`/me/saved?id=${id}`, 0);
			return response;
		} catch (error) {
			throwError('Error posting save');
			return 0;
		}
	}

</script>

<div class="post">
	<h3>{caption}</h3>
	{#if (type = 'song')}
		<Song {content} />
	{:else if (type = 'album')}
		<Album {content} />
	{:else if (type = 'playlist')}
		<Playlist {content} />
	{:else if (type = 'artist')}
		{#await getArtists() then artists}
				<Artist content={artists} />
		{:catch error}
			<p>{error.message}</p>
		{/await}
	{:else}
		<p>Invalid content type</p>
	{/if}
	<h4>Comments:</h4>
	{#await getComments() then comments}
		{#each comments as comment}
			<Comment content={comment} />
		{:else}
			<p>No comments</p>
		{/each}
	{:catch error}
		<p>{error.message}</p>
	{/await}
	<h4>Likes: {likes}</h4>
	<ActionButton type="like" on:click={postLike} />
	<ActionButton type="save" on:click={postSave} />
</div>

<style>
	.post {
		border: 1px solid black;
		padding: 50px;
		margin: 50px;
		border-radius: 10px;
		background-color: grey;
		color: black;
		text-align: center;
	}
</style>
