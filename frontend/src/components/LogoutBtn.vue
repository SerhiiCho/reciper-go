<template>
    <li @click="logout">
        <a :title="$t('auth.logout')">
            <i class="fas fa-sign-out-alt fa-15x left"></i>
            {{ $t('auth.logout') }}
        </a>
    </li>
</template>

<script>
export default {
    methods: {
        logout() {
            if (confirm(this.$t('messages.sure_to_logout'))) {
                this.$axios.post('/logout')
                    .then(_ => window.location = '/')
                    .catch(e => console.error(e))
            }
        },
    },
}
</script>