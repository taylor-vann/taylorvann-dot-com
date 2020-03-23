import * as React from "react";
import { AppContainer } from "./app_container/app_container";
import { ContactButtons } from "./contact/contact_buttons";
import { FooterContainer } from "./footer_container/footer_container";
import { HeaderContainer } from "./header_container/header_container";
import { MainContentContainer } from "./main_content_container/main_container";
import { ScrollContentContainer } from "./scroll_content_container/scroll_content_container";
import { SocialButtons } from "./social/social_buttons";

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
