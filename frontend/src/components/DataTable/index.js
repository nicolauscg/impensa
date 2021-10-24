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
  Paper,
  Tooltip,
  IconButton
} from "@material-ui/core";
import { Edit, Delete } from "@material-ui/icons";

import { objFromListWith } from "../../ramdaHelpers";

const moment = require("moment");

const useStyles = makeStyles({
  table: {
    minWidth: 650,
    tableLayout: "auto"
  },
  tableCellFitContent: {
    width: "1%",
    whiteSpace: "nowrap"
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

export default function DataTable({
  headerNames,
  dataFormatters,
  data,
  onEdit,
  onDelete
}) {
  const classes = useStyles();

  return (
    <TableContainer component={Paper}>
      <Table
        className={classes.table}
        aria-label="simple table"
        fixedHeader={false}
      >
        <TableHead>
          <TableRow>
            {headerNames.map((headerName, i) => (
              <TableCell key={i}>{headerName}</TableCell>
            ))}
            <TableCell />
          </TableRow>
        </TableHead>
        <TableBody>
          {data.map(row => (
            <TableRow key={row.id}>
              {dataFormatters.map((dataFormatter, columnIndex) => (
                <TableCell key={columnIndex}>{dataFormatter(row)}</TableCell>
              ))}
              <TableCell className={classes.tableCellFitContent}>
                <Tooltip title="Edit">
                  <IconButton aria-label="edit" onClick={() => onEdit(row)}>
                    <Edit />
                  </IconButton>
                </Tooltip>
                <Tooltip title="Delete" onClick={() => onDelete(row)}>
                  <IconButton aria-label="delete">
                    <Delete />
                  </IconButton>
                </Tooltip>
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </TableContainer>
  );
}
