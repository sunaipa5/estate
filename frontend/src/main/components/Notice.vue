<template>
  <div>
    <div
      class="card ca-md hover"
      v-for="notice in notices"
      :key="notice.id"
      @click="this.$router.push('/ilan/'+notice.Page_name);"
    >
      <div class="inner">
        <img :src="'/img/' + notice.File_names.split(' ')[0]" alt="Resim" />
        <div class="card-col">
          <p>{{ notice.Title }}</p>
          <p>{{ notice.Centare }} m2 | {{ notice.Room_number }}</p>
          <p>{{ notice.Address }}</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";

export default {
  props: ["pageName"],
  data() {
    return {
      notices: [],
    };
  },
  created() {
    this.getNoticeDetails();
  },
  methods: {
    getNoticeDetails() {
      console.log("here run...");
      const getData = {
        PageName: this.pageName,
      };
      axios
        .post("/getnotice", getData)
        .then((response) => {
          if (response != undefined) {
            console.log(response.data);
            this.notices = response.data;
            document.title = this.notices[0].Title;
          }
        })
        .catch((error) => {
          console.error("Get Error:", error);
        });
    },
  },
};
</script>