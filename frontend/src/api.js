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
export const urlGetAllTransactions = () => ({
  url: `${baseUrl}/transaction`,
  method: "GET",
  type: RESPONSE_TYPE.ARRAY,
  manual: false
});
export const urlPostCreateTransaction = () => ({
  url: `${baseUrl}/transaction`,
  method: "POST",
  type: RESPONSE_TYPE.OBJECT,
  manual: true
});
export const urlGetAllAccounts = () => ({
  url: `${baseUrl}/account`,
  method: "GET",
  type: RESPONSE_TYPE.ARRAY,
  manual: false
});
export const urlGetAllCategory = () => ({
  url: `${baseUrl}/category`,
  method: "GET",
  type: RESPONSE_TYPE.ARRAY,
  manual: false
});

/*
  wrapper for fetching data that returns
  empty object or array as data based on type in requestInfo instead of undefined,
  error message when backend returns invalid error response as normalization
*/
export const useAxiosSafely = requestInfo => {
  const { type, manual, ...rest } = requestInfo;
  const defaultWrappedData = type === RESPONSE_TYPE.OBJECT ? {} : [];
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
      loading,
      error: wrappedError,
      response
    },
    refetch
  ];
};
