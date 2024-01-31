<script>
export default{
  data: function() {
    return {
      errormsg: null,
      userId: localStorage.getItem("auth"),
    }
  },

  methods: {
    async uploadPhoto(){
      let input = document.getElementById("fileUploader")

      const file = input.files[0];
      const reader = new FileReader();
      reader.readAsArrayBuffer(file)

      reader.onload = async () => {
        let response = await this.$axios.post("/users/"+localStorage.getItem("auth")+"/photos", reader.result, {
          headers: {'Content-Type': file.type},
        })
        this.photos.unshift(response.data)
      };
    },
  },

  mounted(){
    this.userId = localStorage.getItem("auth")
    console.log('userId:', this.userId)
  }
}
</script>

<template>
  <div>
    <div>
      <h3>{{userId}}</h3>
    </div>
    <div>
      <h2>Posts</h2>
      <input id="fileUploader" type="file" @change="uploadPhoto" accept=".jpg, .png">
      <label class="btn" for="fileUploader">
        <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#upload"/>Upload</svg>
      </label>
    </div>


    <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
  </div>
</template>

<style>
</style>