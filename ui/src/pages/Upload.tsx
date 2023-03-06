import * as React from "react";
import MenuAppBar from "../components/Header";
import UploadButton from "../components/UploadState";


const Upload = () => {
  // TODO post file data
  // https://zenn.dev/hitoshiasano/articles/0d79429e03bba8
  const [postFileData, setPostFileData] = React.useState({});
  const changeUploadFile = async (event: any) => {
    const { name, files } = event.target;
    setPostFileData({
      ...postFileData,
      [name]: files[0]
    });
    event.target.value = '';
  }
  return (
    <div>
      <MenuAppBar />
      <UploadButton onChange={changeUploadFile}></UploadButton>
    </div>
  )
}

export default Upload;