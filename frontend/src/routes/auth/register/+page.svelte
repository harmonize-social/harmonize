<script lang="ts">
	import Button from '../../../components/Button.svelte';
	import Panel from '../../../components/Panel.svelte';
	import TextInput from '../../../components/TextInput.svelte';
	import { loginpost, throwError } from '../../../fetch';
	import { goto } from '$app/navigation';
	import ErrorPopup from '../../../components/ErrorPopup.svelte';
	import { errorMessage } from '../../../store';

	let confirmPassword: string = '';
	let password: string = '';
	let email: string = '';
	let username: string = '';
	let error: string = '';
	errorMessage.subscribe((value) => {
		error = value;
	});
	const handleRegister = async () => {
		if (password !== confirmPassword) {
			throwError('Passwords do not match');
			return;
		}
		try {
			const response = await loginpost<{ token: string }>('/users/register', {
				email,
				password,
				username
			});
			localStorage.setItem('token', response.token);
			goto('/dashboard');
		} catch (e) {
			throwError(e as string) || 'Register failed. Please try again.';
		}
	};
</script>

<Panel title="Register">
	{#if error}
		<ErrorPopup message={error} />
	{/if}
	<div class="text-input">
		<TextInput placeholder="Username" bind:value={username} />
	</div>
	<div class="text-input">
		<TextInput placeholder="Email" bind:value={email} />
	</div>
	<div class="text-input">
		<TextInput placeholder="Password" type="password" bind:value={password} />
		<TextInput placeholder="Confirm Password" bind:value={confirmPassword} type="password" />
	</div>

	<a class="forgot-password" href="./login">Already have an account?</a>

	<div class="buttonlogin">
		<Button buttonText="Register" action={handleRegister} />
	</div>
</Panel>

<style>
	.buttonlogin {
		margin: 10px;
		display: flex;
		justify-content: center;
	}
	.forgot-password {
		display: flex;
		justify-content: center;
	}
</style>
