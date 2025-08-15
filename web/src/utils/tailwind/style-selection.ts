export const selectBackgroundFrom = (selected: boolean): string => {
  return selected ? 'bg-primary hover:bg-primary-' : 'bg-base-100 hover:bg-base-200';
};

export const selectTextColorFrom = (selected: boolean): string => {
  return selected ? 'text-primary-content' : 'text-base-content';
};
