<template>
    <div class="mt-4" style="min-width:700px">
        <div v-if="loader"
            class="valign-wrapper position-absolute center"
            style="top:50%;left:50%;transform:translate(-50%,-50%)"
        >
            <slot name="loader"></slot>
        </div>
        <line-chart :chart-data="data" :height="120" :options="options" />
    </div>
</template>

<script>
import LineChart from './LineChart.js';

export default {
    data() {
        return {
            data: [],
            options: {},
            loader: true,
        };
    },

    mounted() {
        this.fetchData();
    },

    methods: {
        fetchData() {
            this.$axios.get('/popularity-chart')
                .then(data => {
                    this.loader = false;
                    this.data = data.data;
                    this.options = data.data.options;
                })
                .catch(err => console.error(err));
        },
    },

    components: {
        LineChart,
    },
};
</script>