const button = document.getElementById("button");

button.addEventListener("click", changeBG);

function changeBG() {
    document.body.style.backgroundColor = generateRandomColor();
}

function generateRandomColor()
{
    var randomColor = '#'+Math.floor(Math.random()*16777215).toString(16);
    return randomColor;
}