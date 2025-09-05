import { ProjectCommand, selectSelectProject, useAppSelector } from '@blueskyfish/pierflow/stores';
import { ProjectPath } from '@blueskyfish/pierflow/utils';
import * as React from 'react';
import { ProjectAllow } from './project';
import { Checkout } from '../components/Checkout.tsx';

export const ProjectCheckoutPage: React.FC = () => {
  const project = useAppSelector(selectSelectProject)!;
  return (
    <ProjectAllow command={ProjectCommand.CheckoutRepository} project={project} projectKey={ProjectPath.Checkout}>
      <Checkout project={project} />
    </ProjectAllow>
  );
};
