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
import DataTable, { DataTableFormatter } from "../../components/DataTable";
import { cleanNilFromObject } from "../../ramdaHelpers";
import CreateOrEditModal, {
  FormFields
} from "../../components/CreateOrEditModal";

const TransactionsPage = () => {
  const [newTransactionModelIsOpen, setNewTransactionModelIsOpen] = useState(
    false
  );
  const handleOpenNewTransactionModal = () =>
    setNewTransactionModelIsOpen(true);
  const handleCloseNewTransactionModal = () =>
    setNewTransactionModelIsOpen(false);

  const [
    { data: transactionsData, loading: transactionsLoading },
    refetchTransactions
  ] = useAxiosSafely(urlGetAllTransactions());
  const [{ data: accountsData, loading: accountsLoading }] = useAxiosSafely(
    urlGetAllAccounts()
  );
  const [{ data: categoriesData, loading: categoriesLoading }] = useAxiosSafely(
    urlGetAllCategory()
  );
  const loading = transactionsLoading || accountsLoading || categoriesLoading;

  const [, postCreateTransaction] = useAxiosSafely(urlPostCreateTransaction());

  const formikCreateTransaction = useFormik({
    initialValues: {
      amount: 0,
      description: "",
      dateTime: new Date(),
      account: null,
      category: null,
      picture: null
    },
    onSubmit: values => {
      postCreateTransaction({
        data: cleanNilFromObject(values)
      }).then(() => {
        handleCloseNewTransactionModal();
        refetchTransactions();
      });
    }
  });

  const createOrEditTransactionModalProps = {
    title: "New Transaction",
    loading,
    isOpen: newTransactionModelIsOpen,
    handleClose: handleCloseNewTransactionModal,
    formik: formikCreateTransaction,
    formFields: [
      FormFields.textField({
        label: "Amount",
        type: "number",
        name: "amount"
      }),
      FormFields.dateField({
        label: "Date",
        name: "dateTime"
      }),
      FormFields.textField({
        label: "Description",
        type: "text",
        name: "description"
      }),
      FormFields.selectField({
        label: "Category",
        name: "category",
        options: categoriesData,
        optionDisplayer: R.prop("name")
      }),
      FormFields.selectField({
        label: "Account",
        name: "account",
        options: accountsData,
        optionDisplayer: R.prop("name")
      })
    ]
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
      <CreateOrEditModal {...createOrEditTransactionModalProps} />
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
