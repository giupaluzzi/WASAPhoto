<script setup>
import { RouterLink, RouterView } from 'vue-router'
</script>
<script>
export default {
  data() {
    return {
      logged: false,
    }
  },
  methods: {
    logout(value) {
      this.logged = value
      this.$router.replace("/login")
    },
    login(value) {
      this.logged = value
    },
  },

  created() {
    if (!localStorage.getItem('started')) {
      localStorage.clear()
      localStorage.setItem('started', true)
    }
  },

  mounted() {
    if (!localStorage.getItem('auth')) {
      this.$router.replace("/login")
    } else {
      this.logged = true
    }
  },
}
</script>

<template>
  <div class="container-fluid">
    <div class="row">
      <main class="col-md-9 ms-sm-auto col-lg-10 px-md-4">
        <Sidebar v-if="logged"
                 @logout="logout"/>
        <RouterView
                 @login="login"/>
      </main>
    </div>
  </div>
</template>

<style>
</style>
