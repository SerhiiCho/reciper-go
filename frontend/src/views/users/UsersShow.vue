<template>
    <div>
        <div class="page profile-header">
        <!-- @if ($user->isActive()) -->
            <div class="row">
                <div class="col s12 l6">
                    <div class="mt-2">
                        <div class="profile-image">
                            <!-- <img src="./assets/img/big/users/master" alt="master"> -->
                            <img src="https://proxy.duckduckgo.com/iu/?u=https%3A%2F%2Ftse4.explicit.bing.net%2Fth%3Fid%3DOIP.wd2PrQItsPLW0pLUQqxDLAHaHa%26pid%3DApi&f=1" alt="master">
                        </div>
                    </div>

                    <div class="my-2">
                        <!-- Manage user -->
                        <!-- @if (optional(user())->hasRole('master')) -->
                        <a :href="`/master/manage-users/${$route.params.id}`"
                            class="mr-2"
                            :title="$t('manage-users.manage')"
                        >
                            <i class="fas fa-cog red-text fa-15x"></i>
                        </a>
                        <!-- @endif -->

                        <!-- Streak days -->
                        <div class="tooltipped d-inline-block"
                            :data-tooltip="$t('recipers.streak_days')"
                        >
                            <i class="fas fa-fire fa-15x" style="color:orangered"></i>
                            <b class="px-1">2</b> <!-- TODO: streak days -->
                        </div>

                        <!-- Stars -->
                        <div class="tooltipped d-inline-block"
                            :data-tooltip="$t('recipers.amount_of_favs')"
                        >
                            <i class="fas fa-star fa-15x" style="color:#d49d10"></i>
                            <b class="px-1">20</b> <!-- TODO: favs count -->
                        </div>
                    </div>
                </div>

                <h1 class="header mb-2">Master</h1> <!-- username -->

                <!-- Last visit -->
                <!-- @unless ($user->id === optional(user())->id) -->
                <span class="d-block py-2">
                    <!-- {!! get_online_icon(time_ago($user->online_check)) !!} -->
                    {{ $t('date.online') }}
                    <!-- {{ time_ago($user->online_check, 'online') }} -->
                </span>
                <!-- @endunless -->

                <!-- Registered -->
                <span class="d-block py-2 grey-text">
                    {{ $t('auth.register') }}: 3 days ago <!-- TODO: created at -->
                </span>

            <div class="col s12 l6">
                <div class="bubbles no-select">
                    <!-- Bubbles -->
                    <!-- @foreach ($bubbles as $bubble) -->
                    <!-- <div class="bubbles__single">
                        <i class="fas {{ $bubble->icon }} bubbles__single__icon"></i>
                        <div class="bubbles__single__inner" title="{{ $bubble->title }}">
                                    <span class="bubbles__single__inner__number">
                                        {!! $bubble->value !!}
                                    </span>
                        </div>
                        <span>{{ $bubble->text }}</span>
                    </div> -->
                    <!-- @endforeach -->
                </div>

                <!-- Level bar -->
                <!-- <div class="xp-bar mt-4 z-depth-1 mb-2" data-xp="$t('recipers.xp') {{ $user->xp }} {{ $max_xp }}"> -->
                <div class="xp-bar mt-4 z-depth-1 mb-2" :data-xp="`${$t('recipers.xp')} 233 / 4344`">
                    <!-- <div class="xp-bar__line" style="width:{{ $xp->getPercent() }}%"></div> -->
                    <div class="xp-bar__line" style="width:20%"></div>
                </div>

                <!-- Level Badge -->
                <div class="badge-panel mb-3">
                    <h3 class="mr-3 d-inline-block">{{ $t('recipers.reciper') }}</h3>
                    <!-- <div class="level-badge--wrap mt-5 d-inline-block z-depth-1 {{ $xp->getColor() }}"> -->
                    <div class="level-badge--wrap mt-5 d-inline-block z-depth-1">
                        <!-- <div class="level-badge {{ $xp->getColor() }}"> -->
                        <div class="level-badge">
                            <!-- <span>{{ $xp->getLevel() }}</span> -->
                            <span>2</span>
                        </div>
                    </div>
                    <h3 class="ml-3 d-inline-block">{{ $t('recipers.level') }}</h3>
                </div>

                <div class="center pb-3 pt-2">
                    <!-- <h4>{{ $user->status ?? null }}</h4> -->
                    <h4>Status</h4>
                </div>
            </div>
        </div>
    </div>

        <!-- <div class="divider"></div> -->

        <!--  All my recipes  -->
    <!--    @listOfRecipes(compact('recipes'))-->
    <!--    @slot('no_recipes')-->
    <!--    {{ $t('recipers.reciper_doesnt_have_recipes') }}-->
    <!--    @endslot-->
    <!--    @endlistOfRecipes-->
    <!--    @else-->
    <!--    <div class="center">-->
    <!--        <img src="{{ asset('storage/big/users/not_active.jpg') }}"-->
    <!--             class="profile-image z-depth-1 hoverable"-->
    <!--             alt="{{ $user->getName() }}"-->
    <!--        />-->
    <!--        <p>{{ $t('recipers.reciper_is_not_active') }}</p>-->
    <!--    </div>-->
    <!--    @endif-->
    </div>
</template>

<script>
    export default {
        data() {
            return {
                //
            }
        },

        created() {
            console.log('user with id ' + this.$route.params.id)
        }
    }
</script>
