<template>
    <ul class="sidenav" id="mobile-sidenav">
        <li>
            <div class="user-view">
                <div class="background">
                    <img src="@/assets/img/other/img-small.jpg" alt="reciper">
                </div>

                <!-- guest -->
                    <a href="/login"
                        class="btn waves-effect waves-light min-w mb-2"
                        :title="$t('auth.login')"
                    >
                        {{ $t('auth.login') }}
                    </a>

                    <br>

                    <a href="/register"
                        class="btn waves-effect waves-light min-w mb-2"
                        :title="$t('auth.register')"
                    >
                        {{ $t('auth.register') }}
                    </a>
                <!-- @else -->
                    <a href="/users/master">
                        <img class="circle z-depth-1 hoverable"
                            src="https://www.winhelponline.com/blog/wp-content/uploads/2017/12/user.png"
                            alt="master"
                        >
                    </a>
                    <a href="/users/master">
                        <span class="white-text name">Master</span>
                    </a>
                    <a href="#email">
                        <span class="white-text email">
                            {{ $t('recipers.popularity') }} 
                            <i class="fas fa-crown" style="font-size:0.8em;color:orange"></i> 
                            9999999
                        </span>
                    </a>
                <!-- @endguest -->
            </div>
        </li>

        <li>
            <a href="/" :title="$t('home.home')">
                <i class="fas fa-home fa-15x grey-text left"></i>
                {{ $t('home.home') }}
            </a>
        </li>
        <li>
            <a href="/recipes#new" :title="$t('recipes.recipes')">
                <i class="fas fa-book-open fa-15x grey-text left"></i>
                {{ $t('recipes.recipes') }}
            </a>
        </li>
        <li>
            <a href="/search" :title="$t('pages.search')">
                <i class="fas fa-search fa-15x grey-text left"></i>
                {{ $t('pages.search') }}
            </a>
        </li>
        <li>
            <a href="/docs/help/"
                :title="$t('messages.help')"
                target="_blank"
            >
                <i class="fas fa-question-circle fa-15x grey-text left"></i>
                {{ $t('messages.help') }}
            </a>
            <div class="divider"></div>
        </li>

        <li>
            <a class="subheader grey-text">{{ $t('recipes.categories') }}</a>
        </li>

        <Categories />
    </ul>
</template>

<script>
import Categories from '@/includes/nav/Categories.vue'

export default {
    components: {
        Categories,
    },
}
</script>