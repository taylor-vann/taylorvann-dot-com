console.log("HELLO WORLD!");
const signInForm: HTMLFormElement | null = document.querySelector("#form");
const submitButton: HTMLElement | null = document.querySelector(
  "#submit-button"
);

const handleSubmit = (e: Event) => {
  e.preventDefault();

  if (signInForm) {
    const formData = new FormData(signInForm);
    const jsonFormData = Object.fromEntries(formData); // get entries from iterable
    jsonFormData["environment"] = "DEVELOPMENT";

    console.log(jsonFormData);
  }
};

// add event listener on click
submitButton?.addEventListener("click", handleSubmit);
// on click
// print form data
