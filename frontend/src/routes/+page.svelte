<script lang="ts">
	import { goto } from '$app/navigation';
	import Logo from '../components/Logo.svelte';
	import Button from '../components/Button.svelte';
	import deezerIcon from '../lib/assets/deezer.png';
	import spotifyIcon from '../lib/assets/Spotify_App_Logo.svg.png';
	function loginClicked() {
		goto('/auth/login'); // Leidt naar de loginpagina
	}
</script>

<main>
	<Logo />
	<section>
		<h1>Welcome!</h1>
		<p>
			At Harmonize, we are passionate about music's power to unite people and transcend boundaries.
			We believe that every artist, regardless of their platform or location, should have the
			opportunity to share their unique sound with the world. That's why we've created a
			revolutionary social media platform dedicated to making music sharing and collaboration
			seamless and effortless.
		</p>
	</section>

	<section>
		<h2>Our Vision: Music Without Boundaries</h2>
		<p>
			In a world where music has the potential to touch every corner of the globe, Harmonize seeks
			to eliminate the barriers that often hinder the free flow of musical creativity. We envision a
			future where artists can effortlessly transfer their music between two different platforms,
			bridging the gaps that separate artists, genres, and cultures.
		</p>
	</section>

	<section>
		<h2>Seamless Music Transfer</h2>
		<p>
			Harmonize empowers musicians to seamlessly move their music between two different platforms,
			offering unparalleled flexibility and accessibility. Whether you're a SoundCloud enthusiast, a
			Spotify aficionado, or any other platform user, Harmonize lets you effortlessly connect your
			music to share your creations with the world.
		</p>
	</section>

	<section>
		<h2>Sharing and Collaboration</h2>
		<p>
			Sharing your musical journey with Harmonize is a breeze. Post your tracks, albums, and
			playlists with ease, and watch your audience grow. Connect with fellow musicians and music
			enthusiasts, fostering collaboration, inspiration, and a sense of community that transcends
			geographical boundaries.
		</p>
	</section>

	<section>
		<h2>Automatic Music Preservation</h2>
		<p>
			At Harmonize, we understand the value of preserving your musical legacy. That's why we've
			implemented a unique feature that automatically saves your music from others directly onto
			your platforms. No need to worry about losing tracks or painstakingly transferring them
			manually. Harmonize does the work for you, ensuring your music is safely preserved for future
			generations to enjoy.
		</p>
	</section>

	<section>
		<h2>Join the Harmonize Movement</h2>
		<p>
			Join us on a journey to harmonize the world through music. At Harmonize, we are dedicated to
			fostering a global community of music lovers, creators, and innovators who share a common
			passion for breaking down barriers and making music accessible to all. Together, we can create
			a world where music knows no boundaries.
		</p>
	</section>

	<section>
		<p>
			Experience the future of music sharing and collaboration. Join Harmonize today and become part
			of a vibrant and inclusive musical community that connects the world one song at a time.
		</p>
	</section>

	<div class="button-container">
		<Button buttonText="Get Started" link="./auth/register">Get Started!</Button>
		<div>
			<p>
				Already registered? <a href="/login" on:click|preventDefault={loginClicked}>Log in here</a>
			</p>
		</div>
	</div>
		<div class="icons-container">
			<h4>The platforms we provide our services for (more to come)</h4>
			<div class="icons">
			<img src="{spotifyIcon}" alt="Spotify Icon">
				<img src="{deezerIcon}" alt="Deezer Icon">
			</div>
		</div>
</main>

<style>
	main {
		max-width: 800px;
		margin: 0 auto;
		padding: 20px;
		font-family: 'Roboto', sans-serif;
	}

	h2,
	h1 {
		color: #f8e7e7;
		font-size: 24px;
		margin-top: 20px;
	}
	h4{
		color: #f8e7e7;
		font-size: 15px;
		margin-top: 20px;
	}

	p {
		color: #f8e7e7;
		line-height: 1.5;
		margin-bottom: 15px;
	}

	.button-container {
		margin-top: 30px;
		text-align: center;
	}

	section {
		margin-bottom: 40px;
	}

	.icons-container{
		width: 100%;
		margin-top: 30px;
		margin-bottom: 100px;
	}
	.icons-container h4{
		margin: 0 0 20px 12rem;
	}

	.icons{
		display: flex;
		flex-direction: row;
		justify-content: space-evenly;
	}

	.icons img{
		width: 100px;
		height: 100px;
		border-radius: 50%;
		border: 1px solid black;
	}
</style>
