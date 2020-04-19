import cookie from "cookie";

export const isLoggedIn = () => {
  try {
    const userDataObj = parseAuthPayloadFromCookie();
    if (payloadIsValid(userDataObj)) {
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
    return parseAuthPayloadFromCookie();
  }

  return null;
};

export const clearUserObject = () => {
  delete_cookie("impensa");
};

const delete_cookie = name =>
  (document.cookie = `${name}=; Path=/; Expires=Thu, 01 Jan 1970 00:00:01 GMT;`);

const parseAuthPayloadFromCookie = () =>
  JSON.parse(cookie.parse(document.cookie).impensa);

const payloadIsValid = payload =>
  payload.id && payload.username && payload.token;
