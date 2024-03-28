require('dotenv').config();

var Airtable = require('airtable');
var base = new Airtable({apiKey: process.env.AIRTABLE_API_KEY }).base(process.env.AIRTABLE_BASE_ID);

const table = base(process.env.AIRTABLE_TABLE_NAME);


const createRecord = async(fields) => {
    const createRecord = await table.create(fields);
};

/* Form Submission */
const form = document.getElementById("contact__form");
const firstName = document.getElementById('fName').value;
const lastName = document.getElementById('lName').value;
const email = document.getElementById('email').value;
const url = document.getElementById('website').value;
const message = document.getElementById('message').value;

form.addEventListener("submit", function(event){
    
    event.preventDefault();

    //Submit form data to airtable
    createRecord({
        "first_name": firstName,
        "last_name": lastName,
        "email": email,
        "website": url,
        "message": message
    });
});

