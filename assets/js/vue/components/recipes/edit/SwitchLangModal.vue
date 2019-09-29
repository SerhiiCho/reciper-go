<template>
    <div>
        <div id="switch-lang-modal" class="modal switch-lang" v-if="switchTo == requestedLang" @keydown="cancelSwitching">
            <h3>{{ _recipesWantToSwitchLang }}</h3>

            <div class="switch-lang__btns mt-4">
                <button type="button" class="btn-flat btn-large mr-3" @click="switchLang">
                    <img :src="`/images/other/${switchTo}.png`">
                    {{ _messagesYes }}
                </button>
                <button type="button" class="btn-flat btn-large" @click="cancelSwitching">
                    <img :src="`/assets/images/other/${switchFrom}.png`">
                    {{ _messagesNo }}
                </button>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    data() {
        return {
            requestedLang: null,
        }
    },

    props: [
        '_recipesWantToSwitchLang',
        '_messagesYes',
        '_messagesNo',
        'switchFrom',
        'switchTo',
    ],

    created() {
        const request = window.location.search.substring(1)
        if (request == 'en' || request == 'ru') {
            this.requestedLang = request
        }
    },

    methods: {
        switchLang() {
            this.clearUrl()
            Event.$emit('switch-language', this.switchTo)
        },

        cancelSwitching() {
            Event.$emit('switch-language', this.switchFrom)
            this.clearUrl()
        },

        clearUrl() {
            history.pushState("", document.title, window.location.pathname)
            this.requestedLang = null
        },
    },

    computed: {
        //
    },
}
</script>

<style lang="sass" scoped>
.switch-lang
    &__btns
        display: flex
        flex-direction: row
        justify-content: center
        align-items: center

        button[type=button]
            display: flex
            flex-direction: row
            justify-content: center
            align-items: center

            img
                height: 19px
                margin-right: 7px
</style>
