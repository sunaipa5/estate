<template>
  <div>
    <h2>İlanı düzenle</h2>
    <div v-for="notice in notices" :key="notice.id">
      <div class="col col-lg container">
        <div class="imgArea" v-if="notice.File_names && notice.File_names.trim() !== ''">
          <img
            v-for="(fileName, index) in notice.File_names.trim().split(/\s+/)"
            :key="index"
            :src="'/img/' + fileName"
            alt="Resim"
          />
        </div>
        <div class="addBtn btn-x fll hover">Ekle</div>
      </div>
      <br />
      <div class="col col-lg">
          <form autocomplete="off" @submit.prevent="addNotice">
      <span>Resim ekle</span>
      <input class="fit" type="file" multiple ref="fileInput" />
      <span>
        İlan başlığı
        <span style="color:red;">*</span>
      </span>
      <input
        type="text"
        class="fit"
        placeholder="Bombay Sappire Rezidansları 5.Etap'da daire"
        ref="titlex"
        @input="autoPageName"
        required
      />
      <span>Sayfa ismi</span>
      <input
        type="text"
        class="fit"
        placeholder="bombay-sappire-rezidansları"
        ref="pageNamex"
        disabled
      />
      <br />
      <span>
        Adres
        <span style="color:red;">*</span>
      </span>
      <textarea
        type="text"
        class="fit"
        placeholder="Bombay Sapphire Cad. Sandre Politan / Shizen-United"
        ref="address"
        required
      ></textarea>
      <div>
        <div class="sz-20 fll">
          <span>Oda sayısı</span>
          <input type="text" class="fit" placeholder="3+1" ref="roomNumber" />
        </div>
        <div class="sz-20 fll mgl">
          <span>Metrekare</span>
          <input type="number" class="fit" placeholder="100" ref="centare" />
        </div>
        <div class="sz-20 fll mgl">
          <span>Bulunduğu kat</span>
          <input type="number" class="fit" placeholder="2" ref="floor" />
        </div>
        <div class="sz-20 fll mgl">
          <span>Binadaki kat sayısı</span>
          <input type="number" class="fit" placeholder="8" ref="floorNumber" />
        </div>
      </div>
      <div class="row fit">
        <span class="fll">Açıklama</span>
        <textarea name id cols="30" rows="10" class="fit" placeholder="Açıklama" ref="description"></textarea>
      </div>
      <button type="submit" class="btn-md btb flr">Submit</button>
    </form>
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
    console.log(this.$refs.title);

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
          }
        })
        .catch((error) => {
          console.error("Get Error:", error);
        });
    },
    autoPageName() {
      const title = this.$refs.titlex;
      console.log(title)
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
      let modifiedTitle = title.replace(/[çığıöşüÇIİĞÖŞÜ]/g, function (match) {
        return turkishChars[match];
      });

      modifiedTitle = modifiedTitle.replace(/[^\w\s]/g, "");
      modifiedTitle = modifiedTitle.replace(/\s+/g, " ");

      const words = modifiedTitle.split(" ");
      while (words.join("-").length > 50) {
        words.pop();
      }
      const slugWords = words.map((word) => {
        return word.replace(/[^\w]/g, "").slice(0, 50);
      });

      const slug = slugWords.join("-").toLowerCase(); // Tüm harfleri küçült
      this.$refs.pageNamex.value = slug;
    },
  },
};
</script>

<style>
.imgArea {
  width: fit-content;
}
.imgArea img {
  width: 150px;
  height: 100px;
  border-radius: 7px;
  margin: 2.5px;
}
.addBtn {
  border-radius: 7px;
  width: 150px;
  height: 100px;
  padding: 0;
}
</style>