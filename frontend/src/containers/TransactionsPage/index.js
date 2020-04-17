import React, { useState } from "react";
import { useFormik } from "formik";
import {
  useAxiosSafely,
  urlGetAllTransactions,
  urlPostCreateTransaction,
  urlGetAllAccounts,
  urlGetAllCategory
} from "../../api";

import { makeStyles } from "@material-ui/core/styles";
import {
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  Paper,
  Button,
  Typography,
  Box
} from "@material-ui/core";
import NewTransactionModal from "../../components/CreateTransactionModal";

const useStyles = makeStyles({
  table: {
    minWidth: 650
  }
});

const TransactionsPage = () => {
  const classes = useStyles();

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
      <TableContainer component={Paper}>
        <Table className={classes.table} aria-label="simple table">
          <TableHead>
            <TableRow>
              <TableCell>Date</TableCell>
              <TableCell align="right">Description</TableCell>
              <TableCell align="right">Category</TableCell>
              <TableCell align="right">Account</TableCell>
              <TableCell align="right">Amount</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {transactionsData.map(row => (
              <TableRow key={row.id}>
                <TableCell component="th" scope="row">
                  {row.dateTime}
                </TableCell>
                <TableCell align="right">{row.description}</TableCell>
                <TableCell align="right">{row.category}</TableCell>
                <TableCell align="right">{row.account}</TableCell>
                <TableCell align="right">{row.amount}</TableCell>``
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </TableContainer>
    </>
  );
};

export default TransactionsPage;
