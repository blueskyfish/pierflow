import React, { useEffect, useState } from 'react';
import { HeadLine, ScrollBar, ScrollingDirection } from '@blueskyfish/pierflow/components';
import { ProjectEdit, type ProjectEditData } from './project';
import { setError, updatePageKey, updateProjectDetail, useAppDispatch } from '@blueskyfish/pierflow/stores';
import { type ErrorDto, fetchCreateProject } from '@blueskyfish/pierflow/api';
import { useNavigate } from 'react-router';
import { RouteBuilder } from '@blueskyfish/pierflow/utils';

export const ProjectCreatePage: React.FC = () => {
  const [loading, setLoading] = useState(false);
  const dispatch = useAppDispatch();
  const navigate = useNavigate();

  const project = {
    name: '',
    description: '',
    path: '',
    gitUrl: '',
    user: '',
    token: '',
  };

  const sendProjectCreate = (data: ProjectEditData) => {
    console.log('Create Project', data);
    setLoading(true);
    fetchCreateProject(data)
      .then((project) => {
        dispatch(updateProjectDetail(project));
        setTimeout(() => {
          navigate(RouteBuilder.buildProjectHomePath(project.id));
        }, 100);
      })
      .catch((error: ErrorDto) => {
        dispatch(setError(error));
      })
      .finally(() => setLoading(false));
  };

  useEffect(() => {
    dispatch(updatePageKey('create-project'));
  }, [dispatch]);

  return (
    <div className={'flex flex-col items-stretch h-full overflow-auto'}>
      <HeadLine
        title={`Create New Project`}
        as={'h2'}
        icon={`mdi mdi-new-box`}
        className={'mb-4 px-3 pt-3 flex-shrink-1'}
        loading={loading}
      />
      <ScrollBar direction={ScrollingDirection.Vertical} className={'flex-grow-1'}>
        <div className={'p-3 w-[40rem]'}>
          <ProjectEdit data={project} onSend={(data) => sendProjectCreate(data)} />
        </div>
      </ScrollBar>
    </div>
  );
};
