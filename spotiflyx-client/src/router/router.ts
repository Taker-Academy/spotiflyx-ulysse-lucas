import { createRouter, createWebHistory } from "vue-router";
import axios from 'axios';
import Home from '../views/homepage.vue';
import Signup from '../views/Signup.vue';
import Signin from '../views/Signin.vue';
import ErrorPage from '../views/ErrorPage.vue';
import MediaPage from '../views/MediaPage.vue';
import Account from '../views/Account.vue';

const routes = [
    { path: '/:pathMatch(.*)*', name: 'not-found', component: ErrorPage },
    { path: '/home', name: 'Home', component: Home },
    { path: '/signup', name: 'Signup', component: Signup },
    { path: '/signin', name: 'Signin', component: Signin },
    { path: '/account', name: 'Account', component: Account},
    { path: '/', redirect: '/home' },
    { path: '/media/:type/:id', name: 'Media', component: MediaPage }
]

const x = axios.create({
    baseURL: 'http://localhost:3000/api/',
    timeout: 5000,
});

x.interceptors.response.use(
    response => {
        return response;
    },
    error => {
        console.log("error", error);
        if (error.response && error.response.status === 401) {
            handleUnauthorized();
        }
        return Promise.reject(error);
    }
);

// add token to axios header if already in local storage
const token = localStorage.getItem('token');
if (token)
    x.defaults.headers.common['Authorization'] = `Bearer ${token}`;

const routers = createRouter({
    history: createWebHistory(),
    routes,
    scrollBehavior(to, from, savedPosition) {
        return { top: 0 }
    }
})

routers.beforeEach(async (to, from, next) => {
    // redirect to login page if not logged in and trying to access a restricted page
    const publicPages = ['/signin', '/signup'];
    const authRequired = !publicPages.includes(to.path);

    if (authRequired) {
        try {
            const response = await x.get('auth/authenticated');
            if (response.status === 200) {
                next();
            } else
                next('/signin?redirect=' + to.path);
        } catch (error) {
            console.log(error);
            next('/signin?redirect=' + to.path);
        }
    } else {
        next();
    }
});

function handleUnauthorized() {
    const path = window.location.pathname;
    console.log("Unauthorized request, redirect to signin page.");
    if (path !== '/signin' && path !== '/signup')
        routers.push('/signin?redirect=' + path);
    else
        routers.push('/signin');
}

export default routers
export var ax = x;