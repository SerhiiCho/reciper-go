/**
 * @param {HTMLElement} button 
 */
export default function(button) {
    button.setAttribute('disabled', 'disabled')
    button.classList.add('disabled')
    button.classList.add('btn-loading')
}