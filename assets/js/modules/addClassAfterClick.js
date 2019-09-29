/**
 * Show certain element by adding active class and
 * removing after another click event
 *
 * @param {object} el that u want to show
 * @param {object} btn that u want to click
 * @param {string} className
 * @param {callback} onOpen
 * @param {callback} onClose
 * @return {void}
 */
export default function(el, btn, className = 'active', onOpen, onClose) {
    btn.addEventListener('click', e => {
        if (el.classList.contains(className)) {
            if (Object.prototype.toString.call(onClose) == "[object Function]") {
                onClose(e)
            }
            el.classList.remove(className)
        } else {
            if (Object.prototype.toString.call(onOpen) == "[object Function]") {
                onOpen(e)
            }
            el.classList.add(className)
        }
    })
}
