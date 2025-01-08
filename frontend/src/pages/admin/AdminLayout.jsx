import { Box, Container } from "@mui/material";
import NavBar from "./components_admin_page/Navbar/NavBar";
import MainContent from "./components_admin_page/MainContent/MainContent";

export default function AdminDashboard() {
  return (
    <Box sx={{ flexGrow: 1 }}>
      <NavBar />
      <Container sx={{ mt: 4, mb: 4 }}>
        <MainContent />
      </Container>
    </Box>
  );
}
