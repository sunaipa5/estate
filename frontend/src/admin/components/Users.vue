<template>
  <div>
    <calert v-if="alertVisible" @close="closeAlert" class="col-lg">
      <form @submit.prevent="addUser">
        <h2>Yeni Kullanıcı Ekle</h2>
        <input type="text" class="fit" placeholder="Kullanıcı Adı" ref="username" />
        <input type="password" class="sz-50" placeholder="Parola" ref="password" />
        <input type="password" class="sz-50 mgl" placeholder="Parola tekrar" ref="passwordCheck" />
        <p class="fll" ref="status"></p>
        <button class="btn-md btb flr" type="submit">Ekle</button>
      </form>
    </calert>

    <div class="btn-x hover fll" @click="showAlert">
      <svg xmlns="http://www.w3.org/2000/svg" width="50" height="50" fill="currentColor"
        class="bi bi-file-earmark-plus-fill" viewBox="0 0 16 16">
        <path
          d="M9.293 0H4a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V4.707A1 1 0 0 0 13.707 4L10 .293A1 1 0 0 0 9.293 0M9.5 3.5v-2l3 3h-2a1 1 0 0 1-1-1M8.5 7v1.5H10a.5.5 0 0 1 0 1H8.5V11a.5.5 0 0 1-1 0V9.5H6a.5.5 0 0 1 0-1h1.5V7a.5.5 0 0 1 1 0" />
      </svg>
      <h3 style="padding-top: 10px">Kullanıcı ekle</h3>
    </div>
    <br />
    <div class="col col-lg">
      <table class="fit">
        <thead>
          <tr>
            <th>Kullanıcı adı</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="user in users" :key="user.id">
            <th>{{ user.Username }}</th>
            <button class="btr btn-sm flr" @click="this.deleteUser(user.Username)">Sil</button>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import Calert from "./Calert.vue";

export default {
  components: {
    Calert,
  },
  data() {
    return {
      users: [],
      accessToken: null,
      alertVisible: false,
    };
  },
  created() {
    this.getaccesstoken();
    setInterval(this.getaccesstoken, 888000);
  },
  methods: {
    showAlert() {
      this.alertVisible = true;
    },
    closeAlert() {
      this.alertVisible = false;
    },
    getaccesstoken() {
      console.log("access token succesfully getted");
      axios
        .get("/admin/getaccesstoken")
        .then((response) => {
          if (response.status == 200) {
            this.accessToken = response.data.token;
            this.getUsers();
          }
        })
        .catch((error) => {
          if (error.response.status == 401) {
            window.location.replace("/login")
          }
        });
    },
    getUsers() {
      axios
        .get("/api/getUsers", {
          headers: {
            Authorization: `Bearer ${this.accessToken}`,
          },
        })
        .then((response) => {
          if (response.status == 200) {
            this.users = response.data;
          }
        })
        .catch((error) => {
          if (error.response.status == 401) {
            window.location.replace("/login")
          }
        });
    },
    addUser() {
      const status = this.$refs.status;
      if (this.$refs.password.value == this.$refs.passwordCheck.value) {
        const username = this.$refs.username.value;
        const password = this.$refs.password.value;

        const postData = {
          Username: username,
          Password: password,
        };

        axios
          .post("/api/addUser", postData, {
            headers: {
              Authorization: `Bearer ${this.accessToken}`,
            },
          })
          .then((response) => {
            if (response.status == 200) {
              status.innerText = "Kullanıcı başarılı şekilde eklendi.";
              this.getUsers();
            }
          })
          .catch((error) => {
            if (error.response.status == 401) {
              window.location.replace("/login")
            }
          });
      } else {
        status.innerText = "Parola eşleşmedi!";
      }

      console.log("Add user running");
    },
    deleteUser(username) {
      const postData = {
        Username: username,
      };

      axios
        .post("/api/deleteUser", postData, {
          headers: {
            Authorization: `Bearer ${this.accessToken}`,
          },
        })
        .then((response) => {
          console.log("response status:", response.status)
          alert = "Kullanıcı başarılı şekilde silindi.";
          this.getUsers();
        })
        .catch((error) => {
          if (error.response.status == 401) {
            window.location.replace("/login")
          }
        });
    },
  },
};
</script>