import React, { useState } from "react";
import { useFormik } from "formik";
import * as R from "ramda";

import { Button, Typography, Box } from "@material-ui/core";

import {
  useAxiosSafely,
  urlGetAllTransactions,
  urlCreateTransaction,
  urlGetAllAccounts,
  urlGetAllCategory,
  urlUpdateTransaction,
  urlDeleteTransaction
} from "../../api";
import DataTable, { DataTableFormatter } from "../../components/DataTable";
import {
  cleanNilFromObject,
  transformValuesToUpdatePayload
} from "../../ramdaHelpers";
import CreateOrEditModal, {
  FormFields,
  FormTypes
} from "../../components/CreateOrEditModal";

const TransactionsPage = () => {
  const initialModalData = {
    amount: 0,
    description: "",
    dateTime: new Date(),
    account: null,
    category: null,
    picture: null
  };

  const [newTransactionModelIsOpen, setNewTransactionModelIsOpen] = useState(
    false
  );
  const [modalData, setModalData] = useState(initialModalData);
  const [modalMode, setModalMode] = useState(FormTypes.CREATE);

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
  const [, createTransaction] = useAxiosSafely(urlCreateTransaction());
  const [, updateTransaction] = useAxiosSafely(urlUpdateTransaction());
  const [, deleteTransaction] = useAxiosSafely(urlDeleteTransaction());

  const formikTrasaction = useFormik({
    initialValues: modalData,
    enableReinitialize: true,
    onSubmit: values => {
      const cleanedValues = R.evolve(
        {
          amount: R.ifElse(R.isEmpty, R.always(0), R.identity)
        },
        values
      );
      switch (modalMode) {
        case FormTypes.CREATE:
          createTransaction({
            data: cleanNilFromObject(cleanedValues)
          }).then(() => {
            handleCloseNewTransactionModal();
            refetchTransactions();
          });
          break;
        case FormTypes.UPDATE:
          updateTransaction({
            data: R.pipe(
              cleanNilFromObject,
              transformValuesToUpdatePayload
            )(cleanedValues)
          }).then(() => {
            handleCloseNewTransactionModal();
            refetchTransactions();
          });
          break;
        default:
          throw new Error("unrecognized FormType in modalMode");
      }
    }
  });

  const resetModalData = () => setModalData(initialModalData);
  const handleOpenNewTransactionModal = () =>
    setNewTransactionModelIsOpen(true);
  const handleCloseNewTransactionModal = () => {
    setNewTransactionModelIsOpen(false);
  };
  const handleOnCreate = () => {
    setModalMode(FormTypes.CREATE);
    resetModalData();
    handleOpenNewTransactionModal();
  };
  const handleOnEdit = data => {
    setModalMode(FormTypes.UPDATE);
    setModalData(data);
    handleOpenNewTransactionModal();
  };
  const handleOnDelete = data => {
    deleteTransaction({ data: { ids: Array.of(data.id) } }).then(() =>
      refetchTransactions()
    );
  };

  const loading = transactionsLoading || accountsLoading || categoriesLoading;
  const createOrEditTransactionModalProps = {
    title: "New Transaction",
    data: modalData,
    loading,
    isOpen: newTransactionModelIsOpen,
    handleClose: handleCloseNewTransactionModal,
    formik: formikTrasaction,
    formType: modalMode,
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
  const transactionDataTableProps = {
    headerNames: ["Date", "Description", "Category", "Account", "Amount"],
    dataFormatters: [
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
    ],
    data: transactionsData,
    onEdit: handleOnEdit,
    onDelete: handleOnDelete
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
          onClick={handleOnCreate}
          className="ml-3"
        >
          Add transaction
        </Button>
      </Box>
      <CreateOrEditModal {...createOrEditTransactionModalProps} />
      <DataTable {...transactionDataTableProps} />
    </>
  );
};

export default TransactionsPage;
