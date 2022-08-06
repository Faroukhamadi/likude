let accessToken: string | undefined = '';

export const setAccessToken = (s: string | undefined) => {
	accessToken = s;
};

export const getAccessToken = () => {
	return accessToken;
};
