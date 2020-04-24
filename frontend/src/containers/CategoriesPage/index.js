import React, { useState, useEffect } from "react";
import { useFormik } from "formik";
import * as R from "ramda";

import { Button, Typography, Box } from "@material-ui/core";

import {
  useAxiosSafely,
  urlUpdateCategory,
  urlDeleteCategory,
  urlCreateCategory,
  urlGetAllCategories
} from "../../api";
import DataTable from "../../components/DataTable";
import { transformValuesToUpdateIdsPayload } from "../../ramdaHelpers";
import CreateOrEditModal, {
  FormFields,
  FormTypes
} from "../../components/CreateOrEditModal";

export default function CategoriesPage() {
  const initialModalData = {
    name: ""
  };

  const [newCategoryModelIsOpen, setNewCategoryModelIsOpen] = useState(false);
  const [modalData, setModalData] = useState(initialModalData);
  const [modalMode, setModalMode] = useState(FormTypes.CREATE);

  const [
    { data: categoriesData, loading: categoriesLoading },
    refetchCategories
  ] = useAxiosSafely(urlGetAllCategories());
  useEffect(() => {
    refetchCategories();
  }, []);
  const [, createCategory] = useAxiosSafely(urlCreateCategory());
  const [, updateCategory] = useAxiosSafely(urlUpdateCategory());
  const [, deleteCategory] = useAxiosSafely(urlDeleteCategory());

  const formikCategory = useFormik({
    initialValues: modalData,
    enableReinitialize: true,
    onSubmit: values => {
      switch (modalMode) {
        case FormTypes.CREATE:
          createCategory({
            data: values
          }).then(() => {
            handleCloseNewCategoryModal();
            refetchCategories();
          });
          break;
        case FormTypes.UPDATE:
          updateCategory({
            data: transformValuesToUpdateIdsPayload(values)
          }).then(() => {
            handleCloseNewCategoryModal();
            refetchCategories();
          });
          break;
        default:
          throw new Error("unrecognized FormType in modalMode");
      }
    }
  });

  const resetModalData = () => setModalData(initialModalData);
  const handleOpenNewCategoryModal = () => setNewCategoryModelIsOpen(true);
  const handleCloseNewCategoryModal = () => {
    setNewCategoryModelIsOpen(false);
  };
  const handleOnCreate = () => {
    setModalMode(FormTypes.CREATE);
    resetModalData();
    handleOpenNewCategoryModal();
  };
  const handleOnEdit = data => {
    setModalMode(FormTypes.UPDATE);
    setModalData(data);
    handleOpenNewCategoryModal();
  };
  const handleOnDelete = data => {
    deleteCategory({ data: { ids: Array.of(data.id) } }).then(() =>
      refetchCategories()
    );
  };

  const loading = categoriesLoading;
  const createOrEditCategoryModalProps = {
    title: modalMode === FormTypes.CREATE ? "New Category" : "Edit Category",
    data: modalData,
    loading,
    isOpen: newCategoryModelIsOpen,
    handleClose: handleCloseNewCategoryModal,
    formik: formikCategory,
    formType: modalMode,
    formFields: [
      FormFields.textField({
        label: "Name",
        type: "text",
        name: "name"
      })
    ]
  };
  const categoryDataTableProps = {
    headerNames: ["Name"],
    dataFormatters: [R.prop("name")],
    data: categoriesData,
    onEdit: handleOnEdit,
    onDelete: handleOnDelete
  };

  return (
    <>
      <Box className="d-flex">
        <Typography variant="h3" display="inline">
          Categories
        </Typography>
        <Button
          variant="contained"
          color="primary"
          onClick={handleOnCreate}
          className="ml-3"
        >
          Add Category
        </Button>
      </Box>
      <CreateOrEditModal {...createOrEditCategoryModalProps} />
      <DataTable {...categoryDataTableProps} />
    </>
  );
}
