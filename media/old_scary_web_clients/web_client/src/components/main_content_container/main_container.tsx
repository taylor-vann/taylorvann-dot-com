import * as React from "react";

import * as Styles from "./main_content_container.styles.css";

type Props = {
  children: React.ReactElement | React.ReactElement[];
};

const MainContentContainer = (props: Props) => {
  return (
    <main className={Styles.main_content_container}>{props.children}</main>
  );
};

export { MainContentContainer };
