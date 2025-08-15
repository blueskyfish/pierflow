import * as React from 'react';
import { Link } from 'react-router';

export interface SidebarButtonProps {
  menuKey: string;
  label: string;
  link: string;
  icon: string;
  selected: boolean;
  disabled: boolean;
}

export const SidebarItem: React.FC<SidebarButtonProps> = ({ menuKey, label, link, icon, selected, disabled }) => {
  return (
    <li key={menuKey} className={'mb-1 last-of-type:mb-0'}>
      <Link
        to={link}
        className={`flex flex-row align-items-center p-2 ${selected ? 'menu-active' : ''} ${disabled ? 'menu-disabled' : ''}`}
        aria-disabled={disabled}
      >
        <span className={`app-icon ${icon} flex-shrink-1`}></span>
        <span className={'app-title ml-2 flex-grow-1 text-truncate'}>{label}</span>
        <span className={'mdi mdi-chevron-right flex-shrink-1'}></span>
      </Link>
    </li>
  );
};
