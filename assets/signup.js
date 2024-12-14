const passwordInput = document.querySelector('#password');
const showPass = document.querySelector('#show-password-btn');

showPass.addEventListener("click", () => {
    const inputType = passwordInput.getAttribute('type') === 'password' ? 'text' : 'password';
    passwordInput.setAttribute('type', inputType);
    showPass.textContent = inputType === 'password' ? 'Show' : 'Hide';
})