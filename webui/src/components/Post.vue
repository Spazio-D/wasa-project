<script>

export default {
    props: ["post", "token"],
    data: function () {
        return {

            ownerID: this.post.user.id,
            username: this.post.user.username,
            postID: this.post.id,
            isLiked: this.post.likeCheck,
            image64: this.post.image,
            likesCount: this.post.likesCount,
            commentsCount: this.post.commentsCount,
            timestamp: this.post.timestamp,
            isOwner: false,
            errorMsg: "",

            comments: [],

            newComment: '',
            url: __API_URL__
        };
    },
    methods: {
        deleteImage() {
            window.confirm("Are you sure you want to delete this image?") &&
                this.$axios.delete("/photos?photoId=" + this.post.ID, { headers: { 'Authorization': this.token } })
                    .then((response) => {
                        console.log(response);
                        this.$emit('delete', this.post.ID);
                    })
                    .catch((error) => {
                        console.log(error);
                    })
        },
        toggleLike() {
            console.log(this.token)
            if (this.liked) {
                this.$axios.delete("/photos/" + this.post.ID+ "/likes" , { headers: { 'Authorization': this.token } })
                    .then((response) => {
                        console.log(response);
                        this.liked = !this.liked;
                    })
                    .catch((error) => {
                        console.log(error);
                    })
            }
            else {
                this.$axios.post("/photos/" + this.post.ID+ "/likes", {}, { headers: { 'Authorization': this.token } })
                    .then((response) => {
                        console.log(response);
                        this.liked = !this.liked;
                    })
                    .catch((error) => {
                        console.log(error);
                    })
            }
        },
        submitComment() {
            const payload = {
                comment: this.newComment
            };
            this.$axios.post('/comments?photoId=' + this.post.ID, payload, {
                headers: { 'Authorization': this.token }
            })
                .then((response) => {
                    console.log(response);
                    this.newComment = ''; // Reset comment input
                })
                .catch((error) => {
                    console.log(error);
                });
        },
        deleteComment(id) {
            console.log(this.post)
            this.$axios.delete('/comments?commentId=' + id, {
                headers: { 'Authorization': this.token }
            })
                .then((response) => {
                    console.log(response);
                })
                .catch((error) => {
                    console.log(error);
                });
        }
    }
}
</script>

<template>
    <div v-if="post != null" class="post-container">
        <div class="username-container">
            <RouterLink :to="{ path: 'user/' + post.user.id }" replace>
                {{ post.user.username }}
            </RouterLink>
            <button class="" @click="deleteImage" v-if=isOwner>
                <svg class="feather" :class="this.liked ? 'liked' : ''" >
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
                    <svg class="feather" :class="this.liked ? 'liked' : ''">
                        <use href="/feather-sprite-v4.29.0.svg#heart" />
                    </svg>
                </button>
                {{ this.likesCount + ((this.likesCount > 1) ? " likes" : " like") }}
            </span>
            <span class="date">
                {{ new Date(post.UploadTime).toLocaleString('default', {
                    month: 'long',
                    year: 'numeric',
                    hour: '2-digit',
                    day: '2-digit',
                    hour12: true // change to false if you want 24-hour format
                }) }}
            </span>
        </div>
        <div class="comments-container">
            <ul>
                <li class="comment" v-for="comment in this.comments" :key="comment.ID">
                    <div class="comment-top">

                        <div>
                            <strong>{{ comment.UserName }}</strong> - <small>{{ new
                                Date(comment.CreateTime).toLocaleString()
                            }}</small>
                        </div>
                        <button class="" @click="() => deleteComment(comment.ID)" v-if="comment.UserID == this.token">
                            <svg class="feather">
                                <use href="/feather-sprite-v4.29.0.svg#trash" />
                            </svg>
                        </button>
                    </div>
                    {{ comment.Text }}

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