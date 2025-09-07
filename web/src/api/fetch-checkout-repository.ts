import type { CheckoutPayload } from './entities.ts';
import { errorHandling, fetchingHeaders } from './fetch-helpers.ts';
import axios from 'axios';

export const fetchCheckoutRepository = async (projectId: string, payload: CheckoutPayload): Promise<void> => {
  const options = {
    url: `/api/projects/${projectId}/checkout`,
    method: 'PUT',
    headers: {
      ...fetchingHeaders,
    },
    data: payload,
  };
  try {
    await axios.request(options);
  } catch (error: any) {
    return errorHandling(error, `/api/projects/${projectId}/checkout`);
  }
};
