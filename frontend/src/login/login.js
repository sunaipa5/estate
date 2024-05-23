import { createApp } from 'vue';
import App from './Login.vue';

import '../assets/simplify-1.1.min.css'
import '../assets/simplify-extrait.css'
import '../assets/simplify.js'

import { createRouter, createWebHistory } from 'vue-router';

const routes = [
    {
        path: '/login', component: () => import('./Login.vue'), meta: {
            title: 'Estate Login'
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