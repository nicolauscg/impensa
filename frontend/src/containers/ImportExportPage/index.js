import React, { useCallback } from "react";
import { useFormik } from "formik";
import {
  useAxiosSafely,
  urlImportTransaction,
  urlExportTransaction
} from "../../api";
import { useDropzone } from "react-dropzone";
import { fileToText } from "../../ramdaHelpers";
import { FormFields } from "../../components/CreateOrEditModal";

import { makeStyles } from "@material-ui/core/styles";
import { Button, IconButton } from "@material-ui/core";
import FileIcon from "@material-ui/icons/Description";

const fileDownload = require("js-file-download");
const moment = require("moment");

const useStyles = makeStyles({
  rootIcon: {
    width: "5rem",
    height: "5rem"
  }
});

const ImportExportPage = () => {
  const classes = useStyles();
  const [, postImportTransaction] = useAxiosSafely(urlImportTransaction());
  const [, getExportTransaction] = useAxiosSafely(urlExportTransaction());

  const formikImportTransactions = useFormik({
    initialValues: {
      csv: ""
    },
    enableReinitialize: true,
    onSubmit: async values => {
      postImportTransaction({
        data: {
          csv: await fileToText(values.csv)
        }
      });
    }
  });
  const formikExportTransactions = useFormik({
    initialValues: {
      dateTimeStart: moment()
        .startOf("month")
        .format("YYYY-MM-DD"),
      dateTimeEnd: moment()
        .endOf("month")
        .format("YYYY-MM-DD")
    },
    onSubmit: values => {
      getExportTransaction({
        params: values
      }).then(res => {
        fileDownload(res.data, "impensa_transactions.csv");
      });
    }
  });

  const onDrop = useCallback(acceptedFiles => {
    formikImportTransactions.setFieldValue("csv", acceptedFiles[0]);
  }, []);
  const { getRootProps, getInputProps } = useDropzone({ onDrop });

  return (
    <>
      <h2>Import & Export</h2>

      <div className="d-flex align-items-stretch">
        <form
          onSubmit={formikImportTransactions.handleSubmit}
          className="flex-grow-1 d-flex flex-column align-items-center mr-2 justify-content-end"
        >
          <div {...getRootProps()}>
            <IconButton>
              <FileIcon
                classes={{
                  root: classes.rootIcon
                }}
              />
            </IconButton>
            <div>
              {formikImportTransactions.values.csv
                ? formikImportTransactions.values.csv.name
                : "enter csv file"}
            </div>
            <input {...getInputProps()} />
          </div>
          <Button
            variant="contained"
            color="primary"
            type="submit"
            fullWidth={true}
            className="mt-5"
          >
            Import
          </Button>
        </form>

        <form
          onSubmit={formikExportTransactions.handleSubmit}
          className="flex-grow-1 d-flex flex-column align-items-center ml-2 justify-content-end"
        >
          {FormFields.dateField({
            label: "start date",
            name: "dateTimeStart"
          })(formikExportTransactions)}
          {FormFields.dateField({
            label: "end date",
            name: "dateTimeEnd"
          })(formikExportTransactions)}
          <Button
            variant="contained"
            color="primary"
            type="submit"
            fullWidth={true}
            className="mt-5"
          >
            Export
          </Button>
        </form>
      </div>
    </>
  );
};

export default ImportExportPage;
