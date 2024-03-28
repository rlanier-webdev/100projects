const billAmount = document.getElementById("billAmt");
const tipPercentage = document.getElementById("tipPercentage");

billAmount.value="0.00";

billAmount.addEventListener('keyup', function(){
    
    calc();

});

tipPercentage.addEventListener('change', function(){

    calc();

});

function calc() {
    
    const bill = Number(billAmount.value);
    
    var tipAmount = bill * Number(tipPercentage.value);
    var total = bill + Number(tipAmount);

    if (isNaN(bill)) {
        document.getElementById("tipAmt").value = "0.00";
        document.getElementById("total").value = "0.00";
    } else {
        document.getElementById("tipAmt").value = Number(tipAmount).toFixed(2);
        document.getElementById("total").value = Number(total).toFixed(2);
    }
}

