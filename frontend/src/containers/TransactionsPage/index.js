import React, { useState, useEffect, useRef } from "react";
import { useFormik } from "formik";
import * as R from "ramda";
import { throttle } from "lodash";

import { Button, Typography, Box, IconButton } from "@material-ui/core";
import KeyboardArrowRightIcon from "@material-ui/icons/KeyboardArrowRight";
import KeyboardArrowLeftIcon from "@material-ui/icons/KeyboardArrowLeft";

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
  cleanEmptyFromObject,
  transformValuesToUpdateIdsPayload
} from "../../ramdaHelpers";
import CreateOrEditModal, {
  FormFields,
  FormTypes
} from "../../components/CreateOrEditModal";

const moment = require("moment");

const TransactionsPage = () => {
  const initialModalData = {
    amount: 0,
    description: "",
    dateTime: new Date(),
    account: "",
    category: "",
    picture: "",
    location: ""
  };

  const [newTransactionModelIsOpen, setNewTransactionModelIsOpen] = useState(
    false
  );
  const [modalData, setModalData] = useState(initialModalData);
  const [modalMode, setModalMode] = useState(FormTypes.CREATE);
  const [infiniteScrollData, setInfiniteScrollData] = useState([]);
  const [scrollHasNext, setScrollHasNext] = useState(false);
  const [scrollNextUrl, setScrollNextUrl] = useState(null);
  const [currentMonthViewed, setMonthViewed] = useState(
    moment().startOf("month")
  );
  const setNextMonth = () =>
    setMonthViewed(currentMonthViewed.clone().add(1, "month"));
  const setPreviousMonth = () =>
    setMonthViewed(currentMonthViewed.clone().subtract(1, "month"));
  const getStartOfCurrentMonthAsString = () =>
    currentMonthViewed.format("YYYY-MM-DD");
  const getEndOfCurrentMonthAsString = () =>
    currentMonthViewed
      .clone()
      .endOf("month")
      .format("YYYY-MM-DD");

  const [
    {
      data: transactionsData,
      loading: transactionsLoading,
      paging: transactionPaging
    },
    refetchTransactionsFromAxios
  ] = useAxiosSafely(
    urlGetAllTransactions({
      params: {
        dateTimeStart: getStartOfCurrentMonthAsString(),
        dateTimeEnd: getEndOfCurrentMonthAsString()
      }
    })
  );

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
    refetchTransactionsFromAxios({
      url: scrollNextUrl,
      params: {
        dateTimeStart: getStartOfCurrentMonthAsString(),
        dateTimeEnd: getEndOfCurrentMonthAsString()
      }
    });
  };

  const refetchTransactions = args => {
    const test = R.mergeDeepRight(args, {
      params: {
        dateTimeStart: getStartOfCurrentMonthAsString(),
        dateTimeEnd: getEndOfCurrentMonthAsString()
      }
    });

    setInfiniteScrollData([]);
    setScrollHasNext(null);
    setScrollNextUrl(null);
    refetchTransactionsFromAxios(test);
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
    refetchTransactionsFromAxios({
      params: {
        dateTimeStart: getStartOfCurrentMonthAsString(),
        dateTimeEnd: getEndOfCurrentMonthAsString()
      }
    });
    fetchAccounts();
    fetchCategories();
  }, []);

  const isFirstRun = useRef(true);
  useEffect(() => {
    if (isFirstRun.current) {
      isFirstRun.current = false;
    } else {
      refetchTransactions();
    }
  }, [currentMonthViewed]);

  const [, createTransaction] = useAxiosSafely(urlCreateTransaction());
  const [, updateTransaction] = useAxiosSafely(urlUpdateTransaction());
  const [, deleteTransaction] = useAxiosSafely(urlDeleteTransaction());

  const formikTrasaction = useFormik({
    initialValues: modalData,
    enableReinitialize: true,
    onSubmit: values => {
      const cleanedValues = R.pipe(
        R.evolve({
          amount: R.ifElse(R.isEmpty, R.always(0), R.identity)
        }),
        cleanEmptyFromObject
      )(values);
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
  const handleOnLoad = throttle(() => {
    if (!loading && scrollHasNext && infiniteScrollData.length > 0) {
      loadMoreTransactions();
    }
  }, 250);
  const throttledAndCancel = () => {
    handleOnLoad.cancel();
    handleOnLoad();
  };

  const loading = transactionsLoading || accountsLoading || categoriesLoading;
  const createOrEditTransactionModalProps = {
    title:
      modalMode === FormTypes.CREATE ? "New Transaction" : "Edit Transaction",
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
      }),
      FormFields.placesAutocompleteField({
        label: "Location",
        name: "location"
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
    onLoadMore: throttledAndCancel,
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
      <div className="d-flex flex-row justify-content-center align-items-center">
        <IconButton className="mr-3" onClick={setPreviousMonth}>
          <KeyboardArrowLeftIcon fontSize="large" />
        </IconButton>
        <Typography>{currentMonthViewed.format("MMM YYYY")}</Typography>
        <IconButton className="ml-3" onClick={setNextMonth}>
          <KeyboardArrowRightIcon fontSize="large" />
        </IconButton>
      </div>
      <CreateOrEditModal {...createOrEditTransactionModalProps} />
      <MuiDataTable {...transactionDataTableProps} />
    </>
  );
};

export default TransactionsPage;
