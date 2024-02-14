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
      newCommentText: "",
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
          this.TotalLikes = this.TotalLikes.filter(userid => userid !== localStorage.getItem("auth"))
        } else {
          await this.$axios.put("/users/" + this.userid + "/photos/" + this.photoid + "/likes/" + localStorage.getItem("auth"))
          this.TotalLikes.push(localStorage.getItem("auth"))
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

    // Comment/Uncomment methods
    async postComment() {
      this.errormsg = null
      if (!this.newCommentText.trim()) return
      try {
        let response = await this.$axios.post("/users/" + this.userid + "/photos/" + this.photoid + "/comments/",
            {photoid: this.photoid, userid: localStorage.getItem("auth"), commentText: this.newCommentText})
        this.TotalComments.push(response.data)
        this.newCommentText = ""
      } catch (e) {
        this.errormsg = e.toString()
      }
    },

    async deleteComment(commentid) {
      this.errormsg = null
      try {
        await this.$axios.delete("/users/" + localStorage.getItem("auth") + "/photos/" + this.photoid + "/comments/" + commentid)
        this.TotalComments = this.TotalComments.filter(comment => comment.commentid !== commentid)
      } catch (e) {
        this.errormsg = e.toString()
      }
    },

  },

    mounted() {
      if (this.likes != null) {
        this.TotalLikes = this.likes
        this.isLiked = this.TotalLikes.includes(localStorage.getItem("auth"))
      }

      if (this.comments != null) {
        this.TotalComments = this.comments
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

            <i @click="toggleLike" :class="{ 'like-icon': true, 'active': isLiked }">&#x2665;</i>

            <br>

            <b>Comments: </b>
              <ul v-if="TotalComments.length > 0">
                <li v-for="comment in TotalComments" :key="comment.commentid">
                  <div>
                    {{comment.userid}}: {{ comment.commentText }}
                    <button v-if="isOwner" @click="deleteComment(comment.commentid)">
                      <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#trash-2"/></svg>
                    </button>
                  </div>
                </li>
              </ul>

            <form @submit.prevent="postComment">
              <textarea v-model="newCommentText" placeholder="Add a comment"></textarea>
              <button type="submit">
                Post
              </button>
            </form>

            <hr>
            <button type="button" v-if="isOwner" @click="deletePhoto(photoid)" class="delete-button">
              Delete
            </button>
          </div>

          <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>

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

.like-icon {
  cursor: pointer;
  font-size: 15px;
}

.like-icon.active {
  color: red;
}

</style>