import AsyncLoader from '../../modules/AsyncLoader'

;(function RunAsyncLoaderModules() {
    let holders = document.querySelectorAll('[data-async-load]')
    holders.length > 0 ? new AsyncLoader(holders).start() : null
})()