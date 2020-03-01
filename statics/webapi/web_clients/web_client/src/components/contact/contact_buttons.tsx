import * as React from "react";

const ContactButtons: React.FunctionComponent = () => {
  return (
    <section>
      <ul>
        <li>
          <a href="https://www.linkedin.com/in/brian-vann/" target="_blank">
            linked-in
          </a>
        </li>
        <li>
          <a href="mailto:brian.t.vann@gmail.com">email</a>
        </li>
      </ul>
    </section>
  );
};

export { ContactButtons };
