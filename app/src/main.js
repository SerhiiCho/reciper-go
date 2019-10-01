import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import './registerServiceWorker'
import axios from 'axios'
import VueI18n from 'vue-i18n'
import ru from './locales/ru.json'
import en from './locales/en.json'

Vue.config.productionTip = false
Vue.prototype.$axios = axios
window.Event = new Vue()

Vue.use(VueI18n)

const i18n = new VueI18n({
    fallbackLocale: 'en',
    messages: { en, ru },
})

new Vue({
    router,
    store,
    i18n,
    render: h => h(App)
}).$mount('#app')
