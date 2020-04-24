import React, { useState, useEffect } from "react";
import { useFormik } from "formik";
import * as R from "ramda";

import { Button, Typography, Box } from "@material-ui/core";

import {
  useAxiosSafely,
  urlGetAllAccounts,
  urlUpdateAccount,
  urlDeleteAccount,
  urlCreateAccount
} from "../../api";
import DataTable from "../../components/DataTable";
import { transformValuesToUpdateIdsPayload } from "../../ramdaHelpers";
import CreateOrEditModal, {
  FormFields,
  FormTypes
} from "../../components/CreateOrEditModal";

export default function AccountsPage() {
  const initialModalData = {
    name: ""
  };

  const [newAccountModelIsOpen, setNewAccountModelIsOpen] = useState(false);
  const [modalData, setModalData] = useState(initialModalData);
  const [modalMode, setModalMode] = useState(FormTypes.CREATE);

  const [
    { data: accountsData, loading: accountsLoading },
    refetchAccounts
  ] = useAxiosSafely(urlGetAllAccounts());
  const [, createAccount] = useAxiosSafely(urlCreateAccount());
  const [, updateAccount] = useAxiosSafely(urlUpdateAccount());
  const [, deleteAccount] = useAxiosSafely(urlDeleteAccount());

  useEffect(() => {
    refetchAccounts();
  }, []);
  const formikAccount = useFormik({
    initialValues: modalData,
    enableReinitialize: true,
    onSubmit: values => {
      switch (modalMode) {
        case FormTypes.CREATE:
          createAccount({
            data: values
          }).then(() => {
            handleCloseNewAccountModal();
            refetchAccounts();
          });
          break;
        case FormTypes.UPDATE:
          updateAccount({
            data: transformValuesToUpdateIdsPayload(values)
          }).then(() => {
            handleCloseNewAccountModal();
            refetchAccounts();
          });
          break;
        default:
          throw new Error("unrecognized FormType in modalMode");
      }
    }
  });

  const resetModalData = () => setModalData(initialModalData);
  const handleOpenNewAccountModal = () => setNewAccountModelIsOpen(true);
  const handleCloseNewAccountModal = () => {
    setNewAccountModelIsOpen(false);
  };
  const handleOnCreate = () => {
    setModalMode(FormTypes.CREATE);
    resetModalData();
    handleOpenNewAccountModal();
  };
  const handleOnEdit = data => {
    setModalMode(FormTypes.UPDATE);
    setModalData(data);
    handleOpenNewAccountModal();
  };
  const handleOnDelete = data => {
    deleteAccount({ data: { ids: Array.of(data.id) } }).then(() =>
      refetchAccounts()
    );
  };

  const loading = accountsLoading;
  const createOrEditAccountModalProps = {
    title: modalMode === FormTypes.CREATE ? "New Account" : "Edit Account",
    data: modalData,
    loading,
    isOpen: newAccountModelIsOpen,
    handleClose: handleCloseNewAccountModal,
    formik: formikAccount,
    formType: modalMode,
    formFields: [
      FormFields.textField({
        label: "Name",
        type: "text",
        name: "name"
      })
    ]
  };
  const accountDataTableProps = {
    headerNames: ["Name"],
    dataFormatters: [R.prop("name")],
    data: accountsData,
    onEdit: handleOnEdit,
    onDelete: handleOnDelete
  };

  return (
    <>
      <Box className="d-flex">
        <Typography variant="h3" display="inline">
          Accounts
        </Typography>
        <Button
          variant="contained"
          color="primary"
          onClick={handleOnCreate}
          className="ml-3"
        >
          Add Account
        </Button>
      </Box>
      <CreateOrEditModal {...createOrEditAccountModalProps} />
      <DataTable {...accountDataTableProps} />
    </>
  );
}
