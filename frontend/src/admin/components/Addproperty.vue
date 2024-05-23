<template>
  <div class="center col-lg">
    <div class="imgArea col col-lg cotnainer" v-if="imageUrls.length > 0">
      <div v-for="(imageUrl, index) in imageUrls" :key="index" class="fll">
        <img :src="imageUrl" v-if="imageUrl" />
      </div>
    </div>
    <div class="col col-lg">
      <h2>Yeni ilan ekle</h2>
      <form autocomplete="off" @submit.prevent="addNotice">
        <span>Resim ekle</span>

        <input
          class="fit"
          type="file"
          multiple
          ref="fileInput"
          @change="onFileChange"
          accept="image/*"
        />
        <span>
          İlan başlığı
          <span style="color:red;">*</span>
        </span>
        <input
          type="text"
          class="fit"
          placeholder="Bombay Sappire Rezidansları 5.Etap'da daire"
          ref="title"
          @input="autoPageName"
          required
        />
        <span>Sayfa ismi</span>
        <input
          type="text"
          class="fit"
          placeholder="bombay-sappire-rezidansları"
          ref="pageName"
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
        <div class="col-fluid row">
          <div class="sz-20 fll">
            <span>Durumu</span>
            <select class="inp-md fit" ref="propertyStatus">
              <option value="satılık">Satılık</option>
              <option value="kiralık">Kiralık</option>
            </select>
          </div>
          <div class="sz-20 fll mgl">
            <span>Tipi</span>
            <select class="inp-md fit" ref="propertyType">
              <option value="villa">Villa</option>
              <option value="müstakil">Müstakil</option>
              <option value="apartman">Apartman</option>
              <option value="rezidans">Rezidans</option>
            </select>
          </div>
        </div>
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
          <textarea
            name
            id
            cols="30"
            rows="10"
            class="fit"
            placeholder="Açıklama"
            ref="description"
          ></textarea>
        </div>
        <button type="submit" class="btn-md btb flr">Submit</button>
      </form>
    </div>
  </div>
</template>

<script>
import axios from "axios";

export default {
  data() {
    return {
      notices: [],
      accessToken: null,
      imageUrls: [],
    };
  },
  created() {
    this.getaccesstoken();
    setInterval(this.getaccesstoken, 888000);
  },
  methods: {
    getaccesstoken() {
      console.log("access token succesfully getted");
      axios
        .get("/admin/getaccesstoken")
        .then((response) => {
          this.accessToken = response.data.token;
        })
        .catch((error) => {
          console.log("Post Error:", error);
        });
    },
    onFileChange(event) {
      this.imageUrls = [];
      const files = event.target.files;
      for (let i = 0; i < files.length; i++) {
        const file = files[i];
        if (file.type.startsWith("image/")) {
          const imageUrl = URL.createObjectURL(file);
          this.imageUrls.push(imageUrl);
        }
      }
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
    autoPageName() {
      let modifiedTitle = this.clearTrChars(this.$refs.title.value)

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
      this.$refs.pageName.value = slug;
    },
    addNotice() {
      const title = this.$refs.title.value;
      const pageName = this.$refs.pageName.value;
      const address = this.$refs.address.value;
      const roomNumber = this.$refs.roomNumber.value;
      const centare = this.$refs.centare.value;
      const floor = this.$refs.floor.value;
      const floorNumber = this.$refs.floorNumber.value;
      const description = this.$refs.description.value;
      const propertyType = this.clearTrChars(this.$refs.propertyType.value);
      const propertyStatus = this.clearTrChars(this.$refs.propertyStatus.value);

      const postData = {
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
      console.log(postData)

      const files = this.$refs.fileInput.files;
      const formData = new FormData();

      for (let i = 0; i < files.length; i++) {
        formData.append("files[]", files[i]);
      }

      formData.append("jsonData", JSON.stringify(postData));

      axios
        .post("/api/addProperty", formData, {
          headers: {
            Authorization: `Bearer ${this.accessToken}`,
            "Content-Type": "multipart/form-data",
          },
        })
        .then((response) => {
          console.log(response.data.message);
          alert(response.data.message);
        })
        .catch((error) => {
          console.log("Post Error:", error);
        });
    },
    async uploadImages() {
      const formData = new FormData();

      this.images.forEach((image) => {
        formData.append("images[]", image);
      });

      try {
        const response = await axios.post("/upload", formData, {
          headers: {
            "Content-Type": "multipart/form-data",
          },
        });

        console.log(response.data);
      } catch (error) {
        console.error("Hata:", error);
      }
    },
  },
};
</script>