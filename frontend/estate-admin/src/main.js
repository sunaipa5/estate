import { createApp } from 'vue';
import App from './App.vue';
import router from './router';

import './assets/simplify-1.1.min.css'
import './assets/simplify-extrait.css'
import './assets/simplify.js'


const app = createApp(App);
app.use(router);
app.mount('#app');