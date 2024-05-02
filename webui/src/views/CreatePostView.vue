<script>
export default {
    data() {
        return {
            image: null,
            errorMsg: "",
        };
    },
    emits: ['login-success'],
    methods: {
        async handleFileChange() {
            this.errorMsg = "";
            const file = this.$refs.file.files[0];
            if (file.type !== "image/jpeg") {
                this.errorMsg = "File type not supported, only jpg and jpeg are allowed";
                return
            }
            if (file.size > 5242880) {
                this.errorMsg = "File size is too big. Max size is 5MB";
                return
            }
            this.image = await this.convertFileToBase64(file);
        },
        async submitPost() {
            if (!this.image) {
                alert("Please select a file.");
                return;
            }

            const formData = new FormData();
            formData.append("image", this.image.split(',')[1]);

            try {
                await this.$axios.post(`users/${sessionStorage.token}/posts`, formData, { headers: {'Authorization': `${sessionStorage.token}`, 'Content-Type': 'multipart/form-data'}});
                this.$router.push(`user/${sessionStorage.token}`);
            } catch (e) {
                this.errorMsg = e.toString();
            }
        },
        convertFileToBase64(file) {
            return new Promise((resolve, reject) => {
                const reader = new FileReader();
                reader.readAsDataURL(file);
                reader.onload = () => resolve(reader.result);
                reader.onerror = error => reject(error);
            });
        }
    }
}
</script>

<template>
    <div class="create-post-container">
        <form @submit.prevent="submitPost">
            <h2>Create Post</h2>
            <ErrorMsg v-if="errorMsg" :msg="errorMsg"></ErrorMsg>
            <input type="file" ref="file" accept=".jpg,.jpeg" @change="handleFileChange" />
            <button type="submit">Submit</button>
        </form>
    </div>
</template>
  
  
  
<style>
.create-post-container {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100vh;
    padding: 20px;
    background-color: #f4f4f4;
}

.create-post-container form {
    display: flex;
    flex-direction: column;
    align-items: center;
    padding: 20px;
    background: white;
    border-radius: 8px;
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
    width: 300px;
}

.create-post-container input[type="file"] {
    margin: 15px 0;
    padding: 10px;
    border: 1px solid #ccc;
    border-radius: 5px;
    width: 100%;
}

.create-post-container button {
    padding: 10px 20px;
    border: none;
    border-radius: 5px;
    background-color: #42b983;
    color: white;
    cursor: pointer;
    width: 100%;
    margin-top: 10px;
}

.create-post-container button:hover {
    background-color: #269261;
}
</style>