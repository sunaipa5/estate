<template>
  <div>
    <nav class="navbar">
      <div class="navl">
        <h2 class="vex" style="margin-right:5px;" @click="this.$router.push('/admin');">Estate</h2>
        <p class="vex">Admin</p>
      </div>
      <div class="navr">
        <ul>
          <li @click="this.$router.push('/admin');">Anasayfa</li>
          <li @click="this.$router.push('/admin/ilanlar');">İlanlar</li>
          <li @click="this.$router.push('/admin/ilan-ekle');">İlanlar</li>
          <li @click="changeTheme()">Toggle</li>
        </ul>
      </div>
    </nav>

    <main class="mvp-90 center">
      <br />
      <router-view></router-view>
    </main>
  </div>
</template>

<script>
import { toggleTheme } from "../assets/simplify.js";

export default {
  methods: {
    changeTheme(type) {
      toggleTheme(type);
    },
  },
};
</script>
