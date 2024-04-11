<script>

function userHasLiked(post, token) {
    return post.Likes.some(like => like.UserID == token);
}

export default {
    props: ["post", "token"],
    data: function () {
        return {
            liked: userHasLiked(this.post, this.token),
            newComment: '',
            localPost: JSON.parse(JSON.stringify(this.post)),
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
                        this.localPost.Likes = response.data;
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
                        this.localPost.Likes = response.data;
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
                    this.localPost.Comments.push(response.data);
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
                    this.localPost.Comments = this.localPost.Comments.filter(comment => comment.ID != id);
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
            <RouterLink :to="{ path: 'user/' + post.UserID }" replace>
                {{ post.UserName }}
                <!-- <p>{{ this.test}}</p> -->
            </RouterLink>
            <button class="" @click="deleteImage" v-if="post.UserID == this.token">
                <svg class="feather" :class="this.liked ? 'liked' : ''">
                    <use href="/feather-sprite-v4.29.0.svg#trash" />
                </svg>
            </button>
        </div>
        <div class="image-container">
            <img :src="this.url + '/images/' + post.FileName" :alt="post.FileName">
        </div>
        <div class="likes-container">
            <span>
                <button class="like-btn" @click="toggleLike">
                    <svg class="feather" :class="this.liked ? 'liked' : ''">
                        <use href="/feather-sprite-v4.29.0.svg#heart" />
                    </svg>
                </button>
                {{ this.localPost.Likes.length + ((this.localPost.Likes.length > 1) ? " likes" : " like") }}
            </span>
            <span class="date">
                <!--  display the time with month day year and only the hour-->
                {{ new Date(post.UploadTime).toLocaleString('default', {
                    month: 'long',
                    year: 'numeric',
                    hour: '2-digit',
                    day: '2-digit',
                    hour12: true // change to false if you want 24-hour format
                }) }}
                <!-- {{ new Date(post.UploadTime).toLocaleDateString()  }} -->
            </span>
        </div>
        <div class="comments-container">
            <ul>
                <li class="comment" v-for="comment in this.localPost.Comments" :key="comment.ID">
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