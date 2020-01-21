// Brian Taylor Vann
// taylorvann dot com

// This is the connection between the View and the ViewModel
// In this case, the View is the DOM and the ViewModel is React

import * as React from "react";
import * as ReactDOM from "react-dom";

import { LoginForm } from "./components/login_form";

const AppRoot: React.FunctionComponent = () => {
  return (
    <div>
      <p>Hello there!</p>
      <LoginForm />
    </div>
  );
};

ReactDOM.render(<AppRoot />, document.getElementById("root"));
