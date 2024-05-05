import { createApp } from 'vue';
import App from './Admin.vue';

import '../assets/simplify-1.1.min.css'
import '../assets/simplify-extrait.css'
import '../assets/simplify.js'

import { createRouter, createWebHistory } from 'vue-router';

const routes = [
  { path: '/admin', component: () => import('./components/Home.vue') },
  { path: '/admin/ilanlar', component: () => import('./components/Notices.vue') },
  { path: '/admin/ilan/:pageName', component: () => import('./components/Notice.vue'), props: true },
  { path: '/admin/ilan-ekle', component: () => import('./components/Addnotice.vue')},
  { path: '/admin/send', component: () => import('./components/SendFile.vue')}
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});


const app = createApp(App);
app.use(router);
app.mount('#app');