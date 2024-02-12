<script>
import ErrorMsg from "../components/ErrorMsg.vue";
import Photo from "../components/Photo.vue";

export default {
  components: {Photo, ErrorMsg},

  data: function() {
    return {
      errormsg: null,

      isBanned: false,
      isFollowed: false,

      followersCounter: 0,
      postCounter:  0,

      followers: [],
      following: [],

    }
  },

  methods:{
    // Follow/Unfollow method
    async follow() {
      this.errormsg = null
      try {
        if (this.isFollowed) {
          await this.$axios.delete("/users/" + localStorage.getItem("auth") + "/following/" + this.$route.params.id)
          this.followersCounter -= 1
        } else {
          await this.$axios.put("/users/" + localStorage.getItem("auth") + "/following/" + this.$route.params.id)
          this.followersCounter += 1
        }
        this.isFollowed = !this.isFollowed
      } catch (e) {
        this.errormsg = e.toString()
      }
    },

    // Ban/Unban method
    async ban() {
      this.errormsg = null
      try {
        if (this.isBanned) {
          await this.$axios.delete("/users/" + localStorage.getItem("auth") + "/banned/" + this.$route.params.id)
          this.getProfile()
        } else {
          await this.$axios.put("/users/" + localStorage.getItem("auth") + "/banned/" + this.$route.params.id)
          this.isFollowed = false
        }
        this.isBanned = !this.isBanned
      } catch (e) {
        this.errormsg = e.toString()
      }
    },

    // Method to retrieve information about an user from db
    async getProfile(){
      this.errormsg = null
      try {
        let response = await this.$axios.get("/users/" + this.$route.params.id)
        this.followers = response.data.followers != null ? response.data.followers : []
        this.following = response.data.following != null ? response.data.following : []
        this.followersCounter = response.data.followers != null ? response.data.followers : []
        this.isFollowed = response.data.followers != null ? response.data.followers.find(f => f.userid === localStorage.getItem("auth")) : false
        this.photos = response.data.photos != null ? response.data.photos : []
        this.postCounter = response.data.photos != null ? response.data.photos.length : 0
      } catch(e){
        this.errormsg = e.toString();
      }
    },
  },

  async mounted(){
    await this.getProfile()
  }
}

</script>

<template>
  <div>
    <div>
      <!--  Logged User's info  -->
      <h2>{{this.$route.params.id}}</h2>
      <h6>
        <!--  Follow/Unfollow -->
        <button v-if="!isBanned" @click="follow">
          {{isFollowed ? "Unfollow" : "Follow"}}
        </button>

        <!-- Ban/Unban  -->
        <button @click="ban">
          {{isBanned ? "Unban" : "Ban"}}
        </button>
      </h6>
      <div v-if="!isBanned">
        <hr>
        <h4>Followers: {{followers.length}}</h4>
        <h4>Following: {{following.length}}</h4>
      </div>
      <hr>
    </div>
    <div v-if="!isBanned">
      <h4>Posts: {{postCounter}}</h4>

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
                   :isOwner=false
            />

          </div>
        </div>
      </div>
    </div>
    <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
  </div>
</template>

<style>
</style>