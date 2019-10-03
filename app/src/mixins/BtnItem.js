export default {
    data() {
        return {
            iconClass: '',
            amount: this.items.length,
            userId: 1,
        }
    },

    props: {
        recipeSlug: { required: true, },
        recipeId: { required: true, },
        items: { required: true, },
    },

    created() {
        this.toggleActive()
    },

    methods: {
        /**
         * Got to a given url taken from vue component and
         * get string 'active' if recipe was just added to favs,
         * or null if recipe was removed from list of favs.
         * 
         * @return {void}
         */
        fetchItems() {
            this.$axios.post(this.url)
                .then(data => {
                    if (data.statusText === 'OK') {
                        this.iconClass = data.data
                        data.data == 'active' ? this.amount++ : this.amount--
                    }
                })
                .catch(err => console.error(err))
        },

        /**
         * If user is auth, set this.iconClass value to 'active' if user has
         * current recipe in list of favs, and set to empty string if user
         * doesn't have this recipe in the list of favorite recipes.
         * 
         * @return {void}
         */
        toggleActive() {
            if (this.userId) {
                let that = this
                let checkIfAddedBefore = this.items.filter(item => {
                    return that.recipeId == item.recipe_id && that.userId == item.user_id
                })
                this.iconClass = checkIfAddedBefore.length > 0 ? 'active' : ''
            }
        },
    },
}
