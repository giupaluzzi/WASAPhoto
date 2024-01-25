<script>
export default{
  data: function() {
    return {
      errormsg: null,
    }
  },

  methods: {
    async uploadPhoto(){
      let input = document.getElementById("fileUploader")

      const file = input.files[0];
      const reader = new FileReader();
      reader.readAsArrayBuffer(file)

      reader.onload = async () => {
        let response = await this.$axios.post("/users/"+localStorage.getItem("auth")+"/photos"+reader.result, {
          headers: {'Content-Type': file.type},
        })
      this.photos.unshift(response.data)
      };
    },
  },

  computed:{
    isLogged(){
      return this.$route.params.id === localStorage.getItem('auth')
    },
  },
}
</script>

<template>
  <div>
    <div>
      <h3>{{ this.$route.params.userId }}</h3>
    </div>
    <div>
        <h2>Posts</h2>
        <input id="fileUploader" type="file" @change="uploadPhoto" accept=".jpg, .png">
        <label class="btn" v-if="isLogged" for="fileUploader">
          <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#upload"/>Upload</svg>
        </label>
    </div>


    <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
  </div>
</template>

<style>
</style>