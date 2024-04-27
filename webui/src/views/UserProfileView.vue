<script>
import Post from "../components/Post.vue";
import Modal from "../components/Modal.vue";
export default {
    data() {
        return {
            errorMsg: "",
            userID: parseInt(this.$route.params.userID),
            username: "",
            newUsername: "",
            followersCount: 0,
            followers: [],
            followingsCount: 0,
            followings: [],
            postsCount: 0,
            isFollowed: false,
            usernameValidation: new RegExp('^\\w{3,16}$'),
            isOwner: false,
            isBanned: false,

            // Posts data
            posts: [],
            showPost: false,
            postViewData: {},
            
            followingModalIsVisible: false,
			followersModalIsVisible: false,
            updateNameModalIsVisible: false,

        }
    },
    components: {
        Post,
        Modal
    },
    watch: {
        '$route.params.userID'() {
            this.userID = parseInt(this.$route.params.userID);
            if(!isNaN(this.userID)){
                this.getProfile();
                this.getPosts();
            }
        }
    },
    methods: {  
        handleFollowersToggle() {
			this.followersModalIsVisible = !this.followersModalIsVisible;
		},
		handleFollowingToggle() {
			this.followingModalIsVisible = !this.followingModalIsVisible;
		},
		handleUpdateNameToggle() {
			this.updateNameModalIsVisible = !this.updateNameModalIsVisible;
            this.newUsername = "";
            this.errorMsg = "";
		},
        async getProfile() {
            try {
                let response = await this.$axios.get(`users/${this.userID}`, { headers: { 'Authorization': `${sessionStorage.token}` } })
                this.username = response.data.user.username;
                this.followersCount = response.data.followersCount;
                this.followers = response.data.followers;
                this.followingsCount = response.data.followedCount;
                this.followings = response.data.followed;
                this.postsCount = response.data.postsCount;
                this.isFollowed = response.data.followCheck;
                this.isOwner = sessionStorage.userID == this.userID;
                this.isBanned = response.data.isBanned;
            } catch (e) {
                this.errorMsg = e.toString();
            }
        },
        async getPosts() {
            if(this.isBanned)return;
            try {
                let response = await this.$axios.get(`/users/${this.userID}/posts`, { headers: { 'Authorization': `${sessionStorage.token}` } });
                if (response.data == null) {
                    this.dataAvaible = false;
                    return;
                }
                this.posts.push(...response.data);
            } catch (e) {
                this.errorMsg = e.toString();
            };
        },
        async updateUsername() {
            if(this.newUsername == this.username){
                this.errorMsg = "You must enter a new username";  
                return
            } 
            if(this.newUsername.length < 3 || this.newUsername.length > 16){
                this.errorMsg = "Invalid username, it must contains min 3 characters and max 16 characters";
                return
            }
            if(!this.usernameValidation.test(this.newUsername)){
                this.errorMsg = "Invalid username, it must contain only letters and numbers";
                return
            }
            try{
                let _ = await this.$axios.put(`/users/${this.userID}/username`, { username: this.newUsername }, { headers: { 'Authorization': `${sessionStorage.token}` } })
                this.username = this.newUsername;
                this.errorMsg = "";
                this.handleUpdateNameToggle();
            } catch (e) {
                if(e.response.data == "Username already exist\n"){
                    this.errorMsg = "This username is already taken. Please try another one.";
                }else{
                    this.errorMsg = e.toString();
                }
            }
            
		},
        async follow() {
            if (this.isFollowed) {
                try {
                    let _ = await this.$axios.delete(`users/${sessionStorage.userID}/follows/${this.userID}`, { headers: { 'Authorization': `${sessionStorage.token}` } });
                    this.isFollowed = false;
                    this.followersCount--;
                } catch (e) {
                    this.errorMsg = e.toString();
                }
            } else {
                try {
                    let _ = await this.$axios.put(`users/${sessionStorage.userID}/follows/${this.userID}`, {}, { headers: { 'Authorization': `${sessionStorage.token}` } });
                    this.isFollowed = true;
                    this.followersCount++;
                } catch (e) {
                    this.errorMsg = e.toString();
                }
            }
        },
        openPost(post) {
            this.showPost = true;
            this.postViewData = post;
        },
        exitPost() {
            this.showPost = false;
            this.postViewData = {};
        },
        updateLike(data) {
            this.posts.forEach(post => {
                if (post.postID == data['postID']) {
                    post.liked = data.liked;
                    post.likesCount = data.liked ? post.likesCount + 1 : post.likesCount - 1;
                }
            });
        },
        async banUser() {
            if(!this.isBanned){
                try {
                    let _ = await this.$axios.put(`/users/${sessionStorage.userID}/banned/${this.userID}`, {}, { headers: { 'Authorization': `${sessionStorage.token}` } });
                    this.$router.push(`${sessionStorage.userID}`);
                } catch (e) {
                    this.errorMsg = e.toString();
                }
            }else{
                try {
                    let _ = await this.$axios.delete(`/users/${sessionStorage.userID}/banned/${this.userID}`, { headers: { 'Authorization': `${sessionStorage.token}` } });
                    this.isBanned = false;
                } catch (e) {
                    this.errorMsg = e.toString();
                    console.log(e);
                }
            }
        },
        async deletePost(postID) {
            try {
                let _ = await this.$axios.delete(`profiles/${localStorage.userID}/posts/${postID}`, { headers: { 'Authorization': `${localStorage.token}` } });
                this.posts = this.posts.filter(post => post.postID != postID);
                this.postsCount--;
                this.exitPost();
            } catch (e) {
                this.errorMsg = e.toString();
            }
        },
        updateProfile() {   
            this.getProfile();
            localStorage.propic64 = this.proPic64;
        },
    },
    beforeMount() {
        if (!sessionStorage.token) {
            this.$router.push('/login');
        }
        if (sessionStorage.userID === this.$route.params.userID) {
            this.isOwner = true;
        }
    },

    mounted() {
        this.getProfile();
        if(!this.isBanned){
            this.getPosts();
        }
    },
    
}

</script>


<template>
    <div class="header"> 
        <Modal :show="followersModalIsVisible" @close="handleFollowersToggle" :users="this.followers" title = "followers">
            <template v-slot:header>
                <h3>Followers</h3>
            </template>
        </Modal>
        <Modal :show="followingModalIsVisible" @close="handleFollowingToggle" :users="this.followings" title = "followed">
            <template v-slot:header>
                <h3>Followed</h3>
            </template>
        </Modal>
        <Modal :show="updateNameModalIsVisible" @close="handleUpdateNameToggle" title = "username">
            <template v-slot:header>
                <h3>Update Username</h3>
            </template>
            <template v-slot:body>
                <form class="username-form">
                    <ErrorMsg v-if="errorMsg" :msg="errorMsg"></ErrorMsg>
                    <input type="text" v-model="this.newUsername" placeholder="New username" />
                    <button type="submit" @click.prevent="updateUsername">Update</button>
                </form>
            </template>
        </Modal>
        <div class="top">
            <h1>{{ this.username }}
                <button v-if="isOwner" @click="handleUpdateNameToggle">
                    <svg class="feather edit">
                        <use href="/feather-sprite-v4.29.0.svg#edit" />
                    </svg>
                </button>
            </h1>

            <div v-if="!isOwner">
                <button v-if="!isBanned" class="follow-btn" @click="follow">{{ isFollowed ? "Unfollow" : "Follow" }}</button>
                <button class="ban-btn" @click="banUser">{{ isBanned ? "Unban" : "Ban" }}</button>
            </div>
        </div>
        <div class="bottom"> 
            <h4 class = "post">{{ ((this.postsCount > 1) ? " Posts: " : " Post: ") + this.postsCount }}</h4>
            <h4 class = "followed" @click="handleFollowingToggle">{{ " Followed: " + this.followingsCount }}</h4>
            <h4 class = "followers" @click="handleFollowersToggle">{{ (( this.followersCount > 1) ? " Followers: " : " Follower: ") + this.followersCount}}</h4>
        </div>
    </div>



    <div class="feed" v-for="post in this.posts" :key="post.ID">
        <Post :post="post" :identifier="this.identifier" @delete="() => deleteImage(post.ID)" />
    </div>


</template>