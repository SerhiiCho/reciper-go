<template>
    <section class="ml-2 d-inline-block">
        <a v-if="userId" class="_like-button like-btn" @click="fetchItems">
            <i class="fas fa-heart like-btn__icon" :class="iconClass"></i> 
            <span v-text="amount" class="like-btn__amount grey-text"></span>
        </a>

        <!-- If user is not logged in show this -->
        <a v-else class="like-btn" href="/login">
            <i class="fas fa-heart like-btn__icon"></i> 
            <span v-text="amount" class="like-btn__amount grey-text"></span>
        </a>
    </section>
</template>

<script>
import BtnItem from '../../mixins/BtnItem'

export default {
    data() {
        return {
            url: `/likes/${this.recipeSlug}`,
        }
    },
    mixins: [BtnItem],
}
</script>