<template>
    <div>
        <li style="border-bottom:none" class="position-relative" v-if="lang">
            <select class="px-3 pl-5" @change="changeLanguage(null)" v-model="selectedLang">
                <option value="en" :selected="lang == 'en'">English</option>
                <option value="ru" :selected="lang == 'ru'">Русский</option>
            </select>
            <img :src="require(`@/assets/img/other/${flag}.png`)" class="dropdown__flag">
        </li>
    </div>
</template>

<script>
export default {
    data() {
        return {
            selectedLang: this.lang,
            flag: this.lang,
        }
    },

    created() {
        Event.$on('switch-language', lang => this.changeLanguage(lang))
    },

    props: {
        lang: { required: false, }
    },

    methods: {
        changeLanguage(lang = null) {
            if (lang !== null) {
                this.selectedLang = lang
            }

            this.flag = this.flag === 'ru' ? 'en' : 'ru'

            this.$axios.get(`/invokes/change-language-switcher/${this.selectedLang}`)
                .then(_ => location.reload())
                .catch(e => console.error(e))
        },
    },
}
</script>