<template>
    <div id="app">
        <Sidenav />
        <Navbar />
        <router-view/>
        <Footer />
    </div>
</template>

<script>
import Navbar from '@/includes/Navbar'
import Sidenav from '@/includes/Sidenav'
import Footer from '@/includes/Footer'
import 'materialize-css/dist/js/materialize.min'
import '@/materialize-init'

export default {
    components: {
        Navbar,
        Footer,
        Sidenav,
    },

    created() {
        this.$i18n.locale = 'ru'
    },
}
</script>

<style lang="scss">
@import './assets/sass/app.sass';
</style>
