<template>
  <div style="margin-top:-4px;">
    <div class="imgArea">
      <img src="../../assets/house2.webp" alt class="img" />
      <div class="imgCenter">
        <div class="col col-sm col-b search">
          Ara
          <input type="text" class="fit" />
        </div>
      </div>
    </div>
    <div class="col-lg center">
      <h2>Yeni ilanlar</h2>
    </div>
    <div class="center">
      <div class="container">
        <p>{{ pagec }}</p>
        <div class="card ca-md hover" v-for="notice in notices" :key="notice.id" @click="this.$router.push('/ilan/'+notice.Page_name);">
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
        <li
          @click="changePage(page)"
          :class="{ active: currentPage === page }"
          v-for="page in totalPages"
          :key="page"
        >{{ page }}</li>
      </ul>
    </div>
  </div>
</template>


<style>
.img {
  width: 100vw;
  max-height: 100vh;
  transform: translate(0%, -20%);
  z-index: 99;
}
.imgArea {
  max-width: 100vw;
  max-height: 50vh;
  overflow: hidden;
  position: relative;
  user-select: none;
}

.imgCenter {
  position: absolute; /* Görüntü kutusuna göre konumlandırmak için */
  top: 50%; /* Dikey olarak ortalamak için */
  left: 50%; /* Yatay olarak ortalamak için */
  transform: translate(-50%, -50%); /* Ortalamak için */
}
</style>

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