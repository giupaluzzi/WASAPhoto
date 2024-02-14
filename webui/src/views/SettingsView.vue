<script>
import ErrorMsg from "../components/ErrorMsg.vue";

export default{
  components: {ErrorMsg},

  data: function() {
    return {
      errormsg: null,
      newuserid: null,
    }
  },

  methods: {
    async setUsername() {
      this.errormsg = null
      try{
        await this.$axios.put("/users/"+localStorage.getItem("auth")+"/username", {userid: this.newuserid})
        localStorage.setItem("auth", this.newuserid)
        this.newuserid = ""
      } catch(e) {
        this.errormsg = e.toString()
      }
    }
  },
}
</script>

<template>
  <div>
    <h3>
      Settings
    </h3>
    <br>
    <h5>
      Change your Username
      <input type="text" placeholder="Enter new Username" v-model="newuserid" />
      <button @click="setUsername" :disabled="newuserid == null || newuserid.length >16 || newuserid.length <3 || newuserid.trim().length<3">
        Apply
      </button>
    </h5>

    <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>

  </div>
</template>

<style>

</style>