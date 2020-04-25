import React, { useContext } from "react";

import { makeStyles } from "@material-ui/core/styles";
import clsx from "clsx";
import {
  CardHeader,
  Avatar,
  IconButton,
  MenuItem,
  Popper,
  MenuList,
  Grow,
  Paper,
  ClickAwayListener
} from "@material-ui/core";
import { red } from "@material-ui/core/colors";
import ExpandMoreIcon from "@material-ui/icons/ExpandMore";

import { getUserObject, clearUserObject } from "../../auth";
import { UserContext } from "../../containers/App/index";

const useStyles = makeStyles(theme => ({
  profileBadgeRoot: {
    width: "fit-content",
    position: "absolute",
    top: 0,
    right: 0,
    paddingRight: "inherit",
    marginRight: "inherit"
  },
  cardHeader: {
    padding: "4px"
  },
  cardHeaderAction: {
    marginTop: 0,
    marginRight: 0
  },
  expand: {
    transform: "rotate(0deg)",
    marginLeft: "auto",
    transition: theme.transitions.create("transform", {
      duration: theme.transitions.duration.shortest
    })
  },
  expandOpen: {
    transform: "rotate(180deg)"
  },
  avatar: {
    backgroundColor: red[500]
  }
}));

export default function ProfileBadge({ history }) {
  const username = getUserObject().username;
  const classes = useStyles();
  const { data: userData } = useContext(UserContext);

  const [open, setOpen] = React.useState(false);
  const anchorRef = React.useRef(null);

  const handleToggle = () => {
    setOpen(prevOpen => !prevOpen);
  };

  const handleClose = event => {
    if (anchorRef.current && anchorRef.current.contains(event.target)) {
      return;
    }

    setOpen(false);
  };

  function handleListKeyDown(event) {
    if (event.key === "Tab") {
      event.preventDefault();
      setOpen(false);
    }
  }

  // return focus to the button when we transitioned from !open -> open
  const prevOpen = React.useRef(open);
  React.useEffect(() => {
    if (prevOpen.current === true && open === false) {
      anchorRef.current.focus();
    }

    prevOpen.current = open;
  }, [open]);

  return (
    <div className={classes.profileBadgeRoot}>
      <CardHeader
        ref={anchorRef}
        classes={{
          action: classes.cardHeaderAction
        }}
        avatar={
          <Avatar
            aria-label="recipe"
            className={classes.avatar}
            alt={userData.username}
            src={
              userData.picture
                ? `data:image/png;base64,${userData.picture}`
                : null
            }
          />
        }
        action={
          <IconButton
            className={clsx(classes.expand, {
              [classes.expandOpen]: open
            })}
            onClick={handleToggle}
            aria-expanded={open}
            aria-label="show more"
            aria-controls={open ? "menu-list-grow" : undefined}
            aria-haspopup="true"
          >
            <ExpandMoreIcon />
          </IconButton>
        }
        className={classes.cardHeader}
        title={`${userData.username.substring(0, 10)}${
          username.length > 10 ? "..." : ""
        }`}
      />
      <Popper
        open={open}
        anchorEl={anchorRef.current}
        role={undefined}
        transition
        disablePortal
        style={{
          zIndex: 150
        }}
      >
        {({ TransitionProps, placement }) => (
          <Grow
            {...TransitionProps}
            style={{
              transformOrigin:
                placement === "bottom" ? "center top" : "center bottom"
            }}
          >
            <Paper>
              <ClickAwayListener onClickAway={handleClose}>
                <MenuList
                  autoFocusItem={open}
                  id="menu-list-grow"
                  onKeyDown={handleListKeyDown}
                >
                  <MenuItem
                    onClick={event => {
                      handleClose(event);
                      history.push("/profile");
                    }}
                  >
                    Profile
                  </MenuItem>
                  <MenuItem
                    onClick={() => {
                      clearUserObject();
                      history.push("/");
                    }}
                  >
                    Logout
                  </MenuItem>
                </MenuList>
              </ClickAwayListener>
            </Paper>
          </Grow>
        )}
      </Popper>
    </div>
  );
}
