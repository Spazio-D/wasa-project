<script>
import Post from "../components/Post.vue";
import Modal from "../components/Modal.vue";
export default {
    data() {
        return {
            errorMsg: "",
            // Profile data
            userID: parseInt(this.$route.params.userID),
            username: "",
            followersCount: 0,
            followers: [],
            followingsCount: 0,
            followings: [],
            postsCount: 0,
            isFollowed: false,

            isOwner: false,

            // Buttons Text
            followTextButton: "Follow",

            // Other Data
            textCounter: 0,
            profilesArray: [],
            textHeader: "",
            typeList: "",

            // Posts data
            posts: [],
            showPost: false,
            postViewData: {},

            // Load more data
            busy: false,
            dataAvaible: true,

            // Follower data
            dataGetter: () => { },
            showList: false,

            // Options data
            showOptions: false,

            isLoading: false,
            
            followingModalIsVisible: false,
			followersModalIsVisible: false,
            updateNameModalIsVisible: false,

        }
    },
    components: {
        Post,
        Modal
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
		},
        async getProfile() {
            this.isLoading = true;
            try {
                let response = await this.$axios.get(`users/${this.userID}`, { headers: { 'Authorization': `${sessionStorage.token}` } })
                this.username = response.data.user.username;
                this.followersCount = response.data.followersCount;
                this.followers = response.data.followers;
                this.followingsCount = response.data.followedCount;
                this.followings = response.data.followed;
                this.postsCount = response.data.postsCount;
                this.isFollowed = response.data.followCheck;
                this.followTextButton = this.isFollowed ? "Unfollow" : "Follow";
                this.isOwner = sessionStorage.userID == this.userID;
            } catch (e) {
                this.errorMsg = e.toString();
            }
            this.isLoading = false;
        },
        async getPosts() {
            this.isLoading = true;
            try {
                let response = await this.$axios.get(`/users/${this.userID}/posts`, { headers: { 'Authorization': `${sessionStorage.token}` } });
                if (response.data == null) {
                    this.dataAvaible = false;
                    this.isLoading = false;
                    return;
                }
                this.posts.push(...response.data);
            } catch (e) {
                this.errorMsg = this.$utils.errorToString(e);;
            };
            this.isLoading = false;
        },
        editingUsername() {
            if (this.isOwner) {
                document.querySelectorAll(".top-body-profile-username")[0].style.outline = "auto";
                document.querySelectorAll(".top-body-profile-username")[0].style.outlineColor = "#03C988";
            }
        },
        async saveChangeUsername() {
            if (this.isOwner) {
                document.querySelectorAll(".top-body-profile-username")[0].style.outline = "none";
                if (this.username == "" | this.username.length < 3) {
                    this.username = localStorage.username;
                    return
                }
                this.isLoading = true;
                try {
                    let _ = await this.$axios.put(`/profiles/${this.userID}/username`, { username: this.username }, { headers: { 'Authorization': `${localStorage.token}` } });
                    localStorage.username = this.username;
                } catch (e) {
                    this.errorMsg = this.$utils.errorToString(e);;
                    this.username = localStorage.username;
                }
                this.isLoading = false;
            }
        },
        getFollowers() {
            this.showList = true;
            this.textHeader = "Followers";
            this.typeList = "simple";
            this.dataGetter = async (profilesArray, limit, offset, dataAvaible) => {
                try {
                    let response = await this.$axios.get(`/profiles/${this.userID}/followers?limit=${limit}&offset=${offset}`, { headers: { 'Authorization': `${localStorage.token}` } });
                    if (response.data == null) {
                        dataAvaible = false;
                        return;
                    }
                    profilesArray.push(...response.data);
                } catch (e) {
                    this.errorMsg = this.$utils.errorToString(e);;
                }
            }
        },
        getFollowings() {
            this.showList = true;
            this.textHeader = "Followings";
            this.typeList = "simple";
            this.dataGetter = async (profilesArray, limit, offset, dataAvaible) => {
                try {
                    let response = await this.$axios.get(`/profiles/${this.userID}/followings?limit=${limit}&offset=${offset}`, { headers: { 'Authorization': `${localStorage.token}` } });
                    if (response.data == null) {
                        dataAvaible = false;
                        return;
                    }
                    profilesArray.push(...response.data);
                } catch (e) {
                    this.errorMsg = this.$utils.errorToString(e);;
                }
            }
        },
        freeLists() {
            this.showList = false;
            this.profilesArray = [];
            this.textHeader = "";
        },
        async follow() {
            if (this.isFollowed) {
                try {
                    let _ = await this.$axios.delete(`profiles/${localStorage.userID}/followings/${this.userID}`, { headers: { 'Authorization': `${localStorage.token}` } });
                    this.isFollowed = false;
                    this.followTextButton = "Follow";
                    this.followersCount--;
                } catch (e) {
                    this.errorMsg = this.$utils.errorToString(e);;
                }
            } else {
                try {
                    let _ = await this.$axios.put(`profiles/${localStorage.userID}/followings/${this.userID}`, {}, { headers: { 'Authorization': `${localStorage.token}` } });
                    this.isFollowed = true;
                    this.followTextButton = "Unfollow";
                    this.followersCount++;
                } catch (e) {
                    this.errorMsg = this.$utils.errorToString(e);;
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
        getBans() {
            this.showList = true;
            this.textHeader = "Bans";
            this.typeList = "ban";
            this.dataGetter = async (profilesArray, limit, offset, dataAvaible) => {
                try {
                    let response = await this.$axios.get(`/profiles/${this.userID}/bans?limit=${limit}&offset=${offset}`, { headers: { 'Authorization': `${localStorage.token}` } });
                    if (response.data == null) {
                        dataAvaible = false;
                        return;
                    }
                    profilesArray.push(...response.data);
                } catch (e) {
                    this.errorMsg = this.$utils.errorToString(e);;
                }
            }
        },
        closeOptions() {
            setTimeout(() => {
                this.showOptions = false;
            }, 500);
        },
        async banUser() {
            try {
                let _ = await this.$axios.put(`/profiles/${localStorage.userID}/bans/${this.userID}`, {}, { headers: { 'Authorization': `${localStorage.token}` } });
                this.$router.push(`/profiles/${localStorage.userID}`);
            } catch (e) {
                this.errorMsg = this.$utils.errorToString(e);;
            }
            this.showOptions = false;
        },
        async deletePost(postID) {
            this.isLoading = true;
            try {
                let _ = await this.$axios.delete(`profiles/${localStorage.userID}/posts/${postID}`, { headers: { 'Authorization': `${localStorage.token}` } });
                this.posts = this.posts.filter(post => post.postID != postID);
                this.postsCount--;
                this.exitPost();
            } catch (e) {
                this.errorMsg = this.$utils.errorToString(e);;
            }
            this.isLoading = false;
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
        this.getPosts();
    },

}

</script>


<template>

    <div class="header"> 
        <Modal :show="followersModalIsVisible" @close="handleFollowersToggle" :users="this.followers">
            <template v-slot:header>
                <h3>Followers</h3>
            </template>
        </Modal>
        
        <Modal :show="followingModalIsVisible" @close="handleFollowingToggle" :users="this.followings">
            <template v-slot:header>
                <h3>Following</h3>
            </template>
        </Modal>
        <Modal :show="updateNameModalIsVisible" @close="handleUpdateNameToggle">
            <template v-slot:header>
                <h3>Update Username</h3>
            </template>
            <template v-slot:body>
                <form class="username-form">
                    <input type="text" v-model="this.username" placeholder="New username" />
                    <button type="submit" @click.prevent="updateUsername">Update</button>
                </form>
            </template>
        </Modal>
        <div class="top">
            <h1>{{ this.username }}
                <button v-if="isOwner" class="" @click="handleUpdateNameToggle">
                    <svg class="feather edit">
                        <use href="/feather-sprite-v4.29.0.svg#edit" />
                    </svg>
                </button>
            </h1>

            <div v-if="!isOwner">
                <button class="follow-btn" @click="toggleFollow">{{ isFollowing ? "Unfollow" : "Follow" }}</button>
                <button class="ban-btn" @click="toggleBan">{{ isBanned ? "Unban" : "Ban" }}</button>
            </div>
        </div>
        <div class="bottom"> 
            <h4>{{ ((this.postsCount > 1) ? " Posts: " : " Post: ") + this.postsCount }}</h4>
            <h4 @click="handleFollowingToggle">{{ " Followed: " + this.followingsCount }}</h4>
            <h4 @click="handleFollowersToggle">{{ (( this.followersCount > 1) ? " Followers: " : " Follower: ") + this.followersCount}}</h4>
        </div>
    </div>



    <!-- <div class="feed" v-for="post in this.posts" :key="post.ID">
        <Post :post="post" :identifier="this.identifier" @delete="() => deleteImage(post.ID)" />
    </div> -->


</template>