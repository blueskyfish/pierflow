import axios from 'axios';

export interface BranchDto {
  branch: string;
  place: string;
  active: boolean;
}

export const fetchBranchList = async (projectId: string, refresh: boolean = false): Promise<void> => {
  const userId = localStorage.getItem('blueskyfish.pierflow.userId') ?? '4711'; // FIXME: replace with actual user ID)^^
  const url = refresh ? `/api/projects/${projectId}/branches?refresh=true` : `/api/projects/${projectId}/branches`;
  const options = {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
      'x-pierflow-user': userId,
    },
    url,
  };
  try {
    await axios.request(options);
    return;
  } catch (e) {
    console.error('> Failed to load branch list', e);
    return;
  }
};
