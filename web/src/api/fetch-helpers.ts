/**
 * Returns the user id from local storage, or creates a new one if it doesn't exist.
 *
 * The user id is stored in local storage under the key 'blueskyfish.pierflow.userId'.
 */
const getUserIdByLocalStorage = () => {
  const userId = localStorage.getItem('blueskyfish.pierflow.userId');
  if (!userId) {
    const newUserId = crypto.randomUUID();
    localStorage.setItem('blueskyfish.pierflow.userId', newUserId);
    return newUserId;
  }
  return userId;
};

/**
 * The headers to use for fetching API requests, including the user id.
 */
export const fetchingHeaders = {
  'Content-Type': 'application/json',
  'x-pierflow-user': getUserIdByLocalStorage(),
};

/**
 * Handles errors from API requests, extracting relevant information.
 * @param error The error object from the API request.
 * @param type A string indicating the type or context of the error.
 * @returns A rejected promise with an object containing the type, status, and message of the error.
 */
export const errorHandling = (error: any, type: string): Promise<never> => {
  if (error.response) {
    const { status, data } = error.response;
    let message = '';
    if (data && typeof data === 'object') {
      message = data.message || data.error || data.detail || JSON.stringify(data);
    }
    return Promise.reject({
      type,
      status,
      message,
    });
  }
  return Promise.reject(error);
};
