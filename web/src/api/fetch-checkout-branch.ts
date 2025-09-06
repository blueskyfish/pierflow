import type { CheckoutPayload } from './entities.ts';
import { fetchingHeaders } from './fetch-user.ts';
import axios from 'axios';

export const fetchCheckoutBranch = async (projectId: string, payload: CheckoutPayload): Promise<void> => {
  const options = {
    url: `/api/projects/${projectId}/checkout`,
    method: 'PUT',
    headers: {
      ...fetchingHeaders,
    },
    data: payload,
  };
  await axios.request(options);
};
