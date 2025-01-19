import { Box, Container } from "@mui/material";
import NavBar from "./components_admin_page/Navbar/NavBar";
import MainContent from "./components_admin_page/MainContent/MainContent";
import { Route, Routes } from "react-router-dom";
import { Navigate } from "react-router-dom";
import CreateProduct from "./create_product/CreateProduct";
import CreateCategory from "./create_category/CreateCategory";
import AdminCategoriesTable from "./components_admin_page/AdminCategoriesTable/AdminCategoriesTable";
import AdminProductTable from "./components_admin_page/AdminProductTable/AdminProductTable";
import AdminUserTable from "./components_admin_page/AdminUserTable/AdminUserTable";

export default function AdminDashboard() {
  return (
    <Box sx={{ flexGrow: 1 }}>
      <NavBar />
      <Container sx={{ mt: 4, mb: 4 }}>
        <Routes>
          <Route path="/" element={<MainContent />} />
          <Route path="/create_product" element={<CreateProduct />} />
          <Route path="/create_category" element={<CreateCategory />} />
          {/* <Route path="/edit_product/:id" element={<EditProduct />} /> */}
          <Route path="/table_category" element={<AdminCategoriesTable />} />
          <Route path="/table_product" element={<AdminProductTable />} />
          <Route path="/table_user" element={<AdminUserTable />} />
        </Routes>
      </Container>
    </Box>
  );
}
