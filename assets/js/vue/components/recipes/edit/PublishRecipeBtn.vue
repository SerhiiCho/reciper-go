<template>
    <div class="d-inline-block">
        <input type="checkbox"
            :checked="isReadyToPublish"
            name="ready"
            value="1"
        >

        <button type="button"
            @click="publishTheRecipe"
            :data-tooltip="_tipsPublish"
            class="btn-floating green tooltipped waves-effect waves-light"
            data-publish-recipe-btn
        >
            <i class="fas fa-clipboard-check"></i>
        </button>
    </div>
</template>

<script>
import disableButton from '../../../../modules/disableButton'

export default {
    data() {
        return {
            isReadyToPublish: false,
        }
    },

    props: {
        _tipsPublish: { required: true, },
        _approvesAreYouSureToPublish: { required: true, },
    },

    methods: {
        publishTheRecipe() {
            let button = document.querySelector('[data-publish-recipe-btn]')

            if (confirm(this._approvesAreYouSureToPublish)) {
                this.isReadyToPublish = true

                disableButton(button)
                setTimeout(() => button.closest('form').submit(), 1)
            }
        },
    },
}
</script>

