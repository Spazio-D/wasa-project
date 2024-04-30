<script>

import Post from '../components/Post.vue';

export default {
	data: function () {
		return {
			errorMsg: null,
			feed: null,
			token : sessionStorage.token,
		};
	},
	mounted() {
		if (!sessionStorage.token) {
			this.$router.push('/');
			return
		}
	},
	components: { Post },
    emits: ['login-success'],

}   
</script>

<template>
    
    <div v-if="feed != null" style="margin-top: 30px;">

        <div class="feed" v-for="post in feed" :key="post.ID">
            <Post :post="post" :token="this.token"/>
        </div>

    </div>
    <div v-else class="empty-msg">
        <p>There's nothing here, just follow someone :D</p>
    </div>
</template>
  
<style>
.empty-msg {
    display: flex;
    justify-content: center;
    align-items: center;
    margin-top: 30%;
    font-size: 20px;
};
.feed {
    display: flex;
    flex-direction: column;
    align-items: center;
    width: 400px;
    margin: 0 auto;
};
</style>