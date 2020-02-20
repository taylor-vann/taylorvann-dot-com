import * as React from "react";

import * as Styles from "./footer_container.styles.css";

type Props = {
  children: React.ReactElement | React.ReactElement[];
};

const FooterContainer = (props: Props) => {
  return <footer className={Styles.footer_container}>{props.children}</footer>;
};

export { FooterContainer };
