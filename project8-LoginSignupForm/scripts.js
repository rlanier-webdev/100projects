//both buttons and container
const signUpBtn = document.getElementById('signUp');
const logInBtn = document.getElementById('logIn');
const container = document.getElementById('container');

//When the buttons are clicked we add or remove the .right-panel-active class
signUpBtn.addEventListener('click', () => {
    container.classList.add('right-panel-active');
});

logInBtn.addEventListener('click', () => {
    container.classList.remove('right-panel-active');
});