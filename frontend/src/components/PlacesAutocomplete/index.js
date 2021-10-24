import React from "react";
import usePlacesAutocomplete from "use-places-autocomplete";
import { TextField, Grid, Typography } from "@material-ui/core";
import LocationOnIcon from "@material-ui/icons/LocationOn";
import Autocomplete from "@material-ui/lab/Autocomplete";
import { makeStyles } from "@material-ui/core/styles";
import parse from "autosuggest-highlight/parse";
const useStyles = makeStyles(theme => ({
  icon: {
    color: theme.palette.text.secondary,
    marginRight: theme.spacing(2)
  }
}));

export default function PlacesAutocomplete({
  label,
  name,
  handleValueChange,
  value
}) {
  const classes = useStyles();
  const {
    ready,
    suggestions: { status, data },
    setValue
  } = usePlacesAutocomplete({
    debounce: 300
  });

  return (
    <Autocomplete
      fullWidth
      freeSolo
      disabled={!ready}
      options={status === "OK" ? data : []}
      getOptionLabel={option =>
        typeof option === "string" ? option : option.description
      }
      name={name}
      value={value}
      onChange={(e, val) => {
        setValue(val);
        handleValueChange(val.description);
      }}
      onInputChange={(e, val) => {
        setValue(val);
        handleValueChange(val);
      }}
      renderInput={params => (
        <TextField {...params} label={label} name={name} variant="outlined" />
      )}
      renderOption={option => {
        const matches =
          option.structured_formatting.main_text_matched_substrings;
        const parts = parse(
          option.structured_formatting.main_text,
          matches.map(match => [match.offset, match.offset + match.length])
        );

        return (
          <Grid container alignItems="center">
            <Grid item>
              <LocationOnIcon className={classes.icon} />
            </Grid>
            <Grid item xs>
              {parts.map((part, index) => (
                <span
                  key={index}
                  style={{ fontWeight: part.highlight ? 700 : 400 }}
                >
                  {part.text}
                </span>
              ))}

              <Typography variant="body2" color="textSecondary">
                {option.structured_formatting.secondary_text}
              </Typography>
            </Grid>
          </Grid>
        );
      }}
    />
  );
}
