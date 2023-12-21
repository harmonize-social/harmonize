//https://eckertalex.dev/blog/typescript-fetch-wrapper
//https://github.com/EHB-TI/programming-project-groep-3_brussel-student-guide/blob/main/frontend/src/fetch.ts

import ErrorPopup from './components/ErrorPopup.svelte';

export async function http<T>(path: string, config: RequestInit): Promise<T> {
	// Get token from local storage
	const token = localStorage.getItem('token');
	config.headers = {
		...config.headers,
		Authorization: `Bearer ${token}`
	};
	const request = new Request(path, config);
	const response = await fetch(request);

	if (response.status === 401) {
		console.error('Unauthorized access detected');
		throw new Error('Unauthorized');
	}

	if (!response.ok) {
		throw new Error('Error fetching.');
	}

	try {
		const json = await response.json();
		const object = json.result;
		if (object == null) {
			throw new Error(json.error);
		}
		return object;
	} catch (error) {
		console.error('Error parsing JSON response:', error);
		throw new Error('Error parsing JSON response');
	}
}

export async function rawhttp<T>(request: Request): Promise<T> {
	const response = await fetch(request);

	if (!response.ok) {
		throw new Error('Error fetching.');
	}

	try {
		const json = await response.json();
		const object = json.result;
		if (object == null) {
			throw new Error(json.error);
		}
		return object;
	} catch (error) {
		console.error('Error parsing JSON response:', error);
		throw new Error('Error parsing JSON response');
	}
}

export async function loginpost<T>(url: string, body: Object): Promise<T> {
	const body_string = JSON.stringify(body);
	const length = new TextEncoder().encode(body_string).length;
	const request = new Request(url, {
		method: 'POST',
		body: body_string,
		headers: {
			'Content-Type': 'application/json; charset=UTF-8',
			'Content-Length': length.toString()
		}
	});

	try {
		const response = await fetch(request);

		if (!response.ok) {
			const json = await response.json();
			throw json.error;
		}

		const json = await response.json();
		const object = json.result;
		return object;
	} catch (error) {
		console.error('Internal error, try again later:', error);
		throw new Error('Internal error, try again later.');
	}
}

export async function get<T>(path: string, config?: RequestInit): Promise<T> {
	const init = { method: 'get', ...config };
	return await http<T>(path, init);
}

export async function delete_<T>(path: string, config?: RequestInit): Promise<T> {
	const init = { method: 'delete', ...config };
	return await http<T>(path, init);
}

export async function post<T, U>(path: string, body: T, config?: RequestInit): Promise<U> {
	const init = { method: 'post', body: JSON.stringify(body), ...config };
	return await http<U>(path, init);
}

export async function deleteAccount(url: string): Promise<void> {
	const token = localStorage.getItem('token');
	const request = new Request(url, {
		method: 'DELETE',
		headers: {
			Authorization: `Bearer ${token}`
		}
	});

	try {
		const response = await fetch(request);

		if (!response.ok) {
			const json = await response.json();
			throw new Error(json.error);
		}

		console.log('Account succesvol verwijderd');
		localStorage.removeItem('token');
	} catch (error) {
		console.error('Fout bij het verwijderen van het account:', error);
		throwError('Failed to delete account');
	}
}

export async function logout(url: string): Promise<void> {
	const token = localStorage.getItem('token');

	if (token) {
		localStorage.removeItem('token');
	}

	try {
		const response = await fetch(url, {
			method: 'POST'
		});

		if (!response.ok) {
			const json = await response.json();
			throw new Error(json.error);
		}

		console.log('Uitloggen succesvol');
	} catch (error) {
		console.error('Fout bij uitloggen:', error);
		throwError('Failed to logout');
	}
}

export async function updateUserInfo(
	url: string,
	updatedInfo: { username: string; email: string; password: string }
) {
	const { username, email, password } = updatedInfo;

	const usernamePattern = /^[a-zA-Z0-9]{6,}$/;
	if (!usernamePattern.test(username)) {
		throwError('Username must contain at least 6 characters and only contain letters and numbers');
	}

	const passwordPattern = /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)[a-zA-Z\d]{5,}$/;
	if (!passwordPattern.test(password)) {
		throwError(
			'Password must contain at least 5 characters, 1 uppercase letter, 1 lowercase letter and 1 number'
		);
	}

	const emailPattern = /\S+@\S+\.\S+/;
	if (!emailPattern.test(email)) {
		throwError('Invalid email address');
	}

	const token = localStorage.getItem('token');
	const request = new Request(url, {
		method: 'PUT',
		headers: {
			Authorization: `Bearer ${token}`,
			'Content-Type': 'application/json'
		},
		body: JSON.stringify({ username, email, password })
	});

	try {
		const response = await fetch(request);

		if (!response.ok) {
			const json = await response.json();
			throw new Error(json.error);
		}

		const data = await response.json();
		return data;
	} catch (error) {
		console.error('Fout bij het bijwerken van de gebruikersinformatie:', error);
		throwError('Failed to edit user data');
	}
}

//create a function that renders a pop up with a message that the backend returns when an error occurs
export async function throwError(errorMessage: string) {
	let popup = new ErrorPopup({
		target: document.body as Element,
		props: { message: errorMessage }
	});
	return popup;
}
