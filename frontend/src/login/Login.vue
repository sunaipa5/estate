<template>
  <div>
    <div class="center loginForm">
      <div class="formArea">
        <form @submit.prevent="loginRequest" class="col col-smx">
          <div class="fomrh">
            <button class="tbtn" @click="changeTheme()" type="button">
              <svg id="dark" style="display: block;" xmlns="http://www.w3.org/2000/svg" width="16" height="16"
                fill="#111" class="bi bi-moon-fill" viewBox="0 0 16 16">
                <path
                  d="M6 .278a.768.768 0 0 1 .08.858 7.208 7.208 0 0 0-.878 3.46c0 4.021 3.278 7.277 7.318 7.277.527 0 1.04-.055 1.533-.16a.787.787 0 0 1 .81.316.733.733 0 0 1-.031.893A8.349 8.349 0 0 1 8.344 16C3.734 16 0 12.286 0 7.71 0 4.266 2.114 1.312 5.124.06A.752.752 0 0 1 6 .278z" />
              </svg>
              <svg id="light" style="display: none;" xmlns="http://www.w3.org/2000/svg" width="16" height="16"
                fill="#999" class="bi bi-brightness-high-fill" viewBox="0 0 16 16">
                <path
                  d="M12 8a4 4 0 1 1-8 0 4 4 0 0 1 8 0zM8 0a.5.5 0 0 1 .5.5v2a.5.5 0 0 1-1 0v-2A.5.5 0 0 1 8 0zm0 13a.5.5 0 0 1 .5.5v2a.5.5 0 0 1-1 0v-2A.5.5 0 0 1 8 13zm8-5a.5.5 0 0 1-.5.5h-2a.5.5 0 0 1 0-1h2a.5.5 0 0 1 .5.5zM3 8a.5.5 0 0 1-.5.5h-2a.5.5 0 0 1 0-1h2A.5.5 0 0 1 3 8zm10.657-5.657a.5.5 0 0 1 0 .707l-1.414 1.415a.5.5 0 1 1-.707-.708l1.414-1.414a.5.5 0 0 1 .707 0zm-9.193 9.193a.5.5 0 0 1 0 .707L3.05 13.657a.5.5 0 0 1-.707-.707l1.414-1.414a.5.5 0 0 1 .707 0zm9.193 2.121a.5.5 0 0 1-.707 0l-1.414-1.414a.5.5 0 0 1 .707-.707l1.414 1.414a.5.5 0 0 1 0 .707zM4.464 4.465a.5.5 0 0 1-.707 0L2.343 3.05a.5.5 0 1 1 .707-.707l1.414 1.414a.5.5 0 0 1 0 .708z" />
              </svg>
            </button>

            <h1>Login</h1>
          </div>
          <input type="text" class="fit inp-md" placeholder="username" ref="username" autofocus required />
          <input type="password" class="fit inp-md" placeholder="password" ref="password" required />
          <p ref="loginStatus" class="fll" style="width:160px"></p>
          <button class="btn-md btb flr" type="submit">Login</button>
        </form>
      </div>
    </div>
  </div>
</template>

<style>
.loginForm {
  height: 100vh;
  width: 60vw;
  justify-content: center;

}

.loginForm div {
  display: flex;
  align-items: center;
}

.formArea {
  justify-content: center;
  height: 50vh;

}

.col {
  margin-top: 20vh;
  position: relative;
}

.tbtn {
  background-color: transparent;
  border: none;
  outline: none;
}

.tbtn * {
  position: absolute;
  top: 0;
  right: 0;
  padding: 10px;
  margin: 0;
}

.formh {
  margin: 0;
}
</style>

<script>
import { toggleTheme } from "../assets/simplify.js";
import { themeIcon } from "../assets/themeManager.js";

export default {
  created() {
    toggleTheme("onload");
  },
  mounted() {
    themeIcon();
  },
  methods: {
    changeTheme(type) {
      toggleTheme(type);
      themeIcon();
    },
    loginRequest() {
      const username = this.$refs.username.value;
      const password = this.$refs.password.value;

      const postData = {
        Username: username,
        Password: password,
      };

      fetch("/login", {
        method: "POST",
        redirect: "manual",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(postData),
      })
        .then((response) => {
          if (response.status == 404) {
            this.$refs.loginStatus.innerHTML = "Kullanıcı bulunamadı"
          }
          if (response.headers.get("location")) {
            location = response.headers.get("location");
          }
        })
        .catch((error) => {
          console.error(
            "There was a problem with your fetch operation:",
            error
          );
        });

    },
  },
};
</script>