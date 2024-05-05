/*Simplify v1.1*/
var htmlElement = document.querySelector('html');

window.onload = toggleTheme('onload');
export function toggleTheme(type) {
    const body = document.getElementsByTagName("body")[0];
    if (type == 'onload') {
        if (localStorage.getItem("theme")) {
            const savedTheme = localStorage.getItem("theme");
            if (savedTheme == "light") {
                theme('light')
            } else if (savedTheme == "dark") {
                theme('dark')
            }
        } else {
            if (body.classList.contains("dark")) {
                theme('dark')
            } else if (body.classList.contains("light")) {
                theme('light')
            }
        }
    } else {
        if (body.classList.contains("dark")) {
            theme('light')
        } else if (body.classList.contains("light")) {
            theme('dark')
        } else {
            body
                .classList
                .remove("dark", "light");
            theme('light')
        }
    }
}

export function theme(type) {
    const body = document.getElementsByTagName("body")[0];
    htmlElement.setAttribute('style', 'color-scheme: ' + type + ';');
    if (type == 'dark') {

        body
            .classList
            .remove("light");

    } else {

        body
            .classList
            .remove("dark");

    }
    body
        .classList
        .add(type);
    localStorage.setItem("theme", type)
}
window
    .addEventListener('load', function () {
        document
            .body
            .classList
            .remove('preload')
    });


