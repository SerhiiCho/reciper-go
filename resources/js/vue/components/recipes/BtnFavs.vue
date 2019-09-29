<template>
    <section class="btn-favs">
        <!-- Show star icon that you can click -->
        <a v-if="userId" class="_fav-button fav-btn" @click="fetchItems">
            <i class="fas fa-star fav-btn__icon" :class="iconClass"></i> 
            <span v-text="amount" class="fav-btn__amount grey-text"></span>
        </a>

        <!-- If user is not logged in show this -->
        <a v-else class="fav-btn" href="/login">
            <i class="fas fa-star fav-btn__icon"></i> 
            <span v-text="amount" class="fav-btn__amount grey-text"></span>
        </a>
    </section>
</template>

<script>
import BtnItem from '../../mixins/BtnItem'

export default {
    data() {
        return {
            url: `/favs/${this.recipeSlug}`,
        }
    },
    mixins: [BtnItem],
}
</script>