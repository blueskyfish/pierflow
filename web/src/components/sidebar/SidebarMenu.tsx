import { RouteBuilder, RoutePath } from '@blueskyfish/pierflow/utils';
import * as React from 'react';
import { SidebarItem } from './SidebarItem.tsx';

export interface SideMenuItem {
  id: string;
  name: string;
}

export interface SidebarMenuProps {
  selected?: string;
  menu: SideMenuItem[];
}

export const SidebarMenu: React.FC<SidebarMenuProps> = ({ selected, menu }) => {
  return (
    <ul className={'menu width-100'}>
      <SidebarItem
        menuKey={'dashboard'}
        label={'Dashboard'}
        link={RoutePath.HomePath}
        icon={'mdi mdi-monitor-dashboard'}
        selected={'dashboard' === selected}
        disabled={false}
      />
      <SidebarItem
        menuKey={'project-list'}
        label={'Project Overview'}
        link={RoutePath.ProjectListPath}
        icon={'mdi mdi-list-box-outline'}
        selected={'project-list' === selected}
        disabled={false}
      />
      {menu.map((item) => (
        <SidebarItem
          key={item.id}
          menuKey={item.id}
          label={item.name}
          link={RouteBuilder.buildProjectHomePath(item.id)}
          icon={'mdi mdi-factory'}
          selected={item.id === selected}
          disabled={false}
        />
      ))}
      <SidebarItem
        menuKey={'sse'}
        label={'SSE Test'}
        link={'/sse'}
        icon={'mdi mdi-server'}
        selected={'sse' === selected}
        disabled={false}
      />
    </ul>
  );
};
