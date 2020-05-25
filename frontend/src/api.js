import useAxios from "axios-hooks";
import * as R from "ramda";
import { useEffect, useState } from "react";

const baseUrl = process.env.REACT_APP_BACKENDURL;

const RESPONSE_TYPE = {
  OBJECT: true,
  ARRAY: false
};

export const urlLogin = () => ({
  url: `${baseUrl}/auth/login`,
  method: "POST",
  type: RESPONSE_TYPE.OBJECT,
  manual: true
});
export const urlRegister = () => ({
  url: `${baseUrl}/auth/register`,
  method: "POST",
  type: RESPONSE_TYPE.OBJECT,
  manual: true
});
export const urlGetUser = userId => ({
  url: `${baseUrl}/user/${userId}`,
  method: "GET",
  type: RESPONSE_TYPE.OBJECT,
  manual: true
});
export const urlUpdateUser = () => ({
  url: `${baseUrl}/user`,
  method: "PUT",
  type: RESPONSE_TYPE.OBJECT,
  manual: true
});
export const urlGetAllTransactions = () => ({
  url: `${baseUrl}/transaction`,
  method: "GET",
  type: RESPONSE_TYPE.ARRAY,
  manual: true
});
export const urlGetAllTransactionsForTable = () => ({
  url: `${baseUrl}/transaction/table`,
  method: "GET",
  type: RESPONSE_TYPE.ARRAY,
  manual: true
});
export const urlGetEditTransaction = id => ({
  url: `${baseUrl}/transaction/edit/${id}`,
  method: "GET",
  type: RESPONSE_TYPE.OBJECT,
  manual: true
});
export const urlGetCreateTransaction = () => ({
  url: `${baseUrl}/transaction/create`,
  method: "GET",
  type: RESPONSE_TYPE.OBJECT,
  manual: true
});
export const urlCreateTransaction = () => ({
  url: `${baseUrl}/transaction`,
  method: "POST",
  type: RESPONSE_TYPE.OBJECT,
  manual: true
});
export const urlUpdateTransaction = () => ({
  url: `${baseUrl}/transaction`,
  method: "PUT",
  type: RESPONSE_TYPE.OBJECT,
  manual: true
});
export const urlDeleteTransaction = () => ({
  url: `${baseUrl}/transaction`,
  method: "DELETE",
  type: RESPONSE_TYPE.OBJECT,
  manual: true
});
export const urlGetAllAccounts = () => ({
  url: `${baseUrl}/account`,
  method: "GET",
  type: RESPONSE_TYPE.ARRAY,
  manual: true
});
export const urlCreateAccount = () => ({
  url: `${baseUrl}/account`,
  method: "POST",
  type: RESPONSE_TYPE.OBJECT,
  manual: true
});
export const urlUpdateAccount = () => ({
  url: `${baseUrl}/account`,
  method: "PUT",
  type: RESPONSE_TYPE.OBJECT,
  manual: true
});
export const urlDeleteAccount = () => ({
  url: `${baseUrl}/account`,
  method: "DELETE",
  type: RESPONSE_TYPE.OBJECT,
  manual: true
});
export const urlGetAllCategories = () => ({
  url: `${baseUrl}/category`,
  method: "GET",
  type: RESPONSE_TYPE.ARRAY,
  manual: true
});
export const urlCreateCategory = () => ({
  url: `${baseUrl}/category`,
  method: "POST",
  type: RESPONSE_TYPE.OBJECT,
  manual: true
});
export const urlUpdateCategory = () => ({
  url: `${baseUrl}/category`,
  method: "PUT",
  type: RESPONSE_TYPE.OBJECT,
  manual: true
});
export const urlDeleteCategory = () => ({
  url: `${baseUrl}/category`,
  method: "DELETE",
  type: RESPONSE_TYPE.OBJECT,
  manual: true
});
export const urlGraphTransactionCategory = () => ({
  url: `${baseUrl}/graph/transaction/category`,
  method: "GET",
  type: RESPONSE_TYPE.ARRAY,
  manual: true
});
export const urlGraphTransactionAccount = () => ({
  url: `${baseUrl}/graph/transaction/account`,
  method: "GET",
  type: RESPONSE_TYPE.ARRAY,
  manual: true
});
export const urlImportTransaction = () => ({
  url: `${baseUrl}/transaction/import`,
  method: "POST",
  type: RESPONSE_TYPE.OBJECT,
  manual: true
});
export const urlExportTransaction = () => ({
  url: `${baseUrl}/transaction/export`,
  method: "GET",
  type: RESPONSE_TYPE.OBJECT,
  manual: true
});

/*
  wrapper for fetching data that returns
  empty object or array as data based on type in requestInfo instead of undefined,
  error message when backend returns invalid error response as normalization
*/
export const useAxiosSafely = requestInfo => {
  const { type, manual, ...rest } = requestInfo;
  const defaultWrappedData = type === RESPONSE_TYPE.OBJECT ? {} : [];
  const [wrappedPaging, setWrappedPaging] = useState({});
  const [wrappedData, setWrappedData] = useState(defaultWrappedData);
  const [wrappedError, setWrappedError] = useState({});

  const [{ data, loading, error, response }, refetch] = useAxios(
    rest,
    manual ? { manual } : {}
  );

  useEffect(() => {
    if (R.hasPath(["data"], data)) {
      setWrappedData(data.data);
    } else {
      setWrappedData(defaultWrappedData);
    }
    if (R.hasPath(["paging"], data)) {
      setWrappedPaging(data.paging);
    } else {
      setWrappedPaging({});
    }
    if (R.hasPath(["response", "data", "error"], error)) {
      // caught error
      setWrappedError(error.response.data.error);
    } else if (
      // uncaught error i.e backend crashed
      error !== undefined &&
      data === undefined &&
      !R.hasPath(["response", "data", "error"], error) &&
      !loading
    ) {
      setWrappedError({
        message: "unknown error response format from backend"
      });
    } else {
      // no error or not request not finished yet
      setWrappedError({});
    }
  }, [loading]);

  return [
    {
      data: wrappedData,
      paging: wrappedPaging,
      loading,
      error: wrappedError,
      response
    },
    refetch
  ];
};
