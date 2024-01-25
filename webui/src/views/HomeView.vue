<script>
export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
			stream: [],
      maxVisibleComments: 5,
      liked: false,
		}
	},
	methods: {
		async refresh() {
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.get("/home");
				this.some_data = response.data;
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
		},
    async getStream() {
      try{
        let response = await this.$axios.get("/users"+localStorage.getItem('auth')+"/stream")
        this.stream = response.data;
      } catch(e){
        this.errormsg = e.toString();
      }
    },
    toggleLike() {
      this.liked = !this.liked;
    },
	},
	mounted() {
		this.refresh()
    this.getStream()
	}
}
</script>

<template>
	<div>
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2">Home page</h1>
			<div class="btn-toolbar mb-2 mb-md-0">
				<div class="btn-group me-2">
					<button type="button" class="btn btn-sm btn-outline-secondary" @click="refresh">
						Refresh
					</button>
				</div>
      </div>
		</div>

    <div>
      <div v-if="stream.length === 0">
        <p>Your stream is empty</p>
      </div>

      <div v-else>
        <div v-for="photo in stream" :key="photo.PhotoId">

          <img :src="photo.File" alt="Photo">
          <p @click="getUserProfile(photo.UserId)" class="clickable">{{photo.UserId}}}</p>
          <p>{{ photo.Date }}</p>

          <button @click="toggleLike">
            <i :class="'fa ' + (liked ? 'fa-heart' : 'fa-heart-o')"></i>
          </button>
          <span>{{ photo.Likes.length }}</span>
          <!-- <p>{{photo.Likes.length}} Likes</p> -->
          <!-- <p>{{photo.Comments.length}} Comments</p> -->
          <div v-for="(comment, index) in photo.comments.slice(0, maxVisibleComments)"
               :key="index" @click="getUserProfile(comment.userId)">
            <p><span class="clickable">{{ comment.userId }}</span>: {{ comment.text }}</p>
          </div>
          <button v-if="photo.comments.length > maxVisibleComments" @click="expandComments(photo)"> See other comments </button>
        </div>
      </div>
    </div>

		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
	</div>
</template>

<style>
</style>
