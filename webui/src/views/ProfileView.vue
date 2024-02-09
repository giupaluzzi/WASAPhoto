<script>
import ErrorMsg from "../components/ErrorMsg.vue";

export default{
  components: {ErrorMsg},
  data: function() {
    return {
      errormsg: null,
      userid: localStorage.getItem("auth"),

      newuserid: "",
      editing: false,

      followers: [],
      following: [],
      photos: [],

      photoToUpload: null,
      postCounter: 0,

    }
  },

  methods: {

    editUserid() {
      this.editing = !this.editing
    },

    cancelEdit(){
      this.editing = !this.editing
      this.newuserid = ""
    },

    async setUsername(){
      this.loading = true
      this.errormsg = null
      try{
        let response = await this.$axios.put("/users/" +this.userid+"/username", {newuserid: this.newuserid});
        this.userid = response.data.newuserid
        localStorage.setItem("auth", this.userid)
        this.editing = false
        this.newuserid = ""
      } catch (e) {
        this.errormsg = e.toString()
      }
      this.loading = false
    },

    async uploadPhoto(){
      this.loading = true
      this.errormsg = null
      try{
        let response = await this.$axios.post("/users/" + this.userid + "/photos/", this.photoToUpload)
        // console.log("photoToUpload: ", response.data)
        this.photos.unshift(response.data)
        this.postCounter += 1
      } catch(e) {
        this.errormsg = e.toString()
      }
      this.loading = false
    },

    changeFile() {
      this.photoToUpload = this.$refs.inputFile.files[0];
      // console.log("photoToUpload: ", this.$refs.inputFile.files[0])

    },

    getImg(data) {
      return `data:image/jpeg;base64,${data}`
    },

    async deletePhoto(photoId){
      this.loading = true
      this.errormsg = null
      try{
        let response = await this.$axios.delete("/users/"+this.userid+"/photos/"+photoId)
        this.photos = this.photos.filter(photo => photo.photoid !== photoId);
        this.postCounter -=1
      } catch(e) {
        this.errormsg = e.toString()
      }
      this.loading = false
    },

    async getProfile(){
      this.loading = true
      this.errormsg = null
      try {
        let response = await this.$axios.get("/users/" + this.userid)
        this.followers = response.data.followers != null ? response.data.followers : []
        this.following = response.data.following != null ? response.data.following : []
        this.photos = response.data.photos != null ? response.data.photos : []
        this.photos.forEach(photo => {
          photo.likes = photo.likes != null ? photo.likes : [];
          photo.comments = photo.comments != null ? photo.comments : [];
        })
        this.postCounter = response.data.photos != null ? response.data.photos.length : 0
      } catch(e){
        this.errormsg = e.toString();
        this.loading = false
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
      <h2>{{userid}}
        <svg class="feather" @click="editUserid"> <use href="/feather-sprite-v4.29.0.svg#edit"/></svg>
      </h2>
      <h6>
        <input type="text" v-model="newuserid" v-if="editing" />
        <button v-if="editing" @click="setUsername">Apply</button>
        <button v-if="editing" @click="cancelEdit">Cancel</button>
      </h6>
      <hr>
      <h4>Followers: {{followers.length}}</h4>
      <h4>Following: {{following.length}}</h4>
    </div>
    <hr>
    <div>
      <div>
      <h4>Posts: {{postCounter}}</h4>
        <input type="file" ref="inputFile" accept=".jpg, .png" @change="changeFile"/>
        <button type="button" :disabled="!photoToUpload" @click="uploadPhoto">
          <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#upload"/></svg>
        </button>
      </div>
      <div v-if="photos && photos.length >0" class="grid">
        <ul class="post-list">
          <li v-for="p in photos" :key="p.photoid" class="grid-item">
            <div class="image-container">
              <img :src="getImg(p.file)" alt="User's photo" class="resizable-image" />
              <div class="post-details">
                <b>Uploading date: </b>
                  {{p.date}}
                <br>
                <b>Likes: </b>
                  {{p.likes !== null ? p.likes.length : 0 }}
                <br>
                <b>Comments: </b>
                  {{ p.comments !== null && p.comments.length > 0 ? p.comments : 'No comments yet' }}
              </div>
              <button type="button" @click="deletePhoto(p.photoid)">
                Delete
              </button>
            </div>
          </li>
        </ul>
      </div>
    </div>
    <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
  </div>
</template>

<style>

.grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
  gap: 20px;
  margin-bottom: 30px;
}

.grid-item {
  border: 1px solid #ccc;
  padding: 10px;
  margin-right: 30px;
}

.post-list {
  list-style: none;
  padding: 0;
  display: flex;
  flex-direction: row;
}

.resizable-image {
  max-width: 100%;
  max-height: 100%;
}

.image-container {
  display: flex;
  flex-direction: column;
  height: 100%;
}

.post-details {
  margin-top: auto;
}
</style>