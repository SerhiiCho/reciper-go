import './settings/bootstrap'
import './components'
import App from './App'
import router from './router'
import store from './store'

new Vue({
    router,
    store,
    render: h => h(App),
    el: '#app',
})
