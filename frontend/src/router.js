import Vue from 'vue'
import Router from 'vue-router'
import Home from './views/Home.vue'
import Recipes from './views/Recipes.vue'
import Register from './views/Register.vue'
import Login from './views/Login.vue'
import Contact from './views/Contact.vue'
import UsersShow from './views/users/UsersShow.vue'

Vue.use(Router)

const router = new Router({
    mode: 'history',
    scrollBehavior (to, from, savedPosition) {
        document.querySelector('html').scrollIntoView({
            behavior: 'smooth',
            block: 'start',
        })
    },
    routes: [
        {
            path: '/',
            component: Home,
        },
        {
            path: '/recipes',
            component: Recipes,
        },
        {
            path: '/register',
            component: Register,
        },
        {
            path: '/login',
            component: Login,
        },
        {
            path: '/contact',
            component: Contact,
        },
        {
            path: '/users/:id',
            component: UsersShow,
        },
        // {
        //   path: '/about',
        //   component: () => import(/* webpackChunkName: "about" */ './views/About.vue')
        // }
    ]
})

router.beforeEach((to, from, next) => {
    let bar = document.createElement('div')

    bar.classList.add('page-loading-bar')
    document.body.appendChild(bar)

    setTimeout(() => bar.classList.add('page-loading-bar--run'), 10)
    setTimeout(() => bar.classList.remove('page-loading-bar--run'), 700)

    next()
})



export default router
