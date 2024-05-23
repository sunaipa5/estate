<template>
  <div style="margin-top:-4px;">
    <div class="imgAreax">
      <img src="../../assets/house2.webp" alt class="img" />
      <div class="imgCenter">
      
      </div>
    </div>
    <div class="center">
      <div class="searchZone">
        <div class="container">
          <select class="inp-md sz-20" ref="propertyStatus">
            <option value="satılık">Satılık</option>
            <option value="kiralık">Kiralık</option>
          </select>
          <select class="inp-md sz-20 mgl" ref="propertyType">
            <option value="villa">Villa</option>
            <option value="müstakil">Müstakil</option>
            <option value="apartman">Apartman</option>
            <option value="rezidans">Rezidans</option>
          </select>
          <button class="btn-md btb mgl" @click="searchProperty()">Listele</button>
        </div>
      </div>
    </div>
    <div class="col-xxl center">
      <div class="container homeList">
        <div class="btn-x hover" @click="this.$router.push('/ilanlar/all/villa');">
           <svg width="60" height="60" class="flr">
            <use :xlink:href="svgLocation + '#villa'" />
          </svg>
          <h3>Villa</h3>
        </div>
        <div class="btn-x hover" @click="this.$router.push('/ilanlar/all/mustakil');">
           <svg width="60" height="60" class="flr" >
            <use :xlink:href="svgLocation + '#house'" />
          </svg>
          <h3>Müstakil</h3>
        </div>
        <div class="btn-x hover" @click="this.$router.push('/ilanlar/all/apartman');">
          <svg width="60" height="60" class="flr" >
            <use :xlink:href="svgLocation + '#apartment'" />
          </svg>
          <h3>Apartman</h3>
        </div>
        <div class="btn-x hover" @click="this.$router.push('/ilanlar/all/rezidans');">
          <svg width="60" height="60" class="flr" >
            <use :xlink:href="svgLocation + '#residence'" />
          </svg>
          <h3>Rezidans</h3>
        </div>
      </div>
    </div>
    <br />
    <div class="center col-xxl">
      <h2>Yeni İlanlar</h2>
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
              <p>{{ notice.Title }}</p>
              <p>{{ notice.Centare }} m2 | {{ notice.Room_number }}</p>
              <p>{{ notice.Address }}</p>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>


<style>
.img {
  min-width: 100vw;
  max-width: 250vw;
  min-height: 65vh;
  max-height: 100vh;
  position: relative;
  left: 50%;
  transform: translate(-50%, -15%);
}
.imgAreax {
  max-width: 100vw;
  height: 60vh;
  overflow: hidden;
  position: relative;
  user-select: none;
}

.imgCenter {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
}

@media screen and (max-width: 500px) {
  .imgAreax {
    height: 50vh;
  }
}

.searchZone {
  width: 100vw;
  padding: 15px 0px;
}

.searchZone select{
  min-width:0px;
}

.light .searchZone {
   background-color: #f4f5f6;
  border-bottom: 1px solid #d1d1d1;
}

.dark .searchZone {
  background-color: #0e0e0e;
  border-bottom: 1px solid #333;
}

.homeList .btn-x{
  width: 80px;
}
</style>

<script>
import axios from "axios";
import svgLoc from "../../assets/all.svg";

export default {
  data() {
    return {
      notices: [],
      currentPage: 1,
      totalPages: null,
      svgLocation: svgLoc,
    };
  },
  created() {
    this.getProperties();
    console.log(svgLoc);
  },
  methods: {
    searchProperty(){
      console.log("running")
      this.$router.push('/ilanlar/'+this.clearTrChars(this.$refs.propertyStatus.value)+'/'+this.clearTrChars(this.$refs.propertyType.value))
    
    },
    clearTrChars(text){
      const turkishChars = {
        ç: "c",
        ı: "i",
        ğ: "g",
        ö: "o",
        ş: "s",
        ü: "u",
        Ç: "C",
        İ: "I",
        Ğ: "G",
        Ö: "O",
        Ş: "S",
        Ü: "U",
      };
      let modifiedText = text.replace(/[çığıöşüÇIİĞÖŞÜ]/g, function (match) {
        return turkishChars[match];
      });
      return modifiedText
    },
    getProperties() {
      axios
        .get("/api/getProperties/1/all/all")
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
      console.log("Yeni sayfa:", pageNumber);
    },
  },
};
</script>