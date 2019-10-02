<template>
    <div class="image-bg row mb-0">
        <div class="col s12 m6 offset-m4 form-wrapper my-5 z-depth-1">
            <div class="center pt-2">
                <h1 class="header">{{ $t('auth.register') }}</h1>
            </div>

            <form method="POST" action="/register" class="px-4">
                <div class="input-field">
                    <input type="text"
                        id="username"
                        name="username"
                        required
                    >
                    <label for="username">
                        {{ $t('auth.username') }}
                        <!-- @include('includes.tip', ['tip' => trans('auth.username_desc')]) -->
                    </label>
                    <!-- @include('includes.input-error', ['field' => 'username']) -->
                </div>

                <Visibility :tooltip="$t('forms.pwd_desc')"
                    input-name="password"
                    :input-label="$t('forms.pwd')"
                    input-error="Some error"
                ></Visibility>

                <Visibility :tooltip="$t('forms.pwd2_desc')"
                    input-name="password_confirmation"
                    :input-label="$t('forms.pwd_confirm')"
                    input-error="Some error"
                ></Visibility>

                <blockquote class="mt-5">
                    {!! trans('forms.agree_to_terms', ['btn' => trans('auth.register')]) !!}
                </blockquote>

                <button type="submit" id="register-btn" class="waves-effect waves-light btn mt-3">
                    {{ $t('auth.register') }}
                </button>
            </form>
        </div> <!-- / .col -->
    </div> <!-- / .image-bg -->
</template>

<script>
import Visibility from '@/components/Visibility.vue'

export default {
    components: {
        Visibility,
    },
}
</script>