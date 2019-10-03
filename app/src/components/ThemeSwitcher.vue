<template>
    <li>
        <a :title="$t('messages.dark_mode')">
            <i class="fas fa-moon fa-15x left"></i>
            <div class="switch">
                <label class="ml-1">
                    {{ $t('messages.on') }}

                    <input type="checkbox" :checked="darkThemeIsOn == 'light'" @click="switchTheme($event)">
                    <span class="lever"></span>

                    {{ $t('messages.off') }}
                </label>
            </div>
        </a>
    </li>
</template>

<script>
export default {
    data() {
        return {
            darkThemeIsOn: 'dark',
        }
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