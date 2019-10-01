<template>
    <div class="home-scenes">
        <div class="row">
            <!-- Add account -->
            <div class="col s12 m6 l3">
                <a href="/register" title="@lang('auth.register')">
                    <div class="home-scenes__item z-depth-1">
                        <img src="../../assets/img/other/scenes/add-account.svg" alt="" height="120">
                        <h2 class="home-scenes__item__text">@lang('register.scene_add_account')</h2>
                    </div>
                </a>
            </div>

            <!-- Add recipe -->
            <div class="col s12 m6 l3">
                <a href="#add-recipe-modal" title="@lang('recipes.new_recipe')" class="modal-trigger">
                    <div class="home-scenes__item z-depth-1">
                        <img src="../../assets/img/other/scenes/add-recipe.svg" alt="" height="120">
                        <h2 class="home-scenes__item__text">@lang('register.scene_add_recipe')</h2>
                    </div>
                </a>
            </div>

            <!-- Translate -->
            <div class="col s12 m6 l3">
                <a href="/recipes">
                    <div class="home-scenes__item z-depth-1">
                        <img src="../../assets/img/other/scenes/translate.svg" alt="" height="120">
                        <h2 class="home-scenes__item__text">@lang('register.scene_translate')</h2>
                    </div>
                </a>
            </div>

            <!-- Stats -->
            <div class="col s12 m6 l3">
                <a href="/statistics">
                    <div class="home-scenes__item z-depth-1">
                        <img src="../../assets/img/other/scenes/stats.svg" alt="" height="120">
                        <h2 class="home-scenes__item__text">@lang('register.scene_watch_stats')</h2>
                    </div>
                </a>
            </div>
        </div>
    </div>
</template>