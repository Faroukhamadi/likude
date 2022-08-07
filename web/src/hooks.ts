import type { Handle, GetSession } from '@sveltejs/kit';
import * as cookie from 'cookie';

export const handle: Handle = async function ({ event, resolve }) {
	console.log('hello from handle');
	const cookies = cookie.parse(event.request.headers.get('cookie') || '');
	console.log('cookies: ', cookies);
	const jwt = cookies.jwt && Buffer.from(cookies.jwt, 'base64').toString('utf-8');
	console.log('this is the jwt: ', jwt);
	event.locals.user = jwt ? JSON.parse(jwt) : null;
	console.log('this is the user: ', event.locals.user);
	return await resolve(event);
};

export const getSession: GetSession = function ({ locals }) {
	console.log('get session called');
	return {
		user: locals.user && {
			username: locals.user.username,
			id: locals.user.id
		}
	};
};
