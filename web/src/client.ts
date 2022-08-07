import { HoudiniClient, type RequestHandlerArgs } from '$houdini';

async function fetchQuery({
	fetch,
	text = '',
	variables = {},
	session,
	metadata
}: RequestHandlerArgs) {
	let val = localStorage.getItem('sid') ? `Bearer ${localStorage.getItem('sid')}` : '';
	console.log(val);

	const url = import.meta.env.VITE_GRAPHQL_ENDPOINT || 'http://localhost:4000/query';
	const result = await fetch(url, {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json',
			Authorization: localStorage.getItem('sid') ? `Bearer ${localStorage.getItem('sid')}` : ''
		},
		body: JSON.stringify({
			query: text,
			variables
		})
	});
	return await result.json();
}

export default new HoudiniClient(fetchQuery);
