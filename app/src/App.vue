<template>
    <div id="app" class="page">
        <router-link to="/">Home</router-link> |
        <router-link to="/about">About</router-link>
        <router-view/>
    </div>
</template>

<style lang="scss">
@import './assets/sass/app.sass'
</style>
