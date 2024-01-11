<!-- Login.svelte -->

<script lang="ts">
	import Button from '../../../components/Button.svelte';
	import Panel from '../../../components/Panel.svelte';
	import TextInput from '../../../components/TextInput.svelte';
	import { goto } from '$app/navigation';
	import { loginpost, throwError } from '../../../fetch';
	import { errorMessage } from '../../../store';
	import ErrorPopup from '../../../components/ErrorPopup.svelte';
	import { onMount } from 'svelte';

	let username = '';
	let password = '';
	let error = '';
	errorMessage.subscribe((value) => {
		error = value;
	});

	const handleLogin = async () => {
		try {
			const response = await loginpost<{ token: string }>('/users/login', { username, password });

			// Set token in local storage
			localStorage.setItem('token', response.token);

			// Redirect to dashboard or other protected route on successful login
			goto('/feed');
		} catch (e) {
			throwError('Login failed');
		}
	};
	onMount(async () => {
		errorMessage.set('');
	});
</script>

<div class="panel-container">
	<Panel title="">
		<h2>Login</h2>
		{#if error}
			<ErrorPopup message={error} />
		{/if}
		<div class="text-input">
			<TextInput placeholder="Username/Email" bind:value={username} />
		</div>
		<div class="text-input">
			<TextInput placeholder="Password" type="password" bind:value={password} />
		</div>
		<div class="buttonlogin">
			<Button buttonText="Login" action={handleLogin} />
		</div>
		<a class="forgot-password" href="/auth/forgot-password">Forgot Password?</a>
		<a class="not-registered" href="/auth/register">Not registred yet?</a>
	</Panel>
</div>

<style>
	.panel-container {
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
		height: 80vh;
	}

	.text-input {
		width: 200px; /* Take up all available width */
		margin: 10px;
		display: flex;
		justify-content: center;
	}

	.buttonlogin {
		margin: 10px;
		display: flex;
		justify-content: center;
	}

	.forgot-password {
		display: flex;
		justify-content: center;
		margin-top: 10px;
	}
	.not-registered {
		display: flex;
		justify-content: center;
		margin-top: 10px;
	}
</style>
