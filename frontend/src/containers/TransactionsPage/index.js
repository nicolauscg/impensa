import React, { useState } from "react";
import { useFormik } from "formik";
import * as R from "ramda";

import { Button, Typography, Box } from "@material-ui/core";

import {
  useAxiosSafely,
  urlGetAllTransactions,
  urlPostCreateTransaction,
  urlGetAllAccounts,
  urlGetAllCategory
} from "../../api";
import NewTransactionModal from "../../components/CreateTransactionModal";
import DataTable, { DataTableFormatter } from "../../components/DataTable";

const TransactionsPage = () => {
  const [newTransactionModelIsOpen, setNewTransactionModelIsOpen] = useState(
    false
  );
  const handleOpenNewTransactionModal = () =>
    setNewTransactionModelIsOpen(true);
  const handleCloseNewTransactionModal = () =>
    setNewTransactionModelIsOpen(false);

  const [
    { data: transactionsData, loading: transactionsLoading }
  ] = useAxiosSafely(urlGetAllTransactions());
  const [{ data: accountsData, loading: accountsLoading }] = useAxiosSafely(
    urlGetAllAccounts()
  );
  const [{ data: categoriesData, loading: categoriesLoading }] = useAxiosSafely(
    urlGetAllCategory()
  );
  const loading = transactionsLoading || accountsLoading || categoriesLoading;

  const formikCreateTransaction = useFormik({
    initialValues: {
      amount: 0,
      description: "",
      date: new Date(),
      account: null,
      category: null,
      picture: null
    },
    onSubmit: values => {
      urlPostCreateTransaction({
        data: values
      }).then(handleCloseNewTransactionModal);
    }
  });

  const newTransactionModelProps = {
    loading,
    isOpen: newTransactionModelIsOpen,
    handleClose: handleCloseNewTransactionModal,
    formik: formikCreateTransaction,
    accounts: accountsData,
    categories: categoriesData
  };

  return (
    <>
      <Box className="d-flex">
        <Typography variant="h3" display="inline">
          Transactions
        </Typography>
        <Button
          variant="contained"
          color="primary"
          onClick={handleOpenNewTransactionModal}
          className="ml-3"
        >
          Add transaction
        </Button>
      </Box>
      <NewTransactionModal {...newTransactionModelProps} />
      <DataTable
        headerNames={["Date", "Description", "Category", "Account", "Amount"]}
        dataFormatters={[
          DataTableFormatter.formatDateFrom(R.prop("dateTime")),
          R.prop("description"),
          DataTableFormatter.mapFromLookup(
            R.prop("category"),
            categoriesData,
            R.prop("name")
          ),
          DataTableFormatter.mapFromLookup(
            R.prop("account"),
            accountsData,
            R.prop("name")
          ),
          R.prop("amount")
        ]}
        data={transactionsData}
      />
    </>
  );
};

export default TransactionsPage;
