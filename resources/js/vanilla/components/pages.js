import activeAfterClickBtn from '../../modules/addClassAfterClick'
import disableButton from '../../modules/disableButton'

/**
 * When visitor clicks search button, it will show the search form
 * instead of submitting, js prevents default behavior. After first click
 * it will set varible preventing to false, now after second click, it
 * will submit the search form
 */
;(function HomeHeaderSearchFormAppearsAfterButtonIsClicked() {
    const btn = document.getElementById('home-search-btn')
    const form = document.getElementById('home-search-form')

    if (form && btn) {
        activeAfterClickBtn(form, btn, 'active', e => e.preventDefault())
    }
})()

/**
 * This object auto updates pictures after
 * selecting them via file input
 */
;(function runImageUploader() {
    let target = document.getElementById('target-image')
    const src = document.getElementById('src-image')
    const fr = new FileReader()

    if (target && src) {
        src.addEventListener('change', () => {
            if (src.files.length !== 0) {
                fr.readAsDataURL(src.files[0])
                fr.onload = function () {
                    target.src = this.result
                }
            } else {
                target.src = '/storage/big/recipes/default.jpg'
            }
        })
    }
})()

;(function PreventSubmittingIfNotConfirmed() {
    const btns = document.querySelectorAll('.confirm')

    if (btns) {
        btns.forEach(btn => {
            btn.addEventListener('click', e => {
                if (!confirm(btn.getAttribute('data-confirm'))) {
                    e.preventDefault()
                }
            })
        })
    }
})()

;(function PreventMultipleFormSubmitting() {
    document.querySelectorAll('form').forEach(form => {
        form.addEventListener('submit', e => {
            e.target.querySelectorAll('button').forEach(btn => disableButton(btn))
        }, false)
    })
})()

;(function ShowProgressBarWhenLeavingThePage() {
    let bar = document.createElement('div')

    bar.classList.add('page-loading-bar')
    document.body.appendChild(bar)

    window.onbeforeunload = () => bar.classList.add('page-loading-bar--run')
})()
