import addClassAfterClick from '../../modules/addClassAfterClick'

;(function PopupWindowShowMore() {
    const btn = document.getElementById('popup-window-trigger')
    const target = document.getElementById('popup-window')

    if (!target && !btn) return

    const icon = btn.querySelector('i.fas')

    addClassAfterClick(target, btn, 'active',
        () => swapClass('fa-ellipsis-v', 'fa-times', icon),
        () => swapClass('fa-times', 'fa-ellipsis-v', icon)
    )

    function swapClass(swapFrom, swapTo, elem) {
        elem.classList.remove(swapFrom)
        elem.classList.add(swapTo)
    }
})()
