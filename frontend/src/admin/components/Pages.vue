<template>
  <div>
    <div class="container col-lg">
      <div class="btn-x hover fll">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          width="50"
          height="50"
          fill="currentColor"
          class="bi bi-file-earmark-plus-fill"
          viewBox="0 0 16 16"
        >
          <path
            d="M9.293 0H4a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V4.707A1 1 0 0 0 13.707 4L10 .293A1 1 0 0 0 9.293 0M9.5 3.5v-2l3 3h-2a1 1 0 0 1-1-1M8.5 7v1.5H10a.5.5 0 0 1 0 1H8.5V11a.5.5 0 0 1-1 0V9.5H6a.5.5 0 0 1 0-1h1.5V7a.5.5 0 0 1 1 0"
          />
        </svg>
        <h3 style="padding-top: 10px">Yeni ilan ekle</h3>
      </div>
      <div class="btn-x hover fll">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          width="50"
          height="50"
          fill="currentColor"
          class="bi bi-file-earmark-plus-fill"
          viewBox="0 0 16 16"
        >
          <path
            d="M9.293 0H4a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V4.707A1 1 0 0 0 13.707 4L10 .293A1 1 0 0 0 9.293 0M9.5 3.5v-2l3 3h-2a1 1 0 0 1-1-1M8.5 7v1.5H10a.5.5 0 0 1 0 1H8.5V11a.5.5 0 0 1-1 0V9.5H6a.5.5 0 0 1 0-1h1.5V7a.5.5 0 0 1 1 0"
          />
        </svg>
        <h3 style="padding-top: 10px">Yeni ilan ekle</h3>
      </div>
      <div class="btn-x hover fll">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          width="50"
          height="50"
          fill="currentColor"
          class="bi bi-file-earmark-plus-fill"
          viewBox="0 0 16 16"
        >
          <path
            d="M9.293 0H4a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V4.707A1 1 0 0 0 13.707 4L10 .293A1 1 0 0 0 9.293 0M9.5 3.5v-2l3 3h-2a1 1 0 0 1-1-1M8.5 7v1.5H10a.5.5 0 0 1 0 1H8.5V11a.5.5 0 0 1-1 0V9.5H6a.5.5 0 0 1 0-1h1.5V7a.5.5 0 0 1 1 0"
          />
        </svg>
        <h3 style="padding-top: 10px">Yeni ilan ekle</h3>
      </div>
      <div class="btn-x hover fll">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          width="50"
          height="50"
          fill="currentColor"
          class="bi bi-file-earmark-plus-fill"
          viewBox="0 0 16 16"
        >
          <path
            d="M9.293 0H4a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V4.707A1 1 0 0 0 13.707 4L10 .293A1 1 0 0 0 9.293 0M9.5 3.5v-2l3 3h-2a1 1 0 0 1-1-1M8.5 7v1.5H10a.5.5 0 0 1 0 1H8.5V11a.5.5 0 0 1-1 0V9.5H6a.5.5 0 0 1 0-1h1.5V7a.5.5 0 0 1 1 0"
          />
        </svg>
        <h3 style="padding-top: 10px">Yeni ilan ekle</h3>
      </div>
    </div>
    <div class="col col-lg">
      <h2>Yeni ilan ekle</h2>
      <span>İlan başlığı</span>
      <form autocomplete="off" @submit.prevent="addNotice">
        <input
          type="text"
          class="fit"
          placeholder="Bombay Sappire Rezidansları 5.Etap'da daire"
          ref="title"
        />
        <span>Adres</span>
        <textarea
          type="text"
          class="fit"
          placeholder="Bombay Sapphire Cad. Sandre Politan / Shizen-United"
          ref="address"
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
            <input
              type="number"
              class="fit"
              placeholder="8"
              ref="floorNumber"
            />
          </div>
        </div>
        <div class="row fit">
          <span class="fll">Açıklama</span>
          <textarea
            name=""
            id=""
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
    <br />
    <div class="col col-lg">
      <h2>GetDB</h2>
      <p class="fll">{{ message }}</p>
      <table class="center">
        <tbody>
          <tr>
            <th>Başlık</th>
            <th>Adres</th>
            <th>Oda sayısı</th>
            <th>Metrekare</th>
            <th>Bulunduğu kat</th>
            <th>Binadaki kat sayısı</th>
            <th>Açıklama</th>
          </tr>

          <tr v-for="notice in notices" :key="notice.id">
            <td>{{ notice.Title }}</td>
            <td>{{ notice.Address }}</td>
            <td>{{ notice.RoomNumber }}</td>
            <td>{{ notice.Centare }}</td>
            <td>{{ notice.Floor }}</td>
            <td>{{ notice.Floor_number }}</td>
            <td>{{ notice.Description }}</td>
          </tr>
        </tbody>
      </table>
      <button class="btn-md btb" @click="getDB()">GET</button>
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
    addNotice() {
      const title = this.$refs.title.value;
      const address = this.$refs.address.value;
      const roomNumber = this.$refs.roomNumber.value;
      const centare = this.$refs.centare.value;
      const floor = this.$refs.floor.value;
      const floorNumber = this.$refs.floorNumber.value;
      const description = this.$refs.description.value;

      const postData = {
        Title: title,
        Address: address,
        Room_number: roomNumber,
        Centare: parseInt(centare),
        Floor: parseInt(floor),
        Floor_number: parseInt(floorNumber),
        Description: description,
      };

      axios
        .post("/addnotice", postData, {
          headers: {
            Authorization: `Bearer ${this.accessToken}`,
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
    getDB() {
      axios
        .get("/getnotices", {
          headers: {
            Authorization: `Bearer ${this.accessToken}`,
          },
        })
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
  },
};
</script>