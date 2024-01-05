<script lang="ts">
	import type SongModel from '../models/song';
	import Preview from './Preview.svelte';
	export let content: SongModel;
	export let number: number;
</script>

<div class="content">
	<div class="top">
		<div class="top-wrapper">
			<div class="left">
				{#if content.album && content.album.title}
					<img src={content.album.mediaUrl} alt="Album Cover" />
				{/if}
			</div>
			<div class="right">
				<div class="title">
					<h3>{number}. {content.title}</h3>
				</div>
				<div class="artists">
					{#if content.artists && content.artists}
						<p>
							{#each content.artists as artist, i}
								{#if i > 0},
								{/if}
								{artist.name}
							{/each}
						</p>
					{/if}
				</div>
			</div>
		</div>
	</div>
	<div class="bottom">
		<div class="preview">
			<Preview previewUrl={content.previewUrl} />
		</div>
	</div>
</div>

<style>
	.content {
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
	}

	.top {
		width: 15rem;
	}

	.top-wrapper {
		display: flex;
		flex-direction: row;
		justify-content: space-between;
	}

	.left {
		display: flex;
		width: 50px;
		justify-content: center;
	}

	.left img {
		width: 50px;
		height: 50px;
		margin-top: auto;
		margin-bottom: auto;
	}

	.right {
		width: 100%;
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
	}
</style>
