import { createApp } from 'vue';
import App from './Main.vue';

import '../assets/simplify-1.1.min.css'
import '../assets/simplify-extrait.css'
import '../assets/simplify.js'


import { createRouter, createWebHistory } from 'vue-router';

const routes = [
  {
    path: '/', component: () => import('./components/Home.vue'),
    meta: {
      title: 'Ana Sayfa'
    }
  },
  {
    path: '/ilanlar', component: () => import('./components/Notices.vue'),
    meta: {
      title: 'İlanlar',
    },
  },
  {
    path: "/ilan/:pageName", component: () => import('./components/Notice.vue'), props: true,
    meta: {
      title: "İlan"
    },
  }
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

router.beforeEach((to, from, next) => {
  document.title = to.meta.title;
  next();
});


const app = createApp(App);
app.use(router);
app.mount('#app');