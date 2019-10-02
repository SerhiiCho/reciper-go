<template>
    <div class="home-scenes">
        <div class="row">
            <!-- Add account -->
            <div class="col s12 m6 l3">
                <router-link to="/register" :title="$t('auth.register')">
                    <div class="home-scenes__item z-depth0">
                        <img src="@/assets/img/other/scenes/add-account.svg" alt="" height="120">
                        <h2 class="home-scenes__item__text">{{ $t('register.scene_add_account') }}</h2>
                    </div>
                </router-link>
            </div>

            <!-- Add recipe -->
            <div class="col s12 m6 l3">
                <a href="#add-recipe-modal" :title="$t('recipes.new_recipe')" class="modal-trigger">
                    <div class="home-scenes__item z-depth-1">
                        <img src="@/assets/img/other/scenes/add-recipe.svg" alt="" height="120">
                        <h2 class="home-scenes__item__text">{{ $t('register.scene_add_recipe') }}</h2>
                    </div>
                </a>
            </div>

            <!-- Translate -->
            <div class="col s12 m6 l3">
                <router-link to="/recipes">
                    <div class="home-scenes__item z-depth-1">
                        <img src="@/assets/img/other/scenes/translate.svg" alt="" height="120">
                        <h2 class="home-scenes__item__text">{{ $t('register.scene_translate') }}</h2>
                    </div>
                </router-link>
            </div>

            <!-- Stats -->
            <div class="col s12 m6 l3">
                <a href="/statistics">
                    <div class="home-scenes__item z-depth-1">
                        <img src="@/assets/img/other/scenes/stats.svg" alt="" height="120">
                        <h2 class="home-scenes__item__text">{{ $t('register.scene_watch_stats') }}</h2>
                    </div>
                </a>
            </div>
        </div>
    </div>
</template>