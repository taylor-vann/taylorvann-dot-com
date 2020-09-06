"use strict";
console.log("HELLO WORLD!");
const signInForm = document.querySelector("#form");
const submitButton = document.querySelector("#submit-button");
const handleSubmit = (e) => {
    e.preventDefault();
    if (signInForm) {
        const formData = new FormData(signInForm);
        const jsonFormData = Object.fromEntries(formData); // get entries from iterable
        jsonFormData["environment"] = "DEVELOPMENT";
        console.log(jsonFormData);
    }
};
// add event listener on click
submitButton === null || submitButton === void 0 ? void 0 : submitButton.addEventListener("click", handleSubmit);
// on click
// print form data
