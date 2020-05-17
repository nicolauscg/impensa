import React from "react";
import cyrpto from "crypto";

import { makeStyles } from "@material-ui/core/styles";
import {
  Modal,
  TextField,
  Button,
  Box,
  Typography,
  InputLabel,
  Select,
  MenuItem,
  FormControl
} from "@material-ui/core";
import { KeyboardDatePicker } from "@material-ui/pickers";

import PlacesAutocomplete from "../../components/PlacesAutocomplete";

const useStyles = makeStyles(theme => ({
  box: {
    position: "absolute",
    boxShadow: theme.shadows[5],
    padding: theme.spacing(2, 4, 3),
    top: "50%",
    left: "50%",
    transform: "translate(-50%, -50%)"
  },
  formControl: {
    margin: theme.spacing(1),
    minWidth: 120
  },
  selectEmpty: {
    marginTop: theme.spacing(2)
  }
}));

export const FormTypes = {
  CREATE: 0,
  UPDATE: 1
};

export const FormFields = {
  textField: ({ label, type, name, textFieldProps = {} }) => formik => (
    <TextField
      label={label}
      type={type}
      InputLabelProps={{
        shrink: true
      }}
      name={name}
      fullWidth={true}
      onChange={formik.handleChange}
      value={formik.values[name]}
      {...textFieldProps}
    />
  ),
  dateField: ({ label, name, keyboardDatePickerProps = {} }) => formik => (
    <KeyboardDatePicker
      disableToolbar
      variant="inline"
      format="DD/MM/YYYY"
      margin="normal"
      label={label}
      name={name}
      fullWidth={true}
      onChange={event => {
        formik.setFieldValue(name, event.toISOString());
      }}
      value={formik.values[name]}
      KeyboardButtonProps={{
        "aria-label": "change date"
      }}
      {...keyboardDatePickerProps}
    />
  ),
  selectField: ({
    label,
    name,
    options,
    optionDisplayer,
    formControlProps = {},
    inputLabelProps = {},
    selectProps = {}
  }) => (formik, classes) => {
    const uniqueId = `${label}-${name}-${cyrpto
      .randomBytes(4)
      .toString("hex")}`;

    return (
      <FormControl
        className={classes.formControl}
        fullWidth={true}
        {...formControlProps}
      >
        <InputLabel id={uniqueId} {...inputLabelProps}>
          {label}
        </InputLabel>
        <Select
          labelId={uniqueId}
          value={formik.values[name]}
          name={name}
          onChange={formik.handleChange}
          {...selectProps}
        >
          {options.map(option => (
            <MenuItem value={option.id} key={option.id}>
              {optionDisplayer(option)}
            </MenuItem>
          ))}
        </Select>
      </FormControl>
    );
  },
  placesAutocompleteField: ({ label, name }) => formik => (
    <PlacesAutocomplete
      label={label}
      name={name}
      value={formik.values[name]}
      handleValueChange={val => formik.setFieldValue(name, val)}
    />
  )
};

export default function CreateOrEditModal({
  title,
  loading,
  isOpen,
  handleClose,
  formik,
  formFields,
  formType
}) {
  const classes = useStyles();

  const body = loading ? (
    <></>
  ) : (
    <Box width="50%" bgcolor="white" className={classes.box}>
      <Typography variant="h4">{title}</Typography>
      <form onSubmit={formik.handleSubmit}>
        {formFields.map(formField => formField(formik, classes))}
        <Button
          variant="contained"
          color="primary"
          type="submit"
          fullWidth={true}
          className="mt-5"
        >
          {formType === FormTypes.CREATE ? "Create" : "Update"}
        </Button>
      </form>
    </Box>
  );

  return (
    <div>
      <Modal
        open={isOpen}
        onClose={handleClose}
        aria-labelledby="simple-modal-title"
        aria-describedby="simple-modal-description"
      >
        {body}
      </Modal>
    </div>
  );
}
