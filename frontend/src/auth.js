export const isLoggedIn = () => {
  try {
    const userDataObj = JSON.parse(localStorage.impensa);
    if (userDataObj.id && userDataObj.email && userDataObj.token) {
      return true;
    }
  } catch (err) {
    // pass
  }
  clearUserObject();

  return false;
};

export const getUserObject = () => {
  if (isLoggedIn()) {
    return JSON.parse(localStorage.impensa);
  }

  return null;
};

export const clearUserObject = () => {
  localStorage.removeItem("impensa");
};
