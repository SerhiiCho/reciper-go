<template>
    <div class="home">
        <div class="home-section home-section--contrast">
            <div class="home-section__inner">
                <Scenes />
            </div>
            <div class="home-section home-section--after-scene">
                <div class="home-section__inner">
                    <!-- <random-recipes visitor-id="{{ visitor_id() }}" limit="8"></random-recipes> -->
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import Scenes from '@/includes/home/Scenes.vue'

export default {
    name: 'home',
    components: {
        Scenes
    }
}
</script>
