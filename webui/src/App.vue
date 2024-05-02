<script setup>
import { RouterLink, RouterView } from 'vue-router';
import Modal from './components/Modal.vue';
</script>
<script>
export default {
	data() {
		return {
			searchModalIsVisible: false,
			isLoggedIn: sessionStorage.token ? true : false,
			path: "/user/" + sessionStorage.userID,
			
		}
	},

	methods: {
		handleSearchModalToggle() {
			this.searchModalIsVisible = !this.searchModalIsVisible;
		},
		handleLoginSuccess() {
      		this.isLoggedIn = true;
			this.path = "/user/" + sessionStorage.userID;
    	},
		logout() {
			sessionStorage.clear();
			this.isLoggedIn = false;
			this.$router.push('/');
		},
	},


}
</script>

<template>

	<header class="navbar navbar-dark sticky-top bg-dark flex-md-nowrap p-0 shadow">
		<span class="navbar-brand col-md-3 col-lg-2 me-0 px-4 fs-4">WasaPhoto</span>
	</header>

	<div class="container-fluid">
		<div class="row">
			<nav id="sidebarMenu" class="col-md-3 col-lg-2 d-md-block bg-light sidebar collapse" v-show="isLoggedIn">
				<div class="position-sticky pt-3 sidebar-sticky">
					<Modal :show="searchModalIsVisible" @close="handleSearchModalToggle" title="search">
						<template v-slot:header>
							<h3>Users</h3>
						</template>
					</Modal>
					<ul class="nav flex-column fs-5">
						<li class="nav-item m-2" v-if="isLoggedIn">
							<a class="nav-link" @click="handleSearchModalToggle">
								<svg class="feather">
									<use href="/feather-sprite-v4.29.0.svg#search" />
								</svg> 
								Search
							</a>
						</li>
						<li class="nav-item m-2" v-if="isLoggedIn">
							<RouterLink to="/home" class="nav-link">
								<svg class="feather">
									<use href="/feather-sprite-v4.29.0.svg#list" />
								</svg>
								Feed
							</RouterLink>
						</li>
						
						<li class="nav-item m-2" v-if="isLoggedIn">

							<RouterLink :to="path" class="nav-link">
								<svg class="feather">
									<use href="/feather-sprite-v4.29.0.svg#user" />
								</svg>
								Account
							</RouterLink>
						</li>

						<li class="nav-item m-2" v-if="isLoggedIn">
							<RouterLink class="nav-link" to="/create">
								<svg class="feather">
									<use href="/feather-sprite-v4.29.0.svg#plus-circle" />
								</svg>
								Create Post
							</RouterLink>
						</li>
						<li class="nav-item m-2" v-if="isLoggedIn">
							<a class="nav-link" @click="logout">
								<svg class="feather">
									<use href="/feather-sprite-v4.29.0.svg#log-out" />
								</svg>
								Logout
							</a>
						</li>
						<li class="nav-item m-2" v-else>
							<RouterLink to="/" class="nav-link">
								<svg class="feather">
									<use href="/feather-sprite-v4.29.0.svg#log-in" />
								</svg>
								Login
							</RouterLink>
						</li>
					</ul>
				</div>
			</nav>

			<main class="col-md-9 ms-sm-auto col-lg-10 px-md-4">
				<RouterView @login-success="handleLoginSuccess" />
			</main>
		</div>
	</div>
</template>

<style>
a{
	cursor: pointer
}
span, header h1, h2, h3, h4, p{
	user-select: none;
}
</style>
