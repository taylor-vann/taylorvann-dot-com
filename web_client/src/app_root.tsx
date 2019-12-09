// Brian Taylor Vann
// taylorvann dot com

// This is the connection between the View and the ViewModel
// In this case, the View is the DOM and the ViewModel is React

import * as React from "react";
import * as ReactDOM from "react-dom";
import { App } from "./components/app";

const AppRoot: React.FunctionComponent = () => {
  return <App />;
};

ReactDOM.render(<AppRoot />, document.getElementById("root"));
