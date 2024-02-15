<script>
import {RouterLink} from "vue-router";

export default {
  data (){
    return {
      searchUserId: ""
    }
  },
  components: {RouterLink},
  methods: {
    logout() {
      localStorage.removeItem('auth');
      this.$emit("logout", false)
    },

    search() {
      if (this.searchUserId.trim() !== "") {
        this.$router.push("/users/"+this.searchUserId)
      }
    },
  }
};
</script>

<template>
  <div class="container-fluid">
    <div class="row">
      <nav id="sidebarMenu" class="col-md-3 col-lg-2 d-md-block bg-light sidebar collapse">
        <div class="position-sticky pt-3 sidebar-sticky">
          <ul class="nav flex-column">
            <li class="nav-item">
              <form class="d-flex" @submit.prevent="search">
                <input class="form-control me-2" type="search" placeholder="Enter UserID" v-model="searchUserId">
                <button class="btn btn-outline-primary" type="submit">Search</button>
              </form>
            </li>
            <li class="nav-item">
              <RouterLink to="/home" class="nav-link">
                <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#home"/></svg>
                Home
              </RouterLink>
            </li>

            <li class="nav-item">
              <RouterLink to="/profile" class="nav-link">
                <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#user"/></svg>
                My Profile
              </RouterLink>
            </li>

            <li class="nav-item">
              <RouterLink to="/profile/settings" class="nav-link">
                <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#settings"/></svg>
                Settings
              </RouterLink>
            </li>

            <li class="nav-item">
              <RouterLink to="/login" class="nav-link" @click="logout">
                <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#log-out"/></svg>
                Log out
              </RouterLink>
            </li>

          </ul>
        </div>
      </nav>
    </div>
  </div>
</template>

<style>

</style>