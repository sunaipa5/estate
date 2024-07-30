<template>
  <div>
    <nav class="navbar">
      <div class="navl">
        <h1 @click="this.$router.push('/');">Estate</h1>
      </div>
      <div class="navl">
        <ul>
          <li @click="this.$router.push('/');">Anasayfa</li>
          <li @click="this.$router.push('/ilanlar');">İlanlar</li>
        </ul>
      </div>
      <div class="navr">
        <button class="tbtn" @click="changeTheme()" type="button">
          <svg
            id="dark"
            style="display: block;"
            xmlns="http://www.w3.org/2000/svg"
            width="16"
            height="16"
            fill="#111"
            class="bi bi-moon-fill"
            viewBox="0 0 16 16"
          >
            <path
              d="M6 .278a.768.768 0 0 1 .08.858 7.208 7.208 0 0 0-.878 3.46c0 4.021 3.278 7.277 7.318 7.277.527 0 1.04-.055 1.533-.16a.787.787 0 0 1 .81.316.733.733 0 0 1-.031.893A8.349 8.349 0 0 1 8.344 16C3.734 16 0 12.286 0 7.71 0 4.266 2.114 1.312 5.124.06A.752.752 0 0 1 6 .278z"
            />
          </svg>
          <svg
            id="light"
            style="display: none;"
            xmlns="http://www.w3.org/2000/svg"
            width="16"
            height="16"
            fill="#999"
            class="bi bi-brightness-high-fill"
            viewBox="0 0 16 16"
          >
            <path
              d="M12 8a4 4 0 1 1-8 0 4 4 0 0 1 8 0zM8 0a.5.5 0 0 1 .5.5v2a.5.5 0 0 1-1 0v-2A.5.5 0 0 1 8 0zm0 13a.5.5 0 0 1 .5.5v2a.5.5 0 0 1-1 0v-2A.5.5 0 0 1 8 13zm8-5a.5.5 0 0 1-.5.5h-2a.5.5 0 0 1 0-1h2a.5.5 0 0 1 .5.5zM3 8a.5.5 0 0 1-.5.5h-2a.5.5 0 0 1 0-1h2A.5.5 0 0 1 3 8zm10.657-5.657a.5.5 0 0 1 0 .707l-1.414 1.415a.5.5 0 1 1-.707-.708l1.414-1.414a.5.5 0 0 1 .707 0zm-9.193 9.193a.5.5 0 0 1 0 .707L3.05 13.657a.5.5 0 0 1-.707-.707l1.414-1.414a.5.5 0 0 1 .707 0zm9.193 2.121a.5.5 0 0 1-.707 0l-1.414-1.414a.5.5 0 0 1 .707-.707l1.414 1.414a.5.5 0 0 1 0 .707zM4.464 4.465a.5.5 0 0 1-.707 0L2.343 3.05a.5.5 0 1 1 .707-.707l1.414 1.414a.5.5 0 0 1 0 .708z"
            />
          </svg>
        </button>
      </div>
    </nav>

    <main class="mvp-90 center">
      <br />
      <router-view></router-view>
    </main>
      <div class="center">
        <footer>
          <div class="center">
            <p style="padding:20px 0; font-size:0.8rem;">
              Released under
              <a
                href="https://mit-license.org"
                target="_blank"
                class="link"
              >MIT License</a> | Copyright ©
              <a class="vex hover" href="https://github.com/sunaipa5" target="_blank">
                Sunaipa
                <svg
                  style="margin-left:3px;"
                  xmlns="http://www.w3.org/2000/svg"
                  width="20"
                  height="20"
                  fill="currentColor"
                  class="bi bi-github"
                  viewBox="0 0 16 16"
                >
                  <path
                    d="M8 0C3.58 0 0 3.58 0 8c0 3.54 2.29 6.53 5.47 7.59.4.07.55-.17.55-.38 0-.19-.01-.82-.01-1.49-2.01.37-2.53-.49-2.69-.94-.09-.23-.48-.94-.82-1.13-.28-.15-.68-.52-.01-.53.63-.01 1.08.58 1.23.82.72 1.21 1.87.87 2.33.66.07-.52.28-.87.51-1.07-1.78-.2-3.64-.89-3.64-3.95 0-.87.31-1.59.82-2.15-.08-.2-.36-1.02.08-2.12 0 0 .67-.21 2.2.82.64-.18 1.32-.27 2-.27.68 0 1.36.09 2 .27 1.53-1.04 2.2-.82 2.2-.82.44 1.1.16 1.92.08 2.12.51.56.82 1.27.82 2.15 0 3.07-1.87 3.75-3.65 3.95.29.25.54.73.54 1.48 0 1.07-.01 1.93-.01 2.2 0 .21.15.46.55.38A8.012 8.012 0 0 0 16 8c0-4.42-3.58-8-8-8z"
                  />
                </svg>
              </a>
            </p>
          </div>
        </footer>
      </div>
  </div>
</template>
  
  <script>
import { toggleTheme } from "../assets/simplify.js";
import { themeIcon } from "../assets/themeManager.js";

export default {
  mounted() {
    themeIcon();
  },
  methods: {
    changeTheme(type) {
      toggleTheme(type);
      themeIcon();
    },
  },
};
</script>
  