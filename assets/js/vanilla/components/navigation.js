import axios from 'axios'
import addClassAfterClick from '../../modules/addClassAfterClick'

;(function SearchFromWillAppearAfterBtnWillBeClicked() {
    const btn = document.getElementById('nav-btn-for-search')
    const el = document.getElementById('nav-search-form')

    if (btn && el) {
        const input = el.querySelector('#search-input')
        addClassAfterClick(el, btn, 'active', () => input.focus(), () => input.blur())
    }
})()

;(function MarkNotificationsAsRead() {
    const button = document.getElementById('mark-notifs-as-read')
    const alertIcon = document.querySelector('#mark-notifs-as-read span')

    if (button && alertIcon) {
        button.addEventListener('click', () => {
            alertIcon.classList.remove('small-notif')
            axios.put('/invokes/notifications').catch(e => console.error(e))
        })
    }
})()
