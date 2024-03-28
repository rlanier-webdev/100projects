const btn = document.getElementById("button");

btn.addEventListener('click', function(){

    convert();

});

function convert() {
    var feet = Number(document.getElementById("feet").value);
    var inches = Number(document.getElementById("inches").value);

    var cm = 2.54 * ((feet * 12) + inches);

    document.getElementById("output__area").innerHTML = cm + " cm";

}

