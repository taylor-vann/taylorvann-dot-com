import * as React from "react";

import { ContactButtons } from "./contact/contact_buttons";
import { SocialButtons } from "./social/social_buttons";
import { MainContentContainer } from "./main_content_container/main_container";
import { FooterContainer } from "./footer_container/footer_container";
import { HeaderContainer } from "./header_container/header_container";
import { AppContainer } from "./app_container/app_container";
import { ScrollContentContainer } from "./scroll_content_container/scroll_content_container";

const App: React.FunctionComponent = () => {
  return (
    <AppContainer>
      <HeaderContainer>
        <p>i'm brian</p>
        <ContactButtons />
        <SocialButtons />
      </HeaderContainer>
      <MainContentContainer>
        <ScrollContentContainer>
          <p>this is a website</p>
        </ScrollContentContainer>
        <FooterContainer>
          <p>-</p>
          <p>this is a footer</p>
        </FooterContainer>
      </MainContentContainer>
    </AppContainer>
  );
};

export { App };
