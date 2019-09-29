<template>
    <div class="row">
        <div class="col s12 m6 l3" v-for="recipe in recipes" :key="recipe.id">
            <div class="card">
                <a :href="`/recipes/${recipe.slug}`" :title="recipe.title">
                    <img :src="`/storage/small/recipes/${recipe.image}`"
                        style="width:100%"
                        :alt="recipe.title"
                    />
                </a>

                <div class="card__content card__content--sm" title="recipe.title">
                    <h5 class="card__content__title" v-text="recipe.excerpt"></h5>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    data() {
        return {
            recipes: [],
        }
    },

    created() {
        this.fetchData()
    },

    props: ['visitorId', 'limit'],

    methods: {
        fetchData() {
            this.$axios.get(`/api/random-recipes/${this.visitorId}/${this.limit}`)
                .then(res => this.recipes = res.data.data)
                .catch(err => console.error(err))
        },
    },
}
</script>
