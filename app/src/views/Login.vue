<template>
    <div class="image-bg row mb-0">
        <div class="col s12 m6 offset-m4 form-wrapper my-5 z-depth-1">
            <div class="center pt-2">
                <h1 class="header">{{ $t('auth.login') }}</h1>
            </div>

            <form method="POST" action="/login" class="px-4">
                <div class="input-field">
                    <input id="username"
                        type="text"
                        name="username"
                        class="validate"
                        autofocus
                    >

                    <label for="username">{{ $t('auth.username') }}</label>
                    <!-- @include('includes.input-error', ['field' => 'username']) -->
                </div>

                <Visibility
                    :input-label="$t('forms.pwd')"
                    input-error="error here"
                    input-name="password"
                ></Visibility>

                <div class="mt-3">
                    <label>
                        <input type="checkbox" name="remember">
                        <span>
                            {{ $t('forms.remember_me') }}
                            <!-- @include('includes.tip', ['tip' => trans('tips.remember_info')]) -->
                        </span>
                    </label>
                </div>

                <div class="input-field">
                    <button type="submit" id="go-to-account" class="waves-effect waves-light btn">
                        {{ $t('auth.login') }}
                        <i class="fas fa-sign-in-alt right"></i>
                    </button>
                </div>

                <a href="/password/reset" class="text text-hover">
                    {{ $t('forms.forgot_pwd') }}
                </a>
            </form>
        </div>
    </div>
</template>

<script>
import Visibility from '@/components/Visibility.vue'

export default {
   components: {
       Visibility,
   },
}
</script>