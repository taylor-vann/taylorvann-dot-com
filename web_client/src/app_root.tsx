import * as React from "react";
import * as ReactDOM from "react-dom";
import { App } from "./components/app";

const AppRoot: React.FunctionComponent = () => {
  return <App />;
};

ReactDOM.render(<AppRoot />, document.getElementById("root"));
