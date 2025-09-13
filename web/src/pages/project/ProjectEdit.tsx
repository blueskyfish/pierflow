import React, { type ChangeEvent, type FormEvent, useEffect, useState } from 'react';
import { InputControl, TextAreaControl } from '@blueskyfish/pierflow/components';

export type ProjectEditData = {
  name: string;
  description: string;
  path: string;
  giturl: string;
  user: string;
  token: string;
};

interface ProjectEditProps {
  data: Partial<ProjectEditData>;
  onSend: (data: ProjectEditData) => void;
}

export const ProjectEdit: React.FC<ProjectEditProps> = ({ data, onSend }) => {
  const [formData, setFormData] = useState<ProjectEditData>({
    name: '',
    description: '',
    path: '',
    giturl: '',
    user: '',
    token: '',
  });

  useEffect(() => {
    setFormData((prev) => ({ ...prev, ...data }));
  }, [data]);

  const handleChange = (e: ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
    const { name, value } = e.target;
    setFormData((prev) => ({ ...prev, [name]: value }));
  };

  const handleSubmit = (e: FormEvent) => {
    e.preventDefault();
    onSend(formData);
  };

  return (
    <>
      <form onSubmit={handleSubmit} className={'form'} autoComplete={'off'}>
        <div
          style={{
            display: 'grid',
            gridTemplateColumns: '1fr 1fr',
            gap: '1rem',
            gridTemplateAreas: `
              "name description"
              "path description"
              "gitUrl gitUrl"
              "user token"
            `,
          }}
        >
          <div style={{ gridArea: 'name' }}>
            <InputControl
              type='text'
              name='name'
              label='Project Name'
              value={formData.name}
              placeholder={'A human-friendly name...'}
              required
              helpText={'A human-friendly name for the project.'}
              onChange={handleChange}
            />
          </div>
          <div style={{ gridArea: 'path' }}>
            <InputControl
              type='text'
              name='path'
              label='Project Path'
              value={formData.path}
              placeholder={'Unique project path...'}
              required
              helpText={'Path where the project will be stored. It must be unique in whole system.'}
              onChange={handleChange}
            />
          </div>
          <div style={{ gridArea: 'description' }}>
            <TextAreaControl
              label={'Project Description'}
              name={'description'}
              value={formData.description}
              rows={5}
              helpText={'A brief description of the project.'}
              placeholder={'Describe the project...'}
              onChange={handleChange}
            />
          </div>
          <div style={{ gridArea: 'gitUrl' }}>
            <InputControl
              type='text'
              name='giturl'
              label='Git Repository URL'
              value={formData.giturl}
              placeholder={'HTTPs Git repository URL...'}
              required
              helpText={'HTTPS URL of the Git repository'}
              onChange={handleChange}
            />
          </div>
          <div style={{ gridArea: 'user' }}>
            <InputControl
              type='text'
              name='user'
              label='Git Username'
              value={formData.user}
              placeholder={'Username with access to the repository...'}
              required
              helpText={'Username with access to the repository'}
              onChange={handleChange}
            />
          </div>
          <div style={{ gridArea: 'token' }}>
            <InputControl
              type='password'
              name='token'
              label='Git Access Token'
              value={formData.token}
              placeholder={'Personal Access Token...'}
              required
              helpText={'Personal Access Token with repo access'}
              onChange={handleChange}
            />
          </div>
        </div>
        <div className={'flex items-center justify-end mt-4'}>
          <span className={'text-red-600 font-bold mr-3'}>*</span>
          <span className={'text-neutral-600 text-xs flex-grow-1'}>Required input fields</span>
          <button type='submit' className={'btn btn-soft btn-primary'}>
            <span>{data.name === '' ? 'Create Project' : 'Update Project'}</span>
            <span className={'mdi mdi-chevron-right'} />
          </button>
        </div>
      </form>
    </>
  );
};
