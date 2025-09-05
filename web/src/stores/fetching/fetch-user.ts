const getUserIdByLocalStorage = () => {
  const userId = localStorage.getItem('blueskyfish.pierflow.userId');
  if (!userId) {
    const newUserId = crypto.randomUUID();
    localStorage.setItem('blueskyfish.pierflow.userId', newUserId);
    return newUserId;
  }
  return userId;
};

export const fetchingHeaders = {
  'Content-Type': 'application/json',
  'x-pierflow-user': getUserIdByLocalStorage(),
};
