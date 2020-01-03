import * as React from "react";

import { ContactButtons } from "./contact/contact_buttons";
import { SocialButtons } from "./social/social_buttons";
import { MainContentContainer } from "./main_content_container/main_container";
import { AppContainer } from "./app_container/app_container";

const App: React.FunctionComponent = () => {
  return (
    <AppContainer>
      <MainContentContainer>
        <p>hello</p>
        <p>howdy</p>
      </MainContentContainer>
    </AppContainer>
  );
};

export { App };
