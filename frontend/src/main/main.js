import { createApp } from 'vue';
import App from './Main.vue';
import VueSplide from '@splidejs/vue-splide';

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
    path: '/ilanlar/:propertyStatus?/:propertyType?',
    component: () => import('./components/Properties.vue'),
    props: true,
    meta: {
      title: 'İlanlar',
    },
  },  
  {
    path: '/ilanlar',
    component: () => import('./components/Properties.vue'),
    meta: {
      title: 'İlanlar',
    },
  },
  {
    path: "/ilan/:pageName", component: () => import('./components/Property.vue'), props: true,
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
  window.scrollTo(0, 0);
  document.title = to.meta.title;
  next();
});


const app = createApp(App);
app.use(router);
app.use(VueSplide);
app.mount('#app');