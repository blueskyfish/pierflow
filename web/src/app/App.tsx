import { DashboardPage, ProjectDetailPage, ProjectListPage, RootPage } from '@blueskyfish/pierflow/pages';
import { RoutePath } from '@blueskyfish/pierflow/utils';
import * as React from 'react';
import { Route, Routes } from 'react-router';

export const App: React.FC = () => {
  return (
    <Routes>
      <Route element={<RootPage />}>
        <Route index element={<DashboardPage />} />
        <Route path={RoutePath.ProjectListPath} element={<ProjectListPage />} />
        <Route path={RoutePath.ProjectDetailPath} element={<ProjectDetailPage />} />
      </Route>
    </Routes>
  );
};
