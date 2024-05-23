import { createApp } from 'vue';
import App from './Admin.vue';

import '../assets/simplify-1.1.min.css'
import '../assets/simplify-extrait.css'
import '../assets/simplify.js'

import { createRouter, createWebHistory } from 'vue-router';

const routes = [
  {
    path: '/admin', component: () => import('./components/Home.vue'), meta: {
      title: 'Estate Admin'
    }
  },
  {
    path: '/admin/ilanlar', component: () => import('./components/Properties.vue'), meta: {
      title: 'İlanlar'
    }
  },
  {
    path: '/admin/ilan/:pageName', component: () => import('./components/Property.vue'), props: true, meta: {
      title: 'İlan'
    }
  },
  {
    path: '/admin/ilan-ekle', component: () => import('./components/Addproperty.vue'), meta: {
      title: 'İlan Ekle'
    }
  },
  {
    path: '/admin/kullanicilar', component: () => import('./components/Users.vue'), meta: {
      title: 'Kullanıcılar'
    }
  },
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