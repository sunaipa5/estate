<template>
  <div>
    <div class="container">
      <p>{{ pagec }}</p>
      <div class="card ca-md hover" v-for="notice in notices" :key="notice.id">
        <div class="inner">
          <!--img here-->
          <div class="card-col">
            <p>{{ notice.Title }}</p>
            <p>{{ notice.Centare }} m2 | {{ notice.Room_number }}</p>
            <p>{{ notice.Address }}</p>
          </div>
        </div>
      </div>
    </div>
    <ul class="pageNumberArea center">
      <li
        @click="changePage(page)"
        :class="{ active: currentPage === page }"
        v-for="page in totalPages"
        :key="page"
      >
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
        .get("/api/getnotices/" + pageNumber)
        .then((response) => {
          if (response != undefined) {
            console.log(response.data);
            this.notices = response.data.notices;
            this.totalPages = Math.ceil(response.data.total_count / 4);
          }
        })
        .catch((error) => {
          console.error("Get Error:", error);
        });
    },
    changePage(pageNumber) {
      this.currentPage = pageNumber;
      this.getNotices(pageNumber);
      // Sayfa değiştirildiğinde yapılması gereken işlemleri burada gerçekleştirin
      console.log("Yeni sayfa:", pageNumber);
    },
  },
};
</script>