<template>
  <div>
    <input type="file" ref="fileInput" multiple>
    <button @click="uploadFiles">Dosyaları Yükle</button>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  methods: {
    uploadFiles() {
      const files = this.$refs.fileInput.files;
      const formData = new FormData();

      for(let i = 0; i < files.length; i++) {
        formData.append('files[]', files[i]);
      }

      axios.post('/api/upload', formData, {
        headers: {
          'Content-Type': 'multipart/form-data'
        }
      })
      .then(response => {
        console.log(response.data);
      })
      .catch(error => {
        console.error(error);
      });
    }
  }
}
</script>
