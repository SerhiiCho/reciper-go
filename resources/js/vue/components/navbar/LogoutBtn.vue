<template>
    <li @click="logout">
        <a :title="_authLogout">
            <i class="fas fa-sign-out-alt fa-15x left"></i>
            {{ _authLogout }}
        </a>
    </li>
</template>

<script>
export default {
    props: {
        _authLogout: { required: false, },
        _messagesSureToLogout: { required: false, },
    },

    methods: {
        logout() {
            if (confirm(this._messagesSureToLogout)) {
                this.$axios.post('/logout')
                    .then(res => window.location = '/')
                    .catch(e => console.error(e))
            }
        },
    },
}
</script>