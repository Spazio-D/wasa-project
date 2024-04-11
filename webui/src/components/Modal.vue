<script>
import { RouterLink } from 'vue-router';

export default {
  props: {
    show: Boolean,
    title: String,
    users: Array,
  },

  data() {
    return {
      searchText: '',
      token: sessionStorage.getItem('token'),
    };
  },
  computed: {
    filteredUsers() {
      // return []
      return this.users.filter(user =>
        user.Name.toLowerCase().includes(this.searchText.toLowerCase())
      );
    }
  },
  components: { RouterLink }
}
</script>

<template>
  <Transition name="modal">
    <div v-if="show" class="modal-mask">
      <div class="modal-wrapper">
        <div class="modal-container">
          <div class="modal-header">
            <slot name="header">default header</slot>
            <button class="like-btn" @click="$emit('close')">
              <svg class="feather">
                <use href="/feather-sprite-v4.29.0.svg#x" />
              </svg>
            </button>
          </div>

          <div class="modal-body">
            <!-- loop through users and create a RouterLink for each user -->
            <slot name="body">
              <div class="search-input">
                <input type="text" v-model="searchText" placeholder="Search" />
              </div>
              <div class="search-results">


                <div v-for="user in filteredUsers" :key="user.ID" @click="$emit('close')">
                  <router-link :to="{ name: 'profile', params: { userId: user.ID } }" replace force>
                  <!-- <router-link :to="{ name: 'profile', params: { userId: user.ID } }" replace force> -->
                    <div class="user">

                      <p>{{ user.Name }}</p>
                    </div>
                  </router-link>
                  <!-- <RouterLink :to="{ name: 'Profile', params: { id:user.ID}}">
                  <div class="user">
                    
                    <h3>{{ user.Name }}</h3>
                  </div>
                </RouterLink> -->
                </div>
              </div>
            </slot>
          </div>


        </div>
      </div>
    </div>
  </Transition>
</template>

<style>
.modal-mask {
  position: fixed;
  z-index: 9998;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5);
  display: table;
  transition: opacity 0.3s ease;
}

.modal-wrapper {
  display: table-cell;
  vertical-align: middle;
}

.modal-container {
  width: 300px;
  margin: 0px auto;
  /* padding: 20px 15px 0 15px; */
  background-color: #fff;
  border-radius: 2px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.33);
  transition: all 0.3s ease;
}

.modal-header {
  padding: 20px 15px 10px 15px;
}
.modal-header h3 {
  margin-top: 0;
  font-size: 19px;
  color: #42b983;
}

.modal-header button {
  color: rgb(86, 86, 86);
  background: none;
  border: none;
  padding: 5px;
  line-height: 12px;
  font-size: 15px;
}

.modal-header button svg {
  width: 20px;
  height: 20px;
}

.modal-body {
  /* margin: 20px 0; */
  /* border: 1px solid #eee; */
}

.search-input {
  padding: 0 15px;
  /* border-bottom:
   1px solid #eee; */
  
}

.search-input input {
  width: 100%;
  outline: none;
  /* border: none */
  border-radius: 3px;
  border:1px solid rgb(179, 179, 179)
}

.search-results {
  padding: 10px 15px;
  border-bottom: 1px solid #eee;
  cursor: pointer;
  max-height: 200px;
  overflow-y: scroll;
}

.modal-default-button {
  float: right;
}

.username-form {
  display: flex;
  flex-direction: column;
  padding: 0 15px;
}

.username-form input {
  margin-bottom: 10px;
  margin-top:5px;
  outline: none;
  border-radius: 3px;
  border:1px solid rgb(179, 179, 179)
  /* border: none; */
}

.username-form button {
  margin-bottom: 15px;
}
</style>