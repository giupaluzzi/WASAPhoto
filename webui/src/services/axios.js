import axios from "axios";

const instance = axios.create({
	baseURL: __API_URL__,
	timeout: 1000 * 5
});

instance.interceptors.request.use(
	(config)=> {
		const auth = localStorage.getItem('auth')
		if (auth){
			config.headers['Authorization'] = 'Bearer ' + auth;
		}

		return config
	},

	(error)=> {
		return Promise.reject(error)
	}
);

export default instance;
