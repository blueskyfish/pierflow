import {
  DashboardPage,
  ProjectBuildPage,
  ProjectCheckoutPage,
  ProjectClonePage,
  ProjectDeletePage,
  ProjectDetailPage,
  ProjectHomePage,
  ProjectListPage,
  ProjectPullPage,
  ProjectStartPage,
  ProjectStopPage,
  RootPage
} from '@blueskyfish/pierflow/pages';
import { ProjectPath, RoutePath } from '@blueskyfish/pierflow/utils';
import * as React from 'react';
import { Route, Routes } from 'react-router';

export const App: React.FC = () => {
  return (
    <Routes>
      <Route element={<RootPage />}>
        <Route index element={<DashboardPage />} />
        <Route path={RoutePath.ProjectListPath} element={<ProjectListPage />} />
        <Route path={RoutePath.ProjectHomePath} element={<ProjectHomePage />}>
          <Route index element={<ProjectDetailPage />} />
          <Route path={RoutePath.ProjectCommandPath(ProjectPath.Clone)} element={<ProjectClonePage />} />
          <Route path={RoutePath.ProjectCommandPath(ProjectPath.Checkout)} element={<ProjectCheckoutPage />} />
          <Route path={RoutePath.ProjectCommandPath(ProjectPath.Build)} element={<ProjectBuildPage />} />
          <Route path={RoutePath.ProjectCommandPath(ProjectPath.Start)} element={<ProjectStartPage />} />
          <Route path={RoutePath.ProjectCommandPath(ProjectPath.Stop)} element={<ProjectStopPage />} />
          <Route path={RoutePath.ProjectCommandPath(ProjectPath.Pull)} element={<ProjectPullPage />} />
          <Route path={RoutePath.ProjectCommandPath(ProjectPath.Delete)} element={<ProjectDeletePage />} />
        </Route>
      </Route>
    </Routes>
  );
};
