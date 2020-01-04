import * as React from "react";

import * as Styles from "./app_container.styles.css";

type Props = {
  children: React.ReactElement | React.ReactElement[];
};

const AppContainer = (props: Props) => {
  return <section className={Styles.app_container}>{props.children}</section>;
};

export { AppContainer };
