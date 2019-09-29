import Vue from 'vue'
import axios from 'axios'

window.Vue = Vue
window.Event = new Vue()

Vue.prototype.$axios = axios
Vue.config.productionTip = false

Vue.config.ignoredElements = [
    'statistics-chart'
]
