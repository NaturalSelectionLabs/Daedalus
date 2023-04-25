import {
  FormControl,
  List,
  ListItem,
  ListSubheader,
  ListItemButton,
  ListItemIcon,
  Collapse,
  ListItemText,
  TextField,
} from "@mui/material";
import InboxIcon from "@mui/icons-material/MoveToInbox";
import DraftsIcon from "@mui/icons-material/Drafts";
import SendIcon from "@mui/icons-material/Send";
import ExpandLess from "@mui/icons-material/ExpandLess";
import ExpandMore from "@mui/icons-material/ExpandMore";
import StarBorder from "@mui/icons-material/StarBorder";
import { Add, Remove } from "@mui/icons-material";
import React, { useState } from "react";

const Routes = (props: ModelProps<string[]>) => {
  const [listOpen, setListOpen] = useState<Array<boolean>>([]);

  const handleAddItem = () => {
    props.onChange([...props.value, ""]);
  };

  const handleRemoveItem = (index: number) => {
    const newRoutes = [...props.value];
    newRoutes.splice(index, 1);
    props.onChange(newRoutes);
  };

  const handleItemChange = (index: number, value: string) => {
    const newRoutes = [...props.value];
    newRoutes[index] = value;
    props.onChange(newRoutes);
  };
  return (
    <FormControl fullWidth>
      <List
        sx={{ width: "100%", maxWidth: 480, bgcolor: "background.paper" }}
        component="nav"
        aria-labelledby="nested-list-subheader"
        subheader={
          <ListSubheader component="div" id="nested-list-subheader">
            Routes List
          </ListSubheader>
        }
      >
        {props.value && props.value.length
          ? props.value.map((item, index) => (
              <ListItem
                key={index}
                secondaryAction={
                  <ListItemButton>
                    <Remove onClick={() => handleRemoveItem(index)} />
                  </ListItemButton>
                }
              >
                <TextField
                  sx={{ width: "90%" }}
                  value={item}
                  onChange={(event) =>
                    handleItemChange(index, event.target.value)
                  }
                />
              </ListItem>
            ))
          : null}
        <ListItemButton onClick={handleAddItem}>
          <ListItemIcon>
            <Add />
          </ListItemIcon>
          <ListItemText primary="Add Routes" />
        </ListItemButton>
      </List>
    </FormControl>
  );
};

export default Routes;
