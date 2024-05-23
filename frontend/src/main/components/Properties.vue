<template>
  <div>
    <div class="container">
      <div
        class="card ca-md hover"
        v-for="notice in notices"
        :key="notice.id"
        @click="this.$router.push('/ilan/'+notice.Page_name);"
      >
        <div class="inner">
          <img :src="'/img/' + notice.File_names.split(' ')[0]" alt="Resim" />
          <div class="card-col">
            <svg width="18" height="18" class="flr">
              <use :xlink:href="svgLocation + '#lightning'" />
            </svg>
            <p>{{ notice.Title }}</p>
            <p>{{ notice.Property_status }} | {{ notice.Property_type }}</p>
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
</template>

<script>
import axios from "axios";

export default {
  props: ["propertyType","propertyStatus"],
  data() {
    return {
      notices: [],
      currentPage: 1,
      totalPages: null,
    };
  },
  created() {
    this.getNotices(1);
    console.log(this.propertyType)
    console.log(this.propertyStatus)
  },
  methods: {
    getNotices(pageNumber) {
      var req = "/api/getProperties/"+pageNumber
      if(this.propertyStatus){
        req += "/"+this.propertyStatus
      }else{
        req += "/all"
      }
      if(this.propertyType){
        req += "/"+this.propertyType
      }else{
        req += "/all"
      }
      axios
        .get(req)
        .then((response) => {
          if (response != undefined) {
            console.log(response.data);
            this.notices = response.data.properties;
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