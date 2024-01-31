<script>
import {RouterLink} from "vue-router";

export default {
  components: {RouterLink},
  methods: {
    logout() {
      localStorage.removeItem('auth');
      this.$emit("logout", false)
    },
  },
};
</script>

<template>
  <div class="container-fluid">
    <div class="row">
      <header class="navbar navbar-dark sticky-top bg-dark flex-md-nowrap p-0 shadow">
        <a class="navbar-brand col-md-3 col-lg-2 me-0 px-3 fs-6" href="#/home">WASAPhoto</a>
        <button class="navbar-toggler position-absolute d-md-none collapsed" type="button" data-bs-toggle="collapse" data-bs-target="#sidebarMenu" aria-controls="sidebarMenu" aria-expanded="false" aria-label="Toggle navigation">
          <span class="navbar-toggler-icon"></span>
        </button>
      </header>
      <nav id="sidebarMenu" class="col-md-3 col-lg-2 d-md-block bg-light sidebar collapse">
        <div class="position-sticky pt-3 sidebar-sticky">
          <h6 class="sidebar-heading d-flex justify-content-between align-items-center px-3 mt-4 mb-1 text-muted text-uppercase">
            <span>General</span>
          </h6>
          <ul class="nav flex-column">
            <li class="nav-item">
              <RouterLink to="/home" class="nav-link">
                <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#home"/></svg>
                Home
              </RouterLink>
            </li>
            <li class="nav-item">
              <RouterLink to="/search" class="nav-link">
                <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#search"/></svg>
                Search
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