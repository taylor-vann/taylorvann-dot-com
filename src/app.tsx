import * as React from "react";
import * as ReactDOM from "react-dom";
import { AppRoot } from "./app_root";

console.log('we are here');

const Index = () => {
  console.log('index');
  return <AppRoot title="hello" subject="world" />;
};

ReactDOM.render(<Index />, document.getElementById("root"));
