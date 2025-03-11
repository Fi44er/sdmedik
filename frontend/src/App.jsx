import { Box, Stack, Typography } from "@mui/material";
import { useState } from "react";
import Header from "./global/header";
import Footer from "./global/footer";
import { RouterProvider } from "react-router-dom";
import { router } from "./routers/routers";
import { ToastContainer, toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";

function App() {
  return (
    <Box>
      <RouterProvider router={router}/>
        <ToastContainer />
    </Box>
  );
}

export default App;
