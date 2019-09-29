<template>
    <button type="button"
        @click="submitClosestForm"
        :data-tooltip="_formsSave"
        class="btn-floating green tooltipped waves-effect waves-light"
        data-save-recipe-btn
    >
        <i class="fas fa-save"></i>
    </button>
</template>

<script>
import disableButton from '../../../../modules/disableButton'

export default {
    props: {
        _formsSave: { required: true, },
    },

    methods: {
        submitClosestForm() {
            let button = document.querySelector('[data-save-recipe-btn]')

            disableButton(button)
            button.closest('form').submit()
        },
    },
}
</script>

