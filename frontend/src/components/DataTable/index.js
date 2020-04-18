import React from "react";
import * as R from "ramda";

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

import { objFromListWith } from "../../ramdaCookbook";

const moment = require("moment");

const useStyles = makeStyles({
  table: {
    minWidth: 650
  }
});

const dataTableFormatHelper = {
  createIdToDataMap: objFromListWith(R.prop("id"))
};

export const DataTableFormatter = {
  formatDateFrom: keyFunc =>
    R.pipe(
      keyFunc,
      dateIsoString => moment(dateIsoString).format("DD/MM/YYYY")
    ),
  mapFromLookup: (lookupKeyFunc, lookupData, displayDataFunc) => dataRow =>
    R.applyTo(
      dataTableFormatHelper.createIdToDataMap(lookupData),
      R.pipe(
        R.prop(lookupKeyFunc(dataRow)),
        displayDataFunc
      )
    )
};

export default function DataTable({ headerNames, dataFormatters, data }) {
  const classes = useStyles();

  return (
    <TableContainer component={Paper}>
      <Table className={classes.table} aria-label="simple table">
        <TableHead>
          <TableRow>
            {headerNames.map((headerName, i) => (
              <TableCell key={i}>{headerName}</TableCell>
            ))}
          </TableRow>
        </TableHead>
        <TableBody>
          {data.map(row => (
            <TableRow key={row.id}>
              {dataFormatters.map((dataFormatter, columnIndex) => (
                <TableCell key={columnIndex}>{dataFormatter(row)}</TableCell>
              ))}
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </TableContainer>
  );
}
