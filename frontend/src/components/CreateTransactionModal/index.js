import React from "react";
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

export default function NewTransactionModal({
  loading,
  isOpen,
  handleClose,
  formik,
  accounts,
  categories
}) {
  const classes = useStyles();

  const body = loading ? (
    <></>
  ) : (
    <Box width="50%" bgcolor="white" className={classes.box}>
      <Typography variant="h4">New Transaction</Typography>
      <form onSubmit={formik.handleSubmit}>
        <TextField
          label="Amount"
          type="number"
          InputLabelProps={{
            shrink: true
          }}
          name="amount"
          fullWidth={true}
          onChange={formik.handleChange}
          value={formik.values.amount}
        />
        <KeyboardDatePicker
          disableToolbar
          variant="inline"
          format="DD/MM/YYYY"
          margin="normal"
          id="date-picker-inline"
          label="Date"
          name="date"
          fullWidth={true}
          onChange={event => {
            formik.setValues({ date: event.toISOString() });
          }}
          value={formik.values.date}
          KeyboardButtonProps={{
            "aria-label": "change date"
          }}
        />
        <TextField
          label="Description"
          type="text"
          InputLabelProps={{
            shrink: true
          }}
          name="description"
          fullWidth={true}
          onChange={formik.handleChange}
          value={formik.values.description}
        />
        <FormControl className={classes.formControl} fullWidth={true}>
          <InputLabel id="create-transaction-category-label">
            Category
          </InputLabel>
          <Select
            labelId="create-transaction-category-label"
            value={formik.values.category}
            name="category"
            onChange={formik.handleChange}
          >
            {categories.map(category => (
              <MenuItem value={category.id} key={category.id}>
                {category.name}
              </MenuItem>
            ))}
          </Select>
        </FormControl>
        <FormControl className={classes.formControl} fullWidth={true}>
          <InputLabel id="create-transaction-account-label">Account</InputLabel>
          <Select
            labelId="create-transaction-account-label"
            name="account"
            value={formik.values.account}
            onChange={formik.handleChange}
          >
            {accounts.map(account => (
              <MenuItem value={account.id} key={account.id}>
                {account.name}
              </MenuItem>
            ))}
          </Select>
        </FormControl>
        <Button
          variant="contained"
          color="primary"
          type="submit"
          fullWidth={true}
          className="mt-5"
        >
          Login
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
