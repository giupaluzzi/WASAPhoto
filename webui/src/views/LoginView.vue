<script>
  export default {

    data: function() {
      return {
        errormsg: null,
        loggedUser: "",
      }
    },

    methods: {
      async login() {
        this.errormsg = null;
        try {
          let response = await this.$axios.post("/session", {userId: this.loggedUser.trim()});
          localStorage.setItem('auth', response.data.userId);
          this.$router.replace("/home")
        } catch (e) {
          this.errormsg = e.toString();
        }
      },
    },

    mounted() {
      if (localStorage.getItem('auth')){
        this.$router.replace("/home")
      }
    },

  }
</script>

<template>
  <div>
    <div
        class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
      <h1 class="h2">Login</h1>
      <div class="btn-toolbar mb-2 mb-md-0">
        <div class="btn-group me-2">
          <button type="button" class="btn btn-sm btn-outline-secondary" @click="login">
            Login / Register
          </button>
        </div>

      </div>
    </div>

    <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>

    <form @submit.prevent="login">
      <div class="mb-3">
        <label for="username" class="form-label">Username</label>
        <input type="text" class="form-control" v-model="loggedUser" maxlength="16" minlength="3" placeholder="Your username" />
      </div>

      <div class="col ">
        <button :disabled="loggedUser == null || loggedUser.length >16 || loggedUser.length <3 || loggedUser.trim().length<3">
          Register / Login
        </button>
      </div>

    </form>

  </div>
</template>

<style>
</style>