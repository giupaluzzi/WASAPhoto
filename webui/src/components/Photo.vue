<script>
import ErrorMsg from "./ErrorMsg.vue";

export default {
  components: {ErrorMsg},
  data: function(){
    return{
      errormsg: null,
      TotalLikes: [],
      TotalComments: [],
      isLiked: false,
    }
  },

  props: ['userid','photoid','likes','comments','date','file','isOwner'],

  methods: {
    // Like/Unlike method

    async toggleLike() {
      this.errormsg = null

      if (this.isOwner) {
        return
      }

      try {
        if (this.isLiked) {
          await this.$axios.delete("/users/" + this.userid + "/photos/" + this.photoid + "/likes/" + localStorage.getItem("auth"))
          this.TotalLikes.push(localStorage.getItem("auth"))
        } else {
          await this.$axios.put("/users/" + this.userid + "/photos/" + this.photoid + "/likes/" + localStorage.getItem("auth"))
          this.TotalLikes.pop(localStorage.getItem("auth"))
        }
        this.isLiked = !this.isLiked
      } catch (e) {
        this.errormsg = e.toString()
      }
    },

    // DeletePhoto Method
    async deletePhoto(){
      this.errormsg = null
      try{
        await this.$axios.delete("/users/"+localStorage.getItem("auth")+"/photos/"+this.photoid)
        this.$emit("deletePost",this.photoid)
      } catch(e) {
        this.errormsg = e.toString()
      }
    },

    getImg(data) {
      return `data:image/jpeg;base64,${data}`
    },
/*
    // Comment/Uncomment methods
    async addComment(photo) {
      this.errormsg = null
      if (!this.newCommentText.trim()) return
      try {
        let response = await this.$axios.post("/users/" + this.$route.params.id + "/photos/" + photo.photoid + "/comments/")
        photo.comments.push(response.data)
        this.newCommentText = ""
      } catch (e) {
        this.errormsg = e.toString()
      }
    },

    async deleteComment(photo, commentid) {
      this.errormsg = null
      try {
        let response = await this.$axios.delete("/users/" + localStorage.getItem("auth") + "/photos/" + photo.photoid + "/comments/" + commentid)
        photo.comments = photo.comments.filter(c => c.id !== commentid);
      } catch (e) {
        this.errormsg = e.toString()
      }
    },
*/
  },

    mounted() {
      if (this.likes != null) {
        this.TotalLikes = this.likes
        this.isLiked = this.TotalLikes.includes(localStorage.getItem("auth"))
      }

      if (this.comments != null) {
        this.TotalComments = this.TotalLikes
      }
    }

}
</script>

<template>

        <div class="image-container">
          <img :src="getImg(file)" alt="User's photo" class="resizable-image" />
          <div class="photo-details">

            <b>From: </b>
            {{userid}}
            <br>

            <b>Uploading date: </b>
            {{date}}
            <br>

            <b>Likes:</b>
              {{TotalLikes.length}}
              <i @click="toggleLike" :class="'fa ' + (isLiked ? 'fa-heart' : 'fa-heart-o')"></i>
            <br>

            <b>Comments: </b>
              {{TotalComments.length}}

            <hr>
            <button type="button" v-if="isOwner" @click="deletePhoto(photoid)" class="delete-button">
              Delete
            </button>
          </div>
        </div>

</template>

<style>
.resizable-image {
  max-width: 100%;
  max-height: 400px;
  width: auto;
  height: auto;
}

.image-container {
  display: flex;
  flex-direction: column;
  align-items: center;
}

.photo-details {
  margin-top: auto;
}

.delete-button {
  width: 100%;
}
</style>