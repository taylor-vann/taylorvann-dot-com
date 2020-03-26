import * as React from "react";

type LoginFormStateType = {
  username?: string;
  password?: string;
};

const handleSubmit = (e: React.FormEvent) => {
  e.preventDefault();
  console.log("handled submit yo!");
};

const LoginForm: React.FunctionComponent = () => {
  console.log("Render");

  return (
    <form onSubmit={handleSubmit}>
      <input type="text" name="username" />
      <input type="password" name="password" />
      <input type="submit" />
    </form>
  );
};

export { LoginForm };
