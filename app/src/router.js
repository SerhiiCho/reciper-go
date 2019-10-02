import Vue from 'vue'
import Router from 'vue-router'
import Home from './views/Home.vue'
import Recipes from './views/Recipes.vue'
import Register from './views/Register.vue'

Vue.use(Router)

const router = new Router({
    mode: 'history',
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
