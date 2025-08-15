import { DashboardPage, ProjectListPage, RootPage } from '@blueskyfish/pierflow/pages';
import * as React from 'react';
import { Route, Routes } from 'react-router';
import { ProjectDetailPage } from '../pages/ProjectDetailPage.tsx';
import { RoutePath } from '../utils/routing/route-names.ts';

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
