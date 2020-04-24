import React, { useState, useEffect } from "react";
import { useFormik } from "formik";
import * as R from "ramda";

import { Button, Typography, Box } from "@material-ui/core";

import {
  useAxiosSafely,
  urlGetAllTransactions,
  urlCreateTransaction,
  urlGetAllAccounts,
  urlGetAllCategories,
  urlUpdateTransaction,
  urlDeleteTransaction
} from "../../api";
import MuiDataTable, {
  DataTableFormatter
} from "../../components/MuiDataTable";
import {
  cleanNilFromObject,
  transformValuesToUpdateIdsPayload
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
  const [infiniteScrollData, setInfiniteScrollData] = useState([]);
  const [scrollHasNext, setScrollHasNext] = useState(false);
  const [scrollNextUrl, setScrollNextUrl] = useState(null);

  const [
    {
      data: transactionsData,
      loading: transactionsLoading,
      paging: transactionPaging
    },
    refetchTransactionsFromAxios
  ] = useAxiosSafely(urlGetAllTransactions());

  useEffect(() => {
    if (transactionsData && transactionPaging.nextUrl !== scrollNextUrl) {
      setInfiniteScrollData(infiniteScrollData.concat(transactionsData));
    }
    if (transactionPaging) {
      setScrollHasNext(transactionPaging.hasNext);
      setScrollNextUrl(transactionPaging.nextUrl);
    }
  }, [transactionsData]);

  const loadMoreTransactions = () => {
    refetchTransactionsFromAxios({ url: scrollNextUrl });
  };

  const refetchTransactions = (...args) => {
    setInfiniteScrollData([]);
    setScrollHasNext(null);
    setScrollNextUrl(null);
    refetchTransactionsFromAxios(...args);
  };

  const [
    { data: accountsData, loading: accountsLoading },
    fetchAccounts
  ] = useAxiosSafely(urlGetAllAccounts());
  const [
    { data: categoriesData, loading: categoriesLoading },
    fetchCategories
  ] = useAxiosSafely(urlGetAllCategories());

  useEffect(() => {
    refetchTransactionsFromAxios();
    fetchAccounts();
    fetchCategories();
  }, []);
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
              transformValuesToUpdateIdsPayload
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
  const handleOnSearch = queryString => {
    refetchTransactions({
      params: {
        description: queryString
      }
    });
  };
  const handleOnEdit = data => {
    setModalMode(FormTypes.UPDATE);
    setModalData(data);
    handleOpenNewTransactionModal();
  };
  const handleOnDelete = ids => {
    deleteTransaction({ data: { ids } })
      .then(() => {
        refetchTransactions();

        return true;
      })
      .catch(() => false);
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
    headerNames: [
      {
        name: "dateTime",
        label: "Date",
        options: {
          filter: false,
          sort: false
        }
      },
      {
        name: "description",
        label: "Description",
        options: {
          filter: false,
          sort: false
        }
      },
      {
        name: "category",
        label: "Category",
        options: {
          filter: true,
          sort: false,
          filterType: "checkbox",
          filterOptions: {
            names: categoriesData.map(row => row.name)
          }
        }
      },
      {
        name: "account",
        label: "Account",
        options: {
          filter: true,
          display: "true",
          sort: false,
          filterType: "checkbox",
          filterOptions: {
            names: accountsData.map(row => row.name)
          }
        }
      },
      {
        name: "amount",
        label: "Amount",
        options: {
          filter: false,
          sort: false
        }
      }
    ],
    dataFormatters: {
      dateTime: DataTableFormatter.formatDateFrom(R.prop("dateTime")),
      description: R.prop("description"),
      category: DataTableFormatter.mapFromLookup(
        R.prop("category"),
        categoriesData,
        R.prop("name")
      ),
      account: DataTableFormatter.mapFromLookup(
        R.prop("account"),
        accountsData,
        R.prop("name")
      ),
      amount: R.prop("amount")
    },
    data: infiniteScrollData,
    onEdit: handleOnEdit,
    onDelete: handleOnDelete,
    categoriesData,
    onSearch: handleOnSearch,
    onLoadMore: () => {
      if (!loading && scrollHasNext) {
        loadMoreTransactions();
      }
    },
    hasMore: true
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
      <MuiDataTable {...transactionDataTableProps} />
    </>
  );
};

export default TransactionsPage;
