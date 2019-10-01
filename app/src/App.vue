<template>
    <div id="app">
        <Navbar />
        <router-view/>
    </div>
</template>

<script>
import Navbar from '@/includes/Navbar'

export default {
    components: {
        Navbar,
    },

    created() {
        this.$i18n.locale = 'ru'
    },
}
</script>

<style lang="scss">
@import './assets/sass/app.sass';
</style>
