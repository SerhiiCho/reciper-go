<template>
    <div @click="iconState = !iconState" class="single__items__line">
        <div class="single__items__line__num" :class="classNames">
            <slot v-if="iconState" name="icon-on"></slot>
            <slot v-else name="icon-off"></slot>
        </div>
        <div class="single__items__line__text">
            <slot name="item"></slot>
        </div>
    </div>
</template>

<script>
export default {
    data() {
        return {
            iconState: false,
        }
    },
    props: {
        classNames: {
            default: null,
        },
    },
}
</script>