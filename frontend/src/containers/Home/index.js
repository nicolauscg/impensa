import React from "react";
import useAxios from "axios-hooks";
import { urlGetAllTransactions } from "../../api";

const Home = () => {
  const [
    {
      data: transactionData,
      loading: transactionLoading,
      error: transactionError
    }
  ] = useAxios(urlGetAllTransactions());

  return (
    <>
      <h1>home page</h1>
      <p>data: {JSON.stringify(transactionData)}</p>
      <p>loading: {JSON.stringify(transactionLoading)}</p>
      <p>error: {JSON.stringify(transactionError)}</p>
    </>
  );
};

export default Home;
