<template>
  <div>
    <div class="container">
      <div class="card ca-md hover" v-for="notice in notices" :key="notice.id"
        @click="this.$router.push('/admin/ilan/' + notice.Page_name);">
        <div class="inner">
          <img :src="'/img/' + notice.File_names.split(' ')[0]" alt="Resim">
          <div class="card-col">
            <p>{{ notice.Title }}</p>
            <p>{{ notice.Centare }} m2 | {{ notice.Room_number }}</p>
            <p>{{ notice.Address }}</p>
          </div>
        </div>
      </div>
    </div>
    <ul class="pageNumberArea center">
      <li @click="changePage(page)" :class="{ active: currentPage === page }" v-for="page in totalPages" :key="page" v-show="totalPages > 1">
        {{ page }}
      </li>
    </ul>
  </div>
</template>

<script>
import axios from "axios";

export default {
  data() {
    return {
      notices: [],
      currentPage: 1,
      totalPages: null,
    };
  },
  created() {
    this.getNotices(1);
  },
  methods: {
    getNotices(pageNumber) {
      axios
        .get("/api/getProperties/" + pageNumber + "/all/all")
        .then((response) => {
          if (response != undefined) {
            console.log(response.data);
            this.notices = response.data.properties;
            this.totalPages = Math.ceil(response.data.total_count / 20);
          }
        })
        .catch((error) => {
          console.error("Get Error:", error);
        });
    },
    changePage(pageNumber) {
      this.currentPage = pageNumber;
      this.getNotices(pageNumber);
      console.log("Yeni sayfa:", pageNumber);
    },
  },
};
</script>