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
    }
}
</script>

<style lang="scss">
@import './assets/sass/app.sass';
</style>
