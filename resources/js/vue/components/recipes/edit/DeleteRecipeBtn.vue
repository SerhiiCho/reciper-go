<template>
    <div class="d-inline-block">
        <span v-if="!error">
            <a href="javascript:;"
                id="_delete-recipe"
                class="btn-floating red tooltipped waves-effect waves-light"
                :data-tooltip="_formsDeleting"
                v-on:click="deleteRecipe"
                data-position="top"
            >
                <i class="fas fa-trash"></i>
            </a>
        </span>
    </div>
</template>

<script>
export default {
    data() {
        return {
            error: '',
        }
    },

    props: {
        recipeSlug: { required: true, },
        _recipesAreYouSureToDelete: { required: true, },
        _formsDeleting: { required: true, },
    },

    methods: {
        deleteRecipe() {
            if (confirm(this._recipesAreYouSureToDelete)) {
                this.$axios.delete(`/recipes/${this.recipeSlug}`)
                    .then(data => {
                        data.data === 'success'
                            ? (window.location.href = '/users/other/my-recipes')
                            : (this.error = this.trans('recipes.deleted_fail'))
                    })
                    .catch(err => console.error(err))
            }
        },
    },
}
</script>