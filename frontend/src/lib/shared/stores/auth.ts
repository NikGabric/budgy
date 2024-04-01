import { browser } from '$app/environment';
import type { BudgyUser } from '$lib/models/user';
import { writable, type Writable } from 'svelte/store';

type AuthToken = {
	usr: BudgyUser;
	exp: Date;
};

const parseJwt = (token: string) => {
	const base64Url = token.split('.')[1];
	const base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/');
	const jsonPayload = decodeURIComponent(
		window
			.atob(base64)
			.split('')
			.map(function (c) {
				return '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2);
			})
			.join('')
	);

	return JSON.parse(jsonPayload);
};

export const authTokenString: Writable<string | null> = writable(null);

export const authToken: Writable<AuthToken | null> = writable(
	browser && window.localStorage.getItem('auth-token')
		? parseJwt(window.localStorage.getItem('auth-token')!)
		: null
);

authTokenString.subscribe((val: string | null) => {
	if (browser && val) {
		window.localStorage.setItem('auth-token', val);
		authToken.set(parseJwt(val));
	} else if (browser) {
		window.localStorage.removeItem('auth-token');
		authToken.set(null);
	}
});
