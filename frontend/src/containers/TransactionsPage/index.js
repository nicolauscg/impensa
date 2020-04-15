import React from "react";
import useAxios from "axios-hooks";
import { urlGetAllTransactions } from "../../api";

import { makeStyles } from "@material-ui/core/styles";
import {
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  Paper
} from "@material-ui/core";

const useStyles = makeStyles({
  table: {
    minWidth: 650
  }
});

const TransactionsPage = () => {
  const classes = useStyles();
  const [{ data: transactionData, loading: transactionLoading }] = useAxios(
    urlGetAllTransactions()
  );

  return (
    <>
      <h1>Transactions</h1>
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
            {!transactionLoading &&
              transactionData.data.map(row => (
                <TableRow key={row.id}>
                  <TableCell component="th" scope="row">
                    {row.dateTime}
                  </TableCell>
                  <TableCell align="right">{row.description}</TableCell>
                  <TableCell align="right">{row.category}</TableCell>
                  <TableCell align="right">{row.account}</TableCell>
                  <TableCell align="right">{row.amount}</TableCell>
                </TableRow>
              ))}
          </TableBody>
        </Table>
      </TableContainer>
    </>
  );
};

export default TransactionsPage;
