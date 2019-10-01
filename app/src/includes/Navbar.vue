<template>
    <!-- User Dropdown menu -->
    <div>
        <ul id="user-menu-dropdown" class="dropdown bottom-borders">
                <!-- home -->
                <li class="">
                    <a href="/users/master" title="@lang('recipers.my_account')">
                        <i class="fas fa-user-circle fa-15x left with-red-hover"></i>
                        @lang('recipers.my_account')
                    </a>
                </li>

                <!-- add recipe -->
                <li class="">
                    <a href="#add-recipe-modal" title="@lang('recipes.new_recipe')" class="modal-trigger">
                        <i class="fas fa-plus-circle fa-15x left with-red-hover"></i>
                        @lang('recipes.new_recipe')
                    </a>
                </li>

                <!-- my recipes -->
                <li class="">
                    <a href="/users/other/my-recipes" title="@lang('recipes.my_recipes')">
                        <i class="fas fa-book-open fa-15x left with-red-hover"></i>
                        @lang('recipes.my_recipes')
                    </a>
                </li>

                <!-- users -->
                <li class="">
                    <a href="/users" title="@lang('recipers.recipers')">
                        <i class="fas fa-user-friends fa-15x left with-red-hover"></i>
                        @lang('recipers.recipers')
                    </a>
                </li>

                <!-- favorites -->
                <li class="">
                    <a href="/favs" title="@lang('messages.favorites')">
                        <i class="fas fa-star fa-15x left with-red-hover"></i>
                        @lang('messages.favorites')
                    </a>
                </li>

                    <ul class="collapsible" id="adminka-collapsible">
                        <li title="@lang('messages.adminka')" class="position-relative">
                            <div class="collapsible-header">
                                <i class="fas fa-shield-alt fa-15x mr-2"></i>
                                @lang('messages.adminka')
                            </div>
                            <div class="collapsible-body p-0">
                                <ul>
                                    <!-- approves -->
                                    <li class="position-relative">
                                        <a href="/admin/approves" title="@lang('approves.verifications')">
                                            <i class="fas fa-search fa-15x left with-red-hover"></i>
                                            @lang('approves.verifications')
                                        </a>
                                    </li>

                                    <!-- feedback -->
                                    <li class="position-relative">
                                        <a href="/admin/feedback" title="@lang('feedback.contact_us')">
                                            <i class="fas fa-comment-dots fa-15x left with-red-hover"></i>
                                            @lang('feedback.contact_us')
                                        </a>
                                    </li>

                                    @hasRole('master')
                                        <li class="divider"></li>

                                        <!-- manage-users -->
                                        <li class="position-relative">
                                            <a href="/master/manage-users" title="@lang('manage-users.manage-users')">
                                                <i class="fas fa-user-cog fa-15x left with-red-hover"></i>
                                                @lang('manage-users.management')
                                            </a>
                                        </li>

                                        <!-- visitors -->
                                        <li class="position-relative">
                                            <a href="/master/visitors" title="@lang('visitors.visitors')">
                                                <i class="fas fa-users fa-15x left with-red-hover"></i>
                                                @lang('visitors.visitors')
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
                    <a href="/statistics" title="@lang('pages.statistics')">
                        <i class="fas fa-chart-bar fa-15x left with-red-hover"></i>
                        @lang('pages.statistics')
                    </a>
                </li>

                <!-- settings -->
                <li>
                    <a href="/settings" title="@lang('settings.settings')">
                        <i class="fas fa-cog fa-15x left with-red-hover"></i>
                        @lang('settings.settings')
                    </a>
                </li>

                <!-- <logout-btn
                    _auth-logout="@lang('auth.logout')"
                    _messages-sure-to-logout="@lang('messages.sure_to_logout')"
                ></logout-btn> -->
                <!-- Login -->
                <li>
                    <a href="/login" title="@lang('auth.login')">
                        <i class="fas fa-sign-in-alt fa-15x left with-red-hover"></i>
                        @lang('auth.login')
                    </a>
                </li>
                <!-- Register -->
                <li>
                    <a href="/register" title="@lang('auth.register')">
                        <i class="fas fa-pen-alt fa-15x left with-red-hover"></i>
                        @lang('auth.register')
                    </a>
                </li>

            <!-- <theme-switcher
                _messages-on="@lang('messages.on')"
                _messages-off="@lang('messages.off')"
                _messages-dark-mode="@lang('messages.dark_mode')"
                dark-theme-is-on-prop="dark"
            ></theme-switcher> -->

            <!-- <language-switcher lang="ru"></language-switcher> -->
        </ul>

        <nav class="no-select z-depth-1">
            <div class="nav-wrapper main-navbar">
                <div class="px-3 position-relative">
                    <!-- Logo -->
                    <a href="/" title="@lang('home.home')" class="brand-logo no-select">
                        <img src="@/assets/img/other/logo.png" alt="" height="30" style="min-width:25px; height:1px">
                        <span class="left pl-2">Reciper</span>
                    </a>

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
                        <a href="javascript:" data-target="mobile-demo"
                            class="hide-on-small-only right mr-4 align-to-the-middle"
                            title="@lang('pages.search')"
                            id="nav-btn-for-search"
                        >
                            <i class="fas fa-search fa-15x"></i>
                        </a>

                        <!-- Notifications bell -->
                        <a href="javascript:"
                            id="mark-notifs-as-read"
                            class="right ml-1 mr-4 dropdown-trigger position-relative align-to-the-middle"
                            title="@lang('notifications.notifications')"
                            data-target="notifications-dropdown"
                        >
                            <span style="height:53px">
                                <i class="fas fa-bell fa-15x"></i>
                            </span>
                        </a>
                    </div>

                    <!-- Regular menu -->
                    <ul class="right hide-on-med-and-down right-borders">
                        <li>
                            <a href="/" title="@lang('home.home')">
                                @lang('home.home')
                            </a>
                        </li>

                        <li>
                            <a href="/recipes#new" title="@lang('recipes.recipes')">
                                @lang('recipes.recipes')
                            </a>
                        </li>

                        <li>
                            <a href="help"
                                title="@lang('messages.help')"
                                target="_blank"
                            >
                            @lang('messages.help')
                            </a>
                        </li>

                        <li> <!-- Categories Dropdown Trigger Categories -->
                            <a href="javascript:"
                                class="dropdown-trigger"
                                data-target="categories-dropdown"
                            >
                                @lang('recipes.categories')
                                <i class="fas fa-caret-down fa-15x right"></i>
                            </a>
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
                            placeholder="@lang('pages.search_details')"
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
            @include('includes.nav.categories')
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