import { createRouter, createWebHistory } from 'vue-router';

const routes = [
  { path: '/', component: () => import('./components/Home.vue') },
  { path: '/ilanlar', component: () => import('./components/Pages.vue') },
  { path: '/login', component: () => import('./components/Login.vue') },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});


export default router;
