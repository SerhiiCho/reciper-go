<template>
    <div id="app">
        <Navbar />
        <router-link to="/">Home</router-link> |
        <router-link to="/about">About</router-link>
        <router-view/>
    </div>
</template>

<script>
import Navbar from '@/includes/Navbar'

export default {
    components: {
        Navbar,
    }
}
</script>

<style lang="scss">
@import './assets/sass/app.sass';
</style>
