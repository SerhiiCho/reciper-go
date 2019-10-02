<template>
    <footer class="px-5 pb-5">
        <div class="row pt-0 page">

            <div class="col s12 m6 l3 text-left">
                <ul class="unstyled-list"></ul>
            </div>

            <!--  Random recipes  -->
            <div class="col s12 m6 l3 text-left">
                <ul class="unstyled-list">
                    <li><strong>{{ $t('recipes.recipes') }}</strong></li>
                    <li>
                        <a href="/recipes/slug">
                            <i class="fas fa-angle-right red-text" style="width:7.5px"></i> 
                            Title
                        </a>
                    </li>
                </ul>
            </div>

            <!-- Popular recipes -->
            <div class="col s12 m6 l3 text-left">
                <ul class="unstyled-list">
                    <li><strong>{{ $t('recipes.popular') }}</strong></li>
                    <li>
                        <a href="/recipes/slug">
                            <i class="fas fa-angle-right red-text" style="width:7.5px"></i> 
                            Title
                        </a>
                    </li>
                </ul>
            </div>

            <!-- Navigation -->
            <div class="col s12 m6 l3 text-left row">
                <ul class="px-0 unstyled-list col s12">
                    <li><strong>{{ $t('pages.navigation') }}</strong></li>

                    <!-- Home page -->
                    <router-link to="/" tag="li" class="brand-logo no-select">
                        <a>
                            <i class="fas fa-angle-right red-text" style="width:7.5px"></i> 
                            {{ $t('home.home') }}
                        </a>
                    </router-link>

                    <!-- Recipe -->
                    <router-link to="/recipes#new" tag="li">
                        <a>
                            <i class="fas fa-angle-right red-text" style="width:7.5px"></i> 
                            {{ $t("recipes.recipes") }}
                        </a>
                    </router-link>

                    <!-- Search -->
                    <li>
                        <a href="/search">
                            <i class="fas fa-angle-right red-text" style="width:7.5px"></i>
                            {{ $t('pages.search') }}
                        </a>
                    </li>

                    <!-- Help -->
                    <li>
                        <a href="docs/help" target="_blank">
                            <i class="fas fa-angle-right red-text" style="width:7.5px"></i>
                            {{ $t('messages.help') }}
                        </a>
                    </li>

                    <!-- Contact -->
                    <li>
                        <a href="/contact">
                            <i class="fas fa-angle-right red-text" style="width:7.5px"></i>
                            {{ $t('feedback.contact_us') }}
                        </a>
                    </li>

                    <!-- Documents -->
                    <li>
                        <a href="/docs" target="_blank">
                            <i class="fas fa-angle-right red-text" style="width:7.5px"></i>
                            {{ $t('documents.documents') }}
                        </a>
                    </li>

                    <!-- Login -->
                    <router-link to="/login" tag="li">
                        <a :title="$t('auth.login')">
                            <i class="fas fa-angle-right red-text" style="width:7.5px"></i>
                            {{ $t('auth.login') }}
                        </a>
                    </router-link>

                    <!-- Register -->
                    <router-link to="/register" tag="li">
                        <a :title="$t('auth.register')">
                            <i class="fas fa-angle-right red-text" style="width:7.5px"></i>
                            {{ $t('auth.register') }}
                        </a>
                    </router-link>
                </ul>

                <!-- Documents -->
                <ul class="px-0 unstyled-list col s12">
                    <li>
                        <a href="/docs/" target="_blank">
                            <strong>{{ $t('documents.info') }}</strong>
                        </a>
                    </li>

                    <li>
                        <a href="/docs/terms" target="_blank">
                            <i class="fas fa-angle-right red-text" style="width:7.5px"></i> 
                            {{ $t('documents.terms') }}
                        </a>
                    </li>
                </ul>
            </div>
        </div>

        <!-- Help & Contact us buttons -->
        <div class="center pt-2">
            <div>
                <a href="/contact" class="mr-3">
                    <i class="fas fa-envelope mr-1" style="width:15px"></i>
                    {{ $t('feedback.contact_us') }}
                </a>

                
                <a href="/docs/help" target="_blank">
                    <i class="fas fa-question mr-1" style="width:11px"></i>
                    {{ $t('messages.help') }}
                </a>
            </div>

            <!-- Copyright -->
            <div class="footer-copyright mt-4 row">
                <div class="col s12 m1 l1 offset-m3 offset-l4 m">
                    <img src="@/assets/img/other/logo.png" alt="" height="50">
                </div>

                <div class="col s12 m5 l3 text-left">
                    <div class="valign-wrapper" style="height:70px">
                        <span>{{ $t('pages.title_footer') }}<br> &copy; {{ currentYear }} {{ $t('messages.app_name') }}</span>
                    </div>
                </div>
            </div>
        </div>
    </footer>
</template>

<script>
export default {
    data() {
        return {
            currentYear: new Date().getFullYear()
        }
    },
}
</script>