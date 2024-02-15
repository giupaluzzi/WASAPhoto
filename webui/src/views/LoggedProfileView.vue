<script>
  import ErrorMsg from "../components/ErrorMsg.vue";
  import Photo from "../components/Photo.vue";

  export default{
    components: {Photo, ErrorMsg},
    data: function() {
      return {
        errormsg: null,
        userid: localStorage.getItem("auth"),

        followers: [],
        following: [],
        photos: [],

        photoToUpload: null,
        postCounter: 0,

      }
    },

    methods: {
      // UploadPhoto Methods
      async uploadPhoto(){
        this.errormsg = null
        try{
          let response = await this.$axios.post("/users/" + localStorage.getItem("auth") + "/photos/", this.photoToUpload)
          // console.log("photoToUpload: ", response.data)
          this.photos.unshift(response.data)
          this.postCounter += 1
        } catch(e) {
          this.errormsg = e.toString()
        }
      },

      changeFile() {
        this.photoToUpload = this.$refs.inputFile.files[0];
        // console.log("photoToUpload: ", this.$refs.inputFile.files[0])
      },

      deletePost(photoid){
          this.photos = this.photos.filter(item => item.photoid !== photoid)
          this.postCounter = this.photos.length
        },

      // Method to retrieve information about logged user from db
      async getProfile(){
        this.errormsg = null
        try {
          let response = await this.$axios.get("/users/" + localStorage.getItem("auth"))
          this.followers = response.data.followers != null ? response.data.followers : []
          this.following = response.data.following != null ? response.data.following : []
          this.photos = response.data.photos != null ? response.data.photos : []
          this.postCounter = response.data.photos != null ? response.data.photos.length : 0
        } catch(e){
          this.errormsg = e.toString();
          }
      },
    },

    async mounted(){
      // console.log('userid:', this.userid)
      await this.getProfile()

    }
  }
</script>

<template>
  <div>
    <div>
      <!--  Logged User's info  -->
      <h2>{{userid}}</h2>
      <hr>
      <h4>Followers: {{followers.length}}</h4>
      <h4>Following: {{following.length}}</h4>
      <hr>
    </div>

    <h4>Posts: {{postCounter}}</h4>
      <!--  Upload photo  -->
      <input type="file" ref="inputFile" accept=".jpg, .png" @change="changeFile"/>
      <button type="button" :disabled="!photoToUpload" @click="uploadPhoto">
        <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#upload"/></svg>
      </button>

    <!--  PhotoList  -->
    <div class="row">
      <div class="col">

        <div v-if="postCounter>0">
          <Photo v-for="(photo,index) in photos"
                 :key="index"
                 :userid="photo.userid"
                 :photoid="photo.photoid"
                 :comments="photo.comments"
                 :likes="photo.likes"
                 :date="photo.date"
                 :file="photo.file"
                 :isOwner=true
                 @deletePhoto="deletePost"
          />

        </div>
      </div>
    </div>
    <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
  </div>
</template>

<style>
</style>