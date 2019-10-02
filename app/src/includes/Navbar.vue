<template>
    <!-- User Dropdown menu -->
    <div>
        <ul id="user-menu-dropdown" class="dropdown bottom-borders">
                <!-- home -->
                <li>
                    <a href="/users/master" :title="$t('recipers.my_account')">
                        <i class="fas fa-user-circle fa-15x left with-red-hover"></i>
                        {{ $t('recipers.my_account') }}
                    </a>
                </li>

                <!-- add recipe -->
                <li>
                    <a href="#add-recipe-modal" :title="$t('recipes.new_recipe')" class="modal-trigger">
                        <i class="fas fa-plus-circle fa-15x left with-red-hover"></i>
                        {{ $t('recipes.new_recipe') }}
                    </a>
                </li>

                <!-- my recipes -->
                <li>
                    <a href="/users/other/my-recipes" :title="$t('recipes.my_recipes')">
                        <i class="fas fa-book-open fa-15x left with-red-hover"></i>
                        {{ $t('recipes.my_recipes') }}
                    </a>
                </li>

                <!-- users -->
                <li>
                    <a href="/users" :title="$t('recipers.recipers')">
                        <i class="fas fa-user-friends fa-15x left with-red-hover"></i>
                        {{ $t('recipers.recipers') }}
                    </a>
                </li>

                <!-- favorites -->
                <li>
                    <a href="/favs" :title="$t('messages.favorites')">
                        <i class="fas fa-star fa-15x left with-red-hover"></i>
                        {{ $t('messages.favorites') }}
                    </a>
                </li>

                    <ul class="collapsible" id="adminka-collapsible">
                        <li :title="$t('messages.adminka')" class="position-relative">
                            <div class="collapsible-header">
                                <i class="fas fa-shield-alt fa-15x mr-2"></i>
                                {{ $t('messages.adminka') }}
                            </div>
                            <div class="collapsible-body p-0">
                                <ul>
                                    <!-- approves -->
                                    <li class="position-relative">
                                        <a href="/admin/approves" :title="$t('approves.verifications')">
                                            <i class="fas fa-search fa-15x left with-red-hover"></i>
                                            {{ $t('approves.verifications') }}
                                        </a>
                                    </li>

                                    <!-- feedback -->
                                    <li class="position-relative">
                                        <a href="/admin/feedback" :title="$t('feedback.contact_us')">
                                            <i class="fas fa-comment-dots fa-15x left with-red-hover"></i>
                                            {{ $t('feedback.contact_us') }}
                                        </a>
                                    </li>

                                    @hasRole('master')
                                        <li class="divider"></li>

                                        <!-- manage-users -->
                                        <li class="position-relative">
                                            <a href="/master/manage-users" :title="$t('manage-users.manage-users')">
                                                <i class="fas fa-user-cog fa-15x left with-red-hover"></i>
                                                {{ $t('manage-users.management') }}
                                            </a>
                                        </li>

                                        <!-- visitors -->
                                        <li class="position-relative">
                                            <a href="/master/visitors" :title="$t('visitors.visitors')">
                                                <i class="fas fa-users fa-15x left with-red-hover"></i>
                                                {{ $t('visitors.visitors') }}
                                            </a>
                                        </li>

                                        <!-- Horizon -->
                                        <li>
                                            <a href="/horizon/dashboard" title="Horizon">
                                                <i class="fas fa-h-square fa-15x left"></i>
                                                Horizon
                                            </a>
                                        </li>

                                        <!-- Telescope -->
                                        <li>
                                            <a href="/telescope/requests" title="Telescope">
                                                <i class="fas fa-star-and-crescent fa-15x left"></i>
                                                Telescope
                                            </a>
                                        </li>
                                        <li class="divider"></li>
                                </ul>
                            </div>
                        </li>
                    </ul>

                <!-- statistics -->
                <li>
                    <a href="/statistics" :title="$t('pages.statistics')">
                        <i class="fas fa-chart-bar fa-15x left with-red-hover"></i>
                        {{ $t('pages.statistics') }}
                    </a>
                </li>

                <!-- settings -->
                <li>
                    <a href="/settings" :title="$t('settings.settings')">
                        <i class="fas fa-cog fa-15x left with-red-hover"></i>
                        {{ $t('settings.settings') }}
                    </a>
                </li>

                <!-- <logout-btn
                    _auth-logout="{{ $t('auth.logout') }}"
                    _messages-sure-to-logout="{{ $t('messages.sure_to_logout') }}"
                ></logout-btn> -->
                <!-- Login -->
                <li>
                    <a href="/login" :title="$t('auth.login')">
                        <i class="fas fa-sign-in-alt fa-15x left with-red-hover"></i>
                        {{ $t('auth.login') }}
                    </a>
                </li>
                <!-- Register -->
                <li>
                    <a href="/register" :title="$t('auth.register')">
                        <i class="fas fa-pen-alt fa-15x left with-red-hover"></i>
                        {{ $t('auth.register') }}
                    </a>
                </li>

            <!-- <theme-switcher
                _messages-on="{{ $t('messages.on') }}"
                _messages-off="{{ $t('messages.off') }}"
                _messages-dark-mode="{{ $t('messages.dark_mode') }}"
                dark-theme-is-on-prop="dark"
            ></theme-switcher> -->

            <!-- <language-switcher lang="ru"></language-switcher> -->
        </ul>

        <nav class="no-select z-depth-1">
            <div class="nav-wrapper main-navbar">
                <div class="px-3 position-relative">
                    <!-- Logo -->

                    <router-link to="/" class="brand-logo no-select">
                        <img src="@/assets/img/other/logo.png" alt="" height="30" style="min-width:25px">
                        <span class="left pl-2 pr-2">Reciper</span>
                    </router-link>

                    <!-- Hamburger menu -->
                    <a href="javascript:" data-target="mobile-sidenav" class="sidenav-trigger no-select">
                        <i class="fas fa-bars"></i>
                    </a>

                    <div class="right">
                        <!-- Dropdown Trigger 2 User -->
                        <a href="javascript:"
                            class="right dropdown-trigger position-relative align-to-the-middle"
                            data-target="user-menu-dropdown"
                        >
                            <i class="user-icon-navbar d-block">
                                <img src="https://www.winhelponline.com/blog/wp-content/uploads/2017/12/user.png" alt="" class="z-depth-1 hoverable">
                            </i>
                        </a>

                        <!-- Dropdown Trigger 2  -->
                        <!-- <a href="javascript:"
                            class="right dropdown-trigger position-relative align-to-the-middle"
                            data-target="user-menu-dropdown"
                            id="_hamb-menu"
                        >
                            <i class="fas fa-bars user-icon-navbar z-depth-1 hoverable waves-effect waves-light d-block center"
                                style="line-height:40px; font-size:19px"
                            ></i>
                        </a> -->

                        <!-- Search button -->
                        <a href="javascript:"
                            data-target="mobile-demo"
                            class="hide-on-small-only right mr-4 align-to-the-middle"
                            :title="$t('pages.search')"
                            id="nav-btn-for-search"
                        >
                            <i class="fas fa-search fa-15x"></i>
                        </a>

                        <!-- Notifications bell -->
                        <!-- <a href="javascript:"
                            id="mark-notifs-as-read"
                            class="right ml-1 mr-4 dropdown-trigger position-relative align-to-the-middle"
                            :title="$t('notifications.notifications')"
                            data-target="notifications-dropdown"
                        >
                            <span style="height:53px">
                                <i class="fas fa-bell fa-15x"></i>
                            </span>
                        </a> -->
                    </div>

                    <!-- Regular menu -->
                    <ul class="right hide-on-med-and-down right-borders">
                        <router-link to="/" tag="li"><a>{{ $t("home.home") }}</a></router-link>
                        <router-link to="/recipes" tag="li"><a>{{ $t("recipes.recipes") }}</a></router-link>
                        <router-link to="/help" tag="li"><a>{{ $t("messages.help") }}</a></router-link>

                        <li> <!-- Categories Dropdown Trigger Categories -->
                            <!-- <a href="javascript:"
                                class="dropdown-trigger"
                                data-target="categories-dropdown"
                            >
                                {{ $t("recipes.categories") }}
                                <i class="fas fa-caret-down fa-15x right"></i>
                            </a> -->
                        </li>
                    </ul>
                </div>
            </div>
        </nav>

        <!-- Search navigation -->
        <nav class="nav-search-form" id="nav-search-form">
            <div class="nav-wrapper container">
                <form action="/search" method="get">
                    <div class="input-field">
                        <input id="search-input" type="search"
                            name="for"
                            :placeholder="$t('pages.search_details')"
                            required
                        >
                        <label class="label-icon" for="search-input">
                            <i class="fas fa-search"></i>
                        </label>
                    </div>
                </form>
            </div>
        </nav>

        <!-- Categories Dropdown trigger menu -->
        <ul id="categories-dropdown" class="dropdown bottom-borders">
            <!-- @include('includes.nav.categories') -->
        </ul>
    </div>
</template>

<script>
export default {
    data() {
        return {
            //
        }
    },
}
</script>