<template>
    <div class="d-inline-block">
        <input type="hidden" name="view" v-if="buttonWasCliked">

        <button type="button"
            @click="previewTheRecipe"
            :data-tooltip="_tipsView"
            class="btn-floating green tooltipped waves-effect waves-light"
            data-preview-recipe-btn
        >
            <i class="fas fa-eye"></i>
        </button>
    </div>
</template>

<script>
import disableButton from '../../../../modules/disableButton'

export default {
    data() {
        return {
            buttonWasCliked: false,
        }
    },

    props: {
        _tipsView: { required: true, },
    },

    methods: {
        previewTheRecipe() {
            let button = document.querySelector('[data-preview-recipe-btn]')
            this.buttonWasCliked = true

            disableButton(button)
            setTimeout(() => button.closest('form').submit(), 1)
        },
    },
}
</script>

