<script>
import Photo from "../components/Photo.vue";
import ErrorMsg from "../components/ErrorMsg.vue";

export default {
  components: {Photo, ErrorMsg},
  data: function() {
    return {
      errormsg: null,
      stream: [],
    }
  },
  methods: {
    async getStream() {
      this.errormsg = null
      try{
        let response = await this.$axios.get("/users/"+localStorage.getItem('auth')+"/stream")
        if (response.data != null) {
          this.stream = response.data;
        }
      } catch(e){
        this.errormsg = e.toString();
      }
    },
  },
  mounted() {
    this.getStream()
  }
}
</script>

<template>
  <div>
    <div
        class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
      <h1 class="h2">Home page</h1>
    </div>

    <div>
      <div v-if="stream.length === 0">
        <p>Your stream is empty</p>
      </div>

      <div v-else>
        <div class="row">
          <div class="col">
              <Photo v-for="(photo,index) in stream"
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
