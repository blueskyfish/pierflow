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
