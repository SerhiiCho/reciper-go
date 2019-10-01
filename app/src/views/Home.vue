<template>
    <div>
        <Header />
        <div class="page">
            <div class="home">
                <div class="home-section home-section--contrast">
                    <div class="home-section__inner">
                        <Scenes />
                    </div>
                    <div class="home-section home-section--after-scene">
                        <div class="home-section__inner">
                            <RandomRecipes visitor-id="1" limit="8"></RandomRecipes>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import Scenes from '@/includes/home/Scenes.vue'
import Header from '@/includes/home/Header.vue'
import RandomRecipes from '@/components/RandomRecipes.vue'

export default {
    components: {
        Scenes,
        Header,
        RandomRecipes,
    }
}
</script>
