/*  
    Project: Time to Decimal Hours
    Author: RaShunda Lanier
    Email: rashunda@citygirlcodes.me

    Convert time (hours/minutes) to decimal hours
    1 min = .02 decimal hours, 2 min = .03 decimal hours, 25 mins = .42 decimal hours
        
        Example: 3 hours and 45 minutes
        1. Divide min by 60; 45/60 = 0.75
        2. Add hours; 3 + 0.75 = 3.75 hours

        - hours
        - mins
        - time
        - decimal hours

    Validate Input
    - If input is NAN then display error as the message instead
*/



function calcHours() {

    //1. get hours input
    var hours = parseFloat(document.getElementById("hrs").value); 
    //2. get mins input
    var mins = parseFloat(document.getElementById("mins").value); 

    //Validate Input: If input is not a number then display error as the message instead
    if (isNaN(hours) || isNaN(mins)) {
        document.getElementById("message").style.color = "red";
        document.getElementById("message").style.fontSize = "18px";
        document.getElementById("message").innerHTML = "Error: Please enter a number.";
        
    } else {
        //3. convert to decimal hours
        var dHours = hours + (mins / 60);

        document.getElementById("message").style.color = "#1A4645";
        document.getElementById("message").style.fontSize = "25px";
        document.getElementById("message").innerHTML = dHours.toFixed(2) + " Decimal Hours";
    }
}
