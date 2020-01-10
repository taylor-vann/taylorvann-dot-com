import * as React from "react";

import * as Styles from "./scroll_content_container.styles.css";

type Props = {
  children: React.ReactElement | React.ReactElement[];
};

const ScrollContentContainer = (props: Props) => {
  return (
    <section className={Styles.scroll_content_container}>
      {props.children}
    </section>
  );
};

export { ScrollContentContainer };
