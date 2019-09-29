<template>
    <li>
        <a :title="_messagesDarkMode">
            <i class="fas fa-moon fa-15x left"></i>
            <div class="switch">
                <label class="ml-1">
                    {{ _messagesOn }}

                    <input type="checkbox" :checked="darkThemeIsOn == 'light'" @click="switchTheme($event)">
                    <span class="lever"></span>

                    {{ _messagesOff }}
                </label>
            </div>
        </a>
    </li>
</template>

<script>
export default {
    data() {
        return {
            darkThemeIsOn: this.darkThemeIsOnProp,
        }
    },

    props: {
        _messagesOn: { required: false, },
        _messagesOff: { required: false, },
        _messagesDarkMode: { required: false, },
        darkThemeIsOnProp: { required: false, },
    },

    methods: {
        switchTheme(event) {
            this.darkThemeIsOn === 'light'
                ? this.switchToDark()
                : this.switchFromDark()
        },

        switchToDark() {
            this.$axios.get('/invokes/dark-theme-switcher/dark')
                .catch(e => console.error(e))

            this.darkThemeIsOn = 'dark'
            document.body.classList.add('dark-theme')
        },

        switchFromDark() {
            this.$axios.get('/invokes/dark-theme-switcher/light')
                .catch(e => console.error(e))

            this.darkThemeIsOn = 'light'
            document.body.classList.remove('dark-theme')
        },
    },
}
</script>