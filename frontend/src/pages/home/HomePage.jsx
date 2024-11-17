import { Box, Container } from "@mui/material";
import React from "react";
import PaymantsInfo from "./components/PaymantsInfo";
import TopList from "./components/TopList";

export default function HomePage() {
  return (
    <Box>
      <Container>
        <Box sx={{ mt: "40px" }}>
          <PaymantsInfo />
        </Box>
        <Box sx={{ mt: "40px" }}>
          <TopList />
        </Box>
      </Container>
    </Box>
  );
}
