<template>
    <div class="center py-3 m-auto" :class="className">
        <div class="preloader-wrapper active" :class="size">
            <div class="spinner-layer">
                <div class="circle-clipper left">
                    <div class="circle"></div>
                </div>
                <div class="gap-patch">
                    <div class="circle"></div>
                </div>
                <div class="circle-clipper right">
                    <div class="circle"></div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    props: {
        className: {
            default: '',
        },
        size:{
            default: '',
        }
    },
}
</script>

