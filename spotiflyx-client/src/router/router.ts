import { createRouter, createWebHistory } from "vue-router";
import Home from '../views/homepage.vue';
import Signup from '../views/Signup.vue';
import Signin from '../views/Signin.vue';
import ErrorPage from '../views/ErrorPage.vue';

const routes = [
    { path: '/:pathMatch(.*)*', name: 'not-found', component: ErrorPage },
    { path: '/', name: 'Home', component: Home },
    { path: '/signup', name: 'Signup', component: Signup },
    { path: '/signin', name: 'Signin', component: Signin }
]

const routers = createRouter({
    history: createWebHistory(),
    routes,
    scrollBehavior(to, from, savedPosition) {
        return { top: 0 }
    }
})

export default routers