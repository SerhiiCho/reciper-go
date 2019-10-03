<template>
    <div>
        <a :href="`/recipes#${btn.link}`" v-for="btn in btns"
            :key="btn.link"
            :class="{ 'active': btn.isActive }"
            class="btn btn-sort no-select"
        >
            <i class="fas red-text left" :class="btn.icon"></i>
            <span v-text="btn.title"></span>
        </a>
    </div>
</template>

<script>
export default {
    data() {
        return {
            btns: [
                {
                    title: this.$t('recipes.last'),
                    icon: 'fa-clock',
                    link: 'new',
                    isActive: false,
                },
                {
                    title: this.$t('recipes.popular'),
                    icon: 'fa-heart',
                    link: 'most_liked',
                    isActive: false,
                },
                {
                    title: this.$t('recipes.simple'),
                    icon: 'fa-concierge-bell',
                    link: 'simple',
                    isActive: false,
                },
                {
                    title: this.$t('home.breakfast'),
                    icon: 'fa-utensils',
                    link: 'breakfast',
                    isActive: false,
                },
                {
                    title: this.$t('home.lunch'),
                    icon: 'fa-utensils',
                    link: 'lunch',
                    isActive: false,
                },
                {
                    title: this.$t('home.dinner'),
                    icon: 'fa-utensils',
                    link: 'dinner',
                    isActive: false,
                },
                {
                    title: this.$t('recipes.watched'),
                    icon: 'fa-eye',
                    link: 'my_views',
                    isActive: false,
                },
            ]
        }
    },

    created() {
        this.buttonWasClicked()
    },

    methods: {
        buttonWasClicked() {
            Event.$on('hash-changed', hash => {
                this.btns[0].isActive = true

                if (hash != '') {
                    this.btns.forEach(btn => btn.isActive = btn.link == hash)
                }
            })
        },
    },
}
</script>
