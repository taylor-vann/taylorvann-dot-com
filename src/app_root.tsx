import * as React from "react";

export interface HelloWorldProps {
  title: string;
  subject: string;
}

const AppRoot = (props: HelloWorldProps) => {
  return (
    <div>
      <p>{`${props.title}, ${props.subject}!`}</p>
    </div>
  );
};

export {AppRoot};
