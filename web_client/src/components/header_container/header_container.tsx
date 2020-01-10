import * as React from "react";

import * as Styles from "./header_container.styles.css";

type Props = {
  children: React.ReactElement | React.ReactElement[];
};

const HeaderContainer = (props: Props) => {
  return <header className={Styles.header_container}>{props.children}</header>;
};

export { HeaderContainer };
