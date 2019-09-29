const mix = require('laravel-mix')

const css = 1
const js = 1
const chartjs = 0

if (css == 1) {
    mix.sass('assets/css/app.sass', 'public/assets/app.css').options({
        processCssUrls: false,
    })
}

if (js == 1) {
    mix.js('assets/js/app.js', 'public/assets/app.js')
        .options({
            uglify: {
                uglifyOptions: {
                    compress: {
                        drop_console: true
                    }
                }
            }
        })
        .sourceMaps();
}

if (chartjs == 1) {
    mix.js('assets/js/vue/chart/chart.js', 'public/assets/chart.js')
        .options({
            uglify: {
                uglifyOptions: {
                    compress: {
                        drop_console: true
                    }
                }
            }
        })
}