<script>

export default {
    props: ["post"],
    emits: ["delete-post"],
    data: function () {
        return {

            ownerID: this.post.user.id,
            username: this.post.user.username,
            postID: this.post.id,
            isLiked: this.post.likeCheck,
            image64: this.post.image,
            likesCount: this.post.likesCount,
            commentsCount: this.post.commentsCount,
            comments: this.post.comments,
            timestamp: this.post.timestamp,
            isOwner: false,
            errorMsg: "",
            newComment: '',
  
        };
    },
    methods: {
        deleteImage() {
            window.confirm("Are you sure you want to delete this image?");
            this.$emit('delete-post', this.postID);
        },
        async toggleLike() {
            if (this.isLiked) {
                try{
                    let _ = await this.$axios.delete(`users/${this.ownerID}/posts/${this.postID}/likes/${sessionStorage.userID}`, { headers: { "Authorization": `${sessionStorage.token}` } });
                    this.isLiked = false;
                    this.likesCount--;
                }catch(e){
                    this.errorMsg = e.toString();
                }
            }
            else {
                try{
                    let _ = await this.$axios.put(`users/${this.ownerID}/posts/${this.postID}/likes/${sessionStorage.userID}`, {}, { headers: { "Authorization": `${sessionStorage.token}` } });
                    this.isLiked = true;
                    this.likesCount++;
                }catch(e){
                    this.errorMsg = e.toString();
                }
            }
        },
        async submitComment() {
            const payload = { text: this.newComment };
            try{
                let response = await this.$axios.post(`users/${this.ownerID}/posts/${this.postID}/comments`, payload, { headers: { "Authorization": `${sessionStorage.token}` } });
                this.newComment = '';
                this.comments.push(response.data)
            }catch(e){
                this.errorMsg = e.toString();
            }
        },
        async deleteComment(id) {
            try{
                let _ = await this.$axios.delete(`users/${this.ownerID}/posts/${this.postID}/comments/${id}`, { headers: { "Authorization": `${sessionStorage.token}` } });
                this.comments = this.comments.filter(comment => comment.id != id);
            }catch(e){
                this.errorMsg = e.toString();
            }
        }
    },
    mounted() {
        if (this.ownerID == sessionStorage.userID) {
            this.isOwner = true;
        }
        
    },
}
</script>

<template>
    <div v-if="post != null" class="post-container">
        <div class="username-container">
            <RouterLink :to="{ path: '' + post.user.id }" replace>
                {{ post.user.username }}
            </RouterLink>
            <button class="" v-if="isOwner" @click="deleteImage" >
                <svg class="feather" :class="this.isLiked ? 'liked' : ''" >
                    <use href="/feather-sprite-v4.29.0.svg#trash" />
                </svg>
            </button>
        </div>
        <div class="image-container">
            <img :src="`data:image/jpg;base64,${image64}`">
        </div>
        <div class="likes-container">
            <span>
                <button class="like-btn" @click="toggleLike">
                    <svg class="feather" :class="this.isLiked ? 'liked' : ''">
                        <use href="/feather-sprite-v4.29.0.svg#heart" />
                    </svg>
                </button>
                {{ this.likesCount + ((this.likesCount > 1) ? " likes" : " like") }}
            </span>
            <span class="date">
                {{ new Date(post.timestamp).toLocaleString('default', {
                    year: 'numeric',
                    month: 'long',
                    day: '2-digit', 
                    hour: '2-digit',
                    minute: '2-digit',
                    hour12: false
                }) }}
            </span>
        </div>
        <div class="comments-container">
            <ul>
                <li class="comment" v-for="comment in this.comments" :key="comment.id">
                    <div class="comment-top">
                        <div>
                            <strong>{{ comment.user.username }}</strong> - <small>{{ new
                                Date(comment.timestamp).toLocaleString('default', {
                                year: 'numeric',
                                month: 'long',
                                day: '2-digit', 
                                hour: '2-digit',
                                minute: '2-digit',
                                hour12: false})
                            }}</small> 
                        </div>
                        <button class="" @click="() => deleteComment(comment.id)" v-if="isOwner">
                            <svg class="feather">
                                <use href="/feather-sprite-v4.29.0.svg#trash" />
                            </svg>
                        </button>
                    </div>
                    {{ comment.text }}

                </li>
            </ul>
            <form class="comment-input-container" @submit.prevent="submitComment">
                <input type="text" v-model="newComment" placeholder="Comment" />
                <button type="submit">Post</button>
            </form>
        </div>
    </div>
</template>

<style></style>