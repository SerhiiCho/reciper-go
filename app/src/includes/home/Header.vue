<template>
    <header class="home-header" :style="{'background-image': `url('${bgImageURL}')`}">
        <div class="home-header__overlay"></div>
        <div class="home-header__content">
            <h1 class="home-header__content__title">{{ $t('messages.app_name') }}</h1>

            <div class="home-header__content__meal">
                <div>
                    <a href="/recipes#simple" class="home-header__content__meal__word">
                        {{ $t('home.show_simple_recipes') }}
                    </a>
                </div>

                <div>
                    <a href="/recipes#breakfast" class="home-header__content__meal__word">
                        {{ $t('home.breakfast') }}
                    </a>, 
                    <a href="/recipes#lunch" class="home-header__content__meal__word">
                        {{ $t('home.lunch') }}
                    </a>
                    {{ $t('home.or') }} 
                    <a href="/recipes#dinner" class="home-header__content__meal__word">
                        {{ $t('home.dinner') }}
                    </a>
                </div>

                <a class="modal-trigger main-dark px-3 home-header__content__meal__word"
                    href="#add-recipe-modal"
                >
                    {{ $t('home.add_your_recipe') }}
                </a>
            </div> <!-- / .home-header__content__meal-wrap -->

            <!--  Search Form  -->
            <form action="/search" method="get">
                <div class="position-relative mt-4">
                    <div class="home-header__content__search-wrap" id="home-search-form">
                        <input type="search"
                            name="for"
                            id="header-search-input"
                            :placeholder="$t('pages.search_details')"
                            class="home-header__content__search-wrap__input"
                        />
                    </div> <!-- / .home-header__content__search-wrap -->

                    <!-- Search button -->
                    <button type="submit"
                        id="home-search-btn"
                        class="home-header__content__btn"
                    >
                        <i class="fas fa-search"></i>
                    </button>
                </div>
            </form>
        </div>
    </header>
</template>

<script>
export default {
    data() {
        return {
            bgImageURL: require(`@/assets/img/other/header/by-daria-yakovleva-${Math.floor((Math.random() * 3) + 1)}.jpg`),
        }
    }
}
</script>