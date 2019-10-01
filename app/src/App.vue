<template>
    <div id="app">
        <Navbar />
        <router-view/>
        <Footer />
    </div>
</template>

<script>
import Navbar from '@/includes/Navbar'
import Footer from '@/includes/Footer'

export default {
    components: {
        Navbar,
        Footer,
    },

    created() {
        this.$i18n.locale = 'ru'
    },
}
</script>

<style lang="scss">
@import './assets/sass/app.sass';
</style>
