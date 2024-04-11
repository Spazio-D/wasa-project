import {createRouter, createWebHashHistory} from 'vue-router'
import HomeView from '../views/HomeView.vue'
import LoginView from '../views/LoginView.vue'
import CreatePostView from '../views/CreatePostView.vue'
import UserProfileView from '../views/UserProfileView.vue'

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/', component: LoginView},
		{path: '/home', component: HomeView},
		{path: '/create', component: CreatePostView},
		{path: '/user/:userID', component: UserProfileView},
	]
})

export default router
