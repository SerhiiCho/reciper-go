setTimeout(() => {
    let dropdowns = M.Dropdown.init(document.querySelectorAll('.dropdown-trigger'), {
        coverTrigger: false,
        closeOnClick: false,
        alignment: 'left',
    });

    let menu = document.getElementById('adminka-collapsible')

    if (menu) {
        menu.addEventListener('click', () => dropdowns.forEach(menu => setTimeout(() => menu.recalculateDimensions(), 200)))
    }

    M.Tooltip.init(document.querySelectorAll('.tooltipped'), {
        position: 'top',
    });

    M.Sidenav.init(document.querySelectorAll('.sidenav'), {
        draggable: false,
    });

    M.Collapsible.init(document.querySelectorAll('.collapsible'));
    M.CharacterCounter.init(document.querySelectorAll('.counter'))

    M.Modal.init(document.querySelectorAll('.modal'), {
        opacity: 0.7,
    });
}, 100)