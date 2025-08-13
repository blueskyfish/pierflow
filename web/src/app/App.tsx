import { Brand, MainContent, Sidebar } from '@blueskyfish/pierflow/components';
import { RootPage } from '@blueskyfish/pierflow/pages';
import * as React from 'react';

export const App: React.FC = () => {
  return (
    <RootPage>
      <Sidebar>
        <Brand />
      </Sidebar>
      <MainContent>
        <p>Content</p>
      </MainContent>
    </RootPage>
  );
};
