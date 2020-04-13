const baseUrl = process.env.REACT_APP_BACKENDURL;

export const urlLogin = () => ({
  url: `${baseUrl}/auth/login`,
  method: "POST"
});
export const urlGetAllTransactions = () => ({ url: `${baseUrl}/transaction` });
