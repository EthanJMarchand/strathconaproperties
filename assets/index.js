const menuBtn = document.querySelector('#menu-btn');
const menu = document.querySelector('#menu');

menuBtn.addEventListener("click", () => {
    menu.classList.toggle("h-0");
})
