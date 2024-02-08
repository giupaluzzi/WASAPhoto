<script>
export default{
  data: function() {
    return {
      errormsg: null,
      userid: localStorage.getItem("auth"),

      followers: [],
      following: [],
      photos: [],

      photoToUpload: null,
    }
  },

  methods: {

    async uploadPhoto(){
      this.errormsg = null
      try{
        let response = await this.$axios.post("/users/" + this.userid + "/photos/", this.photoToUpload)
        console.log("photoToUpload: ", response.data)
        this.photos.unshift(response.data)
      } catch(e) {
        this.errormsg = e.toString()
      }
    },

    changeFile() {
      this.photoToUpload = this.$refs.inputFile.files[0];
      console.log("photoToUpload: ", this.$refs.inputFile.files[0])

    },

    async getProfile(){
      this.errormsg = null
      try {
        let response = await this.$axios.get("/users/" + this.userid)
        this.followers = response.data.followers
        this.following = response.data.following
        this.photos = response.data.photos
      } catch(e){
        this.errormsg = e.toString();
      }
    },
  },

  mounted(){
    this.userid = localStorage.getItem("auth")
    // console.log('userid:', this.userid)
    this.getProfile()
  }
}
</script>

<template>
  <div>
    <div>
      <h2>{{this.userid}}</h2>
      <h4>Followers: {{this.followers.length}}</h4>
      <h4>Following: {{this.following.length}}</h4>
    </div>
    <div>
      <h4>Posts</h4>
      <input type="file" ref="inputFile" accept=".jpg, .png" @change="changeFile"/>
      <button type="button" v-if="photoToUpload" @click="uploadPhoto">
        <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#upload"/>Upload</svg>
      </button>
    </div>

    <div v-if="photos && photos.length >0">
      <div v-for="photo in photos" :key="photo.photoid">
        <img :src="photo.file" alt="User's photo">
      </div>
    </div>

    <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
  </div>
</template>

<style>
</style>