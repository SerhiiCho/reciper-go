<template>
    <div @mouseover="stopTimer" @mouseleave="startTimer">
        <transition name="appear" mode="out-in" appear>
            <div class="alert-message" v-if="visible">
                <div class="alert-message__inner">
                    <p v-html="message"></p>
                    <span @click="visible = false"><i class="fas fa-times"></i></span>
                </div> <!-- / .alert-message__inner -->
            </div> <!-- / .alert-message -->
        </transition>
    </div>
</template>

<script>
export default {
    data() {
        return {
            counter: 5,
            initialCounter: 5,
            visible: true,
            timerStopped: false,
            intervalId: 0,
        }
    },

    props: ['state', 'header', 'text'],

    created() {
        this.startTimer()
        this.setCounterValue()
    },

    methods: {
        setCounterValue() {
            const len = this.text.length

            if (len <= 30) {
                this.counter = 3
                this.initialCounter = 3
            } else if (len > 30 && len < 100) {
                this.counter = 5
                this.initialCounter = 5
            } else if (len >= 100) {
                this.counter = 7
                this.initialCounter = 7
            }
        },

        stopTimer() {
            clearInterval(this.intervalId)
            this.timerStopped = true
        },

        startTimer() {
            this.timerStopped = false

            this.intervalId = setInterval(() => {
                this.counter <= 0 ? this.visible = false : this.counter--
            }, 1000)
        }
    },

    computed: {
        message() {
            const statusColor = this.state == 'error' ? 'red-text' : 'green-text'
            const iconClass = this.state == 'error' ? 'fa-exclamation-circle' : 'fa-check'

            return `
                <i class="fas ${iconClass} ${statusColor}"></i> 
                <strong class="${statusColor}">${this.header}.</strong>
                ${this.text}`
        },
    },
}
</script>

<style lang="sass" scoped>
.appear
    &-enter-active,
    &-leave-active
        transition: transform 200ms

    &-enter,
    &-leave-to
        transform: rotateX(90deg)

    &-enter-to,
    &-leave
        transform: rotateX(0)
</style>