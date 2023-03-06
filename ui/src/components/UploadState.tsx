import { Button } from "@mui/material";
import { styled } from "@mui/system";
import * as React from "react";

const Input = styled('input')({
  display: 'none',
});

const UploadButton = (props: any) => {
  return (
    <label htmlFor="contained-button-${props.name}">
      <Input
        accept="application/*"
        id="contained-button-${props.name}"
        multiple
        type="file"
        onChange={props.onChange}
      />
      <Button variant="contained" component="span">
        Upload
      </Button>
    </label>
  )
}
export default UploadButton;