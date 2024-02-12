import {createRouter, createWebHashHistory} from 'vue-router'
import HomeView from '../views/HomeView.vue'
import LoginView from "../views/LoginView.vue";
import LoggedProfileView from "../views/LoggedProfileView.vue";
import NotFoundView from "../views/NotFoundView.vue";
import ProfileView from "../views/ProfileView.vue";
import SettingsView from "../views/SettingsView.vue";

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/', redirect: '/login'},
		{path: '/login', component: LoginView},
		{path: '/home', component: HomeView},
		{path: '/profile', component: LoggedProfileView},
		{path: '/profile/settings', component: SettingsView},
		{path: '/users/:id', component: ProfileView},
		{path: '/:catchAll(.*)', component: NotFoundView},
	]
})

export default router
