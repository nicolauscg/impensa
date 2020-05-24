import React, { useCallback } from "react";
import { useFormik } from "formik";
import { useAxiosSafely, urlImportTransaction } from "../../api";
import { useDropzone } from "react-dropzone";
import { fileToText } from "../../ramdaHelpers";

import { Button } from "@material-ui/core";

const ImportExportPage = () => {
  const [, postImportTransaction] = useAxiosSafely(urlImportTransaction());

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

  const onDrop = useCallback(acceptedFiles => {
    formikImportTransactions.setFieldValue("csv", acceptedFiles[0]);
  }, []);
  const { getRootProps, getInputProps } = useDropzone({ onDrop });

  return (
    <>
      <h2>Import & Export</h2>

      <form onSubmit={formikImportTransactions.handleSubmit}>
        <div {...getRootProps()}>
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
    </>
  );
};

export default ImportExportPage;
