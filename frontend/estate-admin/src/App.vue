<template>
  <nav class="navbar">
    <div class="navl">
      <h1>TEST</h1>
    </div>
    <div class="navr">
      <ul>
        <li @click="this.$router.push('/');">Anasayfa</li>
        <li @click="this.$router.push('/ilanlar');">İlanlar</li>
        <li @click="changeTheme()">Toggle</li>
      </ul>
    </div>
  </nav>

  <main class="mvp-90 center">
    <br>
    <router-view></router-view>
  </main>
</template>

<script>
import { toggleTheme } from './assets/simplify.js'

export default{
  methods: {
        changeTheme(type) {
            toggleTheme(type)
        }
    },
}

</script>