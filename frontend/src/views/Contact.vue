<template>
    <div class="image-bg row mb-0">
        <div class="col s12 m6 offset-m4 form-wrapper my-5 z-depth-1 px-4">
            <div class="center mt-4"><h1 class="header">{{ $t('feedback.feedback') }}</h1></div>

            <!-- @if (session('success'))
                <div class="mt-4">
                    @include('includes.buttons.btns')
                </div>
            @else -->
                <p>{{ $t('contact.intro') }}</p>

                <form action="/feedback" method="post">
                    <div class="input-field">
                        <i class="fas fa-at prefix"></i>
                        <input type="email" name="email" id="email">
                        <label for="email">{{ $t('forms.email') }}</label>
                        <!-- @include('includes.input-error', ['field' => 'email']) -->
                    </div>

                    <div class="input-field">
                        <i class="fas fa-comment-alt prefix"></i>
                        <textarea name="message"
                            id="message"
                            class="materialize-textarea counter"
                            data-length="2000"
                            maxlength="2000"
                            minlength="20"
                            required
                        ></textarea>
                        <label for="message">{{ $t('forms.message') }}</label>
                        <!-- @include('includes.input-error', ['field' => 'message']) -->
                    </div>

                    <div class="input-field">
                        <button type="submit" class="btn btn-main">
                            <i class="fas fa-envelope left w20"></i> {{ $t('forms.send') }}
                        </button>
                    </div>
                </form>
            <!-- @endif -->
        </div>
    </div>
</template>

<script>
export default {
   //
}
</script>