<script>
import ErrorMsg from "../components/ErrorMsg.vue";

export default {
  components: {ErrorMsg},
  data: function() {
    return {
      errormsg: null,
      userid: "",
    }
  },

  methods: {
    async login() {
      this.errormsg = null;
      try {
        let response = await this.$axios.post("/session", {userid: this.userid.trim()});
        // console.log("response:", response)
        // console.log("response.data:", response.data)
        // console.log("response.data.userid:", response.data.userid)
        localStorage.setItem("auth", response.data.userid);
        this.$emit("login", true)
        // console.log("userid:", response.data.userid)
        this.$router.replace("/home")
      } catch (e) {
        this.errormsg = e.toString();
      }
    },
  },

  mounted() {
    if (localStorage.getItem("auth")){
      // console.log("userid:", localStorage.getItem("auth"))
      this.$router.replace("/home")
    }
  },

}
</script>

<template>
  <div>
    <div
        class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
      <h2>Login</h2>
    </div>

    <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>

    <form @submit.prevent="login">
      <div class="mb-3">
        <label for="userid" class="form-label">Username</label>
        <input type="text" class="form-control" v-model="userid" maxlength="16" minlength="3" placeholder="Your username" />
      </div>

      <div class="col ">
        <button :disabled="userid == null || userid.length >16 || userid.length <3 || userid.trim().length<3">
          Register / Login
        </button>
      </div>

    </form>

  </div>
</template>

<style>
</style>