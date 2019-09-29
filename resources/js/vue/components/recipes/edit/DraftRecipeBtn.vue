<template>
    <button type="button"
        @click="draftTheRecipe"
        :data-tooltip="_tipsAddToDrafts"
        class="btn-floating green tooltipped waves-effect waves-light"
        ref="draftRecipe"
    >
        <i class="fas fa-pen"></i>
    </button>
</template>

<script>
import disableButton from '../../../../modules/disableButton'

export default {
    props: {
        _tipsAddToDrafts: { required: false, },
        _recipesAreYouSureToDraft: { required: false, },
    },

    methods: {
        draftTheRecipe() {
            if (confirm(this._recipesAreYouSureToDraft)) {
                disableButton(this.$refs.draftRecipe)
                this.$refs.draftRecipe.closest('form').submit()
            }
        },
    },
}
</script>

