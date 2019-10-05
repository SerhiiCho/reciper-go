<template>
    <div class="page" id="recipes-page">
        <div class="center"><h1 class="header">{{ $t('recipes.recipes') }}</h1></div>

        <div class="my-3" v-cloak>
            <SortButtons/>
        </div>

        <div class="row pt-3">
            <div class="col s12 m6 l3" v-for="recipe in recipes" :key="recipe.id">
                <div class="card">
                    <div class="card__image waves-effect waves-block waves-light">
                        <a :href="`/recipes/${recipe.slug}`" :title="recipe.intro">
                            <!-- Recipe image -->
                            <!-- <img :src="`storage/small/recipes/${recipe.image}`"
                                :alt="recipe.title"
                            > -->
                            <img :src="recipe.image"
                                 :alt="recipe.title"
                            >
                        </a>
                    </div>

                    <div class="card__content" :title="recipe.title">
                        <h5 class="card__content__title">
                            {{ recipe.excerpt }}
                        </h5>

                        <div class="card__content__favs">
                            <!-- Add to Favorites button -->
                            <div class="d-inline-block">
                                <BtnFavs
                                    :recipe-id="recipe.id"
                                    :recipe-slug="recipe.slug"
                                    :items="returnFavs(recipe.id)"
                                ></BtnFavs>
                            </div>

                            <!-- Preparation time -->
                            <span class="card__content__favs__time grey-text">
                                <i class="fas fa-clock"></i>
                                {{ recipe.time }} {{ $t('recipes.min') }}
                            </span>
                        </div> <!-- / .card__content__favs -->
                    </div> <!-- / .card__content -->
                </div> <!-- / .card -->
            </div> <!-- / .col -->
        </div> <!-- / .row -->

        <Preloader v-if="loading"></Preloader>
    </div>
</template>

<script>
import BtnFavs from '@/components/BtnFavs.vue'
import Preloader from '@/components/Preloader.vue'
import SortButtons from '@/components/SortButtons.vue'

export default {
    data() {
        return {
            recipes: [],
            loading: false,
            url: null,
            theEnd: false,
            favs: [],
        }
    },

    created() {
        this.fetchRecipes(true)

        // Call onScroll method if the there are more recipes
        // window.addEventListener('scroll', () => {
        //     if (!this.theEnd) this.onScroll()
        // })

        // Fetch recipes when the hash has changed
        window.onhashchange = () => {
            this.theEnd = false
            this.url = null
            this.fetchRecipes(true)
        }
    },

    methods: {
        /**
         * Make request to the server, get recipes json object and
         * set or concatenate {this.recipes} variable to the response data.
         * Emit the event hash-changed.
         *
         * @param {boolean} reload
         * @return {void}
         */
        fetchRecipes(reload = false) {
            this.loading = true

            const hash = window.location.hash.substring(1)
            // const url = this.url === null ? `http://localhost:8080/api/recipes/${hash}` : this.url
            const url = this.url === null ? `http://localhost:8080/api/recipes` : this.url

            Event.$emit('hash-changed', hash)

            this.$axios.get(url)
                .then(res => {
                    /**
                     * If pagination has the 'next' value not null, set url variable
                     * to that 'next' value. Otherwise set theEnd to true.
                     */
                    // if (res.data.links.next !== null) {
                    //     this.url = res.data.links.next
                    // } else {
                    //     this.theEnd = true
                    // }

                    /**
                     * Set recipes variable to received response data, but if 'reload' variable
                     * is true concatenate recipes with received response data
                     */
                    // this.recipes = reload ? res.data.data : this.recipes.concat(res.data.data)
                    this.recipes = reload ? res.data : this.recipes.concat(res.data)
                    this.loading = false
                    console.log(this.recipes)
                })
                .catch(err => {
                    console.error(err)
                    this.loading = false
                })

            this.loading = false
        },

        /**
         * Fetch recipes if user scrolled the screen to the bottom
         * of the browser window.
         *
         * @return {void}
         */
        onScroll() {
            const wrap = document.getElementById('recipes-page')
            const contentHeight = wrap.offsetHeight
            const yOffset = window.pageYOffset
            const currentPosition = yOffset + window.innerHeight

            if (currentPosition >= contentHeight && !this.loading) {
                this.fetchRecipes()
            }
        },

        /**
         * Return array of favs only with this recipe id
         *
         * @param {number} recipeId
         * @return {object}
         */
        returnFavs(recipeId) {
            return this.favs.filter(fav => fav.recipe_id == recipeId)
        },
    },

    components: {
        Preloader,
        BtnFavs,
        SortButtons,
    }
}
</script>
