import { createRouter, createWebHistory } from "vue-router";
import Home from '../views/homepage.vue';
import Signup from '../views/Signup.vue';
import ErrorPage from '../views/ErrorPage.vue';

const routes = [
    { path: '/:pathMatch(.*)*', name: 'not-found', component: ErrorPage },
    { path: '/', name: 'Home', component: Home },
    { path: '/signup', name: 'Signup', component: Signup }
]

const routers = createRouter({
    history: createWebHistory(),
    routes,
    scrollBehavior(to, from, savedPosition) {
        return { top: 0 }
    }
})

export default routers