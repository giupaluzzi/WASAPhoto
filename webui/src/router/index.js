import {createRouter, createWebHashHistory} from 'vue-router'
import HomeView from '../views/HomeView.vue'
import LoginView from "../views/LoginView.vue";
import ProfileView from "../views/ProfileView.vue";

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/', redirect: '/login'},
		{path: '/login', component: LoginView},
		{path: '/home', component: HomeView /*, meta: {requiresAuth: true}*/ },
//		{path: '/users/:id', component: SearchView, meta: {requiresAuth: true}},
		{path: '/profile', component: ProfileView /*, meta: {requiresAuth: true}*/},
//		{path: '/settings', component: SettingsView, meta: {requiresAuth: true}}
	]
})
/*
router.beforeEach((to, from, next) => {
	const isLogged = checkLogin();
	if (to.matched.some(record => record.meta.requiresAuth) && !isLogged) {
		next({component: LoginView});
	} else {
		next();
	}
});

function checkLogin() {
	const auth = localStorage.getItem('auth')
	return !!auth
}
*/
export default router
