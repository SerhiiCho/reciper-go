<template>
    <div class="input-field">
        <input :type="type"
            :id="inputName"
            :name="inputName"
            class="with-icon"
            autocomplete="off"
            required
        >

        <!-- Eye icon -->
        <div class="visibility-icon waves-effect waves-green" :class="`_${inputName}-icon`" @click="changeType">
            <i :class="icon" class="fas no-select"></i>
        </div>

        <!-- Label -->
        <label :for="inputName">
            {{ inputLabel }}
            <a class="tooltipped ml-1" v-if="tooltip" style="cursor:pointer" :data-tooltip="tooltip">
                <i class="fas fa-question-circle main-text round"></i>
            </a>
        </label>
        <span class="helper-text red-text" v-if="inputError" data-error="wrong">
            {{ inputError }}
        </span> 
    </div>
</template>

<script>
export default {
    data() {
        return {
            type: 'password',
            icon: 'fa-eye',
        }
    },

    props: {
        tooltip: {
            required: false,
            type: String,
        },
        inputName: {
            required: true,
            type: String,
        },
        inputLabel: {
            required: true,
            type: String,
        },
        inputError: {
            required: false,
            type: String,
        },
    },

    methods: {
        changeType() {
            if (this.type === 'password') {
                this.type = 'text'
                this.icon = 'fa-eye-slash'
            } else {
                this.type = 'password'
                this.icon = 'fa-eye'
            }
        },
    },
}
</script>