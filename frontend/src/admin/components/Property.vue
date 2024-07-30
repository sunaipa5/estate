<template>
  <div>
    <calert v-if="alertVisible" @close="closeAlert" class="col-lg">
        <h3>İlanı silmek istediğinize eminmisiniz ?</h3>
        <button class="btb btn-sm flr" @click="this.closeAlert()" style="margin-left:5px;">İptal</button>
        <button class="btn-sm btr flr" @click="this.deleteProperty()">Sil</button>
    </calert>
    <h2>İlanı düzenle</h2>
    <div v-for="notice in notices" :key="notice.id">
      <div class="col col-lg container">
        <div class="imgArea" v-if="notice.File_names && notice.File_names.trim() !== ''">
          <img v-for="(fileName, index) in notice.File_names.trim().split(/\s+/)" :key="index" :src="'/img/' + fileName"
            alt="Resim" />
        </div>
        <div class="addBtn btn-x fll hover">Ekle</div>
      </div>
      <br />
      <div class="col col-lg">
        <button class="btr btn-sm" @click="this.showAlert()">İlanı Kaldır</button>
        <form autocomplete="off" @submit.prevent="updateProperty">
          <input class="fit" type="file" multiple id="fileInput" @change="onFileChange" accept="image/*" />
          <span>
            İlan başlığı
            <span style="color:red;">*</span>
          </span>
          <input type="text" class="fit" placeholder="Bombay Sappire Rezidansları 5.Etap'da daire" id="title"
            @input="autoPageName" :value="notice.Title" required />
          <span>Sayfa ismi</span>
          <input type="text" class="fit" placeholder="bombay-sappire-rezidansları" id="pageName"
            :value="notice.Page_name" disabled />
          <br />
          <span>
            Adres
            <span style="color:red;">*</span>
          </span>
          <textarea type="text" class="fit" placeholder="Bombay Sapphire Cad. Sandre Politan / Shizen-United"
            id="address" :value="notice.Address" required></textarea>
          <div class="col-fluid row">
            <div class="sz-20 fll">
              <span>Durumu</span>&nbsp;
              <select class="inp-md fit" id="propertyStatus" :value="notice.Property_status">
                <option value="satilik">Satılık</option>
                <option value="kiralik">Kiralık</option>
              </select>
            </div>
            <div class="sz-20 fll mgl">
              <span>Tipi</span>&nbsp;
              <select class="inp-md fit" id="propertyType" :value="notice.Property_type">
                <option value="villa">Villa</option>
                <option value="mustakil">Müstakil</option>
                <option value="apartman">Apartman</option>
                <option value="rezidans">Rezidans</option>
              </select>
            </div>
          </div>
          <div>
            <div class="sz-20 fll">
              <span>Oda sayısı</span>
              <input type="text" class="fit" placeholder="3+1" id="roomNumber" :value="notice.Room_number" />
            </div>
            <div class="sz-20 fll mgl">
              <span>Metrekare</span>
              <input type="number" class="fit" placeholder="100" id="centare" :value="notice.Centare" />
            </div>
            <div class="sz-20 fll mgl">
              <span>Bulunduğu kat</span>
              <input type="number" class="fit" placeholder="2" id="floor" :value="notice.Floor" />
            </div>
            <div class="sz-20 fll mgl">
              <span>Binadaki kat sayısı</span>
              <input type="number" class="fit" placeholder="8" id="floorNumber" :value="notice.Floor_number" />
            </div>
          </div>
          <div class="row fit">
            <span class="fll">Açıklama</span>
            <textarea cols="30" rows="10" class="fit" placeholder="Açıklama" id="description"
              :value="notice.Description"></textarea>
          </div>
          <button type="submit" class="btn-md btb flr">Güncelle</button>
        </form>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import Calert from "./Calert.vue";

export default {
  props: ["pageName"],
  components: {
    Calert,
  },
  data() {
    return {
      notices: [],
      accessToken: null,
      propertyId: null,
      alertVisible: false,
    };
  },
  created() {
    this.getaccesstoken();
    setInterval(this.getaccesstoken, 888000);
    this.getPropertyDetails();
  },
  methods: {
    showAlert() {
      this.alertVisible = true;
    },
    closeAlert() {
      this.alertVisible = false;
    },
    getaccesstoken() {
      axios
        .get("/admin/getaccesstoken")
        .then((response) => {
          this.accessToken = response.data.token;
        })
        .catch((error) => {
          console.log("Post Error:", error);
        });
    },
    getPropertyDetails() {
      const getData = {
        PageName: this.pageName,
      };
      axios
        .post("/api/getProperty", getData)
        .then((response) => {
          if (response != undefined) {
            console.log(response.data);
            this.notices = response.data;
            document.title = this.notices[0].Title;
            this.propertyId = this.notices[0].Id;
          }
        })
        .catch((error) => {
          console.error("Get Error:", error);
        });
    },
    autoPageName() {
      const title = document.getElementById("title").value;
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
      document.getElementById("pageName").value = slug;
    },
    updateProperty() {
      const id = this.propertyId;
      const title = document.getElementById("title").value;
      const pageName = document.getElementById("pageName").value;
      const address = document.getElementById("address").value;
      const roomNumber = document.getElementById("roomNumber").value;
      const centare = document.getElementById("centare").value;
      const floor = document.getElementById("floor").value;
      const floorNumber = document.getElementById("floorNumber").value;
      const description = document.getElementById("description").value;
      const propertyType = document.getElementById("propertyType").value;
      const propertyStatus = document.getElementById("propertyStatus").value;

      const postData = {
        Id: id,
        Title: title,
        Address: address,
        Room_number: roomNumber,
        Centare: parseInt(centare),
        Floor: parseInt(floor),
        Floor_number: parseInt(floorNumber),
        Description: description,
        Page_name: pageName,
        Property_type: propertyType,
        Property_status: propertyStatus,
      };
      console.log(postData);

      axios
        .post("/api/updateProperty", postData, {
          headers: {
            Authorization: `Bearer ${this.accessToken}`,
          },
        })
        .then((response) => {
          console.log(response.data.message);
          alert("İlan başarılı şekilde güncellendi");
        })
        .catch((error) => {
          console.log("Post Error:", error);
        });
    },
    deleteProperty() {
      const postData = {
        Id: this.propertyId,
      };
      axios
        .post("/api/deleteProperty", postData, {
          headers: {
            Authorization: `Bearer ${this.accessToken}`,
          },
        })
        .then((response) => {
          console.log(response.data.message);
          alert("İlan Başarılı bir şekilde silindi");
          this.$router.push('/admin/');
        })
        .catch((error) => {
          console.log("Post Error:", error);
        });
    }
  },
};
</script>

<style>
.addBtn {
  border-radius: 7px;
  width: 150px;
  height: 100px;
  padding: 0;
}
</style>